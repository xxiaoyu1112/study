package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 发送HTTP请求获取网页内容
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 用goquery解析HTML页面
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 提取网页标题
	title := doc.Find("title").Text()
	fmt.Println("title:", title)

	// 提取所有链接
	//doc.Find("a").Each(func(i int, s *goquery.Selection) {
	//	link, exists := s.Attr("href")
	//	if exists {
	//		fmt.Println(link)
	//	}
	//})
}
