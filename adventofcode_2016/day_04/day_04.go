package day_04

import (
	"adventofcode_2016/errutil"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	rs, err := ParseRooms(input)
	errutil.ExitOnErr(err)
	var total int
	var real int
	var sectorSum int
	for _, r := range rs {
		total++
		if IsRealRoom(r) {
			real++
			sectorSum += r.SectorID
		}
	}
	fmt.Printf("%d of %d are real: sec-sum: %d \n", real, total, sectorSum)
}

func Part2() {
	rs, err := ParseRooms(input)
	errutil.ExitOnErr(err)
	for _, room := range rs {
		if !IsRealRoom(room) {
			continue
		}
		dec, err := DecodeRoomName(room)
		errutil.ExitOnErr(err)
		if strings.Contains(dec, "north") && strings.Contains(dec, "stor") {
			fmt.Printf("%d: %q\n", room.SectorID, dec)
		}
	}
}

func ParseRooms(in string) ([]Room, error) {
	var rs []Room
	sl := strings.Split(in, "\n")
	for _, s := range sl {
		s = strings.Trim(s, " \r\n\t")
		if s == "" {
			continue
		}
		r, err := ParseRoom(s)
		if err != nil {
			return nil, errors.Wrapf(err, "parse room %q", s)
		}
		rs = append(rs, r)
	}
	return rs, nil
}

type Room struct {
	NameEnc  string
	SectorID int
	Checksum string
}

func ParseRoom(s string) (Room, error) {
	cs := s[len(s)-7:]
	rest := s[:len(s)-7]
	if cs[0] != '[' || cs[6] != ']' {
		return Room{}, errors.Errorf("invalid checksum %q", cs)
	}

	idxLastDash := strings.LastIndexByte(rest, '-')
	if idxLastDash < 0 {
		return Room{}, errors.Errorf("found no sector-id sep in rest %q", rest)
	}
	secIDStr := rest[idxLastDash+1:]
	secID, err := strconv.ParseInt(secIDStr, 10, 64)
	if err != nil {
		return Room{}, errors.Wrapf(err, "parse sectorID %q", secIDStr)
	}
	nameEnc := rest[:idxLastDash]
	return Room{
		NameEnc:  nameEnc,
		SectorID: int(secID),
		Checksum: cs[1:6],
	}, nil
}

func IsRealRoom(r Room) bool {
	type occ struct {
		letter rune
		count  int
	}
	var occs []occ

	//count letters
	m := map[rune]int{}
	for _, r := range r.NameEnc {
		if r == '-' {
			continue
		}
		m[r]++
	}
	for r, c := range m {
		occs = append(occs, occ{letter: r, count: c})
	}
	sort.Slice(occs, func(i, j int) bool {
		if occs[i].count == occs[j].count {
			return occs[i].letter < occs[j].letter
		}
		return occs[i].count > occs[j].count
	})

	for i, csl := range r.Checksum {
		if csl != occs[i].letter {
			return false
		}
	}

	return true
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func idxOf(r rune) int {
	for i, ar := range alphabet {
		if ar == r {
			return i
		}
	}
	return -1
}

func Rotate(r rune, times int) (rune, error) {
	idx := idxOf(r)
	if idx < 0 {
		return '#', errors.Errorf("%q is not in alphabet", string(r))
	}
	rIdx := (idx + times) % len(alphabet)
	return rune(alphabet[rIdx]), nil
}

func DecodeRoomName(room Room) (string, error) {
	var dec string
	for _, r := range room.NameEnc {
		if r == '-' {
			dec += " "
			continue
		}
		rr, err := Rotate(r, room.SectorID)
		if err != nil {
			return "", err
		}
		dec += string(rr)
	}
	return dec, nil
}
