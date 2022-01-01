package day_04

import "testing"

const inputTest = `
aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa
`

const inputTestExt = `
abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio 
`

func TestValidPassphraseCount(t *testing.T) {
	res := ValidPassphraseCount(inputTest)
	exp := 2
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestValidPassphraseCountExt(t *testing.T) {
	res := ValidPassphraseCountExt(inputTestExt)
	exp := 3
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
