package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// go语言的一个爬虫框架
func main() {
	c := colly.NewCollector(colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"))
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", "baidu.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Referer", "http://www.baidu.com")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")

		fmt.Println("Visiting", r.URL)
	})
	c.OnScraped(func(response *colly.Response) {
		fmt.Println("OnScraped")
	})
	// 对响应的HTML元素处理
	c.OnHTML("title", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		fmt.Println("title:", e.Text)
	})
	// c.OnHTML("p", func(e *colly.HTMLElement) {
	// 	space := strings.TrimSpace(e.Text)
	// 	if space != "" {
	// 		fmt.Printf("p标签:%s\n", space)
	// 	}
	// })

	// c.OnHTML("body", func(e *colly.HTMLElement) {
	// 	// <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
	// 	e.ForEach(".hotnews a", func(i int, el *colly.HTMLElement) {
	// 		band := el.Attr("href")
	// 		// title := el.Text
	// 		// fmt.Printf("新闻 %d : %s - %s\n", i, title, band)
	// 		e.Request.Visit(band)
	// 	})
	// })
	//
	// // 发现并访问下一个连接
	// c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	//
	// // extract status code
	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Printf("URL:%v code:%v\n", r.Ctx.Get("url"), r.StatusCode)
	// })

	// 对visit的线程数做限制，visit可以同时运行多个
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	err := c.Visit("http://news.baidu.com")
	if err != nil {
		log.Println("first err:", err)
	}
	fmt.Println("------------")
	err = c.Visit("http://www.vaptcha.com")
	if err != nil {
		log.Println("second err:", err)
	}

}
