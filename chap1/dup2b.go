package main

import (
  "fmt"
  "bufio"
  "os"
)

// Remember the files  along with the string

func main() {
  files := os.Args[1:]
  counts := make(map[string]int)
  file_names := make(map[string]string[])

  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {  // loop over the files
    for _, file := range files {
      fp, err := os.Open(file)
      if err == nil {
        countLines(fp,counts, file)
      } else {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
    }
  }
  for line, n := range counts {
    fmt.Printf("%d:\t%s\n",n,line )
  }
}

func countLines(f *os.File, counts map[string]int, filename ) {
  input := bufio.NewScanner(f)
  for input.Scan(){
    counts[input.Text()]++
  }
}
