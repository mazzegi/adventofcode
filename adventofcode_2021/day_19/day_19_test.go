package day_19

import (
	"adventofcode_2021/testutil"
	"strings"
	"testing"
)

func TestNumBeacons(t *testing.T) {
	res, err := numBeacons(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = 79
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestPart2MainFunc(t *testing.T) {
	res, err := part2MainFunc(inputTest)
	testutil.CheckUnexpectedError(t, err)
	var exp int = -42
	if exp != res {
		t.Fatalf("want %d, have %d", exp, res)
	}
}

func TestOrientations(t *testing.T) {
	os := allOrientations
	pt := p(1, 2, 3)
	var sl []string
	for _, o := range os {
		sl = append(sl, o.apply(pt).String())
	}
	//sort.Strings(sl)
	log("%s", strings.Join(sl, "\n"))

	//
	log("*** inv ***")
	os = allInvOrientations
	pt = p(1, 2, 3)
	var sli []string
	for _, o := range os {
		sli = append(sl, o.apply(pt).String())
	}
	//sort.Strings(sl)
	log("%s", strings.Join(sli, "\n"))
}

func TestOrientations2(t *testing.T) {
	sc := &scanner{
		beacons: []point{
			p(-1, -1, 1),
			p(-2, -2, 2),
			p(-3, -3, 3),
			p(-2, -3, 1),
			p(5, 6, -4),
			p(8, 0, 7),
		},
	}
	oscs := sc.orientations()

	for i, osc := range oscs {
		log("--- orient. %d ---", i)
		for _, pt := range osc.beacons {
			log("%s", pt.String())
		}
	}
}

// func TestCommon(t *testing.T) {
// 	//scs, err := parseScanners(inputTest)
// 	scs, err := parseScanners(input)
// 	testutil.CheckUnexpectedError(t, err)

// 	zeroPt := p(0, 0, 0)
// 	_ = zeroPt
// 	var total []point
// 	_ = total
// 	offsets := map[int]point{}
// 	offsets[0] = p(0, 0, 0)

// 	for {
// 		//uups := false
// 		update := false
// 		for i1, sc1 := range scs {
// 			for i2, sc2 := range scs {
// 				if i1 >= i2 {
// 					continue
// 				}
// 				_, o1ok := offsets[i1]
// 				_, o2ok := offsets[i2]
// 				if !o1ok && !o2ok {
// 					//TODO: remember for extra loop
// 					log("uups %d-%d", i1, i2)
// 					//uups = true
// 					continue
// 				}
// 				if o1ok && o2ok {
// 					// both are already offseted - skip
// 					continue
// 				}
// 				update = true
// 				if o1ok && !o2ok {
// 					ov, off, or := overlapping2(sc1, sc2)
// 					if len(ov) < 12 {
// 						continue
// 					}
// 					osc2 := sc2.oriented(or)
// 					scs[i2] = osc2
// 					offsets[i2] = off.add(offsets[i1])
// 					log("add %d", i2)
// 					log("%d - offset: %s", i2, offsets[i2])
// 					log("%d - %d overlap (%d) (%s): %v", i1, i2, len(ov), off, ov)
// 				} else { //!o1ok && o2ok
// 					ov, off, or := overlapping2(sc2, sc1)
// 					if len(ov) < 12 {
// 						continue
// 					}
// 					osc1 := sc1.oriented(or)
// 					scs[i1] = osc1
// 					offsets[i1] = off.add(offsets[i2])
// 					log("add %d", i1)
// 					log("%d - offset: %s", i1, offsets[i1])
// 					log("%d - %d overlap (%d) (%s): %v", i2, i1, len(ov), off, ov)
// 				}
// 			}
// 		}
// 		log("total-offsets: %d", len(offsets))
// 		if !update {
// 			break
// 		}
// 	}

// 	// sort.Slice(total, func(i, j int) bool {
// 	// 	return total[i].less(total[j])
// 	// })
// 	// for _, p := range total {
// 	// 	log("%s", p)
// 	// }

// 	// sc0 := scs[0]
// 	// //sc1 := scs[1]
// 	// sc1 := scs[1]

// 	// overlap, off := overlapping2(sc0, sc1)
// 	// log("overlap: %v", overlap)
// 	// log("offset: %s", off)
// }
