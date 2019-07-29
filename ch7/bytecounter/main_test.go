package main

import (
	"testing"
)

func TestByteCounter_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		c       *ByteCounter
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByteCounter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ByteCounter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordCounter_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		w       *WordCounter
		args    args
		want    int
		wantErr bool
	}{
		{"no words", new(WordCounter), args{[]byte("")}, 0, false},
		{"two Words", new(WordCounter), args{[]byte("One, Two")}, 2, false},
		{"three words", new(WordCounter), args{[]byte("One, Two, Three")}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.w.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("WordCounter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WordCounter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineCounter_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		l       *LineCounter
		args    args
		want    int
		wantErr bool
	}{
		{"no lines", new(LineCounter), args{[]byte("")}, 0, false},
		{"two lines", new(LineCounter), args{[]byte("\n\n")}, 2, false},
		{"three line", new(LineCounter), args{[]byte("\n\n\n")}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("LineCounter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LineCounter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
