package day_14

import (
	"fmt"
	"testing"
)

func TestTripletChars(t *testing.T) {
	tests := []struct {
		in     string
		tchars []byte
	}{
		{
			in:     "baaac",
			tchars: []byte{'a'},
		},
		{
			in:     "baacfsdjkfjlsdf",
			tchars: []byte{},
		},
		{
			in:     "xxxfhjkdsfskzzz",
			tchars: []byte{'x', 'z'},
		},
		{
			in:     "xxxxfhjkdsssfskzzzzzz",
			tchars: []byte{'x', 's', 'z'},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			tchars := tripletChars(test.in)
			if string(tchars) != string(test.tchars) {
				t.Fatalf("want %q, got %q", string(test.tchars), string(tchars))
			}
		})
	}
}

func TestKeyIndexes(t *testing.T) {
	salt := "abc"
	kis := keyIndexes(salt, 64, hashSalted)

	t.Logf("kis: %v", kis)
	if kis[0] != 39 {
		t.Fatalf("first: want %d, have %d", 39, kis[0])
	}
	if kis[63] != 22728 {
		t.Fatalf("last: want %d, have %d", 22728, kis[63])
	}
}

func TestKeyIndexesHash2016(t *testing.T) {
	salt := "abc"
	kis := keyIndexes(salt, 64, hash2016)

	t.Logf("kis: %v", kis)
	if kis[0] != 10 {
		t.Fatalf("first: want %d, have %d", 39, kis[0])
	}
	if kis[63] != 22551 {
		t.Fatalf("last: want %d, have %d", 22728, kis[63])
	}
}

func TestHash2016(t *testing.T) {
	salt := "abc"
	tests := []struct {
		idx  int
		hash string
	}{
		{
			idx:  0,
			hash: "a107ff634856bb300138cac6568c0f24",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			h := hash2016(salt, test.idx)
			if h != test.hash {
				t.Fatalf("want %q, got %q", test.hash, h)
			}
		})
	}
}
