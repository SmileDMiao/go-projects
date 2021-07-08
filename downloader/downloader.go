package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/schollz/progressbar/v3"
)

var bar *progressbar.ProgressBar
var etag string

type Downloader struct {
	concurrency int
}

func NewDownloader(concurrency int) *Downloader {
	return &Downloader{concurrency: concurrency}
}

// if support use multidownloader or not use singledownloader
func (d *Downloader) Download(strURL, filename string) error {
	if filename == "" {
		filename = path.Base(strURL)
	}

	resp, error := http.Head(strURL)

	if error != nil {
		return error
	}

	etag = resp.Header.Get("Etag")

	fmt.Println(etag)
	// 支持部分请求
	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return d.multiDownload(strURL, filename, int(resp.ContentLength))
	}

	// 不支持部分请求
	return d.singleDownload(strURL, filename, int(resp.ContentLength))
}

// concurrency download
func (d *Downloader) multiDownload(strURL, filename string, contentLen int) error {
	bar = barProcessor(contentLen)
	partSize := contentLen / d.concurrency

	// 创建部分文件的存放目录
	// 0777: 权限
	// 结束后删除文件夹
	partDir := d.getPartDir(filename)
	os.Mkdir(partDir, 0777)
	defer os.RemoveAll(partDir)

	// 控制并发, 等待一组Goroutine返回
	var wg sync.WaitGroup
	wg.Add(d.concurrency)

	rangeStart := 0

	for i := 0; i < d.concurrency; i++ {
		// 并发请求
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize
			// 最后一部分，总长度不能超过 ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen
			}

			d.downloadPartial(strURL, filename, rangeStart, rangeEnd, i)
		}(i, rangeStart)

		rangeStart += partSize + 1
	}

	wg.Wait()

	// 合并文件
	d.merge(filename)

	return nil
}

// single download
func (d *Downloader) singleDownload(strURL, filename string, contentLen int) error {
	bar = barProcessor(contentLen)

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	io.Copy(io.MultiWriter(f, bar), resp.Body)

	return nil
}

// partial download
func (d *Downloader) downloadPartial(strURL, filename string, rangeStart, rangeEnd, i int) {
	doneName := d.getPartFilename(filename, i, true)
	if exists(doneName) {
		percent := int(rangeEnd - rangeStart)
		bar.Add(percent)
		return
	}

	if rangeStart >= rangeEnd {
		return
	}

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 打开文件flags
	// CREATE: 创建并打开一个新文件
	// WRONLY: 以只写的方式打开
	flags := os.O_CREATE | os.O_WRONLY
	name := d.getPartFilename(filename, i, false)
	partFile, err := os.OpenFile(name, flags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer partFile.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(io.MultiWriter(partFile, bar), resp.Body, buf)
	if err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	} else {
		doneName := d.getPartFilename(filename, i, true)
		d.tagPartFilename(name, doneName)
	}
}

// getPartDir 部分文件存放的目录
func (d *Downloader) getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

// getPartFilename 构造部分文件的名字
// done: 完成时的文件名不一样添加标记
func (d *Downloader) getPartFilename(filename string, partNum int, done bool) string {
	partDir := d.getPartDir(filename)
	if done {
		return fmt.Sprintf("%s/%s-%d-%s", partDir, filename, partNum, "Done")
	} else {
		return fmt.Sprintf("%s/%s-%d", partDir, filename, partNum)
	}
}

// 完成时添加标记 重命名
func (d *Downloader) tagPartFilename(undone, done string) {
	os.Rename(undone, done)
}

func (d *Downloader) merge(filename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(filename, i, true)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		io.Copy(destFile, partFile)
		partFile.Close()
		os.Remove(partFileName)
	}

	return nil
}

func barProcessor(contentLen int) *progressbar.ProgressBar {
	return progressbar.NewOptions(contentLen,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription("Dwonloading"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}

// 判断文件是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
