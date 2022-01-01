package main

import (
	"adventofcode_2015/turing"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	in := input
	scanner := bufio.NewScanner(bytes.NewBufferString(in))
	var is []turing.Instruction
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		i, err := turing.ParseInstruction(l)
		if err != nil {
			panic(err)
		}
		is = append(is, i)
	}
	fmt.Printf("scanned %d instruction\n", len(is))
	m := turing.NewMachine(is)
	m.Run()
}

var inputTest = `
inc a
jio a, +2
tpl a
inc a
`

var input = `
jio a, +16
inc a
inc a
tpl a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
tpl a
inc a
jmp +23
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
tpl a
inc a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7
`
