## downloader
1. 支持下载进度显示
2. 支持则多线程下载不支持则直接下载
3. 断点续传

#### 原理
Accept-Ranges: 服务器通过该头来标识自身支持部分请求（partial requests），也叫范围请求。如果服务端支持部分请求，我们就可以实现并发下载。该头有两个可能的值:
`bytes` and `none`
1. none: 不支持任何部分请求单位，由于其等同于没有返回此头部，因此很少使用。不过一些浏览器，比如 IE9，会依据该头部去禁用或者移除下载管理器的暂停按钮。
2. bytes: 部分请求的单位是 bytes （字节）

#### 支持参数
1. url必填
2. output file name
3. concurrency 多线程数量

## bookstore
> go-zero framework practice

## go-learn
> Gin framework practice

## personal-utils
> cobra command line tools for myself
