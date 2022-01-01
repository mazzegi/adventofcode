package shuttle

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type RawInput struct {
	EarliestTime string
	Shuttles     string
}

type Input struct {
	EarliestTime int
	Shuttles     []int
}

func ParseInput(ri RawInput) (Input, error) {
	et, err := strconv.ParseInt(ri.EarliestTime, 10, 64)
	if err != nil {
		return Input{}, errors.Wrap(err, "parse-earliest-time")
	}
	in := Input{
		EarliestTime: int(et),
	}
	sl := strings.Split(ri.Shuttles, ",")
	for _, s := range sl {
		if s == "x" {
			continue
		}
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return Input{}, errors.Wrap(err, "parse-shuttle-id")
		}
		in.Shuttles = append(in.Shuttles, int(id))
	}
	return in, nil
}

func ParseBusses(ri RawInput) ([]Bus, error) {
	var bs []Bus
	sl := strings.Split(ri.Shuttles, ",")
	offset := 0
	for _, s := range sl {
		if s == "x" {
			offset++
			continue
		}
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "parse-shuttle-id")
		}
		bs = append(bs, Bus{
			ID:     int(id),
			Offset: offset,
		})
		offset++
	}
	return bs, nil
}

func (i Input) String() string {
	return fmt.Sprintf("earliest=%d, shuttles=%v", i.EarliestTime, i.Shuttles)
}

type Result struct {
	BusID    int
	DepartAt int
}

func (in Input) FindEarliestDeparture() (Result, error) {
	t := in.EarliestTime
	max := 100000
	for i := 0; i < max; i++ {
		for _, st := range in.Shuttles {
			if t%st == 0 {
				return Result{
					BusID:    st,
					DepartAt: t - in.EarliestTime,
				}, nil
			}
		}
		t++
	}
	return Result{}, errors.Errorf("reached-max (%d)", max)
}

//
type Bus struct {
	ID     int
	Offset int
}

func FindEarliestSubsequent(bs []Bus) (int, error) {
	for _, b := range bs {
		fmt.Printf("(%d, %d) ", b.ID, b.Offset)
	}
	fmt.Printf("\n")

	max := 2 * int(1e9)
	t := 0
outer:
	for i := 0; i < max; i++ {
		for _, b := range bs {
			if (t+b.Offset)%b.ID != 0 {
				t++
				continue outer
			}
		}
		return t, nil
	}
	return -1, errors.Errorf("reached-max (%d)", max)
}

func FindEarliestSubsequentIntelli(bs []Bus) (int, error) {
	// for _, b := range bs {
	// 	fmt.Printf("(%d, %d) ", b.ID, b.Offset)
	// }
	// fmt.Printf("\n")

	max := 2 * int(1e9)

	//start := make([]int, len(bs))
	t := 0
	step := 0
	hits := make([]int, len(bs))
	for i := 0; i < len(hits); i++ {
		hits[i] = 0
	}
	//hits[0] = step
outer:
	for i := 0; i < max; i++ {

		for k, b := range bs {
			if (t+b.Offset)%b.ID != 0 {
				if step > 0 {
					t += step
				} else {
					t++
				}
				continue outer
			} else if t > 0 {
				if hits[k] == 0 {
					hits[k] = t
					fmt.Printf("%d: assign hit %d\n", k, t)
				} else if hits[k] > 0 {
					diff := t - hits[k]
					step = diff
					hits[k] = -1
					fmt.Printf("%d (%d): increase step by %d => %d\n", k, t, diff, step)
				}
			}
		}
		return t, nil
	}
	return -1, errors.Errorf("reached-max (%d)", max)
}

func PrintSubsequentIntelli(bs []Bus, maxCnt int) {

	max := 2 * int(1e9)

	t := 0
	step := bs[0].ID
	cnt := 0
outer:
	for i := 0; i < max; i++ {
		for _, b := range bs {
			if (t+b.Offset)%b.ID != 0 {
				t += step
				continue outer
			}
		}
		fmt.Printf("%d\n", t)
		cnt++
		if cnt > maxCnt {
			return
		}
		t += step
	}
}
