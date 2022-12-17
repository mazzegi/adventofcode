package day_15

import (
	"fmt"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/grid"
	"github.com/mazzegi/adventofcode/list"
	"github.com/mazzegi/adventofcode/mathutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/set"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func fatal(pattern string, args ...interface{}) {
	panic(fmt.Sprintf(pattern+"\n", args...))
}

var (
	skip1 = true
)

func Part1() {
	if skip1 {
		return
	}
	res, err := part1MainFunc(input, 2000000)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input, 4000000)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type Sensor struct {
	Pos    grid.Point
	Beacon grid.Point
	Dist   int
}

func (sen Sensor) covers(p grid.Point) bool {
	return sen.Pos.ManhattenDistTo(p) <= sen.Dist
}

type Interval struct {
	Low, High int
}

func (i Interval) Contains(n int) bool {
	return n >= i.Low && n <= i.High
}

type Intervals struct {
	list *list.List[*Interval]
}

func NewIntervals(is ...Interval) *Intervals {
	ivs := &Intervals{
		list: list.New[*Interval](),
	}
	for _, i := range is {
		i := i
		ivs.list.PushBack(&i)
	}
	return ivs
}

func (is *Intervals) Has(i Interval) bool {
	for n := is.list.First; n != nil; n = n.Next {
		if *n.Value == i {
			return true
		}
	}
	return false
}

func (is *Intervals) Slice() []Interval {
	var sl []Interval
	is.list.Each(func(i *Interval) {
		sl = append(sl, *i)
	})
	return sl
}

func (is *Intervals) Join(i Interval) {
	mfirst := is.list.Find(nil, func(ti *Interval) bool {
		return i.Low <= ti.Low || i.Low <= ti.High
	})
	if mfirst == nil {
		is.list.PushBack(&i)
		return
	}
	if i.High < mfirst.Value.Low {
		is.list.PushFront(&i)
		return
	}
	if i.High <= mfirst.Value.High {
		mfirst.Value.Low = mathutil.Min(i.Low, mfirst.Value.Low)
		mfirst.Value.High = mathutil.Max(i.High, mfirst.Value.High)
		return
	}

	//find next
	mnext := mfirst
	marker := mfirst.Next
	for ; mnext != nil; mnext = mnext.Next {
		marker = mnext.Next
		if mnext.Next == nil {
			break
		}
		if i.High < mnext.Next.Value.Low {
			break
		}
	}
	var high int
	if mnext != nil {
		high = mathutil.Max(i.High, mnext.Value.High)
	} else {
		high = i.High
	}
	mfirst.Value.Low = mathutil.Min(i.Low, mfirst.Value.Low)
	mfirst.Value.High = high

	//remove until marker
	rem := mfirst.Next
	for rem != marker {
		nextRem := rem.Next
		is.list.Remove(rem)
		rem = nextRem
	}

	n := is.list.First
	for n != nil {
		if n.Next == nil {
			break
		}
		nn := n.Next
		if n.Value.High+1 == n.Next.Value.Low {
			n.Next.Value.Low = n.Value.Low
			is.list.Remove(n)
		}
		n = nn
	}
}

func (is *Intervals) contains(v int) bool {
	for n := is.list.First; n != nil; n = n.Next {
		if n.Value.Contains(v) {
			return true
		}
	}
	return false
}

// func (is *Intervals) Join(i Interval) {
// 	for ix, ei := range is.values {
// 		if i.Low >= ei.Low && i.Low <= ei.High {
// 			// search high
// 			for hix := ix; hix < len(is.values); hix++ {
// 				if i.High <= is.values[hix].High {
// 					ei.High = is.values[hix].High
// 					vs := is.values[:ix+1]
// 					vs = append(vs, is.values[hix+1:]...)
// 					is.values = vs
// 					return
// 				}
// 			}
// 			ei.High = i.High
// 			is.values = is.values[:ix+1]
// 			return
// 		}
// 	}
// 	is.values = append(is.values, &i)
// }

func (sen Sensor) rowCoverInterval(y int, lowLimit, highLimit int) (Interval, bool) {
	if y > sen.Pos.Y+sen.Dist || y < sen.Pos.Y-sen.Dist {
		return Interval{}, false
	}
	left := sen.Dist - mathutil.Abs(sen.Pos.Y-y)
	if left <= 0 {
		return Interval{}, false
	}
	i := Interval{
		Low:  mathutil.Max(lowLimit, sen.Pos.X-left),
		High: mathutil.Min(highLimit, sen.Pos.X+left),
	}
	return i, true
}

func mustParseSensor(s string) Sensor {
	//Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	var sen Sensor
	_, err := fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sen.Pos.X, &sen.Pos.Y, &sen.Beacon.X, &sen.Beacon.Y)
	if err != nil {
		fatal(err.Error())
	}
	sen.Dist = sen.Pos.ManhattenDistTo(sen.Beacon)
	return sen
}

func part1MainFunc(in string, row int) (int, error) {
	var sensors []Sensor
	for _, line := range readutil.ReadLines(in) {
		sen := mustParseSensor(line)
		sensors = append(sensors, sen)
	}
	if len(sensors) == 0 {
		fatal("no sensors")
	}

	beacons := set.New[grid.Point]()
	var minXCovered int
	var maxXCovered int
	for i, sen := range sensors {
		if i == 0 {
			minXCovered = sen.Pos.X - sen.Dist
			maxXCovered = sen.Pos.X + sen.Dist
			continue
		}
		if sen.Pos.X-sen.Dist < minXCovered {
			minXCovered = sen.Pos.X - sen.Dist
		}
		if sen.Pos.X+sen.Dist > maxXCovered {
			maxXCovered = sen.Pos.X + sen.Dist
		}
		beacons.Insert(sen.Beacon)
	}

	isCoveredByAny := func(p grid.Point) bool {
		for _, sen := range sensors {
			if sen.covers(p) {
				return true
			}
		}
		return false
	}

	cnt := 0
	for x := minXCovered; x <= maxXCovered; x++ {
		p := grid.Pt(x, row)
		if beacons.Contains(p) {
			continue
		}
		if isCoveredByAny(p) {
			cnt++
		}
	}

	return cnt, nil
}

func part2MainFunc(in string, limit int) (int, error) {
	var sensors []Sensor
	beaconPoss := set.New[grid.Point]()
	sensorPoss := set.New[grid.Point]()
	for _, line := range readutil.ReadLines(in) {
		sen := mustParseSensor(line)
		sensors = append(sensors, sen)
		beaconPoss.Insert(sen.Beacon)
		sensorPoss.Insert(sen.Pos)
	}
	if len(sensors) == 0 {
		fatal("no sensors")
	}

	// isCoveredByAny := func(p grid.Point) bool {
	// 	for _, sen := range sensors {
	// 		if sen.covers(p) {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	// coveredXIntervals := func(y int) []Interval {
	// 	for _, sen := range sensors {
	// 		iv := sen.coversXInRow(y, 0, limit)
	// 	}
	// }

	var dpos grid.Point
	found := false
	whole := Interval{0, limit}
outer:
	for y := 0; y <= limit; y++ {
		if y%100 == 0 {
			log("y = %d", y)
		}
		cis := NewIntervals()
		for _, sen := range sensors {
			i, ok := sen.rowCoverInterval(y, 0, limit)
			if !ok {
				continue
			}
			//log("sen %02d: y=%d: [%d, %d]", si, y, i.Low, i.High)
			cis.Join(i)
		}
		if cis.Has(whole) {
			continue
		}

		for x := 0; x <= limit; x++ {
			if !cis.contains(x) {
				found = true
				dpos = grid.Pt(x, y)
				break outer
			}

			// p := grid.Pt(x, y)
			// if sensorPoss.Contains(p) || beaconPoss.Contains(p) {
			// 	continue
			// }
			// if !isCoveredByAny(p) {
			// 	found = true
			// 	dpos = p
			// 	break outer
			// }
		}
	}
	if !found {
		fatal("not found")
	}
	sig := 4000000*dpos.X + dpos.Y

	return sig, nil
}
