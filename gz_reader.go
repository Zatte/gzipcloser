package gzipcloser

import (
	"compress/gzip"
	"io"
)

var _ io.Reader = (*Reader)(nil)

// Reader is a Gzip reader that once closed will also close the underlying reader (which the standard gz reader does not)
type Reader struct {
	*gzip.Reader
	base io.Reader
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
	if r, ok := gz.base.(closer); ok {
		return r.Close()
	}
	return nil
}
