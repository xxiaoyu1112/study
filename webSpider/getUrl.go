package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func main() {
	links := findLinks("https://baidu.com")
	for _, link := range links {
		method, _ := getRequestMethodAndBody(link)
		fmt.Printf("%s %s\n", method, link)
		//fmt.Printf("Response body: %s\n", body)
	}
}

/*
*
使用 findLinks 函数来获取 HTML 页面中的所有链接。
该函数首先使用 http.Get 方法来发送 HTTP GET 请求，获取页面的 HTML 内容。
然后，它使用 html.Parse 方法来解析 HTML 文档，生成 DOM 树。
接下来，它遍历 DOM 树，查找包含链接的节点，并将链接添加到链接列表中。
最后，它对链接进行规范化，以去除重复的链接和无效的链接，并返回链接列表。
*/
func findLinks(url string) []string {
	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []string{}
	}
	defer resp.Body.Close()

	// Parse HTML document
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []string{}
	}

	// Find links in HTML document
	links := []string{}
	var findLinks func(*html.Node, string)
	findLinks = func(n *html.Node, base string) {
		if n.Type == html.ElementNode {
			if n.Data == "a" {
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						link := normalizeLink(attr.Val, base)
						if link != "" && !contains(links, link) {
							links = append(links, link)
						}
						break
					}
				}
			} else if n.Data == "link" || n.Data == "img" || n.Data == "script" || n.Data == "iframe" {
				for _, attr := range n.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						link := normalizeLink(attr.Val, base)
						if link != "" && !contains(links, link) {
							links = append(links, link)
						}
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLinks(c, base)
		}
	}
	findLinks(doc, url)

	return links
}

func normalizeLink(link string, base string) string {
	// Ignore empty links
	if link == "" {
		return ""
	}

	// Ignore mailto: and tel: links
	if strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "tel:") {
		return ""
	}

	// Ignore javascript: links
	if strings.HasPrefix(link, "javascript:") {
		return ""
	}

	// Ignore anchor links
	if strings.HasPrefix(link, "#") {
		return ""
	}

	// Normalize relative links
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		if strings.HasPrefix(link, "/") {
			link = base + link
		} else {
			link = base + "/" + link
		}
	}

	return link
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func getRequestMethodAndBody(url string) (string, string) {
	// Send HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "", ""
	}
	defer resp.Body.Close()

	// Get request method and response body
	method := resp.Request.Method
	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("Error: %s\n", err)
	//	return "", ""
	//}
	//bodyString := string(bodyBytes)

	return method, ""
}
