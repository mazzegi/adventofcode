package testutil

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, want interface{}, have interface{}) {
	if reflect.DeepEqual(want, have) {
		return
	}
	t.Fatalf("want %v, have %v", want, have)
}
