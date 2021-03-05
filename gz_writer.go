package gzipcloser

import (
	"compress/gzip"
	"io"
)

var _ io.Writer = (*Writer)(nil)

// Writer creates a gzip file which closes the underlying stream as well as the gzip stream on close
type Writer struct {
	*gzip.Writer
	base io.Writer
}

// NewWriter acts like a compress/gzip.NewWriter but any Close And Flushes calls will be cascaded to underlying writer.
// If w it implements Flush() or Close() they will be called after the gzip call to Flush()/Close().
// if w does not implement Flush() or Close(); Flush()/Close() will only be applied to the gzip stream
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer: gzip.NewWriter(w),
		base:   w,
	}
}

// Write writes data to the gzip stream
func (gz *Writer) Write(p []byte) (int, error) {
	return gz.Writer.Write(p)
}

// Flush flushes the gzip stream AND the underlying writer
func (gz *Writer) Flush() error {
	if err := gz.Writer.Flush(); err != nil {
		return err
	}
	if flusher, ok := gz.base.(flusher); ok {
		return flusher.Flush()
	}
	return nil
}

// Close flushes and closes the gzip writer then does the same to the underlying writer.
func (gz *Writer) Close() error {
	if err := gz.Flush(); err != nil {
		return err
	}

	if err := gz.Writer.Close(); err != nil {
		return err
	}

	if closer, ok := gz.base.(closer); ok {
		return closer.Close()
	}

	return nil
}
