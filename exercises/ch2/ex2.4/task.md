# EXercise 2.4

## Description

Write a version of PopCount that count bits by shifting its
arguments through 64 bit positions, testing the rightmost bit
each time. Compare its performance to the tabe-lookup version.

## Implementation

See ex2.3

## Test

```
goos: windows
goarch: amd64
pkg: github.com/arnoks/learngo/exercises/ch2/ex2.3
BenchmarkPopCount-4   	50000000	        24.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoop-4       	30000000	        44.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkShift-4      	20000000	        71.5 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	github.com/arnoks/learngo/exercises/ch2/ex2.3	5.267s
Success: Benchmarks passed.
```
