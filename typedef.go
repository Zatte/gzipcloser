package gzipcloser

type flusher interface {
	Flush() error
}

type closer interface {
	Close() error
}
