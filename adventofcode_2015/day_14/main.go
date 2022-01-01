package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	in := input
	scanner := bufio.NewScanner(bytes.NewBufferString(in))
	var rds []ReindeerDescriptor
	for scanner.Scan() {
		s := strings.Trim(scanner.Text(), " \r\n\t")
		if s == "" {
			continue
		}
		rd, err := ParseReindeerDescriptor(s)
		if err != nil {
			panic(err)
		}
		rds = append(rds, rd)
	}
	fmt.Printf("scanned %d reindeer descriptors\n", len(rds))

	var rs []*Reindeer
	var maxPos int
	for _, rd := range rds {
		r := NewReindeerAndNext(rd)
		rs = append(rs, r)
		if r.Position > maxPos {
			maxPos = r.Position
		}
	}
	for _, r := range rs {
		if r.Position == maxPos {
			r.LeadPoints++
		}
	}

	secs := 2503
	//secs := 1000
	for i := 1; i < secs; i++ {
		var maxPos int
		for _, r := range rs {
			r.Next()
			if r.Position > maxPos {
				maxPos = r.Position
			}
		}
		for _, r := range rs {
			if r.Position == maxPos {
				r.LeadPoints++
			}
		}
	}
	//fmt.Printf("winner is %s with %d km\n", lead.Name, lead.Position)
	for _, r := range rs {
		fmt.Printf("%s: %d lead points\n", r.Name, r.LeadPoints)
	}
}

func ParseReindeerDescriptor(s string) (ReindeerDescriptor, error) {
	var rd ReindeerDescriptor
	_, err := fmt.Sscanf(s, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
		&rd.Name, &rd.Speed, &rd.RunningDuration, &rd.RestDuration)
	if err != nil {
		return ReindeerDescriptor{}, err
	}
	return rd, nil
}

type Reindeer struct {
	ReindeerDescriptor
	Position     int
	Running      bool
	CurrDuration int
	LeadPoints   int
}

func NewReindeerAndNext(rd ReindeerDescriptor) *Reindeer {
	return &Reindeer{
		ReindeerDescriptor: rd,
		Position:           rd.Speed,
		Running:            true,
		CurrDuration:       1,
	}
}

func (r *Reindeer) Next() {
	if r.Running {
		if r.CurrDuration+1 > r.RunningDuration {
			r.Running = false
			r.CurrDuration = 1
		} else {
			r.Position += r.Speed
			r.CurrDuration += 1
		}
	} else {
		if r.CurrDuration+1 > r.RestDuration {
			r.Running = true
			r.Position += r.Speed
			r.CurrDuration = 1
		} else {
			r.CurrDuration += 1
		}
	}
}

type ReindeerDescriptor struct {
	Name            string
	Speed           int
	RunningDuration int
	RestDuration    int
}

var inputTest = `
Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
`

var input = `
Rudolph can fly 22 km/s for 8 seconds, but then must rest for 165 seconds.
Cupid can fly 8 km/s for 17 seconds, but then must rest for 114 seconds.
Prancer can fly 18 km/s for 6 seconds, but then must rest for 103 seconds.
Donner can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 11 km/s for 12 seconds, but then must rest for 125 seconds.
Comet can fly 21 km/s for 6 seconds, but then must rest for 121 seconds.
Blitzen can fly 18 km/s for 3 seconds, but then must rest for 50 seconds.
Vixen can fly 20 km/s for 4 seconds, but then must rest for 75 seconds.
Dancer can fly 7 km/s for 20 seconds, but then must rest for 119 seconds.
`
