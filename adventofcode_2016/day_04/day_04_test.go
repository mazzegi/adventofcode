package day_04

import (
	"fmt"
	"testing"
)

func TestParseRoom(t *testing.T) {
	tests := []struct {
		in   string
		room Room
		real bool
	}{
		{
			in: "aaaaa-bbb-z-y-x-123[abxyz]",
			room: Room{
				NameEnc:  "aaaaa-bbb-z-y-x",
				SectorID: 123,
				Checksum: "abxyz",
			},
			real: true,
		},
		{
			in: "a-b-c-d-e-f-g-h-987[abcde]",
			room: Room{
				NameEnc:  "a-b-c-d-e-f-g-h",
				SectorID: 987,
				Checksum: "abcde",
			},
			real: true,
		},
		{
			in: "not-a-real-room-404[oarel]",
			room: Room{
				NameEnc:  "not-a-real-room",
				SectorID: 404,
				Checksum: "oarel",
			},
			real: true,
		},
		{
			in: "totally-real-room-200[decoy]",
			room: Room{
				NameEnc:  "totally-real-room",
				SectorID: 200,
				Checksum: "decoy",
			},
			real: false,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			room, err := ParseRoom(test.in)
			if err != nil {
				t.Fatalf("failed but no error expected")
			}
			if room != test.room {
				t.Fatalf("parse: expect %v, got %v", test.room, room)
			}

			real := IsRealRoom(room)
			if test.real != real {
				t.Fatalf("is-real: expect %t, got %t", test.real, real)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		in    rune
		times int
		out   rune
	}{
		{
			in:    'a',
			times: 1,
			out:   'b',
		},
		{
			in:    'm',
			times: 8,
			out:   'u',
		},
		{
			in:    'v',
			times: 12,
			out:   'h',
		},
		{
			in:    'v',
			times: 12 + 26,
			out:   'h',
		},
		{
			in:    'v',
			times: 11 + 34*26,
			out:   'g',
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			rr, err := Rotate(test.in, test.times)
			if err != nil {
				t.Fatalf("failed but no error expected")
			}
			if rr != test.out {
				t.Fatalf("rotate: expect %q, got %q", string(test.out), string(rr))
			}
		})
	}
}
