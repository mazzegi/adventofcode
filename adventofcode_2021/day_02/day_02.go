package day_02

import (
	"fmt"

	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func Part1() {
	pos, err := ProcessInput(input, Position2D{}, AppliedCommand)
	errutil.ExitOnErr(err)
	fmt.Printf("pos = %v; product = %d\n", pos, pos.Depth*pos.X)
}

func Part2() {
	pos, err := ProcessInput(input, Position2D{}, AppliedCommandAim)
	errutil.ExitOnErr(err)
	fmt.Printf("aimed: pos = %v; product = %d\n", pos, pos.Depth*pos.X)
}

type Command interface{}

type ForwardCommand struct {
	steps int
}

type UpCommand struct {
	steps int
}

type DownCommand struct {
	steps int
}

func ParseCommand(s string) (Command, error) {
	var cmd string
	var steps int
	_, err := fmt.Sscanf(s, "%s %d", &cmd, &steps)
	if err != nil {
		return nil, errors.Wrapf(err, "scan command %q", s)
	}
	switch cmd {
	case "forward":
		return ForwardCommand{steps: steps}, nil
	case "up":
		return UpCommand{steps: steps}, nil
	case "down":
		return DownCommand{steps: steps}, nil
	default:
		return nil, errors.Errorf("unknown command %s", cmd)
	}
}

type Position2D struct {
	X     int
	Depth int
	Aim   int
}

func AppliedCommand(cmd Command, pos Position2D) Position2D {
	switch cmd := cmd.(type) {
	case ForwardCommand:
		pos.X += cmd.steps
	case UpCommand:
		pos.Depth -= cmd.steps
	case DownCommand:
		pos.Depth += cmd.steps
	}
	return pos
}

func AppliedCommandAim(cmd Command, pos Position2D) Position2D {
	switch cmd := cmd.(type) {
	case ForwardCommand:
		pos.X += cmd.steps
		pos.Depth += pos.Aim * cmd.steps
	case UpCommand:
		pos.Aim -= cmd.steps
	case DownCommand:
		pos.Aim += cmd.steps
	}
	return pos
}

func ProcessInput(in string, pos Position2D, appFunc func(cmd Command, pos Position2D) Position2D) (Position2D, error) {
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		cmd, err := ParseCommand(line)
		if err != nil {
			return Position2D{}, errors.Wrapf(err, "parse-command %q", line)
		}
		//pos = AppliedCommand(cmd, pos)
		pos = appFunc(cmd, pos)
	}
	return pos, nil
}
