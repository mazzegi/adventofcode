package stringutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazzegi/adventofcode/errutil"
)

func MustStringsToInts(sl []string) []int {
	var ns []int
	for _, s := range sl {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		errutil.FatalWhen(err)
		ns = append(ns, n)
	}
	return ns
}

func StringsToInts(sl []string) ([]int, error) {
	var ns []int
	for _, s := range sl {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("atoi %q: %w", s, err)
		}
		ns = append(ns, n)
	}
	return ns, nil
}
