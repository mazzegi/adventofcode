package day_25

type move string

const (
	moveLeft  move = "left"
	moveRight move = "right"
)

type instruction struct {
	write         uint8
	move          move
	continueState string
}

type state struct {
	id  string
	if0 instruction
	if1 instruction
}

type blueprint struct {
	startStateID string
	steps        int
	states       map[string]state
}

var inputBlueprint = blueprint{
	startStateID: "A",
	steps:        12317297,
	states: map[string]state{
		"A": {
			id: "A",
			if0: instruction{
				write:         1,
				move:          moveRight,
				continueState: "B",
			},
			if1: instruction{
				write:         0,
				move:          moveLeft,
				continueState: "D",
			},
		},
		"B": {
			id: "B",
			if0: instruction{
				write:         1,
				move:          moveRight,
				continueState: "C",
			},
			if1: instruction{
				write:         0,
				move:          moveRight,
				continueState: "F",
			},
		},
		"C": {
			id: "C",
			if0: instruction{
				write:         1,
				move:          moveLeft,
				continueState: "C",
			},
			if1: instruction{
				write:         1,
				move:          moveLeft,
				continueState: "A",
			},
		},
		"D": {
			id: "D",
			if0: instruction{
				write:         0,
				move:          moveLeft,
				continueState: "E",
			},
			if1: instruction{
				write:         1,
				move:          moveRight,
				continueState: "A",
			},
		},
		"E": {
			id: "E",
			if0: instruction{
				write:         1,
				move:          moveLeft,
				continueState: "A",
			},
			if1: instruction{
				write:         0,
				move:          moveRight,
				continueState: "B",
			},
		},
		"F": {
			id: "F",
			if0: instruction{
				write:         0,
				move:          moveRight,
				continueState: "C",
			},
			if1: instruction{
				write:         0,
				move:          moveRight,
				continueState: "E",
			},
		},
	},
}
