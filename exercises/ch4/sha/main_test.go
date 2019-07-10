package main

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_readInput(t *testing.T) {
	testfile := "testdata.txt"
	rd, err := os.Open(testfile)
	if err != nil {
		t.Errorf("Could not open testfile %s", testfile)
	}
	defer rd.Close()

	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Reading from a string",
			args{strings.NewReader("Hello World")}, // I need to create a string reader here
			[]byte("Hello World"),
		},
		{
			"Reading from a file",
			args{rd}, // using the above created file reader rd
			[]byte("Hello File"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(tt.args.rd); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("readInput() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func Test_hash(t *testing.T) {
	type args struct {
		b    []byte
		algo string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"default hash",
			args{[]byte("Hello World"), "sha256"},
			[]byte{165, 145, 166, 212, 11, 244, 32, 64, 74, 1, 23, 51,
				207, 183, 177, 144, 214, 44, 101, 191, 11, 205, 163,
				43, 87, 178, 119, 217, 173, 159, 20, 110},
		},
		{
			"sha384 hash",
			args{[]byte("Hello World"), "sha384"},
			[]byte{153, 81, 67, 41, 24, 107, 47, 106, 228, 161, 50, 158,
				126, 230, 198, 16, 167, 41, 99, 99, 53, 23, 74, 198, 183,
				64, 249, 2, 131, 150, 252, 200, 3, 208, 233, 56, 99, 167,
				195, 217, 15, 134, 190, 238, 120, 47, 79, 63},
		},
		{
			"sha512 hash",
			args{[]byte("Hello World"), "sha512"},
			[]byte{44, 116, 253, 23, 237, 175, 216, 14, 132, 71, 176, 212,
				103, 65, 238, 36, 59, 126, 183, 77, 210, 20, 154, 10, 177,
				185, 36, 111, 179, 3, 130, 242, 126, 133, 61, 133, 133, 113,
				158, 14, 103, 203, 218, 13, 170, 143, 81, 103, 16, 100, 97,
				93, 100, 90, 226, 122, 203, 21, 191, 177, 68, 127, 69, 155},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(&tt.args.b, &tt.args.algo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hash: %s(%s)\ngot:  %v\nwant: %v\n", tt.args.algo, tt.args.b, got, tt.want)
			}

		})
	}
}
