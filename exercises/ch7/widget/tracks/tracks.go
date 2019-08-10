package tracks

import (
	"io"
	"log"
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
	{"Go Gentle ", "Robbie Williams", "Swing Both Ways (Deluxe)", 2013, length("4m31s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track, w io.Writer) error {
	t, err := template.ParseFiles("../tracks/tracks.html")
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

func stableLess(i, j int) bool {
	return false
}

// SortTrackList sort the track list by the given method. The sort should be
// stable.
func SortTrackList(w io.Writer, sortBy string) {
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
