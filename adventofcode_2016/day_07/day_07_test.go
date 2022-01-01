package day_07

import (
	"fmt"
	"testing"
)

func TestIPSuportsTLS(t *testing.T) {
	tests := []struct {
		in  string
		tls bool
	}{
		{
			in:  "abba[mnop]qrst",
			tls: true,
		},
		{
			in:  "abcd[bddb]xyyx",
			tls: false,
		},
		{
			in:  "aaaa[qwer]tyui",
			tls: false,
		},
		{
			in:  "ioxxoj[asdfgh]zxcvbn",
			tls: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			tls := IPSupportsTLS(test.in)
			if tls != test.tls {
				t.Fatalf("supports-tls: expect %t, got %t", test.tls, tls)
			}
		})
	}
}

func TestIsABBA(t *testing.T) {
	tests := []struct {
		in   string
		abba bool
	}{
		{
			in:   "abba",
			abba: true,
		},
		{
			in:   "aaaa",
			abba: false,
		},
		{
			in:   "abbc",
			abba: false,
		},
		{
			in:   "iotr",
			abba: false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			abba := isABBA(test.in)
			if abba != test.abba {
				t.Fatalf("is-abba: expect %t, got %t", test.abba, abba)
			}
		})
	}
}

func TestContainsABBA(t *testing.T) {
	tests := []struct {
		in   string
		abba bool
	}{
		{
			in:   "abba",
			abba: true,
		},
		{
			in:   "aaaa",
			abba: false,
		},
		{
			in:   "abbccb",
			abba: true,
		},
		{
			in:   "iotrwqer",
			abba: false,
		},
		{
			in:   "ioxxoj",
			abba: true,
		},
		{
			in:   "ioxzojjo",
			abba: true,
		},
		{
			in:   "iooitureito",
			abba: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			abba := containsABBA(test.in)
			if abba != test.abba {
				t.Fatalf("contains-abba: expect %t, got %t", test.abba, abba)
			}
		})
	}
}

func TestIPSuportsSSL(t *testing.T) {
	tests := []struct {
		in  string
		tls bool
	}{
		{
			in:  "aba[bab]xyz",
			tls: true,
		},
		{
			in:  "xyx[xyx]xyx",
			tls: false,
		},
		{
			in:  "aaa[kek]eke",
			tls: true,
		},
		{
			in:  "zazbz[bzb]cdb",
			tls: true,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			tls := IPSupportsSSL(test.in)
			if tls != test.tls {
				t.Fatalf("supports-ssl: expect %t, got %t", test.tls, tls)
			}
		})
	}
}
