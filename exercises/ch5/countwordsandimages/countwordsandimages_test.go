package main

import (
	"fmt"
	"testing"
)

func TestCountWordsAndImages1(t *testing.T) {
	site := "http://intranet.deutsche-boerse.de"
	words, images, err := CountWordsAndImages(site)
	if err != nil {
		t.Errorf("Error parsing test site %v", err)
		return
	}
	fmt.Printf("%s has %d words and %d images", site, words, images)
}

func TestCountWordsAndImages2(t *testing.T) {
	site := "https://heise.de"
	words, images, err := CountWordsAndImages(site)
	if err != nil {
		t.Errorf("Error parsing test site %v", err)
		return
	}
	fmt.Printf("%s has %d words and %d images", site, words, images)
}
