package main

import (
	"fmt"
	"github.com/gocolly/colly/extensions"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Activate extension
	extensions.RandomUserAgent(c)

	// OnResponse callback
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response received", r.StatusCode)
	})

	// OnRequest callback
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Request URL:", r.URL.String())
		fmt.Println("Request Method:", r.Method)
	})

	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("Response URL:", r.Request.URL.String())
	//	fmt.Println("Response Method:", r.Request.Method)
	//	fmt.Println("Response Headers:", r.Headers)
	//})

	err := c.Visit("https://www.bing.com")
	if err != nil {
		return
	}
}
