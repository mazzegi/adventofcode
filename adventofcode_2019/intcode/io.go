package intcode

import (
	"github.com/mazzegi/adventofcode/slices"
	"github.com/pkg/errors"
)

// Reader

func NewIntSliceReader(values []int) *IntSliceReader {
	return &IntSliceReader{
		values: slices.Clone(values),
	}
}

type IntSliceReader struct {
	values []int
}

func (r *IntSliceReader) Read() (int, error) {
	if len(r.values) == 0 {
		return 0, errors.Errorf("nothing more to read")
	}
	n := r.values[0]
	r.values = r.values[1:]
	return n, nil
}

// Writer

func NewIntSliceWriter() *IntSliceWriter {
	return &IntSliceWriter{}
}

type IntSliceWriter struct {
	values []int
}

func (w *IntSliceWriter) Write(v int) error {
	w.values = append(w.values, v)
	return nil
}

func (w *IntSliceWriter) Values() []int {
	return w.values
}

//

func NewIntChannelReader(buffer int) *IntChannelReader {
	return &IntChannelReader{
		C: make(chan int, buffer),
	}
}

type IntChannelReader struct {
	C chan int
}

func (r *IntChannelReader) Close() {
	close(r.C)
}

func (r *IntChannelReader) Provide(n int) {
	r.C <- n
}

func (r *IntChannelReader) Read() (int, error) {
	v, ok := <-r.C
	if !ok {
		return 0, errors.Errorf("reader is closed")
	}
	return v, nil
}

//

func NewIntChannelWriter(buffer int) *IntChannelWriter {
	return &IntChannelWriter{
		C: make(chan int, buffer),
	}
}

type IntChannelWriter struct {
	C chan int
}

func (r *IntChannelWriter) Close() {
	close(r.C)
}

func (w *IntChannelWriter) Get() (int, bool) {
	v, ok := <-w.C
	return v, ok
}

func (w *IntChannelWriter) Write(v int) error {
	w.C <- v
	return nil
}
