package day_08

import (
	"testing"
)

// ###....
// ###....
// .......

func TestScreen(t *testing.T) {
	screen := NewScreen(7, 3)

	iss := []string{
		"rect 3x2",
		"rotate column x=1 by 1",
		"rotate row y=0 by 4",
		"rotate column x=1 by 1",
	}

	inss, err := ParseInstructions(iss)
	if err != nil {
		t.Fatalf("do not expect parse to fail: %v", err)
	}
	screen.Apply(inss)

	onc := screen.PixelsOn()
	if onc != 6 {
		t.Fatalf("on: want %d, have %d", 6, onc)
	}
}
