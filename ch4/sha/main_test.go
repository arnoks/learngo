package main

import (
	"io"
	"reflect"
	"testing"
)

func Test_readInput(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name string
		args args
		want *[]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(tt.args.rd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
