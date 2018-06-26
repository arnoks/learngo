// fetch prints the content found at url to stdout
// fetch <url> > content
package main

/*
fetch 17 is using io.Copy(dst,src)instead of ioutil.ReadAll to copy the
response body.

os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to
check the error result of io.Copy
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		src := Open(url)
		dst := os.Stdio
		_, err := io.Copy(dst, src)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
