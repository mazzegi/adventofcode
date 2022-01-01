/*
--- Day 5: Binary Boarding ---
You board your plane only to discover a new problem: you dropped your boarding pass! You aren't sure which seat is yours, and all of the flight attendants are busy with the flood of people that suddenly made it through passport control.

You write a quick program to use your phone's camera to scan all of the nearby boarding passes (your puzzle input); perhaps you can find your seat through process of elimination.

Instead of zones or groups, this airline uses binary space partitioning to seat people. A seat might be specified like FBFBBFFRLR, where F means "front", B means "back", L means "left", and R means "right".

The first 7 characters will either be F or B; these specify exactly one of the 128 rows on the plane (numbered 0 through 127). Each letter tells you which half of a region the given seat is in. Start with the whole list of rows; the first letter indicates whether the seat is in the front (0 through 63) or the back (64 through 127). The next letter indicates which half of that region the seat is in, and so on until you're left with exactly one row.

For example, consider just the first seven characters of FBFBBFFRLR:

Start by considering the whole range, rows 0 through 127.
F means to take the lower half, keeping rows 0 through 63.
B means to take the upper half, keeping rows 32 through 63.
F means to take the lower half, keeping rows 32 through 47.
B means to take the upper half, keeping rows 40 through 47.
B keeps rows 44 through 47.
F keeps rows 44 through 45.
The final F keeps the lower of the two, row 44.
The last three characters will be either L or R; these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7). The same process as above proceeds again, this time with only three steps. L means to keep the lower half, while R means to keep the upper half.

For example, consider just the last 3 characters of FBFBBFFRLR:

Start by considering the whole range, columns 0 through 7.
R means to take the upper half, keeping columns 4 through 7.
L means to take the lower half, keeping columns 4 through 5.
The final R keeps the upper of the two, column 5.
So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.

Every seat also has a unique seat ID: multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.

Here are some other boarding passes:

BFFFBBFRRR: row 70, column 7, seat ID 567.
FFFBBBFRRR: row 14, column 7, seat ID 119.
BBFFBBFRLL: row 102, column 4, seat ID 820.
As a sanity check, look through your list of boarding passes. What is the highest seat ID on a boarding pass?
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	buf := bytes.NewBufferString(input)
	scanner := bufio.NewScanner(buf)
	var maxID int
	var seatIDs []int
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \n\r\t")
		if s == "" {
			continue
		}
		seat, err := DecodeSeat(s)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%q: %s\n", s, seat)
		if seat.ID() > maxID {
			maxID = seat.ID()
		}
		seatIDs = append(seatIDs, seat.ID())
	}
	fmt.Printf("max seat-id: %d\n", maxID)
	sort.Ints(seatIDs)
	fmt.Printf("seat ids: %v\n", seatIDs)

	var mySeatID int
	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] == seatIDs[i]+2 {
			mySeatID = seatIDs[i] + 1
			break
		}
	}
	fmt.Printf("my seat id is %d\n", mySeatID)
}

type Seat struct {
	Row int
	Col int
}

func (s Seat) ID() int {
	return s.Row*8 + s.Col
}

func (s Seat) String() string {
	return fmt.Sprintf("row %d, col %d, ID %d", s.Row, s.Col, s.ID())
}

func DecodeSeat(s string) (Seat, error) {
	if len(s) != 10 {
		return Seat{}, errors.Errorf("invalid code %q - must be 10 chars long", s)
	}
	rowCode := s[:7]
	colCode := s[7:]
	// totalRows := 128
	rmin := 0
	rmax := 127
	for _, r := range rowCode {
		switch r {
		case 'F':
			rmax = rmin + (rmax-rmin)/2
		case 'B':
			rmin = rmax - (rmax-rmin)/2
		default:
			return Seat{}, errors.Errorf("invalid code element %q", string(r))
		}
	}
	if rmin != rmax {
		return Seat{}, errors.Errorf("something went completely wrong (min,max = %d, %d)", rmin, rmax)
	}

	cmin := 0
	cmax := 7
	for _, r := range colCode {
		switch r {
		case 'L':
			cmax = cmin + (cmax-cmin)/2
		case 'R':
			cmin = cmax - (cmax-cmin)/2
		default:
			return Seat{}, errors.Errorf("invalid code element %q", string(r))
		}
	}
	if cmin != cmax {
		return Seat{}, errors.Errorf("something went completely wrong (min,max = %d, %d)", cmin, cmax)
	}

	return Seat{
		Row: rmin,
		Col: cmin,
	}, nil
}

var inputTest = `
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`

var input = `
FFBFFFBLLL
BFBFBBFRLR
FFFBFFBLLL
FFBFBBBRRL
BFFFBBFLRR
BFBFBFBLLL
BFFFBFFLLL
FFFFFBFLLR
FBFBFFBRLL
BBFFBFBLRL
FFFBBBBLRL
BFFFFFFRLL
FBBBBFFRLL
FBFBFBFLLL
BBFFBFFLRR
BFBBFFBLRL
FBFFFBBLLR
BFFFFFBLRL
BBFBFBBLLL
FFFBBBBLRR
FBFFBFFLRR
BFFBFBFLRL
BFBBFFFLRR
FBFFBBBRLR
FBFFBBBRRR
BBFBFBFLRR
BFFBBFFLLL
BFFBBFBRRL
FFFBFFBLRR
FBBBBBBLLL
FFBFFFFRRL
BBFBFBBRLR
BFFBFFFLRR
BBFFFFBLRL
BBFFBFBLLL
BBFFFFFRLR
BFBFBBFLRL
FFFBFFBRLR
FFFFBFFLLL
FFBBFBFRLL
BFFBBFBRLL
BFFFBFFRLL
BFFFBBFRLR
BFFFFFBRRR
BFBBBFFRRL
FFBFBFFRRL
FBBFBBBRRR
FFFBBBBRLR
FBFBFBFRLR
FFBFFFFLLL
FBBFFFBLRR
BFBFFBBLLR
FBFFBBBRLL
FBFFBBFRRL
FBBBFFBLRL
BBFBFBFLLL
BFFBFBBLLL
BBFFBFBRRR
FFBBFFFLRL
BFFBBBBLRL
FFFBBFFLRL
FFBBBBBLRR
FBFFFBBRLR
BBFBFFFRLL
FFBFFBFLRL
FFBFBFFLLR
FBBBBFBRRR
FBFBFBFRRR
FBBFBBBLRL
BBFFFBBLRL
BFFBFBBLLR
FBFBBFFLRR
FFBBFFFRLR
FFFFBBFLLR
FBFFBFBRRR
FFFBFFFRLR
FBFFBFBLRL
FBBFFFFRRR
BFFFBFBLLR
FFFBBFFRLL
FFBFBFBLLL
FFBFBBFRLL
BFFBBBBRRR
BFFBFFFLRL
FFFFFFBRRL
FFBBFBBRRR
FFBFBFFRLR
FBBFFBBLRR
BFFBFFBRLR
BBFBFFFLRL
BFBFBFFLRR
BFFBBFFRLL
FFBFBBFLLR
BBFBFFFLRR
FFFFBBBLRL
FFFBFBFRLR
BFBFFFBLRL
BFFFBBFLRL
FFFFBFFRRL
BFBBBFBRLL
BFBBBFFRRR
BFFFBFBLRL
BFFBFFBRRL
FBBBFBFLLL
FBBFBBBRLR
FFBBBFBRLR
FFFBBFBLLL
BFFFFFBRLL
FBFBFFBLRR
FBFFFBFLLR
FBBFFBFRRL
BFFFFFFRLR
FFBFBBBRLL
BFFFFFBLLR
FFFBFBBLRL
FBBBFFBLRR
FBBFFBFRLR
FBBFBBFRRR
FFBFFBFLLL
BFBBBFBLRL
FBFBBBBRRR
FBBBFBFRRL
FBBFBBBRRL
BBFBFFFRLR
BFBFBBBRLL
FFBBFFFLLR
FBFFFFFRRR
BFBBFFBRRR
FBBBBFBLRL
BFBBBFBLLR
FBFBFBFLRR
FBBBFBBLLR
BFFFBFFRRR
BBFBFBFRLR
BBFFBFBRRL
FFBBBFFRRL
FBFFBFBLRR
BBFFBBBRLL
FBFFBFFLLR
FFFBFBFLLL
FFBBBFBLLL
BFBFFBFRRR
FBFBBFBLLR
BBFFFFFLLR
BBFBBFFRLL
FFFFBFBLLR
BBFFBBBLRL
FBBFFFBRLR
BFFFFBFLRL
FFBBFFBRRR
BBFFFBFLLR
BBFFFBBRRR
FBBFFBFRLL
BFFBFFBLRR
FFFBBFBRLL
FBBFFBBRLL
BFFFBFBRRR
FBFBFFFLRL
BFBBFBBRRR
FBBBFFBLLL
FBBFFBFRRR
FBFBFFFRRR
FBFFFFFRLR
FBBBFBBLLL
FFFFBFFLRR
BBFFFFBLLR
FFBFBBFLRL
FBFFFFFLRR
BFFBBFBLLL
FFBBFFFLRR
BFFFFFFLRL
BFBFBBBLRL
BBFBFFBRLL
FFBFBFBRLL
FFFFBBBLLL
FFBBBFFLRR
BFFBBBBLLR
FFFBFFBLRL
FBBBFBFRLR
FFBBFBFLLL
BBFFFBBLRR
BFBBFFFRRL
FFBBFFFRRL
BFBBFFFLRL
FBBFFFFLLL
BBFFFFBRRL
FFBFBFFLRL
FFBBFFBLLR
BFFFFBFRLR
BBFFBBBRLR
FBFBBBFRRR
FBFBFFFRRL
BBFBBFFRLR
FBFBBBFLLR
BFBFBBBRLR
BFBBBBFLRR
BFBBBBBLRL
FBFFFFFRLL
BFBBFBBLLR
BFFBFFBLRL
FBBFBBFRLR
FBBBFFBRRL
BFFFFFBLRR
BBFFFBFRRR
BFBFFFFLRR
BFBFBFFRRL
BFBBBFBRRL
FFFFFBFLLL
FBBFFBFLLR
FFBFBFBLLR
FBFFBBFLLL
BFFBFFFLLL
FFBFBFBRRR
FBBFBFBLRL
FBFFFFFRRL
FFFBBFBRLR
BBFBFFBLLL
FFBFFFBLRL
FFBBBBFLRL
FFBFFBFRRL
FBFFFFBLRR
FBBBBBBRRR
FFFFFBBLLL
FBFBFBBLRR
FBFFBBBRRL
FFBFBBBLRR
BFBFFBBRRL
BFBFFFBRLL
FFBBBFFRRR
BFFFFBBRRR
FFFBBBBLLR
BFBFBFBRRR
FFFBBFFLRR
FFFBFBFRRR
FBBBBFFLLL
BBFBFBBLRR
BBFBFBFRRL
BBFFFBBRLL
FBBBBFBLRR
FBBFBBBLLL
FFBBBFBLRL
FBBFBFFRRL
FBFFFFBRRL
BFFFFFFLRR
BBFFFBFLLL
FFFFBBFRLR
FBFFBFFLLL
FFBFBFBLRR
FFBBBFBLRR
FFFFBFBLRR
BBFBFFBLRR
BBFFFFFRRR
FBBBFFBRLL
FBBFBBFRRL
FFBFBBBRRR
BFFBFBBRLR
FFFBFFFRRL
BFBFFFFRLR
FFBFFFFRLL
FBFBBFFRLL
FBFBBBBLLL
BBFFBBFRRL
FFFFFBBLRR
FFBBFBBLLL
FFFBBBBRLL
FFFBFBBRLL
BFBFFBFRLR
FBBBFBBRLR
FFFBFBFLRR
BFFFBBBRLR
FFBBFBBLRR
FFBFBBFRLR
FFFFFBFRLL
FBBFBFFLLR
FFFFBFBRLL
BFFBBFFLLR
BFBBBBBLRR
FBFBBFBRRL
FBFBBFFRRL
BBFFBBBLLL
BFFBBBBRRL
FBBBFBBRRL
BBFBBFFRRL
FBBFBFBRRL
FBFFBFFRRR
FFFFBBBRRL
FFFBBBBRRL
BFBBFFFRRR
FBBBFBBRLL
FFBFFBFRRR
FFFFBFBLLL
FFBFFBBLLR
FBBBFBFLRL
FFFFBBFRLL
BFFFFBFLLL
FFBBFBFLRL
FBFBFBBRLL
FFBBBFFLLR
FFFBFFBRRR
FBFFBBBLLR
FFFBBFBLLR
BFBBBFFLLL
FFFFBFBRRL
BFFFBBBLRR
FFFBBBBRRR
BBFBFBBLRL
FBFFBFFRLL
FBFBBBBRRL
FBBBBBBRLR
FFFBFBFLLR
FFFFBBFRRR
BFBFBFBRRL
FFFBFFBRLL
FFBBBBFRRR
BFFBBBFRRL
FBBFFBBRRR
FBFFBBFLRR
FBFFFFBLRL
FBBBFBBLRR
BBFBFBBRRL
FBBBFFBLLR
FBFBBBFRRL
BFFFBFBRLL
BFFBBFFRRL
FFBBBBFLRR
FFBBFBBLRL
FFBFFFFLRR
FBBFBFFRLL
FFBBBBFLLL
FFBFFFBRRL
BFFBFFFLLR
BBFFFBBRLR
FFBFBFFRLL
FBFBFBBRRL
FBBBFBBLRL
FBFBFFBRLR
FBFFBBFRRR
FBFBFBBLLR
FBBBBBFLRL
FBFBBFFRRR
FBFFBBBLRR
BFFBBFFLRL
FBFBFFBRRR
FFFFBFBRRR
BFBFBFBLLR
BFFFFFBRLR
BFBBBBBLLL
FFBBFBFLRR
BFBBBFFLRR
FFBBBFFLLL
FBBFFFFLRL
BFBFBFFRLR
FFFBFBBRRL
FBFFBBBLRL
FBFBFBBRRR
FBFFFBFLLL
BFFBFFFRRL
BFFBBBBRLR
FFFFFBFRRR
FBBBFFFLRL
BFFFBFBRRL
FFFBBBFLLL
FFBFFFBLRR
FBFBBFFLLL
FBBBBFBRLR
FFFBFFBLLR
FFBFFBBRLR
BFFFBBFRRR
BBFFFFFLRL
FBFFFFBRLR
BFBFBBBLLL
FBFBFBFLLR
BFFFFFBLLL
BFFBFBFRLL
BFFBFBBRRR
BFBBFBBRRL
BBFBBFBRRL
BFBBFFBRLL
BBFBFBBRLL
FFBFFFFLRL
FFBBFBFRLR
FBBFFFFRRL
BFFFFFFRRR
FBBFFBFLLL
BFFBBFBLRR
FBFBFFBLRL
BFBFFFBRLR
BFBFFBFLLL
FFBBBFBRRL
BFFFBFFLLR
FBFBFBBLLL
BBFFFBBLLR
BFBFFBFLRR
FBFFFBBLRL
FBFBBBFRLR
BFFFBFFRLR
BBFFBBFRLR
FBBFBBFLLR
BFFFBBFLLR
FFBFFFBLLR
FBBBBBFRRR
FFFBFFFLLL
BFBFFFBRRL
BFBBBBFRLL
FFBFBBBLLR
BFBBFBFLLR
FBFFBBFLRL
FFFBFFFLLR
BFBFBFBLRL
BBFFBBBRRR
FBFBBFBRLL
FFFFBFFRLR
FBFFBBFRLL
FBFFFBFLRR
FBBFFFBLRL
FBBBBBFLLR
FFFBFBBLLR
BBFFFFBRRR
BFBFFBBRRR
BFFBFBBRRL
BBFFBBFRLL
BFFFFBFRRL
BFFFBBBRRL
FFBBFFFLLL
FFFFBBBLLR
BBFFBBFLLR
FFFBFBBRRR
BBFFFBFRRL
FFBBBBFRLR
FBBFFFBRRR
BFFFBFBLLL
FFBFBFBLRL
FBFFFFBLLR
BBFBFFBRRL
BFBBBFFRLL
FBFBFFFLRR
FBFBFFFLLR
BFBFBBBRRL
FBFFFFFLLR
BFBFFFFLRL
FBBFBBBRLL
BFBFBFFLLL
BFBBFBBLRL
FBFFFBBLLL
BFFFFFFLLL
BBFBFFBLLR
FFBFBBBLLL
FBBBBBBRLL
BFFFBBBLLR
FBFBBBBLRR
FBBFFFBLLL
FFFBFBFRRL
FFFFBBFRRL
FBFFFBFRRL
BBFFFBFRLR
FFFBBFFLLL
FFFFBFFLRL
FBFBBBBLLR
FBBBFFFRRR
FFFBBBFLRR
BFBFBBFRRR
BBFBFFBRLR
BBFBBFBRLL
BFFBBFBLRL
FBFBFBFRRL
BFFFFFFRRL
BBFFFFBRLL
BFFBBFBLLR
BFFBFBFLRR
FBFBFBBRLR
BFBFFBBLRL
FFFBBBFLRL
FBBBFFFLRR
BBFFFFFRRL
FFFBFBBLRR
FFBFFBBRLL
FFBBBBBLLR
BFBBFBBLLL
BBFBFFBRRR
FBFFFFBRRR
FFBBBBBRRR
FBFBBFBLRL
FFFBBBFRLR
BBFFBFFLLR
FFFFFBBRLL
BFFFFFBRRL
FBBBFFFRLL
FFBBBFBRRR
BBFFFBBRRL
FFFBFFFRRR
FFFBBBBLLL
BFBFFBFLRL
BFBBBBFRLR
FBBFBFBLRR
BFBFFFBLRR
FBBBBFFRRR
BFBFBFFRRR
BBFBFBFRRR
FBBFBFBLLR
FFFBBFBRRL
BFFBFFBLLR
BFBFFFFRRR
FBFFBFFRLR
FBFFFBBRLL
FBFFBFFRRL
FBFBFFBLLR
FFFBBFFRRL
FBFBFFFLLL
BFFBFBFLLR
FBBFBBFLRR
BBFFFFBRLR
BFFFBFFLRL
FBBFBFFLRL
FFBFFFFRLR
FFBFFFFLLR
FBBFBFBRLR
FFFFBBFLRL
FFFFBBFLLL
FFFFBBBRLR
BFBFBFFLRL
BFBBFFBRLR
FBFFFFBRLL
BFFFBBFRLL
FBFBFBFRLL
FFFBBFBLRR
FFBFBBBLRL
BFFBFFFRLL
BFBBFFBLLR
BBFFBFBLRR
BFFBBBFLRL
FBBFFFFLRR
BFBBBBBLLR
FFFFFBBLRL
FFBBBBBRLR
FFFFFBBRRR
BFBBBBBRLL
FBBBBBBLRL
BBFFBBBRRL
BFBBFBFLLL
BBFFFBFRLL
FFFBBFFRLR
FFFBFFBRRL
BFFFFBFRLL
BFFBBFFRRR
BFFFBBFRRL
FBBFFFFRLR
BFFBFFBRRR
BFBBFFBRRL
BFFFBFBLRR
BFFBBBBLRR
FBBBFBFLRR
BFBBBFBRLR
FFBBBFFLRL
FFFFBFFLLR
FBBFFBBRLR
BFFBFFBLLL
FBBBBFFLLR
FBBFBFBLLL
BFBFFFFLLL
FBBBFBFRRR
FBFFFBFRLL
FFBBBBBLLL
BBFBBFFLLR
FFBFBBBRLR
BBFBFFFLLL
FFFFFBBRRL
BFFBBBFLLL
BFFBFFFRLR
FFBBBBFRRL
BFBBBBBRLR
BFFBFBFLLL
BFFBBFFLRR
FBBFBBFRLL
BFBFBBFLLR
FBFFFBFRRR
FBFBBBFLRL
BFFBFBBLRL
BBFBBFBLRR
BBFBFFBLRL
BBFBFBBRRR
FBBBBFBRLL
FBFBFBBLRL
BBFBFBFLLR
FBFFBFBRLR
FFBBFFBLRR
FFFBFBFLRL
BFFBFBBRLL
FBBBBFFLRR
BFBBFFFLLR
BFBFFBFRLL
BBFBBFBLLL
FBBBBFFLRL
FBFBFFBLLL
BFFBBFFRLR
FBBFFFBLLR
BFFBFBFRLR
FBBFFBFLRL
FFBFFBFRLL
BFFFBFFRRL
BFBBFBFRRR
FFBBBFFRLL
FBBBBBFLLL
BFBBFFBLRR
FFFFFBBLLR
FFFFFFBRLR
FBBFBBFLRL
BFBFBBFLRR
BBFFBBBLRR
FFFFFBFRRL
BFFFFBFRRR
BFFFBFBRLR
BFFFFBBLLR
FBBBFBFRLL
BFFFBBFLLL
BFBBBFFLLR
BBFFBFFRLR
FBFBFBFLRL
FFBBFFBLRL
BFBBBFBLLL
FFFFBFBRLR
FBFBFFBRRL
BFFFFBFLRR
BFFBBBFRLR
FFFFBBBRLL
FBBBBFFRLR
BFBFBFFRLL
FBBBBBFRLL
BFBFFFBRRR
FBBBBBFRLR
BFFBFFFRRR
FFBBFFBRRL
BFBBBBFLRL
FFFBBBFRRL
FFFFBFFRRR
FFFBBFBRRR
BBFBFFFLLR
BFBBFBBRLL
BFFFFBBLRR
BBFFFBBLLL
FFFBBBFRRR
FFBFBBFRRL
FBFBBBFLRR
FBFFBFBLLL
BFFBFBBLRR
FFBFFFBRLL
FBBBBFFRRL
BBFBFFFRRL
FFBFFBFLRR
BFFBBFBRLR
FFBBBFBLLR
FBBFBFFRLR
BBFBBFBRLR
BFFBFFBRLL
FFBFFFFRRR
BBFFFFFLRR
FFBFBFBRRL
BBFFBFBRLR
FBBBBFBLLL
FBBFFBBLLL
FFBFFBBLLL
BFBFBBFRRL
FBFFFBBLRR
FBFFBFFLRL
BFBBBFFLRL
FBFBBFBLRR
FBFBBBFLLL
FFBBFBBLLR
FFFFBFFRLL
BBFFFFFLLL
FFBBFBFRRL
BFFFBFFLRR
FFBFBBFLRR
FBBFBFBRLL
BFFBBFBRRR
FBFBBFFLLR
BFBBBFFRLR
BBFFBFFRRL
BBFFBBFRRR
BFBBFFBLLL
BFBFBFBRLR
FBFFFBFRLR
FBFBBFBRLR
FFBBBBBLRL
BFFFBBBLLL
FBFFBFBRLL
FBFFFBBRRL
BBFBBFFLLL
BBFFBFBLLR
FBFBBFBRRR
FBFBBBBRLR
FBFBBFFRLR
BFFBBBFLRR
FBBBFFBRLR
BFBBBBFLLL
FFFBFBFRLL
BBFBFBBLLR
FFBFFBBLRL
FBBBFFFRLR
FFBBFBBRRL
BFBFBBBLRR
BBFBBFFRRR
FBBBBBFLRR
BBFBFBFLRL
BFFFFBFLLR
FBFBFFFRLR
FBBBBFBRRL
BFBFFFFLLR
FFBFBFFLRR
BFBFFFBLLR
FFBFBBFLLL
BBFBBFFLRR
BBFFFBFLRR
BFBBFFFRLL
FFFFBBBLRR
BFBFBBFRLL
BFBFFBBRLL
FFBBBBBRLL
BBFFBFFLLL
BFBBBBFRRL
FBBBFFFLLL
FFBFBFFRRR
BFBFFBBRLR
FFFFBFBLRL
FBFFFBBRRR
BFBFFBFRRL
BFBBFBFRLR
FBFBBBBRLL
FFBBFFFRRR
BFBFFFBLLL
FFFBBFFLLR
FFBFBFBRLR
FBBBBBBRRL
FFBBFBBRLL
FFBFFFBRLR
FBBFFFBRLL
FBBBFBBRRR
BFFFBBBRLL
FBBBFBFLLR
BBFFBFFLRL
FBBFBBBLLR
FBBBBBBLRR
BFBBBBBRRR
BFBFBBBRRR
BFBBFFFLLL
FFBBFFBRLL
BBFBBFBLLR
FBFFBBFLLR
FFFFFBFRLR
BBFFFBFLRL
FBBFFBFLRR
FFBFFBFRLR
FBFFBFBLLR
FFFBFBBLLL
BBFFBBFLRR
BBFFBFFRLL
FBBFBFFLLL
BFFFFBBRLL
FFBBFFBLLL
FFBFFBBRRL
BFFBBBFRRR
FBBFBFFRRR
BFFBBBBLLL
FFBBFBBRLR
FFBBFFBRLR
BFFFFBBRRL
FFBBBBBRRL
FBFFBFBRRL
BFFBFBFRRL
BBFBBBFLLL
FBBFFBBLRL
BFBFBFBRLL
FBFFFFFLRL
FFFBFFFLRR
FBBFBBBLRR
FFFFFBFLRR
FBBBBFBLLR
BFFBFBFRRR
BFFFFBBLLL
FFFFBBFLRR
BFBBFBFLRR
FFBFFBBLRR
BBFFFFFRLL
BFBFBFFLLR
BBFBFBFRLL
BBFFBBFLLL
FBFBBBFRLL
FBBFBBFLLL
BFBBFFFRLR
BFBFFBBLLL
FBBBBBFRRL
FFBBFBFRRR
FBFFFFBLLL
FFFFFBBRLR
BFBBFBBRLR
BFBBBBFLLR
BFFBBBFLLR
FFBBFBFLLR
FFBBBFBRLL
FFFFBBBRRR
FBBFFFFRLL
FFBFFFBRRR
FBBFBFBRRR
FFBBBFFRLR
BFBFFFFRLL
FBFFFBFLRL
FFBFFBBRRR
FBFBBFBLLL
BFBFBBFLLL
FFBBBBFLLR
FBFBBBBLRL
BFBBFBFLRL
BFFFBBBLRL
BFBBBBFRRR
FFFBBBFRLL
FFBFBFFLLL
BFBFFFFRRL
FFFBFBBRLR
BFBFBFBLRR
FBFFBBFRLR
BBFFFFBLLL
BFFBBBBRLL
BFFFFBBLRL
BBFBBFBRRR
BFFFFFFLLR
BBFFBBBLLR
FBBFFFFLLR
BBFFBFBRLL
FBBBFFBRRR
FBBFFBBLLR
BBFFFFBLRR
FBFFFFFLLL
FBBFBFFLRR
BFBBFBFRLL
BFBBBFBRRR
FFBFBBFRRR
FBFBFFFRLL
FFFBBFFRRR
FFFFFBFLRL
BFBFFBFLLR
FBBBFFFRRL
FBBBBBBLLR
FBBFFFBRRL
FFBBFFFRLL
BBFBFFFRRR
FFFBBFBLRL
BFFBBBFRLL
FFFFFFBRRR
BFBFFBBLRR
BBFFBBFLRL
FFFBBBFLLR
BFBBBFBLRR
FBBFFBBRRL
FFBFFBFLLR
BBFBBFBLRL
FFFBFFFRLL
BFBFBBBLLR
FFFBFFFLRL
BFFFFBBRLR
BBFFBFFRRR
BFBBFBFRRL
BBFBBFFLRL
FFBBBBFRLL
FBFFBBBLLL
FBFBBFFLRL
BFFFBBBRRR
BFBBBBBRRL
FBBBFFFLLR
`
