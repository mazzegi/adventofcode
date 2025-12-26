package day_10

import (
	"fmt"

	"github.com/mazzegi/adventofcode/equations"
	"github.com/mazzegi/adventofcode/rat"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/slices"
)

func part2MainFunc(in string) (int, error) {
	lines := readutil.ReadLines(in)
	var ms []machine
	for _, line := range lines {
		m, err := parseMachine(line)
		if err != nil {
			return 0, fmt.Errorf("parse machine %q: %w", line, err)
		}
		ms = append(ms, m)
	}
	var sum int
	for i, m := range ms {

		np, err := findFewestButtonPressesForMachineJoltage(m)
		if err != nil {
			return 0, fmt.Errorf("find_for_machine: %w", err)
		}
		log("part2: machine %d -> %d presses", i+1, np)
		sum += np
	}

	return sum, nil
}

func findFewestButtonPressesForMachineJoltage(m machine) (int, error) {
	dim := len(m.buttons)
	sys, err := equations.NewSystem(dim)
	if err != nil {
		return 0, fmt.Errorf("new_system: %w", err)
	}
	maxJoltage := slices.Max(m.targetJoltages)
	// add eqs
	bounds := slices.Repeat(maxJoltage, dim)

	for i, targetJoltage := range m.targetJoltages {
		cs := make([]int, dim)
		for btnIdx, btn := range m.buttons {
			// does button contribute to joltage?
			if slices.Contains(btn, i) {
				cs[btnIdx] = 1
				if targetJoltage < bounds[btnIdx] {
					bounds[btnIdx] = targetJoltage
				}
			} else {
				cs[btnIdx] = 0
			}
		}
		//
		sys.AddEquation(mkEq(-targetJoltage, cs...))
	}
	//
	sols, err := sys.AllPositiveIntegerSolutions(bounds)

	var minSum int
	for i, sol := range sols {
		fmt.Printf("solution %d: %v\n", i+1, sol)
		sum := slices.Sum(sol)
		if i == 0 || sum < minSum {
			minSum = sum
		}
	}
	return minSum, nil
}

func mkRat(v int) rat.Number {
	return rat.R(v, 1)
}

func mkCoeffs(vs ...int) []rat.Number {
	rns := make([]rat.Number, len(vs))
	for i, v := range vs {
		rns[i] = mkRat(v)
	}
	return rns
}

func mkExpr(constant int, vs ...int) equations.Expression {
	return equations.Expression{
		VariableCoeffs: mkCoeffs(vs...),
		Constant:       mkRat(constant),
	}
}

func mkEq(constant int, vs ...int) equations.Equation {
	return equations.Equation{mkExpr(constant, vs...)}
}
