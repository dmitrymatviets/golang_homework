package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const defaultUrl = "https://goods.ru"
const imgTag = "img"
const srcAttr = "src"

func main() {
	for i, link := range getImagesFromUrl(getUrlToParse()) {
		fmt.Println(i, link)
	}
}

func getUrlToParse() string {
	var urlToParse = defaultUrl
	if len(os.Args) > 1 {
		urlToParse = strings.TrimSpace(os.Args[1])
		if urlToParse == "" {
			panic(fmt.Errorf("Передан пустой URL в параметре"))
		}
	}
	return urlToParse
}

func getImagesFromUrl(url string) []string {
	fmt.Println("Будет выполнен парсинг изображений на", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(fmt.Errorf("Ошибка загрузки %s: %s", url, err))
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("Ошибка в коде ответа %s: %s", url, resp.Status))
	}
	doc, err := html.Parse(resp.Body)

	if err != nil {
		panic(fmt.Errorf("Ошибка парсинга %s как HTML: %v", url, err))
	}

	imgMap := iterateNodes(make(map[string]bool), doc, resp.Request.URL)

	var links = make([]string, 0)

	for key := range imgMap {
		links = append(links, key)
	}
	return links
}

func iterateNodes(links map[string]bool, n *html.Node, url *url.URL) map[string]bool {
	if n.Type == html.ElementNode && n.Data == imgTag {
		for _, a := range n.Attr {
			if a.Key == srcAttr {
				if a.Val == "" {
					continue
				}

				src, err := url.Parse(a.Val)

				if err != nil {
					fmt.Println("Ошибка парсинга", a.Val, err)
					continue
				}

				links[src.String()] = true
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = iterateNodes(links, c, url)
	}
	return links
}
