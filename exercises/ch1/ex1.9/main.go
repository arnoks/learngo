// fetch prints the content found at url to stdout
// fetch <url> > content and adds the protocol prefix if requried.
// The status code is printed to Stderr

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		switch {
		case strings.HasPrefix(url, "http://"):
			break
		case strings.HasPrefix(url, "https://"):
			break
		default:
			url = "http://" + url
		}

		resp, err := http.Get(url)
		// Always print the status code to Stderr
		fmt.Fprintf(os.Stderr, "http status code: %d\n", resp.StatusCode)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
