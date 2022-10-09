package day_14

import (
	"testing"

	"github.com/mazzegi/adventofcode/testutil"
)

const inputTest1 = `
10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL
`

const inputTest2 = `
157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT
`

const inputTest3 = `
2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF
`

const inputTest4 = `
171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX
`

func TestPart1MainFunc(t *testing.T) {
	t.Run("test_01", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest1)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 31
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_02", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest2)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 13312
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_03", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest3)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 180697
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_04", func(t *testing.T) {
		//t.Skip()
		res, err := part1MainFunc(inputTest4)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 2210736
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
}

func TestPart2MainFunc(t *testing.T) {
	t.Run("test_02", func(t *testing.T) {
		t.Skip()
		res, err := part2MainFunc(inputTest2)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 82892753
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_03", func(t *testing.T) {
		t.Skip()
		res, err := part2MainFunc(inputTest3)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 5586022
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})

	t.Run("test_04", func(t *testing.T) {
		//t.Skip()
		res, err := part2MainFunc(inputTest4)
		testutil.CheckUnexpectedError(t, err)
		var exp int = 460664
		if exp != res {
			t.Fatalf("want %d, have %d", exp, res)
		}
	})
}
