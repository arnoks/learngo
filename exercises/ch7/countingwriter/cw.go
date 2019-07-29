package countingwriter

import "io"

// Write a function CountingWriter with the signature below that, given
// an io.Writer, returns a new Writer that wraps the original, and a
// pointer to an int64 variable that at any moment contains the
// number of bytes written to the new Writer.

type countingWriter struct {
	c *int64
	w io.Writer
}

func (cw countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p) // call the original writer
	*cw.c += int64(n)
	return n, err
}

// CountingWriter returns a new Writer that wraps the original, and a
// pointer to an int64 variable that at any moment contains the
// number of bytes written to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw countingWriter
	var c int64
	cw.w = w
	cw.c = &c
	return cw, cw.c
}
