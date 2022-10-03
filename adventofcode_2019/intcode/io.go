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
