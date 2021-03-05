package gzipcloser

import (
	"compress/gzip"
	"io"
)

// BaseReader is an alias for io.Reader
type BaseReader = io.Reader

// Reader is a Gzip reader that once closed will also close the underlying reader (which the standard gz reader does not)
type Reader struct {
	*gzip.Reader
	BaseReader
}

// NewReader treats the r-stream as a gzip stream and returns a Reader
func NewReader(r io.Reader) (*Reader, error) {
	gzReader, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &Reader{gzReader, r}, nil
}

// Close closes the gzip reader as well as the underlying reader
func (gz *Reader) Close() error {
	if err := gz.Reader.Close(); err != nil {
		return err
	}
	if r, ok := gz.BaseReader.(closer); ok {
		return r.Close()
	}
	return nil
}
