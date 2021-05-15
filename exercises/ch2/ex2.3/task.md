# Ex2.2

## Description

Rewrite PopCount to use a loop instead a single expression. Compare
the performance of the two versions.

## Implementation

a.  Using a simple loop to go voer all byte of the input variable x
b. Add package tests benchmark to compare the performance of both versions.

## Results

The implementation with a loop is about half as efficient as the original version.

```cmd
Running tool: D:\Go\bin\go.exe test -benchmem -run=^$ github.com\arnoks\learngo\exercises\ch2\ex2.3 -bench . -coverprofile=C:\Users\vn401\AppData\Local\Temp\vscode-goU810AF\go-code-cover

goos: windows
goarch: amd64
pkg: github.com/arnoks/learngo/exercises/ch2/ex2.3
go version .14 vs .11
.14 BenchmarkPopCount-4    	  1000000000	         0.357 ns/op	       0 B/op	       0 allocs/op
.11 BenchmarkPopCount-4       	50000000	        21.5   ns/op	       0 B/op	       0 allocs/op
.14 BenchmarkLoop-4         	52170963	        23.5   ns/op	       0 B/op	       0 allocs/op
.11 BenchmarkPopCountLoop-4   	50000000	        39.2   ns/op	       0 B/op	       0 allocs/op
.14 BenchmarkRecursion-4     	11427918	        94.1   ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	github.com/arnoks/learngo/exercises/ch2/ex2.3	3.308s
Success: Benchmarks passed.

```
