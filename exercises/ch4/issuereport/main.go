// issue a report of open issues
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/arnoks/learngo/ch4/github"
)

func templ() string {
	templFile := "report.tmpl"
	s, err := ioutil.ReadFile(templFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening template file: %s", templFile)
		os.Exit(1)
	}
	return string(s)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ()))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
