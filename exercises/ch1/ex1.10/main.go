//fetchall fetches all given urls reporting sizes and times
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	// create go routine for each cmd argument and pass in both the url and the channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	// read the reports from the channel and print to Stdout
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	// print total time
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // report get error
		return
	}
	defer resp.Body.Close()

	// create canonical file prefix based on the url

	var pattern string
	switch {
	case strings.HasPrefix(url, "https://"):
		pattern = strings.ReplaceAll(strings.TrimPrefix(url, "https://"), "/", "_")
	case strings.HasPrefix(url, "http://"):
		pattern = strings.ReplaceAll(strings.TrimPrefix(url, "http://"), "/", "_")
	default:
		pattern = strings.ReplaceAll(url, "/", "_")
	}
	pattern += "_*.html"

	destination, err := ioutil.TempFile("", pattern)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer destination.Close()
	nbytes, err := io.Copy(destination, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // report read error
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url) // report size and time
}
