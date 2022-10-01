package jolt

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mazzegi/adventofcode/adventofcode_2020/ints"
)

func matches(v, prev int) bool {
	if v <= prev {
		return false
	}
	return v-prev <= 3
}

func FindChain(in []int) ([]int, bool) {
	if len(in) <= 2 {
		panic("nonsense")
	}
	sort.Ints(in)
	for i := 1; i < len(in); i++ {
		if !matches(in[i], in[i-1]) {
			return nil, false
		}
	}

	return in, true
}

func validChain(in []int, skipIdx ...int) ([]int, bool) {
	mustSkip := func(i int) bool {
		for _, si := range skipIdx {
			if si == i {
				return true
			}
		}
		return false
	}

	start := 1
	for mustSkip(start) {
		start++
	}

	prev := in[0]
	sub := []int{0}
	for i := start; i < len(in); i++ {
		if mustSkip(i) {
			sub = append(sub, -1)
			continue
		}
		val := in[i]
		sub = append(sub, val)

		if !matches(val, prev) {
			return nil, false
		}
		prev = val
	}
	return sub, true
}

func combinations(cnt, min, max int, blacklist map[int]bool) <-chan []int {
	res := make(chan []int)
	go func() {
		blacklisted := func(v int) bool {
			_, ok := blacklist[v]
			return ok
		}
		for blacklisted(min) {
			min++
		}
		for blacklisted(max) {
			max--
		}

		defer close(res)
		vec := make([]int, cnt)
		for i := 0; i < len(vec); i++ {
			vec[i] = min
		}
		send := func() {
			rvec := make([]int, cnt)
			copy(rvec, vec)
			res <- rvec
		}

		sortedStrict := func(sl []int) bool {
			if len(sl) <= 1 {
				return true
			}
			for i := 1; i < len(sl); i++ {
				if sl[i] <= sl[i-1] {
					return false
				}
			}
			return true
		}

		// allMax := func(sl []int) bool {
		// 	for _, v := range sl {
		// 		if v != max {
		// 			return false
		// 		}
		// 	}
		// 	return true
		// }

		done := false
		for !done {
			if sortedStrict(vec) {
				//sortedStrict(vec)
				//if !ints.HasDuplicate(vec) {
				send()
			}
			// if allMax(vec) {
			// 	return
			// }
			for i := 0; i < len(vec); i++ {
				if i < len(vec)-1 && vec[i]+1 >= vec[i+1] {
					continue
				}

				vec[i]++
				for blacklisted(vec[i]) {
					vec[i]++
				}
				if vec[i] <= max {
					break
				} else if i < len(vec)-1 {
					vec[i] = min
				} else {
					done = true
				}
			}
		}
	}()
	return res
}

func uniqueCombis(cnt int, vals []int) <-chan []int {
	res := make(chan []int)
	go func() {
		defer close(res)
		if cnt == 0 {
			return
		}
		if len(vals) == 0 {
			return
		}
		if cnt > len(vals) {
			return
		}
		vec := make([]int, cnt)
		idx := make([]int, cnt)
		for i := 0; i < cnt; i++ {
			vec[i] = vals[i]
			idx[i] = i
		}
		send := func() {
			rvec := make([]int, cnt)
			copy(rvec, vec)
			res <- rvec
		}

		maxIdx := len(vals) - 1
		lastDigit := len(idx) - 1
		next := func() bool {
			if idx[lastDigit]+1 <= maxIdx {
				idx[lastDigit]++
				return true
			}
			for d := len(idx) - 2; d >= 0; d-- {
				if idx[d]+1 < idx[d+1] {
					idx[d]++
					for id := d + 1; id < len(idx); id++ {
						idx[id] = idx[id-1] + 1
						if idx[id] > maxIdx {
							return false
						}
					}
					return true
				}
			}

			return false
		}

		for {
			send()
			n := next()
			if n == false {
				return
			}
			for i, ix := range idx {
				vec[i] = vals[ix]
			}
		}
	}()
	return res
}

func Arrangements(in []int) int {
	sort.Ints(in)
	if _, ok := validChain(in); !ok {
		return 0
	}
	//fmt.Printf("find arrangements for (%d): %v\n", len(in), in)

	// gather numbers, which may definitely not removed
	//blacklist := map[int]bool{}
	var whitelist []int
	var whitelistVals []int
	// var bl []int
	// var blvs []int
	for i := 1; i < len(in)-1; i++ {
		if _, ok := validChain(in, i); !ok {
			// blacklist[i] = true
			// bl = append(bl, i)
			// blvs = append(bl, in[i])
		} else {
			whitelist = append(whitelist, i)
			whitelistVals = append(whitelistVals, in[i])
		}

	}
	// fmt.Printf("whitelist        (%d): %v\n", len(whitelist), whitelist)
	// fmt.Printf("whitelist-values (%d): %v\n", len(whitelistVals), whitelistVals)

	// find out which combis of 2 or 3 may not be removed
	// for 2
	var nrCombis2 int
	for i := 0; i < len(whitelist)-1; i++ {
		test := []int{whitelist[i], whitelist[i+1]}
		if _, ok := validChain(in, test...); !ok {
			nrCombis2++
			//fmt.Printf("nr-2: %v\n", test)
		}
	}

	n := len(whitelist)
	var notValids3 [][]int
	var nrCombis3 int
	var correct3 int
	for i := 0; i < len(whitelist)-2; i++ {
		test := []int{whitelist[i], whitelist[i+1], whitelist[i+2]}
		if _, ok := validChain(in, test...); !ok {
			corr := sumBinCoeff(n-3-nrCombis3*3, 0)

			// if nrCombis3 > 0 {
			// 	corr -= sumBinCoeff(n-3-nrCombis3*3, 0)
			// }
			correct3 += corr
			//fmt.Printf("correct + %d => %d\n", corr, correct3)

			nrCombis3++
			//fmt.Printf("nr-3: %v\n", test)
			notValids3 = append(notValids3, test)

		}
	}

	//crux(len(notValids3), n)

	//sub2 := sumBinCoeff(n-2, 0)
	//sub3 := sumBinCoeff(n-3, 0)

	// fmt.Printf("correct3: %d\n", correct3)
	// fmt.Printf("not-rem-combis-2: %d (sum = %d)\n", nrCombis2, sub2)
	// fmt.Printf("not-rem-combis-3: %d (sum = %d)\n", nrCombis3, sub3)

	//cntBF := sumBinCoeff(n, 0)
	// fmt.Printf("bc: %d\n", cntBF)
	// fmt.Printf("bc-corr: %d\n", cntBF-nrCombis2*sub2-correct3)
	//return 0

	// fmt.Printf("blacklist (%d): %v\n", len(bl), bl)
	// fmt.Printf("blacklist-values (%d): %v\n", len(bl), blvs)

	// maxRem := len(in)/3 + 1
	// for rem := 1; rem <= maxRem; rem++ {
	// 	for comb := range combinations(rem, 1, len(in)-2, blacklist) {
	// 		if set, ok := validChain(in, comb...); ok {
	// 			cnt++
	// 			//fmt.Printf("%v -> %v\n", comb, set)
	// 			fmt.Printf("(%d / %d) (%v) -> %vcnt: %d\n", rem, maxRem, comb, set, cnt)
	// 		}
	// 	}
	// }

	//maxRem := len(in)/3 + 1
	maxRem := len(in)
	cnt := 1 // the whole chain itself
	totCombis := 1
	notVals3Counts := make([]int, len(notValids3))
	notValHits := make([]int, len(notValids3))
	var notValAny int
	for rem := 1; rem <= maxRem; rem++ {
		for comb := range uniqueCombis(rem, whitelist) {
			totCombis++
			_, ok := validChain(in, comb...)
			if ok {
				cnt++
				// if cnt%1000 == 0 {
				// 	fmt.Printf("(%d/%d): -> %d\n", rem, maxRem, cnt)
				// }
				//fmt.Printf("%v -> %v\n", comb, set)
				//fmt.Printf("(%d / %d) (%v) -> %vcnt: %d\n", rem, maxRem, comb, set, cnt)
			}
			//fl.WriteString(dumpIntSlice(comb) + fmt.Sprintf(" => %t\n", ok))
			hits := 0
			for i, nv := range notValids3 {
				if ints.Contains(comb, nv) {
					if !ok {
						notVals3Counts[i]++
						hits++
					}
				}
			}
			if hits > 0 {
				notValAny++
				notValHits[hits-1]++
				if hits == 3 {
					//fmt.Printf("hit-3: %v\n", comb)
				}
			}
		}
	}
	// fmt.Printf("not-valids3: %v\n", notVals3Counts)
	// fmt.Printf("not-valids3-hits: %v\n", notValHits)
	// fmt.Printf("not-valids3-any-hits: %d\n", notValAny)
	// fmt.Printf("bcount - hitany: %d\n", cntBF-notValAny)
	// fmt.Printf("combis: total=%d, valid=%d\n", totCombis, cnt)

	return cnt
}

func cruxOne(n, num int, size int) int {
	fac := binCoeff(num, n)
	sign := 1
	sum := 0
	bs := (num - n + 1) * 3
	for k := 0; k <= num-n; k++ {
		c := binCoeff(num-n, k)
		p := sumBinCoeff(bs, 0)
		sum += sign * c * p
		sign *= -1
		bs -= 3
	}
	return fac * sum
}

func crux(num int, size int) int {
	fmt.Printf("crux: num=%d, size=%d\n", num, size)
	sum := 0
	for n := 1; n <= num; n++ {
		co := cruxOne(n, num, size)
		fmt.Printf("cx(%d) => %d\n", n, co)
		sum += co
	}
	return sum
}

func dumpIntSlice(vs []int) string {
	sl := make([]string, len(vs))
	for i, v := range vs {
		sl[i] = fmt.Sprintf("%02d", v)
	}
	return strings.Join(sl, " ")
}

func ArrangementsRec(in []int) int {
	sort.Ints(in)
	fmt.Printf("find arrangements for: %v\n", in)
	cnt := chains(true, 0, in[1:])
	return cnt
}

// assume in is already sorted
func chains(top bool, prev int, in []int) int {
	//fmt.Printf("chains (prev=%02d): %v\n", prev, in)
	if len(in) < 2 {
		if matches(in[0], prev) {
			//fmt.Printf("chain found\n")
			return 1
		}
		return 0
	}
	var c int
	i := 0
	for {
		v := in[i]
		if !matches(v, prev) {
			break
		}
		sc := chains(false, v, in[i+1:])
		c += sc
		//if c%10 == 0 {
		if top {
			fmt.Printf("found %d\n", c)
		}
		//}
		i++
	}
	return c
}

func ArrangementsEx(in []int) int {
	sort.Ints(in)
	if _, ok := validChain(in); !ok {
		return 0
	}
	fmt.Printf("find arrangements for (%d): %v\n", len(in), in)

	remOne := func(in *[]int) (int, bool) {
		if len(*in) < 3 {
			return -1, false
		}
		for i := 1; i < len(*in)-1; i++ {
			if (*in)[i+1]-(*in)[i-1] <= 3 {
				rem := (*in)[i]
				*in = append((*in)[:i], (*in)[i+1:]...)
				return rem, true
			}
		}
		return -1, false
	}

	var removed []int
	var rem int
	var didRem bool
	for {
		rem, didRem = remOne(&in)
		if !didRem {
			break
		}
		removed = append(removed, rem)
	}

	fmt.Printf("min-chain : %v\n", in)
	fmt.Printf("removed (%d): %v\n", len(removed), removed)

	cnt := 0

	n := len(removed)
	for k := 0; k <= n; k++ {
		bc := binCoeff(n, k)
		//fmt.Printf("add (%d, %d) -> %d\n", n, k, bc)
		cnt += bc
	}

	return cnt
}

func sumBinCoeff(n int, startK int) int {
	sum := 0
	for k := startK; k <= n; k++ {
		bc := binCoeff(n, k)
		sum += bc
	}
	return sum
}

func binCoeff(n, k int) int {
	f := float64(fac(n)) / (float64(fac(k)) * float64(fac(n-k)))
	return int(f)
}

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

// ////
type OneGroup struct {
	values []int
}

func (g *OneGroup) Values() []int {
	return g.values
}

func (g *OneGroup) CanAdd(n int) bool {
	if len(g.values) == 0 {
		return true
	}
	if n-g.values[len(g.values)-1] == 1 {
		return true
	}
	return false
}

func (g *OneGroup) Add(n int) {
	g.values = append(g.values, n)
}

func (g *OneGroup) Permutations() int64 {
	if len(g.values) <= 2 {
		return 1
	}
	return int64(Arrangements(g.values))
}

func Find1Groups(in []int) []*OneGroup {
	if len(in) == 0 {
		return []*OneGroup{}
	}
	var ogs []*OneGroup
	curr := &OneGroup{}
	for _, n := range in {
		if curr.CanAdd(n) {
			curr.Add(n)
			continue
		}
		if len(curr.values) > 0 {
			ogs = append(ogs, curr)
			curr = &OneGroup{}
		}
		curr.Add(n)
	}
	if len(curr.values) > 0 {
		ogs = append(ogs, curr)
	}
	return ogs
}
