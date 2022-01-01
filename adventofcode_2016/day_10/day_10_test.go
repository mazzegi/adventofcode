package day_10

import "testing"

func TestFactoryProcess(t *testing.T) {
	const in = `
	value 5 goes to bot 2
	bot 2 gives low to bot 1 and high to bot 0
	value 3 goes to bot 1
	bot 1 gives low to output 1 and high to bot 0
	bot 0 gives low to output 2 and high to output 0
	value 2 goes to bot 2
	`

	fac, err := Parse(in)
	if err != nil {
		t.Fatalf("failed where no error was expected")
	}
	fac.Process()
}
