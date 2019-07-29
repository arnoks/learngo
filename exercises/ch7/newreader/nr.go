package newreader

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// Reader for String
type Reader struct {
	S string
	I int64
}

// Read from a string implementing the Reader interface
func (r *Reader) Read(p []byte) (n int, err error) {
	if r.I >= int64(len(r.S)) {
		return 0, io.EOF
	}
	n = copy(p, r.S[r.I:])
	r.I += int64(n)
	return n, nil
}

func newReader(s string) io.Reader {
	var r Reader
	r.S = s
	return &r
}

// LimitedReader keep track of the remaing bytes to read from the
// underlying Reader
type LimitedReader struct {
	R io.Reader
	N int64 // Bytes remaining
}

// Read from a string implementing the LimitedReader interface
func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	lmt := l.N
	if lmt > int64(cap(p)) {
		lmt = int64(cap(p))
	}
	n, err = l.R.Read(p[:lmt])
	if n == 0 || err == io.EOF {
		l.N = -1
	}
	l.N -= int64(n)
	return n, err
}

// newLimitReader returns a limitreader similar to io.NewLimitedReader
func newLimitReader(r io.Reader, n int) io.Reader {
	var l LimitedReader
	l.R = r
	l.N = int64(n)
	return &l
}

func parseString(s string) {
	doc, err := html.Parse(newReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit append to links each link found in n abd returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
