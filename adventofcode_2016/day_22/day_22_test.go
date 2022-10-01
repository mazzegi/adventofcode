package day_22

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2016/testutil"
)

var inputTest = `
/dev/grid/node-x0-y0   10T    8T     2T   80%
/dev/grid/node-x0-y1   11T    6T     5T   54%
/dev/grid/node-x0-y2   32T   28T     4T   87%
/dev/grid/node-x1-y0    9T    7T     2T   77%
/dev/grid/node-x1-y1    8T    0T     8T    0%
/dev/grid/node-x1-y2   11T    7T     4T   63%
/dev/grid/node-x2-y0   10T    6T     4T   60%
/dev/grid/node-x2-y1    9T    8T     1T   88%
/dev/grid/node-x2-y2    9T    6T     3T   66%
`

func TestFewestSteps(t *testing.T) {
	ns, err := ParseNodes(inputTest)
	testutil.CheckUnexpectedError(t, err)
	//n := fewestSteps(ns)
	n := fewestStepsSimple(ns, 20, 3, 3)
	exp := 7
	if exp != n {
		t.Fatalf("want %d, have %d", exp, n)
	}
}
