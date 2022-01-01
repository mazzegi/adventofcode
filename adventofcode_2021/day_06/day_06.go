package day_06

import (
	"adventofcode_2021/errutil"
	"adventofcode_2021/readutil"
	"fmt"
	"math"

	"github.com/pkg/errors"
)

func Part1() {
	// res, err := FishCount(input, 80)
	// errutil.ExitOnErr(err)
	// fmt.Printf("part1: result = %d\n", res)
}

func PartTest() {
	//res, err := FishCountSharded(inputTest, 256) //26984457539
	//res, err := FishCountSharded("1", 256) //26984457539
	//res, err := FishCountSharded(inputTest, 80) //5934
	//res, err := FishCountSharded(inputTest, 18) //26
	//res, err := FishCountSharded(inputTest, 5) // => 10
	res, err := FishCountLookup256(input)
	errutil.ExitOnErr(err)
	fmt.Printf("part-test: result = %d\n", res)
	//calcSingle()
}

func calcSingle() {
	for i := 1; i <= 5; i++ {
		res, err := FishCountSharded(fmt.Sprintf("%d", i), 256)
		if err != nil {
			errutil.ExitOnErr(err)
		}
		fmt.Printf("%d: %d\n", i, res)
	}
}

func Part2() {
	PartTest()
	// res, err := FishCountSharded(input, 256) //26984457539
	// errutil.ExitOnErr(err)
	// fmt.Printf("part-test: result = %d\n", res)
}

func FishCount(in string, afterDays int) (int, error) {
	ns, err := readutil.ReadInts(in, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	var uns []uint8
	for _, n := range ns {
		uns = append(uns, uint8(n))
	}

	for day := 0; day < afterDays; day++ {
		var appendCount int
		for i, n := range uns {
			if n > 0 {
				n -= 1
			} else {
				n = 6
				appendCount++
			}
			uns[i] = n
		}
		for i := 0; i < appendCount; i++ {
			uns = append(uns, 8)
		}
		fmt.Printf("after day %d: size = %d\n", day, len(uns))
	}

	return len(uns), nil
}

const divideAfter = 20 * 1024 * 1024

func FishCountSharded(in string, afterDays int) (int, error) {
	ns, err := readutil.ReadInts(in, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	var uns []uint8
	for _, n := range ns {
		uns = append(uns, uint8(n))
	}

	total := CountShard(0, 0, uns, afterDays, 0)

	return total, nil
}

func cloneUint8s(uns []uint8) []uint8 {
	cuns := make([]uint8, len(uns))
	copy(cuns, uns)
	return cuns
}

func CountShard(rec int, sid int, uns []uint8, afterDays int, dayOffset int) int {
	for day := 0; day < afterDays; day++ {
		var appendCount int
		for i, n := range uns {
			if n > 0 {
				n -= 1
			} else {
				n = 6
				appendCount++
			}
			uns[i] = n
		}
		for i := 0; i < appendCount; i++ {
			uns = append(uns, 8)
		}

		//fmt.Printf("shard %q: after day %d: size = %d\n", sid, day, len(uns))
		//
		if day < afterDays-1 && len(uns) >= divideAfter {

			split := len(uns) / 2
			uns1 := cloneUint8s(uns[:split])
			uns2 := cloneUint8s(uns[split:])
			r1 := CountShard(rec+1, 0, uns1, afterDays-day-1, dayOffset+day)
			r2 := CountShard(rec+1, 1, uns2, afterDays-day-1, dayOffset+day)
			//fmt.Printf("day-offset: %d, rec %d, shard %d, day %d: after split shard: %d + %d, %d, %d\n", dayOffset, rec, sid, day, r1, r2, len(uns1), len(uns2))

			return r1 + r2
		}

	}

	return len(uns)
}

func NumDescendants(gen int, afterDays int) int {
	daysConsidered := afterDays - gen
	return int(math.Pow(2, math.Floor(float64(daysConsidered)/7.0)))
}

func FishCountCalced(in string, afterDays int) (int, error) {
	ns, err := readutil.ReadInts(in, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	var total int
	for _, n := range ns {
		ds := NumDescendants(n, afterDays)
		fmt.Printf("%d => %d\n", n, ds)
		total += ds
	}

	return total, nil
}

var lookup256 = map[int]int{
	1: 6206821033,
	2: 5617089148,
	3: 5217223242,
	4: 4726100874,
	5: 4368232009,
}

func FishCountLookup256(in string) (int, error) {
	ns, err := readutil.ReadInts(in, ",")
	if err != nil {
		return 0, errors.Wrap(err, "read-ints")
	}
	var total int
	for _, n := range ns {
		ts, ok := lookup256[n]
		if !ok {
			return 0, errors.Errorf("%d is not in lookup", n)
		}
		total += ts
	}

	return total, nil
}
