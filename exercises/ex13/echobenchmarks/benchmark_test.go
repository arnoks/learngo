// Benchmark the two implementatations

package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkLoop(b *testing.B) {
	b.Log("loop benchmark")
	args := []string{"eins", "zwei", "drei", "vier", "fuenf", "eins", "zwei", "drei", "vier", "fuenf"}
	for i := 0; i < b.N; i++ {
		var w bytes.Buffer
		Loopargs(args, &w)
	}
}

func BenchmarkJoin(b *testing.B) {
	b.Log("join benchmark")
	args := []string{"eins", "zwei", "drei", "vier", "fuenf", "eins", "zwei", "drei", "vier", "fuenf"}
	for i := 0; i < b.N; i++ {
		var w bytes.Buffer
		Joinargs(args, &w)
	}
}

func TestLoop(t *testing.T) {
	args := []string{"eins", "zwei", "drei", "vier", "fuenf", "eins", "zwei", "drei", "vier", "fuenf"}
	want := bytes.NewBufferString(strings.Join(args, " "))
	var have bytes.Buffer
	Loopargs(args, &have)
	if bytes.Compare(have.Bytes(), want.Bytes()) != 0 {
		t.Errorf("Buffers are not equal have/want \n[%v]\n[%v]\n", have, want)
	}
}

func TestJoin(t *testing.T) {
	args := []string{"eins", "zwei", "drei", "vier", "fuenf", "eins", "zwei", "drei", "vier", "fuenf"}
	want := bytes.NewBufferString(strings.Join(args, " "))
	var have bytes.Buffer
	Joinargs(args, &have)
	if bytes.Compare(have.Bytes(), want.Bytes()) != 0 {
		t.Errorf("Buffers are not equal have/want \n[%v]\n[%v]\n", have, want)
	}
}
