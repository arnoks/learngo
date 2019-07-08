package main

import (
	"log"
	"thumbnail"
)

func makeThumpnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.Imagefile(f); err != nil {
			log.Println(err)
		}
	}
}
