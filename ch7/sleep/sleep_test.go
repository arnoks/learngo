package main

import (
	"os"
	"testing"
)

// test to Sleep

func TestSleep(t *testing.T) {
	type args struct {
		cmdl []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1 Secs", args{[]string{"-period", "1s"}}},
		{"2 Secs", args{[]string{"-period", "2s"}}},
		{"3 Secs", args{[]string{"-period", "3s"}}},
		{"4 Secs", args{[]string{"-period", "4s"}}},
		{"0.1 Min", args{[]string{"-period", "0.1m"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = append(os.Args[:1], tt.args.cmdl...)
			main()
		})
	}
}
