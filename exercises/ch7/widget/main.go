package main

import (
	"io"
	"log"
	"os"
	"sort"
	"text/template"
	"time"
)

// Track a sample audio track data structure
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var goTracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track, w io.Writer) error {
	t, err := template.ParseFiles("tracks.html")
	if err != nil {
		return err
	}
	if err := t.Execute(w, tracks); err != nil {
		return err
	}
	return nil
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func htmlOut() io.WriteCloser {
	wc, err := os.Create("sortedtrack.html")
	if err != nil {
		log.Fatalf("Error creating output file %v", err)
	}
	return wc
}

func sortTrackList(w io.Writer, sortBy string) {
	var sortFunc func(x, y *Track) bool

	switch sortBy {
	case "title":
		sortFunc = func(x, y *Track) bool { return x.Title < y.Title }
	case "artist":
		sortFunc = func(x, y *Track) bool { return x.Artist < y.Artist }
	case "album":
		sortFunc = func(x, y *Track) bool { return x.Album < y.Album }
	case "year":
		sortFunc = func(x, y *Track) bool { return x.Year < y.Year }
	case "length":
		sortFunc = func(x, y *Track) bool { return x.Length < y.Length }
	default:
		if err := printTracks(goTracks, w); err != nil {
			log.Fatalf("Error generating html %v", err)
		}
		return
	}

	sort.Sort(customSort{goTracks, sortFunc})
	if err := printTracks(goTracks, w); err != nil {
		log.Fatalf("Error generating html %v", err)
	}
}

func main() {
	wc := htmlOut()
	defer wc.Close()
	sortTrackList(wc, "")
}
