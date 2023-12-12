package day_12

import (
	"fmt"
	"testing"
)

func TestKnorcation(t *testing.T) {
	kns := IntKnorcations(3, 5)
	for _, kn := range kns {
		fmt.Println(kn)
	}
}

func TestRuneKnorcation(t *testing.T) {
	kns := RuneKnorcations(3, 5, '#', '.')
	for _, kn := range kns {
		fmt.Println(string(kn))
	}
}
