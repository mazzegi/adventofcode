package day_13

import (
	"adventofcode_2017/errutil"
	"adventofcode_2017/readutil"
	"fmt"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := tripSeverity(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := minDelay(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

type layer struct {
	depth      int
	rng        int
	scannerPos int
	effSize    int
	//scannerDir int
}

func newLayer() *layer {
	return &layer{
		//scannerDir: 1,
	}
}

func parseLayer(s string) (*layer, error) {
	l := newLayer()
	var d, r int
	_, err := fmt.Sscanf(s, "%d: %d", &d, &r)
	if err != nil {
		return nil, errors.Wrapf(err, "scan-layer %q", s)
	}
	l.depth = d
	l.rng = r
	l.effSize = r + r - 2
	return l, nil
}

func parseLayers(in string) ([]*layer, error) {
	var ls []*layer
	lines := readutil.ReadLines(in)
	for _, line := range lines {
		l, err := parseLayer(line)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}
	if len(ls) == 0 {
		return nil, errors.Errorf("no data")
	}
	return ls, nil
}

func tripSeverity(in string) (int, error) {
	ls, err := parseLayers(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-layers")
	}

	var maxDepth int
	lsMap := map[int]*layer{}
	for _, l := range ls {
		lsMap[l.depth] = l
		if l.depth > maxDepth {
			maxDepth = l.depth
		}
	}

	moveScanners := func() {
		for _, l := range ls {
			l.scannerPos++
			if l.scannerPos >= l.effSize {
				l.scannerPos = 0
			}

			// if l.scannerDir == 1 && l.scannerPos == l.rng-1 {
			// 	l.scannerDir = -1
			// }
			// if l.scannerDir == -1 && l.scannerPos == 0 {
			// 	l.scannerDir = 1
			// }
			// l.scannerPos += l.scannerDir
		}
	}

	var totalSev int
	for pos := 0; pos <= maxDepth; pos++ {
		if l, ok := lsMap[pos]; ok {
			if l.scannerPos == 0 {
				totalSev += l.depth * l.rng
			}
		}
		moveScanners()
	}

	return totalSev, nil
}

func minDelay(in string) (int, error) {
	ls, err := parseLayers(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-layers")
	}

	var maxDepth int
	lsMap := map[int]*layer{}
	for _, l := range ls {
		lsMap[l.depth] = l
		if l.depth > maxDepth {
			maxDepth = l.depth
		}
	}

	moveScanners := func() {
		for _, l := range ls {
			l.scannerPos++
			if l.scannerPos >= l.effSize {
				l.scannerPos = 0
			}
			// if l.scannerDir == 1 && l.scannerPos == l.rng-1 {
			// 	l.scannerDir = -1
			// }
			// if l.scannerDir == -1 && l.scannerPos == 0 {
			// 	l.scannerDir = 1
			// }
			// l.scannerPos += l.scannerDir
		}
	}

	resetScanners := func() {
		for _, l := range ls {
			l.scannerPos = 0
		}
	}

	moveScannersTimes := func(times int) {
		for _, l := range ls {
			l.scannerPos += times
			if l.scannerPos >= l.effSize {
				l.scannerPos = l.scannerPos % l.effSize
			}
		}
	}

	// steps until back 2*rng-2
	// rng -> rng -1
	// back -> rng -1

	isCaught := func(delay int) bool {
		resetScanners()
		moveScannersTimes(delay)
		// for i := 0; i < delay; i++ {
		// 	moveScanners()
		// }

		for pos := 0; pos <= maxDepth; pos++ {
			if l, ok := lsMap[pos]; ok {
				if l.scannerPos == 0 {
					//log("caught by %d", pos)
					return true
				}
			}
			moveScanners()
		}
		return false
	}

	del := 0
	for {
		//log("*** probe delay %d ***", del)
		ic := isCaught(del)
		if !ic {
			break
		}
		if del%1000 == 0 {
			log("tried %d", del)
		}
		del++
	}

	return del, nil
}
