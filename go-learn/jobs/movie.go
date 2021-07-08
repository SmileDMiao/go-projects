package job

import (
	"fmt"
	"go-learn/global"
	"go-learn/model"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Page struct
type Page struct {
	Page int
	URL  string
}

// BaseURL douban
var BaseURL = "https://movie.douban.com/top250"

// GetPages 获取分页
func GetPages(url string) []Page {
	res := generateRequest(BaseURL)
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return ParsePages(doc)
}

func generateRequest(url string) *http.Response {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Host", "douban.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")

	res, _ := client.Do(request)
	return res
}

// ParsePages 分析分页
func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, URL: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			URL:  url,
		})
	})

	fmt.Println(pages)
	return pages
}

// ParseMovies 分析电影数据
func ParseMovies(doc *goquery.Document) (movies []model.Movie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		movie := model.Movie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		log.Printf("i: %d, movie: %v", i, movie)

		movies = append(movies, movie)
	})

	return movies
}

// Add 新增数据
func Add(movies []model.Movie) {
	for index, movie := range movies {
		if err := global.DB.Create(&movie).Error; err != nil {
			log.Fatal(index, err)
		}
	}
}

// Start 开始爬取
func Start() {
	var movies []model.Movie

	pages := GetPages(BaseURL)
	for _, page := range pages {
		url := strings.Join([]string{BaseURL, page.URL}, "")
		res := generateRequest(url)
		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, ParseMovies(doc)...)
	}

	Add(movies)
}
