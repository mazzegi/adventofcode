package readutil

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ReadLines(s string) []string {
	var lines []string
	sl := strings.Split(s, "\n")
	for _, s := range sl {
		s = strings.Trim(s, " \r\n\t")
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}
	return lines
}

func ReadLinesNoTrim(s string) []string {
	var lines []string
	sl := strings.Split(s, "\n")
	for _, s := range sl {
		if s == "" {
			continue
		}
		lines = append(lines, s)
	}
	return lines
}

func ReadLinesKeepEmpty(s string) []string {
	var lines []string
	sl := strings.Split(s, "\n")
	for _, s := range sl {
		s = strings.Trim(s, " \r\n\t")
		lines = append(lines, s)
	}
	return lines
}

func ReadString(s string) string {
	return strings.Join(ReadLines(s), "")
}

func ReadStrings(s string, sep string) []string {
	var sl []string
	sls := strings.Split(s, sep)
	for _, v := range sls {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		sl = append(sl, v)
	}

	return sl
}

func ReadInts(s string, sep string) ([]int, error) {
	s = strings.Trim(s, " \r\n\t")
	var ns []int
	sns := strings.Split(s, sep)
	for _, sn := range sns {
		sn = strings.Trim(sn, " \r\n\t")
		if sn == "" {
			continue
		}

		n, err := strconv.ParseInt(sn, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-int %q", sn)
		}
		ns = append(ns, int(n))
	}
	return ns, nil
}
