package day_02

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	res, err := Checksum(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: res = %d\n", res)
}

func Part2() {
	res, err := ChecksumEvenlyDiv(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: res = %d\n", res)
}

type SpreadsheetRow struct {
	values []int
}

type Spreadsheet struct {
	rows []*SpreadsheetRow
}

func ParseSpreadsheetRow(in string) (*SpreadsheetRow, error) {
	row := &SpreadsheetRow{}
	sl := strings.Fields(in)
	for _, s := range sl {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-int %q", s)
		}
		row.values = append(row.values, int(n))
	}

	return row, nil
}

func ParseSpreadsheet(in string) (*Spreadsheet, error) {
	lines := readutil.ReadLines(in)
	sp := &Spreadsheet{}
	for _, line := range lines {
		row, err := ParseSpreadsheetRow(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-row %q", line)
		}
		sp.rows = append(sp.rows, row)
	}
	return sp, nil
}

func Checksum(in string) (int, error) {
	sp, err := ParseSpreadsheet(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse spreadsheet")
	}

	var sum int
	for _, row := range sp.rows {
		if len(row.values) == 0 {
			continue
		}
		sort.Ints(row.values)
		d := row.values[len(row.values)-1] - row.values[0]
		sum += d
	}

	return sum, nil
}

func ChecksumEvenlyDiv(in string) (int, error) {
	sp, err := ParseSpreadsheet(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse spreadsheet")
	}

	var sum int
outer:
	for _, row := range sp.rows {
		for i1, n1 := range row.values {
			for i2, n2 := range row.values {
				if i2 == i1 {
					continue
				}
				if n2 == 0 {
					continue
				}
				if n1%n2 == 0 {
					sum += n1 / n2
					continue outer
				}
			}
		}
	}

	return sum, nil
}
