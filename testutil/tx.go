package testutil

import (
	"fmt"
	"reflect"
	"testing"
)

func NewTx(t *testing.T) *Tx {
	return &Tx{
		T:              t,
		Name:           t.Name(),
		CurrentContext: "default",
	}
}

func NewTxWithName(t *testing.T, name string) *Tx {
	return &Tx{
		T:              t,
		Name:           name,
		CurrentContext: "default",
	}
}

type Tx struct {
	T              *testing.T
	Name           string
	CurrentContext string
}

func (tx *Tx) WithContext(s string) {
	tx.CurrentContext = s
}

func (tx *Tx) Assert(v bool) {
	if v {
		return
	}
	tx.T.Helper()
	tx.T.Fatalf("%s: assert: false", tx.Name)
}

func (tx *Tx) AssertEqual(want, have any) {
	tx.T.Helper()
	if reflect.DeepEqual(want, have) {
		return
	}
	tx.T.Fatalf("%s (%s): want %v, have %v", tx.Name, tx.CurrentContext, want, have)
}

func (tx *Tx) AssertNoErr(err error) {
	tx.T.Helper()
	if err == nil {
		return
	}
	tx.T.Fatalf("%s (%s): error is not-nil but: %v", tx.Name, tx.CurrentContext, err)
}

func (tx *Tx) AssertErr(err error) {
	tx.T.Helper()
	if err != nil {
		return
	}
	tx.T.Fatalf("%s (%s): expect err; got none", tx.Name, tx.CurrentContext)
}

func (tx *Tx) Log(pattern string, args ...any) {
	tx.T.Helper()
	fmt.Printf(tx.Name+": "+pattern+"\n", args...)
}

func (tx *Tx) Fatalf(pattern string, args ...any) {
	tx.T.Helper()
	tx.T.Fatalf(tx.Name+": "+pattern, args...)
}
