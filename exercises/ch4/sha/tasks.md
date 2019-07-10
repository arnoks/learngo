# Exercise 4.2

## Description

Write a program that prints the sha256 hash of its standard input
by default but support a command-line flag to print the SHA384 and
SHA512 hash instead.

## Implementation

In order to make the applicaiton better testable the hash funtion has
been turned into a function returning the hashes. 

Since th hash functions return specific types [xx]byte those array had 
been turned into slices in to have a common signature for the hash function:

```go
func hash(...)[]byte {
...
  return h[:]
}
```

The table driven tests have been generated using gotests. To get the tables
right was quiete some challenge.

## Test

```
$ go build .
$ ./sha.exe < main.go
hash(sha256) = [158 14 173 220 247 47 19 3 156 226 245 244 84 120 43 139 250 254 171 184 164 76 189 179 32 105 157 114 226 86 222 235]
```
