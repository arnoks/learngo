package newreader

import (
	"fmt"
	"testing"
)

func Test_parseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{"case 1", args{"Hello World\n<a href=\"www.golem.de\">link text</a>\n\n"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseString(tt.args.s)
		})
	}
}

func Test_Reader(t *testing.T) {
	type args struct {
		s string
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Read 10 bytes", args{"Hello World\n<a href=\"www.golem.de\">link text</a>\n\n", make([]byte, 10)}, 10},
		{"Read beyond eof", args{"Hello World", make([]byte, 100)}, len("Hello World")},
		{"Read to short buffer", args{"Hello World", make([]byte, 5)}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newReader(tt.args.s)
			gotN, gotErr := r.Read(tt.args.b)
			if gotErr != nil {
				t.Errorf("Expecting nil got %v\n", gotErr)
			}
			if gotN != tt.want {
				t.Errorf("Expecting %v got %v\n", gotN, tt.want)
			}
		})
	}
}

func Test_LimitReader(t *testing.T) {
	type args struct {
		s     string
		limit int
		b     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr error
	}{
		{"Buf equals limit", args{"Hello Worlds", 5, make([]byte, 5)}, 5, nil},
		{"Buf larger limit", args{"Hello Worlds", 5, make([]byte, 10)}, 5, nil},
		{"Limit larger buffer", args{"Hello Worlds", 5, make([]byte, 2)}, 2, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newReader(tt.args.s)
			l := newLimitReader(r, tt.args.limit)
			gotN, gotErr := l.Read(tt.args.b)
			if gotN != tt.wantN || gotErr != tt.wantErr {
				t.Errorf("Expected to read %v bytes iwitth error: %v - reading: %v bytes error %v\n", gotN, gotErr, tt.wantN, tt.wantErr)
			}
		})
	}

}

func Test_Arrays(t *testing.T) {
	var b []byte
	fmt.Printf("len: %v cap: %v\n", len(b), cap(b))
	b = make([]byte, 10, 20)
	fmt.Printf("len: %v cap: %v\n", len(b), cap(b))
	fmt.Printf("len: %v cap: %v\n", len(b), cap(b[:10]))

}
