package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func PermuteStrings(input []string) [][]string {
	output := [][]string{}

	var generator func(int, []string)
	generator = func(k int, A []string) {
		if k == 1 {
			tmp := make([]string, len(A))
			copy(tmp, A)
			output = append(output, tmp)
		} else {
			for i := 0; i < k; i++ {
				generator(k-1, A)
				if k%2 == 1 {
					A[i], A[k-1] = A[k-1], A[i]
				} else {
					A[0], A[k-1] = A[k-1], A[0]
				}
			}
		}
	}
	generator(len(input), input)
	return output
}

func MustInt(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Errorf("error converting %s to int: %w", input, err))
	}
	return val
}

func GetLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func ReadInputLines(filename string) []string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return []string{}
	}
	return GetLines(string(fileBytes))
}
