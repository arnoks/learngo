# Chapter 7 exercises

## Open

7.1 Using the ideas from ByteCounter, implement counters for words and
    for lines. You will find bufio.Scanwords useful.

7.2 Write a function CountingWriter with the signature below that, given
    an io.Writer, returns a new Writer that wraps the original, and a 
    pointer to an int64 variable that at any moment contains the 
    number of byted written to the new Writer. 
    func CountingWriter(w io.Writer)(io.Writer, int64)

7.3 Write a String MEthod for the *tree type in ch4/treesort that reveals
    the sequence of values in the tree.

7.4 The strins.NewReader function returns a avalue that satisfies the io.Reader
    interface (and others) by reading from its argument, a string. Implement a 
    simple version of NewReader yourself, and use it to make the HTML parser 
    take input from a string

7.5 The LimitReader function in the io package accepts an io.Reader r and a
    number of bytes n, and returns another Reader that reads from r but reports
    and end-of-file condition after n bytes. Implement it. 

7.6 Add support for Kelvin temperatures to tempflag

7.7 Explain why the help messages contains °C when the default value of 20.0 
    does not.

7.8 Many Guis provide a table widget with a stateful multi-tier sort: the
    primary sort key is most recently clicked column head, the secondary sort
    key is the second-most recently clicked column head, and so on. Define an
    implementation of sort.Interface for use by such a table. Compare that 
    approach with repeated sorting using sort.Stable

## Closed