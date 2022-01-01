package day_09

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"

	"github.com/pkg/errors"
)

func Part1() {
	score, _, err := groupScore(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part1: result = %d\n", score)
}

func Part2() {
	_, garbage, err := groupScore(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", garbage)
}

type keyChar rune

const (
	keyGroupOpen     keyChar = '{'
	keyGroupClose    keyChar = '}'
	keyGarbageOpen   keyChar = '<'
	keyGarbageClose  keyChar = '>'
	keyGarbageIgnore keyChar = '!'
	keyGroupSep      keyChar = ','
)

type stream struct {
	data      []rune
	pos       int
	inGarbage bool
}

func newStream(data []rune) *stream {
	return &stream{
		data:      data,
		pos:       0,
		inGarbage: false,
	}
}

func (s *stream) scanGarbage() int {
	var scanned int
	for {
		if s.pos >= len(s.data) {
			return scanned
		}
		c := s.data[s.pos]
		switch keyChar(c) {
		case keyGarbageClose:
			s.pos++
			return scanned
		case keyGarbageIgnore:
			s.pos += 2
		default:
			scanned++
			s.pos++
		}
	}
}

func (s *stream) scanGroup(level int) (score int, garbage int) {
	for {
		if s.pos >= len(s.data) {
			return score, garbage
		}
		c := s.data[s.pos]
		switch keyChar(c) {
		case keyGroupClose:
			s.pos++
			return score, garbage
		case keyGroupOpen:
			s.pos++
			sc, gs := s.scanGroup(level + 1)
			score += sc
			score += level
			garbage += gs
		case keyGroupSep:
			s.pos++
		case keyGarbageOpen:
			s.pos++
			gs := s.scanGarbage()
			garbage += gs
		default:
			s.pos++
		}
	}
}

func (s *stream) iterate() (score int, garbage int) {
	s.pos = 0
	if s.data[0] != rune(keyGroupOpen) {
		return 0, 0
	}
	s.pos++
	score = 1
	sc, gs := s.scanGroup(2)
	score += sc
	return score, gs
}

func groupScore(in string) (int, int, error) {
	streamData := readutil.ReadString(in)
	if streamData == "" {
		return 0, 0, errors.Errorf("no entries")
	}

	stream := newStream([]rune(streamData))
	score, garbage := stream.iterate()

	return score, garbage, nil
}
