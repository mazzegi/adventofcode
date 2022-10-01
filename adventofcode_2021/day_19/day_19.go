package day_19

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2021/day_19/rot3d"
	"github.com/mazzegi/adventofcode/adventofcode_2021/errutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/intutil"
	"github.com/mazzegi/adventofcode/adventofcode_2021/readutil"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	probe2()
	//probe3()
	//fuck()
	// res, err := numBeacons(input)
	// errutil.ExitOnErr(err)
	// fmt.Printf("part1: result = %d\n", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part2: result = %d\n", res)
}

type point struct {
	x, y, z int
}

func p(x, y, z int) point {
	return point{x: x, y: y, z: z}
}

func (p point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

func (p point) less(op point) bool {
	if p.x < op.x {
		return true
	}
	if p.x > op.x {
		return false
	}
	if p.y < op.y {
		return true
	}
	if p.y > op.y {
		return false
	}
	return p.z < op.z
}

func (pt point) add(op point) point {
	return p(pt.x+op.x, pt.y+op.y, pt.z+op.z)
}

func (pt point) sub(op point) point {
	return p(pt.x-op.x, pt.y-op.y, pt.z-op.z)
}

func parsePoint(s string) (point, error) {
	var p point
	_, err := fmt.Sscanf(s, "%d,%d,%d", &p.x, &p.y, &p.z)
	if err != nil {
		return point{}, errors.Wrapf(err, "scan-point %q", s)
	}

	return p, nil
}

type scanner struct {
	beacons []point
	//offsetTo0   point
	//orientation orientation
}

func parseScanners(in string) ([]*scanner, error) {
	var scs []*scanner
	lines := readutil.ReadLines(in)
	var curr *scanner
	for _, line := range lines {
		if strings.HasPrefix(line, "---") {
			if curr != nil {
				scs = append(scs, curr)
			}
			curr = &scanner{}
			continue
		}
		if curr == nil {
			return nil, errors.Errorf("scaned point but no scanner is initialized yet")
		}
		p, err := parsePoint(line)
		if err != nil {
			return nil, errors.Wrap(err, "parse-point")
		}
		curr.beacons = append(curr.beacons, p)
	}
	if len(curr.beacons) > 0 {
		scs = append(scs, curr)
	}
	if len(scs) == 0 {
		return nil, errors.Errorf("no data")
	}
	return scs, nil
}

type orientation []rot3d.Matrix

func (o orientation) apply(pt point) point {
	rp := rot3d.Point{
		X: float64(pt.x),
		Y: float64(pt.y),
		Z: float64(pt.z),
	}

	for _, m := range o {
		rp = m.Rotate(rp)
	}

	return p(int(math.Round(rp.X)), int(math.Round(rp.Y)), int(math.Round(rp.Z)))
}

func (o orientation) applyF(pt rot3d.Point) rot3d.Point {
	for _, m := range o {
		pt = m.Rotate(pt)
	}
	return pt
}

var allOrientations []orientation
var allInvOrientations []orientation

func init() {
	for x := 0; x <= 3; x++ {
		for y := 0; y <= 3; y++ {
			for z := 0; z <= 3; z++ {
				allOrientations = append(allOrientations, []rot3d.Matrix{
					rot3d.XMatrix(float64(x) * rot3d.Rot90),
					rot3d.YMatrix(float64(y) * rot3d.Rot90),
					rot3d.ZMatrix(float64(z) * rot3d.Rot90),
				})
				allInvOrientations = append(allOrientations, []rot3d.Matrix{
					rot3d.ZMatrix(-float64(z) * rot3d.Rot90),
					rot3d.YMatrix(-float64(y) * rot3d.Rot90),
					rot3d.XMatrix(-float64(x) * rot3d.Rot90),
				})
			}
		}
	}
}

// func allOrientationsFnc() []orientation {

// 	angles := [][3]float64{
// 		{0, 0, 0},
// 		{rot3d.Rot90, 0, 0},
// 		{rot3d.Rot180, 0, 0},
// 		{rot3d.Rot270, 0, 0},

// 		{0, 0, rot3d.Rot180},
// 		{rot3d.Rot90, 0, rot3d.Rot180},
// 		{rot3d.Rot180, 0, rot3d.Rot180},
// 		{rot3d.Rot270, 0, rot3d.Rot180},

// 		// +/- z
// 		{0, rot3d.Rot90, 0},
// 		{0, rot3d.Rot90, rot3d.Rot90}, //***
// 		{0, rot3d.Rot90, rot3d.Rot180},
// 		{0, rot3d.Rot90, rot3d.Rot270},

// 		{0, rot3d.Rot270, 0},
// 		{0, rot3d.Rot270, rot3d.Rot90},
// 		{0, rot3d.Rot270, rot3d.Rot180},
// 		{0, rot3d.Rot270, rot3d.Rot270},

// 		// +/- z
// 		{0, 0, rot3d.Rot90},
// 		{0, rot3d.Rot90, rot3d.Rot90},
// 		{0, rot3d.Rot180, rot3d.Rot90},
// 		{0, rot3d.Rot270, rot3d.Rot90},

// 		{0, 0, rot3d.Rot270},
// 		{0, rot3d.Rot90, rot3d.Rot270},
// 		{0, rot3d.Rot180, rot3d.Rot270},
// 		{0, rot3d.Rot270, rot3d.Rot270},
// 	}
// 	var os []orientation
// 	for _, as := range angles {
// 		os = append(os, []rot3d.Matrix{
// 			rot3d.XMatrix(as[0]),
// 			rot3d.YMatrix(as[1]),
// 			rot3d.ZMatrix(as[2]),
// 		})
// 	}
// 	return os

// 	// return []orientation{
// 	// 	// +/- x
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(0), rot3d.ZMatrix(0)},
// 	// 	{rot3d.XMatrix(rot3d.Rot90), rot3d.YMatrix(0), rot3d.ZMatrix(0)},
// 	// 	{rot3d.XMatrix(rot3d.Rot180), rot3d.YMatrix(0), rot3d.ZMatrix(0)},
// 	// 	{rot3d.XMatrix(rot3d.Rot270), rot3d.YMatrix(0), rot3d.ZMatrix(0)},

// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(0), rot3d.ZMatrix(rot3d.Rot180)},
// 	// 	{rot3d.XMatrix(rot3d.Rot90), rot3d.YMatrix(0), rot3d.ZMatrix(rot3d.Rot180)},
// 	// 	{rot3d.XMatrix(rot3d.Rot180), rot3d.YMatrix(0), rot3d.ZMatrix(rot3d.Rot180)},
// 	// 	{rot3d.XMatrix(rot3d.Rot270), rot3d.YMatrix(0), rot3d.ZMatrix(rot3d.Rot180)},

// 	// 	// +/- z
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot90), rot3d.ZMatrix(0)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot90), rot3d.ZMatrix(rot3d.Rot90)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot90), rot3d.ZMatrix(rot3d.Rot180)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot90), rot3d.ZMatrix(rot3d.Rot270)},

// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot270), rot3d.ZMatrix(0)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot270), rot3d.ZMatrix(rot3d.Rot90)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot270), rot3d.ZMatrix(rot3d.Rot180)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot270), rot3d.ZMatrix(rot3d.Rot270)},

// 	// 	// +/- z
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(0), rot3d.ZMatrix(rot3d.Rot90)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot90), rot3d.ZMatrix(rot3d.Rot90)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot180), rot3d.ZMatrix(rot3d.Rot90)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot270), rot3d.ZMatrix(rot3d.Rot90)},

// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(0), rot3d.ZMatrix(rot3d.Rot270)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot90), rot3d.ZMatrix(rot3d.Rot270)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot180), rot3d.ZMatrix(rot3d.Rot270)},
// 	// 	{rot3d.XMatrix(0), rot3d.YMatrix(rot3d.Rot270), rot3d.ZMatrix(rot3d.Rot270)},
// 	// }
// }

// func allInverseOrientations() []orientation {
// 	angles := [][3]float64{
// 		{0, 0, 0},
// 		{rot3d.Rot90, 0, 0},
// 		{rot3d.Rot180, 0, 0},
// 		{rot3d.Rot270, 0, 0},

// 		{0, 0, rot3d.Rot180},
// 		{rot3d.Rot90, 0, rot3d.Rot180},
// 		{rot3d.Rot180, 0, rot3d.Rot180},
// 		{rot3d.Rot270, 0, rot3d.Rot180},

// 		// +/- z
// 		{0, rot3d.Rot90, 0},
// 		{0, rot3d.Rot90, rot3d.Rot90},
// 		{0, rot3d.Rot90, rot3d.Rot180},
// 		{0, rot3d.Rot90, rot3d.Rot270},

// 		{0, rot3d.Rot270, 0},
// 		{0, rot3d.Rot270, rot3d.Rot90},
// 		{0, rot3d.Rot270, rot3d.Rot180},
// 		{0, rot3d.Rot270, rot3d.Rot270},

// 		// +/- z
// 		{0, 0, rot3d.Rot90},
// 		{0, rot3d.Rot90, rot3d.Rot90},
// 		{0, rot3d.Rot180, rot3d.Rot90},
// 		{0, rot3d.Rot270, rot3d.Rot90},

// 		{0, 0, rot3d.Rot270},
// 		{0, rot3d.Rot90, rot3d.Rot270},
// 		{0, rot3d.Rot180, rot3d.Rot270},
// 		{0, rot3d.Rot270, rot3d.Rot270},
// 	}
// 	var os []orientation
// 	for _, as := range angles {
// 		os = append(os, []rot3d.Matrix{
// 			rot3d.ZMatrix(-as[2]),
// 			rot3d.YMatrix(-as[1]),
// 			rot3d.XMatrix(-as[0]),
// 		})
// 	}
// 	return os
// }

func pointOrientations(p point) []point {
	var pos []point
	os := allOrientations
	for _, o := range os {
		pos = append(pos, o.apply(p))
	}
	return pos
}

func (s *scanner) orientations() []*scanner {
	var oscs []*scanner
	os := allOrientations
	for _, o := range os {
		osc := &scanner{
			//orientation: o,
		}
		for _, pt := range s.beacons {
			osc.beacons = append(osc.beacons, o.apply(pt))
		}

		oscs = append(oscs, osc)
	}
	return oscs
}

func (s *scanner) oriented(or orientation) *scanner {
	osc := &scanner{
		//orientation: or,
	}
	for _, pt := range s.beacons {
		osc.beacons = append(osc.beacons, or.apply(pt))
	}
	return osc
}

func clonePoints(pts []point) []point {
	cpts := make([]point, len(pts))
	copy(cpts, pts)
	return cpts
}

func overlapping(sc1, sc2 *scanner) []point {
	pt1sorted := clonePoints(sc1.beacons)
	sort.Slice(pt1sorted, func(i, j int) bool {
		return pt1sorted[i].less(pt1sorted[j])
	})

	pt2sorted := clonePoints(sc2.beacons)
	sort.Slice(pt2sorted, func(i, j int) bool {
		return pt2sorted[i].less(pt2sorted[j])
	})

	// assume the first defines the offset
	pt10 := pt1sorted[0]
	pt20 := pt2sorted[0]
	offset := pt10.sub(pt20)
	outOfRange := func(pt2 point) bool {
		if (intutil.AbsInt(pt2.x) > 1000) || (intutil.AbsInt(pt2.y) > 1000) {
			return true
		}
		return false
	}

	s1Contains := func(pt point) bool {
		for _, pt1 := range pt1sorted {
			if pt1 == pt {
				return true
			}
		}
		return false
	}

	var ops []point
	for _, pt2 := range pt2sorted {
		pt2off := pt2.add(offset)
		if outOfRange(pt2off) {
			continue
		}
		if !s1Contains(pt2off) {
			return []point{}
		}
		ops = append(ops, pt2off)
	}

	return ops
}

func numBeacons(in string) (int, error) {
	scs, err := parseScanners(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse scanners")
	}

	for i1, sc1 := range scs {
		for i2, sc2 := range scs {
			if i1 == i2 {
				continue
			}
			for _, sc2o := range sc2.orientations() {
				ovl := overlapping(sc1, sc2o)
				_ = ovl
			}
		}
	}

	return 0, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}

const RANGE = 1000

/*
func (s *scanner) orientations() []*scanner {
	var oscs []*scanner
	os := allOrientations()
	for _, o := range os {
		osc := &scanner{
			//orientation: o,
		}
		for _, pt := range s.beacons {
			osc.beacons = append(osc.beacons, o.apply(pt))
		}

		oscs = append(oscs, osc)
	}
	return oscs
}

func (s *scanner) oriented(or orientation) *scanner {
	osc := &scanner{
		//orientation: or,
	}
	for _, pt := range s.beacons {
		osc.beacons = append(osc.beacons, or.apply(pt))
	}
	return osc
}
*/

func overlapping2(sc1, sc2 *scanner) (values []point, values2 []point, offset point, or orientation) {
	invOs := allInvOrientations
	var maxOv []point
	var maxOv2 []point
	var maxOff point
	var maxOr orientation
	var maxOrIdx int
	allOrs := allOrientations
	for i, or := range allOrs {
		sc2o := sc2.oriented(or)
		//log("test or %d: %v", i, sc2o.beacons)
		ovl, ovl2, off := overlapPattern(sc1.beacons, sc2o.beacons)
		// if len(ovl) >= 12 {
		if maxOv == nil || (len(ovl) > len(maxOv)) {
			maxOv = ovl
			maxOv2 = ovl2
			maxOff = off
			maxOr = or
			maxOrIdx = i
		}
		// }
	}

	// for i, sc2o := range sc2.orientations() {
	// 	ovl, off := overlapPattern(sc1.beacons, sc2o.beacons)
	// 	if len(ovl) >= 12 {
	// 		if maxOv == nil || (len(ovl) > len(maxOv)) {
	// 			maxOv = ovl
	// 			maxOff = off
	// 			maxOr = sc2o.orientation
	// 			maxOrIdx = i
	// 		}
	// 		//invO := invOs[i]

	// 		//return ovl, off
	// 		//return ovl, invO.apply(off)
	// 		//return ovl, sc2o.orientation.apply(off)
	// 	}
	// }

	_ = maxOr
	_ = maxOrIdx
	// if len(maxOv) >= 12 {
	// 	log("off=%s; oridx=%d; appl=%s; applinv=%s", maxOff, maxOrIdx, maxOr.apply(maxOff), invOs[maxOrIdx].apply(maxOff))
	// }

	//rotate maxOvs2 back
	invOr := invOs[maxOrIdx]
	var imaxOv2 []point
	for _, pt := range maxOv2 {
		imaxOv2 = append(imaxOv2, invOr.apply(pt))
	}

	//_ = invOs
	//return maxOv, maxOv2, maxOff, maxOr
	return maxOv, imaxOv2, maxOff, maxOr
}

func null(n int) {}

func overlapPattern(pts1 []point, pts2 []point) (values []point, values2 []point, offset point) {
	var maxCommon []point
	var maxCommon2 []point
	var maxOff point
	for _, pt1 := range pts1 {
		for i, pt2 := range pts2 {
			common, common2 := commonPoints(pts1, pts2, pt1, pt2)
			// if len(common) >= 12 {
			if maxCommon == nil || len(common) > len(maxCommon) {
				maxCommon = common
				maxCommon2 = common2
				//pt2Inv := or.apply(pt2)

				maxOff = pt1.sub(pt2)
				//maxOff = pt1.sub(pt2Inv)
			}

			//return common, pt1.sub(pt2)
			// }
			null(i)
		}
	}

	return maxCommon, maxCommon2, maxOff
}

func commonPoints(pts1 []point, pts2 []point, cp1, cp2 point) ([]point, []point) {
	pts2Contains := func(pt point) bool {
		for _, pt2 := range pts2 {
			if pt2 == pt {
				return true
			}
		}
		return false
	}

	var common []point
	var common2 []point
	for _, pt1 := range pts1 {
		diff := pt1.sub(cp1)
		tpt2 := cp2.add(diff)
		if pts2Contains(tpt2) {
			common = append(common, pt1)
			common2 = append(common2, tpt2)
		}
	}
	return common, common2
}

func probe() {
	scs, err := parseScanners(input)
	errutil.ExitOnErr(err)
	//testutil.CheckUnexpectedError(t, err)

	zeroPt := p(0, 0, 0)
	_ = zeroPt
	var total []point
	_ = total
	offsets := map[int]point{}
	offsets[0] = p(0, 0, 0)

	for {
		//uups := false
		update := false
		for i1, sc1 := range scs {
			for i2, sc2 := range scs {
				if i1 >= i2 {
					continue
				}
				_, o1ok := offsets[i1]
				_, o2ok := offsets[i2]
				if !o1ok && !o2ok {
					//TODO: remember for extra loop
					log("uups %d-%d", i1, i2)
					//uups = true
					continue
				}
				if o1ok && o2ok {
					// both are already offseted - skip
					continue
				}
				update = true
				if o1ok && !o2ok {
					ov, _, off, or := overlapping2(sc1, sc2)
					if len(ov) < 12 {
						if len(ov) > 0 {
							log("too less overlap %d-%d: %d", i1, i2, len(ov))
						}
						continue
					}
					osc2 := sc2.oriented(or)
					scs[i2] = osc2
					offsets[i2] = off.add(offsets[i1])
					log("add %d", i2)
					log("%d - offset: %s", i2, offsets[i2])
					log("%d - %d overlap (%d) (%s): %v", i1, i2, len(ov), off, ov)
				} else { //!o1ok && o2ok
					ov, _, off, or := overlapping2(sc2, sc1)
					if len(ov) < 12 {
						if len(ov) > 0 {
							log("too less overlap %d-%d: %d", i2, i1, len(ov))
						}
						continue
					}
					osc1 := sc1.oriented(or)
					scs[i1] = osc1
					offsets[i1] = off.add(offsets[i2])
					log("add %d", i1)
					log("%d - offset: %s", i1, offsets[i1])
					log("%d - %d overlap (%d) (%s): %v", i2, i1, len(ov), off, ov)
				}
			}
		}
		log("total-offsets: %d", len(offsets))
		if !update {
			break
		}
	}
}

type scannerResult struct {
	sc       *scanner
	sco      *scanner
	offset   point
	offsetT0 point
	refbase  int
	or       orientation
}

func probe2() {
	scs, err := parseScanners(input)
	errutil.ExitOnErr(err)
	//testutil.CheckUnexpectedError(t, err)

	zeroPt := p(0, 0, 0)
	_ = zeroPt
	offsets := map[int]point{}
	offsets[0] = p(0, 0, 0)

	results := make([]*scannerResult, len(scs))

	startIdx := 0
	results[startIdx] = &scannerResult{
		sc:       scs[startIdx],
		sco:      scs[startIdx],
		offset:   p(0, 0, 0),
		offsetT0: p(0, 0, 0),
		or:       orientation{},
		refbase:  startIdx,
	}

	firstNilResultIdx := func() int {
		for i, res := range results {
			if res == nil {
				return i
			}
		}
		return -1
	}
	_ = firstNilResultIdx()

	ext := 0
	var total []point
	total = append(total, results[0].sco.beacons...)
	for {
		act := false
		for ir, res := range results {
			if res == nil {
				continue
			}
			for si, sc := range scs {
				if results[si] != nil {
					continue
				}

				//ov, _, off, or := overlapping2(res.sc, sc)
				ov, _, off, or := overlapping2(res.sco, sc)
				if len(ov) < 12 {
					// log("test %d-%d - NO overlap", ir, si)
					// if ir == 4 && si == 7 {
					// 	log("dbg: %d: %v", ir, res.sc.beacons)
					// 	log("dbg: %d: %v", si, sc.beacons)
					// }
					continue
				}
				osc := sc.oriented(or)

				//offo := or.apply(off)
				offTo0 := off.add(res.offsetT0)

				var totalOffsetted []point
				for _, p := range osc.beacons {
					totalOffsetted = append(totalOffsetted, p.add(off).add(res.offsetT0))
				}
				total = append(total, totalOffsetted...)
				//offTo0 := offo.add(res.offsetT0)
				results[si] = &scannerResult{
					sc:       sc,
					sco:      osc,
					offset:   off,
					offsetT0: offTo0,
					or:       or, // maybe this should be the orientation relative to zero
					refbase:  res.refbase,
				}
				log("%d (base = %d, ref = %d) - offset: %s; offsett0 = %s", si, ir, res.refbase, off, offTo0)
				act = true
			}
		}
		if !act {
			ext--
			if ext < 0 {
				break
			}
			//pick first nil
			// if fni := firstNilResultIdx(); fni > -1 {
			// 	results[fni] = &scannerResult{
			// 		sc:      scs[fni],
			// 		offset:  p(0, 0, 0),
			// 		refbase: fni,
			// 		or:      orientation{},
			// 	}

			// } else {
			// 	break
			// }
			//break
		}
	}

	for i, res := range results {
		if res != nil {
			log("%02d: (ref = %d) offset = %s offsetto0 = %s", i, res.refbase, res.offset, res.offsetT0)
		} else {
			log("%02d: nil", i)
		}
	}

	log("*** total (%d) ***", len(total))
	sort.Slice(total, func(i, j int) bool {
		return total[i].less(total[j])
	})

	uniqueMap := map[point]bool{}
	for _, p := range total {
		uniqueMap[p] = true
		log("%s", p)
	}

	var unique []point
	for p := range uniqueMap {
		unique = append(unique, p)
	}

	var maxDist int
	for _, res1 := range results {
		for _, res2 := range results {
			p1 := res1.offsetT0
			p2 := res2.offsetT0
			dist := intutil.AbsInt(p1.x-p2.x) + intutil.AbsInt(p1.y-p2.y) + intutil.AbsInt(p1.z-p2.z)
			if dist > maxDist {
				log("max %s + %s => %d", p1, p2, dist)
				maxDist = dist
			}
		}
	}

	log("total unique: %d (maxdist = %d)", len(uniqueMap), maxDist)
}

func probe3() {
	scs, err := parseScanners(input)
	errutil.ExitOnErr(err)

	idx1, idx2 := 13, 30

	sc1 := scs[idx1]
	sc2 := scs[idx2]

	log("%d: %v", idx1, sc1.beacons)
	log("%d: %v", idx2, sc2.beacons)

	ov, ov2, off, or := overlapping2(sc1, sc2)
	log("%d-%d: ov=%d, off=%s", idx1, idx2, len(ov), off)
	log("%d-%d: ov: %v", idx1, idx2, ov)
	log("%d-%d: ov2: %v", idx1, idx2, ov2)

	ov, ov2, off, or = overlapping2(sc2, sc1)
	log("%d-%d: ov=%d, off=%s", idx2, idx1, len(ov), off)
	log("%d-%d: ov: %v", idx2, idx1, ov)
	log("%d-%d: ov2: %v", idx2, idx1, ov2)

	_ = or

}

/*
13: [(-711,624,322) (-804,740,-434) (832,427,-610)
30: [(-341,687,-600) (415,780,-484) (591,-856,-797)
*/

var f1 = `
--- scanner 0 ---
-711,624,322
-804,740,-434
832,427,-610

--- scanner 1 ---
-341,687,-600
415,780,-484
591,-856,-797
`

func fuck() {
	scs, _ := parseScanners(f1)

	idx1, idx2 := 0, 1

	sc1 := scs[idx1]
	sc2 := scs[idx2]

	ov, ov2, off, or := overlapping2(sc1, sc2)
	log("%d-%d: ov=%d, off=%s", idx1, idx2, len(ov), off)
	log("%d-%d: ov: %v", idx1, idx2, ov)
	log("%d-%d: ov2: %v", idx1, idx2, ov2)

	log("\n")
	ov, ov2, off, or = overlapping2(sc2, sc1)
	log("%d-%d: ov=%d, off=%s", idx2, idx1, len(ov), off)
	log("%d-%d: ov: %v", idx2, idx1, ov)
	log("%d-%d: ov2: %v", idx2, idx1, ov2)

	_ = or
}
