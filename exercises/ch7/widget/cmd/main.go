package main

import (
	"os"

	"github.com/arnoks/learngo/exercises/ch7/widget/tracks"
)

func main() {
	tracks.SortTrackList(os.Stdout, "album")
}
