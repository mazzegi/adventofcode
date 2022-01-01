package day_09

import (
	"adventofcode_2016/errutil"
	"adventofcode_2016/readutil"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Part1() {
	enc := readutil.ReadString(input)
	dec, err := Decompress(enc)
	errutil.ExitOnErr(err)
	fmt.Printf("part-1: decoded len: %d\n", len(dec))
}

func Part2() {
	enc := readutil.ReadString(input)
	decLen, err := Decompress2Len(enc)
	errutil.ExitOnErr(err)
	fmt.Printf("part-2: decoded len: %d\n", decLen)
}

type marker struct {
	amount int
	repeat int
}

func Decompress(in string) (string, error) {
	var dcs string
	pos := 0
	scanMarker := func() (marker, error) {
		var ms string
		for {
			r := in[pos]
			pos++
			if r == ')' {
				var m marker
				_, err := fmt.Sscanf(ms, "%dx%d", &m.amount, &m.repeat)
				if err != nil {
					return marker{}, errors.Wrapf(err, "marker-scanf %q", ms)
				}
				return m, nil
			} else {
				ms += string(r)
			}
		}
	}

	for {
		r := in[pos]
		pos++
		if r == '(' {
			m, err := scanMarker()
			if err != nil {
				return "", errors.Wrapf(err, "scan-marker at %d", pos)
			}
			rdata := in[pos : pos+m.amount]
			dcs += strings.Repeat(rdata, m.repeat)
			pos += m.amount

		} else {
			dcs += string(r)
		}

		if pos >= len(in) {
			return dcs, nil
		}
	}
}

func Decompress2Len(in string) (int, error) {
	//var dcs string
	declen := 0
	pos := 0
	scanMarker := func() (marker, error) {
		var ms string
		for {
			r := in[pos]
			pos++
			if r == ')' {
				var m marker
				_, err := fmt.Sscanf(ms, "%dx%d", &m.amount, &m.repeat)
				if err != nil {
					return marker{}, errors.Wrapf(err, "marker-scanf %q", ms)
				}
				return m, nil
			} else {
				ms += string(r)
			}
		}
	}

	for {
		r := in[pos]
		pos++
		if r == '(' {
			m, err := scanMarker()
			if err != nil {
				return 0, errors.Wrapf(err, "scan-marker at %d", pos)
			}
			rdata := in[pos : pos+m.amount]
			if strings.Index(rdata, "(") > -1 {
				rdataDecLen, err := Decompress2Len(rdata)
				if err != nil {
					return 0, err
				}
				declen += rdataDecLen * m.repeat
			} else {
				declen += len(rdata) * m.repeat
			}
			pos += m.amount

		} else {
			declen += 1
		}

		if pos >= len(in) {
			return declen, nil
		}
	}
}
