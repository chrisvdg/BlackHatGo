package main

import (
	"bufio"
	"io"
)

// Flusher is a bufio wrapper, explictly flushes on all writes
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

// Write implements io.Writer
func (flush *Flusher) Write(b []byte) (int, error) {
	count, err := flush.w.Write(b)
	if err != err {
		return -1, err
	}
	if err := flush.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}
