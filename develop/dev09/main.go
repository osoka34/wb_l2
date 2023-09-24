package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//"https://stackoverflow.com/questions/23190311/reverse-a-map-in-value-key-format-in-golang"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	baseURL := os.Args[1]

	downloadSite(baseURL, baseURL)
}

func downloadSite(baseURL, currentURL string) {
	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP status code", resp.StatusCode)
		return
	}

	// Извлекаем имя файла из URL
	fileName := "index.html"
	if parts := strings.Split(currentURL, "/"); len(parts) > 0 {
		fileName = parts[len(parts)-1]
	}

	file, err := os.Create(strings.Join([]string{fileName, ".html"}, ""))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Downloaded", fileName)

	// Парсим HTML и обрабатываем ссылки
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link := attr.Val
					if !strings.HasPrefix(link, "http") {
						link = baseURL + link
					}
					downloadSite(baseURL, link)
				}
			}
		}
	}

	var traverse func(n *html.Node)
	traverse = func(n *html.Node) {
		visitNode(n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
}
