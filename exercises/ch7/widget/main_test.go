package main

import (
	"io"
	"os"
	"testing"
)

func Test_printTracks(t *testing.T) {
	wc, err := os.Create("track.html")
	if wc == nil {
		t.Errorf("Could not create: %v", err)
		return
	}
	defer wc.Close()

	type args struct {
		tracks []*Track
		w      io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Print Tracks", args{goTracks, wc}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErr := printTracks(tt.args.tracks, tt.args.w); (gotErr != nil) != tt.wantErr {
				t.Errorf("Error executing %v", gotErr)
			}
		})
	}
}

func Test_sortTrackList(t *testing.T) {
	wc, err := os.Create("sortedTrackList.html")
	if wc == nil {
		t.Errorf("Could not create: %v", err)
		return
	}
	defer wc.Close()

	type args struct {
		sortBy string
		w      io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{"Default sort", args{"default", wc}},
		{"Titel sort", args{"titel", wc}},
		{"Album sort", args{"album", wc}},
		{"Length sort", args{"length", wc}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortTrackList(tt.args.w, tt.args.sortBy)
		})
	}
}
