# Chapter 7 exercises

## Open
7.1 ch7/bytecounter 

7.2 [Counting Writer](examples/ch7/countingwriter/cw.go)

7.3 [tree.String()](ch4/tresort/main.go)

7.4 [StringReader](exercises/ch7/newreader/nr.go)

7.5 [LimtReader](exercises/ch7/newreader/nr.go)

7.6 [tempflag]ch7/tempflag/main.go

7.7 see Closed

7.8 Many Guis provide a table widget with a stateful multi-tier sort: the
    primary sort key is most recently clicked column head, the secondary sort
    key is the second-most recently clicked column head, and so on. Define an
    implementation of sort.Interface for use by such a table. Compare that 
    approach with repeated sorting using sort.Stable

7.9 Use the html template package to replace the printTracks with a function
    that displays tracks as HTML table. USe the solution to the previous exercise
    to arrange that each click on a custom head makes an HTTP request to sort 
    the table

7.10 The sort.Interaface can be adapted to other uses. Write a function 
     IsPalindrom (s. Interface ) bool that reports whether the sequence s is
     a palindrome, in other words, reversing the sequence would not change
     it. Assume that the elements at indices i and j are equal if !sLess (i,j)
     && !s.Less(j,i).
     

## Closed

7.1 ch7/bytecounter
    Using the ideas from ByteCounter, implement counters for words and
    for lines. You will find bufio.Scanwords useful.

7.2 Write a function CountingWriter with the signature below that, given
    an io.Writer, returns a new Writer that wraps the original, and a 
    pointer to an int64 variable that at any moment contains the 
    number of bytes written to the new Writer. 
    func CountingWriter(w io.Writer)(io.Writer, *int64)

7.3 ch4/tresort/main.go 
    Write a String Method for the *tree type in ch4/treesort that reveals
    the sequence of values in the tree.

7.4 The strings.NewReader function returns a avalue that satisfies the io.Reader
    interface (and others) by reading from its argument, a string. Implement a 
    simple version of NewReader yourself, and use it to make the HTML parser 
    (5.2) take input from a string

7.5 The LimitReader function in the io package accepts an io.Reader r and a
    number of bytes n, and returns another Reader that reads from r but reports
    and end-of-file condition after n bytes. Implement it. 

7.6 Added support for Kelvin temperatures to tempflag in ch 7

7.7 Explain why the help messages contains °C when the default value of 20.0 
    does not.
    The default is just a float while the help message prints the string!

