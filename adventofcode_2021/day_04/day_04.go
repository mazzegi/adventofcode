package day_04

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"

	"github.com/pkg/errors"
)

func Part1() {
	score, err := WinningScore(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: score = %d\n", score)
}

func Part2() {
	score, err := LastWinningScore(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: score = %d\n", score)
}

type BoardPoint struct {
	number int
	marked bool
}

type BoardRow struct {
	points []*BoardPoint
}

func (r *BoardRow) Mark(number int) {
	for _, p := range r.points {
		if p.number == number {
			p.marked = true
		}
	}
}

func (r *BoardRow) AllMarked() bool {
	for _, p := range r.points {
		if !p.marked {
			return false
		}
	}
	return true
}

func (r *BoardRow) UnmarkedSum() int {
	var sum int
	for _, p := range r.points {
		if !p.marked {
			sum += p.number
		}
	}
	return sum
}

type Board struct {
	rows []*BoardRow
	won  bool
}

func (b *Board) Mark(number int) {
	for _, row := range b.rows {
		row.Mark(number)
	}
	if !b.won && b.IsWin() {
		b.won = true
	}
}

func (b *Board) IsWin() bool {
	for _, row := range b.rows {
		if row.AllMarked() {
			return true
		}
	}
	firstRow := b.rows[0]
	for colIdx := range firstRow.points {
		if b.ColAllMarked(colIdx) {
			return true
		}
	}

	return false
}

func (b *Board) ColAllMarked(colIdx int) bool {
	for _, row := range b.rows {
		if !row.points[colIdx].marked {
			return false
		}
	}
	return true
}

func (b *Board) UnmarkedSum() int {
	var sum int
	for _, row := range b.rows {
		sum += row.UnmarkedSum()
	}
	return sum
}

//

func ParseBoard(srows []string, rowSize int) (*Board, error) {
	board := &Board{}
	for _, srow := range srows {
		ns, err := readutil.ReadInts(srow, " ")
		if err != nil {
			return nil, errors.Wrapf(err, "read-row %q", srow)
		}
		if len(ns) != rowSize {
			return nil, errors.Errorf("invalid row-size. want %d, have %d", rowSize, len(ns))
		}

		br := &BoardRow{}
		for _, n := range ns {
			br.points = append(br.points, &BoardPoint{
				number: n,
				marked: false,
			})
		}
		board.rows = append(board.rows, br)
	}
	return board, nil
}

func ParseBoards(in string, rowSize int, colSize int) ([]*Board, error) {
	var bs []*Board
	lines := readutil.ReadLines(in)
	for i := 0; i < len(lines); i += colSize {
		bls := lines[i : i+colSize]
		b, err := ParseBoard(bls, rowSize)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-board %v", bls)
		}
		bs = append(bs, b)
	}
	return bs, nil
}

func WinningScore(in Input) (int, error) {
	numbers, err := readutil.ReadInts(in.numbers, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-numbers")
	}
	boards, err := ParseBoards(in.boards, 5, 5)
	if err != nil {
		return 0, errors.Wrap(err, "parse-boards")
	}

	for _, num := range numbers {
		for _, b := range boards {
			b.Mark(num)
			if b.IsWin() {
				return num * b.UnmarkedSum(), nil
			}
		}
	}

	return 0, nil
}

func LastWinningScore(in Input) (int, error) {
	numbers, err := readutil.ReadInts(in.numbers, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-numbers")
	}
	boards, err := ParseBoards(in.boards, 5, 5)
	if err != nil {
		return 0, errors.Wrap(err, "parse-boards")
	}

	var lastWinBoard *Board
	var lastWinNumber int

	for _, num := range numbers {
		for _, b := range boards {
			if b.won {
				continue
			}

			b.Mark(num)
			if b.won {
				lastWinBoard = b
				lastWinNumber = num
			}
		}
	}

	if lastWinBoard != nil {
		return lastWinNumber * lastWinBoard.UnmarkedSum(), nil
	}
	return 0, errors.Errorf("no board won")
}
