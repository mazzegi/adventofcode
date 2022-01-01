package day_08

import (
	"adventofcode_2016/errutil"
	"adventofcode_2016/readutil"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	lines := readutil.ReadLines(input)
	ins, err := ParseInstructions(lines)
	errutil.ExitOnErr(err)
	screen := NewScreen(50, 6)
	screen.Apply(ins)

	onc := screen.PixelsOn()
	fmt.Printf("part1: %d pixels are on\n", onc)
	fmt.Printf("%s", screen.Img())

}

func Part2() {
	//"RURUCEOEIL" is the answer
}

type Instruction interface{}

type RectInstruction struct {
	Width  int
	Height int
}

type RotateRowInstruction struct {
	Row int
	By  int
}

type RotateColInstruction struct {
	Col int
	By  int
}

func ParseInstructions(ins []string) ([]Instruction, error) {
	var is []Instruction
	for _, in := range ins {
		i, err := ParseInstruction(in)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-instruction %q", in)
		}
		is = append(is, i)
	}
	return is, nil
}

func ParseInstruction(s string) (Instruction, error) {
	switch {
	case strings.HasPrefix(s, "rect "):
		var i RectInstruction
		_, err := fmt.Sscanf(s, "rect %dx%d", &i.Width, &i.Height)
		if err != nil {
			return nil, errors.Wrapf(err, "scan-rect %q", s)
		}
		return i, nil
	case strings.HasPrefix(s, "rotate row "):
		var i RotateRowInstruction
		_, err := fmt.Sscanf(s, "rotate row y=%d by %d", &i.Row, &i.By)
		if err != nil {
			return nil, errors.Wrapf(err, "scan-rotate-row %q", s)
		}
		return i, nil
	case strings.HasPrefix(s, "rotate column "):
		var i RotateColInstruction
		_, err := fmt.Sscanf(s, "rotate column x=%d by %d", &i.Col, &i.By)
		if err != nil {
			return nil, errors.Wrapf(err, "scan-rotate-col %q", s)
		}
		return i, nil
	default:
		return nil, errors.Errorf("invalid instruction %q", s)
	}
}

type ScreenRow []bool

type Screen struct {
	Rows          []*ScreenRow
	width, height int
}

func NewScreen(width, height int) *Screen {
	s := &Screen{
		width:  width,
		height: height,
	}
	for y := 0; y < height; y++ {
		row := make(ScreenRow, width)
		s.Rows = append(s.Rows, &row)
	}
	return s
}

func (s *Screen) Img() string {
	var img string
	for _, row := range s.Rows {
		for _, v := range *row {
			if v {
				img += "#"
			} else {
				img += "."
			}
		}
		img += "\n"
	}
	img += "\n"
	return img
}

func (s *Screen) PixelsOn() int {
	var cnt int
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			if s.Get(x, y) {
				cnt++
			}
		}
	}
	return cnt
}

func (s *Screen) SetGet(x, y int, val bool) bool {
	if y < 0 || y >= len(s.Rows) {
		return false
	}
	row := s.Rows[y]
	if x < 0 || x >= len(*row) {
		return false
	}
	pval := (*row)[x]
	(*row)[x] = val
	return pval
}

func (s *Screen) Get(x, y int) bool {
	if y < 0 || y >= len(s.Rows) {
		return false
	}
	row := s.Rows[y]
	if x < 0 || x >= len(*row) {
		return false
	}
	return (*row)[x]
}

func (s *Screen) ApplyOne(i Instruction) {
	switch i := i.(type) {
	case RectInstruction:
		for x := 0; x < i.Width; x++ {
			for y := 0; y < i.Height; y++ {
				s.SetGet(x, y, true)
			}
		}
	case RotateRowInstruction:
		for b := 0; b < i.By; b++ {
			next := s.Get(s.width-1, i.Row)
			for x := 0; x < s.width; x++ {
				next = s.SetGet(x, i.Row, next)
			}
		}
	case RotateColInstruction:
		for b := 0; b < i.By; b++ {
			next := s.Get(i.Col, s.height-1)
			for y := 0; y < s.height; y++ {
				next = s.SetGet(i.Col, y, next)
			}
		}
	}
}

func (s *Screen) Apply(is []Instruction) {
	fmt.Printf("%s", s.Img())
	for _, i := range is {
		s.ApplyOne(i)
		fmt.Printf("%s", s.Img())
	}
}
