package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2018/errutil"

	"github.com/pkg/errors"
)

var replacer *strings.Replacer

func main() {
	if len(os.Args) < 2 {
		errutil.ExitOnErr(errors.Errorf("missing name argument"))
	}
	name := os.Args[1]
	fmt.Printf("generate packages for %q\n", name)
	replacer = strings.NewReplacer("{name}", name)

	//
	err := os.Mkdir(name, os.ModePerm)
	errutil.ExitOnErr(err)

	// create main
	mainPath := filepath.Join(name, "cmd", "main.go")
	err = os.MkdirAll(filepath.Dir(mainPath), os.ModePerm)
	errutil.ExitOnErr(err)

	err = os.WriteFile(mainPath, contentMain(), os.ModePerm)
	errutil.ExitOnErr(err)

	// package file
	pkgFilePath := filepath.Join(name, name+".go")
	err = os.WriteFile(pkgFilePath, pkgFileContent(), os.ModePerm)
	errutil.ExitOnErr(err)

	// test file
	testFilePath := filepath.Join(name, name+"_test.go")
	err = os.WriteFile(testFilePath, testFileContent(), os.ModePerm)
	errutil.ExitOnErr(err)

	// input file
	inputFilePath := filepath.Join(name, "input.go")
	err = os.WriteFile(inputFilePath, inputFileContent(), os.ModePerm)
	errutil.ExitOnErr(err)
}

var mainTpl = `
package main

import "github.com/mazzegi/adventofcode/adventofcode_2018/{name}"

func main() {
	{name}.Part1()
	{name}.Part2()
}
`

var pkgFileTpl = `
package {name}

import (
	"fmt"	
	"github.com/mazzegi/adventofcode/adventofcode_2018/errutil"	
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

func Part1() {
	res, err := part1MainFunc(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

func part1MainFunc(in string) (int, error){
	return 0, nil
}

func part2MainFunc(in string) (int, error){
	return 0, nil
}
`

var testFileTpl = `
package {name}

import (	
	"testing"
	"github.com/mazzegi/adventofcode/adventofcode_2018/testutil"
)

const inputTest = ""

func TestPart1MainFunc(t *testing.T) {
	res, err := part1MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}
`

var inputFileTpl = `
package {name}

const input = ""
`

func contentMain() []byte {
	return []byte(replacer.Replace(mainTpl))
}

func pkgFileContent() []byte {
	return []byte(replacer.Replace(pkgFileTpl))
}

func testFileContent() []byte {
	return []byte(replacer.Replace(testFileTpl))
}

func inputFileContent() []byte {
	return []byte(replacer.Replace(inputFileTpl))
}
