// fetch prints the content found at url to stdout
// fetch <url> > content
package main

/*  Exercises:
1.7 The function call io.Copy(dst,src) reads from src and writes to dst.
Use it instead of ioutil.ReadAll to Copy the response body to os.Stdout
without requiring a buffer large enough to hold the entire stream. Be sure to
check the error result of io.Copy

1.8 Modify fetch to add the prefix http:// to each argument URL if it is
missing. You might want to use strings.HasPrefix

1.9 Modify fetch to also print the HTTP status code, found in resp.Status

*/

import (
	"fmt"
	"io/ioutil"
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
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
