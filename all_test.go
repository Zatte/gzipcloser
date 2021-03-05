package gzipcloser_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zatte/gzipcloser"
)

type testBuffer struct {
	*bytes.Buffer
	closeCalls int
	flushCalls int

	closeErr error
	flushErr error
}

func (b *testBuffer) Close() error {
	b.closeCalls++
	return b.closeErr
}

func (b *testBuffer) Flush() error {
	b.flushCalls++
	return b.flushErr
}

func TestWriterClose(t *testing.T) {

	// A buffer doesn't have Close/Flush methods and as such they are NOP.
	b := &testBuffer{
		Buffer: bytes.NewBuffer(nil),
	}

	// The closing a writer
	assert.Equal(t, b.closeErr, gzipcloser.NewWriter(b).Close(), "When closing a stream the underlying error should be reported <nil>")
	assert.Equal(t, 1, b.closeCalls, "When closing a gzip-stream-writer the underlying Close() method should have been called")
	assert.Equal(t, 1, b.flushCalls, "When closing a gzip-stream-writer the underlying Flush() method should have been called")

	b.closeErr = fmt.Errorf("some-err")
	assert.Equal(t, b.closeErr, gzipcloser.NewWriter(b).Close(), "When closing a stream the underlying error should be reported <some-err>")
}

func TestWriterFlush(t *testing.T) {

	// A buffer doesn't have Close/Flush methods and as such they are NOP.
	b := &testBuffer{
		Buffer: bytes.NewBuffer(nil),
	}

	// The closing a writer
	assert.Equal(t, b.closeErr, gzipcloser.NewWriter(b).Flush(), "When closing a stream the underlying error should be reported <nil>")
	assert.Equal(t, 1, b.flushCalls, "When flushing a gzip-stream-writer the underlying Flush() method should have been called")

	b.closeErr = fmt.Errorf("some-err")
	assert.Equal(t, b.closeErr, gzipcloser.NewWriter(b).Close(), "When closing a stream the underlying error should be reported <some-err>")
}

func TestReaderClose(t *testing.T) {

	// We need to create some content or the gzip NewReader
	// methos will not recognize the stream as gzip
	getTestReader := func() *testBuffer {
		b := &testBuffer{
			Buffer: bytes.NewBuffer(nil),
		}

		w := gzipcloser.NewWriter(b)
		_, err := w.Write([]byte("test_string"))
		require.NoError(t, err)
		err = w.Close()
		require.NoError(t, err)
		return b
	}

	// The closing a reader
	// No Error case
	b := getTestReader()
	r, err := gzipcloser.NewReader(b)
	require.NoError(t, err)
	assert.Equal(t, b.closeErr, r.Close(), "When closing a stream the underlying error should be reported <nil>")

	// Simulated error case
	b = getTestReader()
	b.closeErr = fmt.Errorf("some-err")
	r, err = gzipcloser.NewReader(b)
	require.NoError(t, err)
	assert.Equal(t, b.closeErr, r.Close(), "When closing a stream the underlying error should be reported <nil>")
}
