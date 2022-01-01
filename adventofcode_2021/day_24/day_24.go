package day_24

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	res, err := largestModelNumber2(input)
	errutil.ExitOnErr(err)
	log("part1: result = %d", res)
}

func Part2() {
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d", res)
}

//
type varValueFunc func(v variable) int

type number interface {
	value(fnc varValueFunc) int
	String() string
}

type fixedNumber int

func (fn fixedNumber) value(fnc varValueFunc) int {
	return int(fn)
}

func (fn fixedNumber) String() string {
	return strconv.FormatInt(int64(fn), 10)
}

type varNumber variable

func (vn varNumber) value(fnc varValueFunc) int {
	return fnc(variable(vn))
}

func (vn varNumber) String() string {
	return variable(vn).String()
}

type variable int

const (
	varW variable = iota
	varX
	varY
	varZ
)

func (v variable) String() string {
	switch v {
	case varW:
		return "w"
	case varX:
		return "x"
	case varY:
		return "y"
	case varZ:
		return "z"
	default:
		return "<na>"
	}
}

type instruction interface {
	VarAndNum() (variable, number)
}

type inp struct {
	v variable
}

func (i inp) VarAndNum() (variable, number) {
	return i.v, fixedNumber(0)
}

type add struct {
	v   variable
	num number
}

func (i add) VarAndNum() (variable, number) {
	return i.v, i.num
}

type mul struct {
	v   variable
	num number
}

func (i mul) VarAndNum() (variable, number) {
	return i.v, i.num
}

type div struct {
	v   variable
	num number
}

func (i div) VarAndNum() (variable, number) {
	return i.v, i.num
}

type mod struct {
	v   variable
	num number
}

func (i mod) VarAndNum() (variable, number) {
	return i.v, i.num
}

type eql struct {
	v   variable
	num number
}

func (i eql) VarAndNum() (variable, number) {
	return i.v, i.num
}

//
func parseInstruction(s string) (instruction, error) {
	cvar := func(s string) (variable, error) {
		switch s {
		case "w":
			return varW, nil
		case "x":
			return varX, nil
		case "y":
			return varY, nil
		case "z":
			return varZ, nil
		default:
			return 0, errors.Errorf("invalid variable %q", s)
		}
	}

	varnum := func(a, b string) (variable, number, error) {
		v, err := cvar(a)
		if err != nil {
			return 0, nil, err
		}
		if n, err := strconv.ParseInt(b, 10, 64); err == nil {
			return v, fixedNumber(n), nil
		}
		bv, err := cvar(b)
		if err != nil {
			return 0, nil, err
		}
		return v, varNumber(bv), nil
	}

	var a, b string
	var v variable
	var err error
	switch {
	case strings.HasPrefix(s, "inp "):
		_, err = fmt.Sscanf(s, "inp %s", &a)
		if err == nil {
			if v, err = cvar(a); err == nil {
				return inp{v: v}, nil
			}
		}
	case strings.HasPrefix(s, "add "):
		_, err = fmt.Sscanf(s, "add %s %s", &a, &b)
		if err == nil {
			if v, num, err := varnum(a, b); err == nil {
				return add{v: v, num: num}, nil
			}
		}
	case strings.HasPrefix(s, "mul "):
		_, err = fmt.Sscanf(s, "mul %s %s", &a, &b)
		if err == nil {
			if v, num, err := varnum(a, b); err == nil {
				return mul{v: v, num: num}, nil
			}
		}
	case strings.HasPrefix(s, "div "):
		_, err = fmt.Sscanf(s, "div %s %s", &a, &b)
		if err == nil {
			if v, num, err := varnum(a, b); err == nil {
				return div{v: v, num: num}, nil
			}
		}
	case strings.HasPrefix(s, "mod "):
		_, err = fmt.Sscanf(s, "mod %s %s", &a, &b)
		if err == nil {
			if v, num, err := varnum(a, b); err == nil {
				return mod{v: v, num: num}, nil
			}
		}
	case strings.HasPrefix(s, "eql "):
		_, err = fmt.Sscanf(s, "eql %s %s", &a, &b)
		if err == nil {
			if v, num, err := varnum(a, b); err == nil {
				return eql{v: v, num: num}, nil
			}
		}
	default:
		return nil, errors.Errorf("invalid prefix in %q", s)
	}
	return nil, err
}

func parseInstructions(s string) ([]instruction, error) {
	var iss []instruction
	lines := readutil.ReadLines(s)
	for _, line := range lines {
		i, err := parseInstruction(line)
		if err != nil {
			return nil, errors.Wrap(err, "parse-instruction")
		}
		iss = append(iss, i)
	}
	if len(iss) == 0 {
		return nil, errors.Errorf("no data")
	}
	return iss, nil
}

//
func splitInSubs(iss []instruction) [][]instruction {
	subs := [][]instruction{}
	currSub := []instruction{}
	for _, is := range iss {
		switch is.(type) {
		case inp:
			if len(currSub) > 0 {
				subs = append(subs, currSub)
				currSub = []instruction{}
			}
		}
		currSub = append(currSub, is)
	}
	if len(currSub) > 0 {
		subs = append(subs, currSub)
	}
	return subs
}

//
type alu struct {
	iss        []instruction
	w, x, y, z int
}

func (a *alu) run(in []int, initZ int) {
	a.w, a.x, a.y = 0, 0, 0
	a.z = initZ
	getV := func(v variable) int {
		switch v {
		case varW:
			return a.w
		case varX:
			return a.x
		case varY:
			return a.y
		case varZ:
			return a.z
		default:
			panic("invalid variable")
		}
	}

	setV := func(v variable, val int) {
		switch v {
		case varW:
			a.w = val
		case varX:
			a.x = val
		case varY:
			a.y = val
		case varZ:
			a.z = val
		default:
			panic("invalid variable")
		}
		//log("set %q = %d", v.String(), val)
	}

	inPos := 0
	nextIn := func() int {
		if inPos >= len(in) {
			panic("next")
		}
		defer func() { inPos++ }()
		return in[inPos]
	}

	for _, is := range a.iss {
		switch is := is.(type) {
		case inp:
			vin := nextIn()
			setV(is.v, vin)
		case add:
			vw := getV(is.v)
			nv := is.num.value(getV)
			setV(is.v, vw+nv)
		case mul:
			vw := getV(is.v)
			nv := is.num.value(getV)
			setV(is.v, vw*nv)
		case div:
			vw := getV(is.v)
			nv := is.num.value(getV)
			if nv == 0 {
				panic("nv is 0")
			}
			setV(is.v, vw/nv)
		case mod:
			vw := getV(is.v)
			nv := is.num.value(getV)
			if vw < 0 || nv <= 0 {
				panic("mod exception")
			}
			setV(is.v, vw%nv)
		case eql:
			vw := getV(is.v)
			nv := is.num.value(getV)
			var rv int
			if vw == nv {
				rv = 1
			} else {
				rv = 0
			}
			setV(is.v, rv)
		}
	}
}

//

//c0 = 1, 26
//c1 = 15, 14, 11, -13, 14, 15, -7, 10, -12, 15, -16, -9, -8, -8
//c2 = 4 16 14 3 11 13 11 7 12 15 13 1  15 4

//z0 % 26 + C1 == in ? z = z0 / C0: z = z0 / C0 * 26 + in + C2
// c0 is always 1 or 26

//
func extraFunc(c0, c1, c2, z0, n int) int {
	if ((z0 % 26) + c1) == n {
		return z0 / c0
	}
	return (z0/c0)*26 + n + c2
}

type resPair struct {
	n  int
	z0 int
}

const eps = 0.0000001

func isMultiple(a int, of int) (int, bool) {
	m := float64(a) / float64(of)
	if math.Abs(m-math.Floor(m)) < eps {
		return int(math.Round(m)), true
	}
	return 0, false
}

func solveExtra(c0, c1, c2, result int) []resPair {
	var rps []resPair

	for n := 1; n <= 9; n++ {
		//first branch
		// z0 / c0 == result ? z0 = result*c0
		if c0 == 1 {
			z0 := result
			if ((z0 % 26) + c1) == n {
				rps = append(rps, resPair{
					n:  n,
					z0: z0,
				})
			}

		} else if c0 == 26 {
			for z0 := result * 26; z0 < (result+1)*26; z0++ {
				if ((z0 % 26) + c1) == n {
					rps = append(rps, resPair{
						n:  n,
						z0: z0,
					})
				}
			}
		}
		matchFirstBranch := func(z0 int) bool {
			return ((z0 % 26) + c1) == n
		}

		//second branch
		// (z0/c0)*26 + n + c2 = result
		if result == n+c2 {
			for z0 := 0; z0 < c0; z0++ {
				if matchFirstBranch(z0) {
					continue
				}

				rps = append(rps, resPair{
					n:  n,
					z0: z0,
				})
			}
		} else if result > n+c2 {
			//
			if c0 == 1 {
				// z0*26 + n + c2 = result
				m, ok := isMultiple(result-(n+c2), 26)
				if ok {
					if matchFirstBranch(m) {
						continue
					}
					rps = append(rps, resPair{
						n:  n,
						z0: m,
					})
				}

				// z0 := (result - (n + c2)) / 26
				// if z0 > 0 {
				// 	rps = append(rps, resPair{
				// 		n:  n,
				// 		z0: z0,
				// 	})
				// }
			} else if c0 == 26 {
				// (z0/26)*26 + n + c2 = result
				//z0 = i*26   => i*26 = result-(n+c2)

				m, ok := isMultiple(result-(n+c2), 26)
				if ok {
					for z := m * 26; z < (m+1)*26; z++ {
						if matchFirstBranch(z) {
							continue
						}
						rps = append(rps, resPair{
							n:  n,
							z0: z,
						})
					}
				}
			}
		}

	}

	return rps
}

// 4, 5, 15
func subParams(sub *alu) (c0, c1, c2 int) {
	i0 := sub.iss[4]
	i1 := sub.iss[5]
	i2 := sub.iss[15]

	_, n0 := i0.VarAndNum()
	_, n1 := i1.VarAndNum()
	_, n2 := i2.VarAndNum()

	dmy := func(variable) int { return 0 }
	c0 = n0.value(dmy)
	c1 = n1.value(dmy)
	c2 = n2.value(dmy)
	return
}

func execSub(sub *alu, in int, z0 int) (z1 int) {
	sub.run([]int{in}, z0)
	return sub.z
}

func largestModelNumber2(in string) (int, error) {
	iss, err := parseInstructions(in)
	if err != nil {
		return 0, errors.Wrap(err, "parse-instructions")
	}
	subs := splitInSubs(iss)
	var subAlus []*alu
	for _, sub := range subs {
		subAlus = append(subAlus, &alu{
			iss: sub,
		})
	}

	maxZ0 := 10000
	findValidInput := func(sub *alu, targetZ1 int) (n int, z0 int, ok bool) {
		for nin := 9; nin >= 1; nin-- {
			for z0 := 0; z0 <= maxZ0; z0++ {
				z1 := execSub(sub, nin, z0)
				if z1 == targetZ1 {
					return nin, z0, true
				}
			}
		}
		return 0, 0, false
	}

	//calc last
	var ins []int
	targetZ1 := 0
	for i := len(subAlus) - 1; i >= 0; i-- {
		sub := subAlus[i]
		n, z0, ok := findValidInput(sub, targetZ1)
		if !ok {
			return 0, errors.Errorf("found no valid input")
		}
		ins = append(ins, n)
		targetZ1 = z0
	}
	log("ins: %v", ins)

	return 0, nil
}

func probeSubs(in string) error {
	iss, err := parseInstructions(in)
	if err != nil {
		return errors.Wrap(err, "parse-instructions")
	}
	subs := splitInSubs(iss)
	var subAlus []*alu
	for _, sub := range subs {
		subAlus = append(subAlus, &alu{
			iss: sub,
		})
	}

	for i, subAlu := range subAlus {
		log("\n*** sub-alu %d ***", i)
		for in := 1; in <= 9; in++ {
			z0 := 0
			hitZs := []int{}
			for {
				z1 := execSub(subAlu, in, z0)
				if z1 == 0 {
					hitZs = append(hitZs, z0)
					if len(hitZs) == 4 {
						break
					}
				}
				z0++
			}
			log("in = %d, z0s = %v", hitZs)
		}
	}

	return nil
}

func compareSubs(in string) error {
	iss, err := parseInstructions(in)
	if err != nil {
		return errors.Wrap(err, "parse-instructions")
	}
	subs := splitInSubs(iss)
	var subAlus []*alu
	for _, sub := range subs {
		subAlus = append(subAlus, &alu{
			iss: sub,
		})
	}

	diff := func(sa1, sa2 *alu) error {
		if len(sa1.iss) != len(sa2.iss) {
			log("different lengths")
			return errors.Errorf("diff lengths")
		}
		for i1, is1 := range sa1.iss {
			is2 := sa2.iss[i1]
			if reflect.TypeOf(is1) != reflect.TypeOf(is2) {
				log("%d: different types: %T != %T", i1, is1, is2)
				return errors.Errorf("diff types")
			}
			v1, n1 := is1.VarAndNum()
			v2, n2 := is2.VarAndNum()
			if v1 != v2 {
				log("%d: different vars: %s != %s", i1, v1, v2)
				return errors.Errorf("diff vars")
			}
			if n1.String() != n2.String() {
				if i1 == 4 || i1 == 5 || i1 == 15 {
					if i1 == 4 {
						if (n1.String() != "1") && (n1.String() != "26") {
							log(n1.String())
							return errors.Errorf("unexpected 4(1)")
						}
						if (n2.String() != "1") && (n2.String() != "26") {
							log(n2.String())
							return errors.Errorf("unexpected 4(2)")
						}
					}

					log("%d: different nums: %s != %s", i1, n1, n2)
					continue
				} else {
					log("%d: different nums: %s != %s", i1, n1, n2)
					return errors.Errorf("diff nums unexpected")
				}
			}
			log("%d: equal", i1)
		}
		return nil
	}

	for i1, subAlu1 := range subAlus {
		for i2, subAlu2 := range subAlus {
			if i2 <= i1 {
				continue
			}
			log("\n*** compare %d and %d ***", i1, i2)
			err := diff(subAlu1, subAlu2)
			if err != nil {
				panic("diff")
			}
		}
	}

	return nil
}

func testExtra(in string) error {
	iss, err := parseInstructions(in)
	if err != nil {
		return errors.Wrap(err, "parse-instructions")
	}
	subs := splitInSubs(iss)
	var subAlus []*alu
	for _, sub := range subs {
		subAlus = append(subAlus, &alu{
			iss: sub,
		})
	}

	funcs := []func(z0, n int) int{}
	for i, subAlu := range subAlus {
		c0, c1, c2 := subParams(subAlu)
		log("%d: params c0=%d, c1=%d, c2=%d", i, c0, c1, c2)
		funcs = append(funcs, func(z0, n int) int {
			return extraFunc(c0, c1, c2, z0, n)
		})
	}

	for i, subAlu := range subAlus {
		for n := 1; n <= 9; n++ {
			for z := 0; z <= 100; z++ {
				subRes := execSub(subAlu, n, z)
				fncRes := funcs[i](z, n)
				if subRes != fncRes {
					log("%d: diff: z=%d, n=%d => sub=%d, fnc=%d", i, z, n, subRes, fncRes)
				}
			}
		}
	}

	return nil
}

func testSolveExtra(in string) error {
	iss, err := parseInstructions(in)
	if err != nil {
		return errors.Wrap(err, "parse-instructions")
	}
	subs := splitInSubs(iss)
	var subAlus []*alu
	for _, sub := range subs {
		subAlus = append(subAlus, &alu{
			iss: sub,
		})
	}

	funcs := []func(z0, n int) int{}
	for i, subAlu := range subAlus {
		c0, c1, c2 := subParams(subAlu)
		log("%d: params c0=%d, c1=%d, c2=%d", i, c0, c1, c2)
		funcs = append(funcs, func(z0, n int) int {
			return extraFunc(c0, c1, c2, z0, n)
		})
	}

	containsPair := func(pairs []resPair, n, z int) bool {
		for _, p := range pairs {
			if p.n == n && p.z0 == z {
				return true
			}
		}
		return false
	}

	for i, fnc := range funcs {
		subAlu := subAlus[i]
		c0, c1, c2 := subParams(subAlu)
		for n := 1; n <= 9; n++ {
			for z := 0; z <= 100; z++ {
				if i == 3 &&
					n == 1 &&
					z == 40 {
					log("check this")
				}
				res := fnc(z, n)

				solvs := solveExtra(c0, c1, c2, res)
				log("%d: solvers (n=%d, z=%d): %v", i, n, z, solvs)
				//check
				for _, solv := range solvs {
					rs := fnc(solv.z0, solv.n)
					if rs != res {
						log("uuuuppppss")
					}
				}

				if !containsPair(solvs, n, z) {
					log("%d: (n=%d, z=%d) not contains", i, n, z)
				}
			}
		}
	}

	return nil
}

func reverseSolve(in string) error {
	iss, err := parseInstructions(in)
	if err != nil {
		return errors.Wrap(err, "parse-instructions")
	}
	subs := splitInSubs(iss)
	var subAlus []*alu
	for _, sub := range subs {
		subAlus = append(subAlus, &alu{
			iss: sub,
		})
	}

	funcs := []func(z0, n int) int{}
	for i, subAlu := range subAlus {
		c0, c1, c2 := subParams(subAlu)
		log("%d: params c0=%d, c1=%d, c2=%d", i, c0, c1, c2)
		funcs = append(funcs, func(z0, n int) int {
			return extraFunc(c0, c1, c2, z0, n)
		})
	}

	// last
	type treePair struct {
		n      int
		z0     int
		childs []*treePair
		parent *treePair
		level  int
	}

	root := &treePair{
		n:      0,
		z0:     0,
		childs: []*treePair{},
		level:  len(funcs),
	}
	nodes := []*treePair{root}
	for i := len(funcs) - 1; i >= 0; i-- {
		c0, c1, c2 := subParams(subAlus[i])

		newNodes := []*treePair{}
		for _, node := range nodes {
			spairs := solveExtra(c0, c1, c2, node.z0)
			for _, p := range spairs {
				newNode := &treePair{
					p.n,
					p.z0,
					nil,
					node,
					node.level - 1,
				}
				node.childs = append(node.childs, newNode)
				newNodes = append(newNodes, newNode)
			}
		}
		nodes = newNodes
	}

	var collect func(node *treePair) [][]int
	collect = func(node *treePair) [][]int {
		matches := [][]int{}
		for _, c := range node.childs {
			if c.level == 0 {
				matches = append(matches, []int{c.n})
			} else {
				cms := collect(c)
				matches = append(matches, cms...)
			}
		}
		for i := range matches {
			matches[i] = append(matches[i], node.n)
		}

		return matches
	}

	matches := [][]int{}
	for _, c := range root.childs {
		cms := collect(c)
		matches = append(matches, cms...)
	}

	for _, m := range matches {
		log("%s", fmtDigits(m))
	}
	log("collected %d matches", len(matches))

	last := matches[len(matches)-1]
	alu := &alu{iss: iss}
	alu.run(last, 0)
	log("last-z: %d", alu.z)

	return nil
}

func fmtDigits(ds []int) string {
	var s string
	for _, d := range ds {
		s += strconv.FormatInt(int64(d), 10)
	}
	return s
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
