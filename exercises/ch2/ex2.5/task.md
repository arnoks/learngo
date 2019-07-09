# Exercise 2.5

## Description

The expression x&(x-1) clears the rightmost non-zero bit of x.
Write a version of PopCount that counts bits by using this
fact, and assess its performance.

## Implementation 

```go
func Eliminate(x uint64) int {
	c := 0
	for ; x != 0; c++ {
		x = x & (x - 1)
	}
	return c
}
```

## Test

This option provides 
```
Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ github.com\arnoks\learngo\exercises\ch2\ex2.3 -bench . -coverprofile=C:\Users\vn401\AppData\Local\Temp\vscode-goU810AF\go-code-cover

goos: windows
goarch: amd64
pkg: github.com/arnoks/learngo/exercises/ch2/ex2.3
BenchmarkPopCount-4    	100000000	        22.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoop-4        	30000000	        37.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkShift-4       	20000000	        69.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkEliminate-4   	30000000	        80.4 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	github.com/arnoks/learngo/exercises/ch2/ex2.3	7.495s
Success: Benchmarks passed.
```
