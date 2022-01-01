package day_25

import (
	"adventofcode_2016/assembunny"
	"adventofcode_2016/errutil"
	"fmt"
)

var registers = []string{"a", "b", "c", "d"}

func Part1() {
	iss, err := assembunny.ParseInstructions(input)
	errutil.ExitOnErr(err)

	maxInit := 1000
	for init := 0; init <= maxInit; init++ {
		isClock := Try(iss, init, 100)
		if isClock {
			fmt.Printf("init %d produces clock signal\n", init)
			return
		}
	}
}

func Try(iss []assembunny.Instruction, initVal int, loops int) bool {
	comp := assembunny.NewComputer(registers)
	comp.SetReg("a", initVal)

	cnt := 0
	//var sig []int
	// 0 => 0
	// 1 => 1
	isClock := true
	outFunc := func(out int) assembunny.CbResult {
		var exp int
		if cnt%2 == 0 {
			exp = 0
		} else {
			exp = 1
		}
		if out != exp {
			isClock = false
			return assembunny.CbQuit
		}
		cnt++
		if cnt > loops {
			return assembunny.CbQuit
		}
		return assembunny.CbProceed
	}

	comp.Execute(iss, outFunc)
	return isClock
}
