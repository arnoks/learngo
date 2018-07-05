//
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Errorf("Error opening url %s, %v\n", url, err)
			continue
		}
		fmt.Printf("url: %s: Words %d, Images: %d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		fmt.Printf("%s, %v\n", url, err)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images += 1
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
		n, _, err := bufio.ScanWords(n.Data, true)
		fmt.Printf("Count %d, Error: %v\n")
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
