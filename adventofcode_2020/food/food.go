package food

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

//mxmxvkd kfcds sqjhc nhms (contains dairy, fish)

var contains = "(contains"

func Parse(s string) (Food, error) {
	cix := strings.Index(s, contains)
	if cix < 0 {
		return Food{}, errors.Errorf("found no allergens (intended?)")
	}

	slincs := strings.Split(s[:cix], " ")
	var incs []string
	for _, inc := range slincs {
		inc = strings.Trim(inc, " ")
		if inc == "" {
			continue
		}
		incs = append(incs, inc)
	}

	slalgs := strings.Split(s[cix+len(contains):len(s)-1], ", ")
	var algs []string
	for _, alg := range slalgs {
		alg = strings.Trim(alg, " ")
		if alg == "" {
			continue
		}
		algs = append(algs, alg)
	}

	f := Food{
		Ingredients: incs,
		Allergens:   algs,
	}
	return f, nil
}

func intersect(sl1, sl2 []string) []string {
	in2 := func(s string) bool {
		for _, s2 := range sl2 {
			if s2 == s {
				return true
			}
		}
		return false
	}

	var isl []string
	for _, s1 := range sl1 {
		if in2(s1) {
			isl = append(isl, s1)
		}
	}
	return isl
}

func removeAll(sl []string, rem string) ([]string, int) {
	var rsl []string
	var remCnt int
	for _, s := range sl {
		if s != rem {
			rsl = append(rsl, s)
		} else {
			remCnt++
		}
	}
	return rsl, remCnt
}

func CheckAllergens(fs []Food) {
	algMapping := map[string][]string{}
	for _, f := range fs {
		for _, alg := range f.Allergens {
			if sl, ok := algMapping[alg]; ok {
				algMapping[alg] = intersect(sl, f.Ingredients)
			} else {
				algMapping[alg] = f.Ingredients
			}
		}
	}

	uniqueAlgMapping := map[string]string{}
	for alg, incs := range algMapping {
		fmt.Printf("%s => %v\n", alg, incs)
		if len(incs) == 1 {
			uniqueAlgMapping[alg] = incs[0]
			delete(algMapping, alg)
		}
	}

	//try to make unique
	fmt.Printf("*** make unique ***\n")
	changed := true
	for changed {
		changed = false
		for _, uinc := range uniqueAlgMapping {
			for alg, incs := range algMapping {
				if rincs, cnt := removeAll(incs, uinc); cnt > 0 {
					if len(rincs) == 1 {
						uniqueAlgMapping[alg] = rincs[0]
						delete(algMapping, alg)
					} else {
						algMapping[alg] = rincs
					}
					changed = true
				}
			}
		}
	}

	type mappedInc struct {
		inc string
		alg string
	}

	var mappedIncsSlice []mappedInc
	mappedIncs := map[string]string{}
	for alg, inc := range uniqueAlgMapping {
		fmt.Printf("%s => %s\n", alg, inc)
		mappedIncs[inc] = alg
		mappedIncsSlice = append(mappedIncsSlice, mappedInc{
			inc: inc,
			alg: alg,
		})
	}

	//
	var nCnt int
	for _, f := range fs {
		for _, inc := range f.Ingredients {
			if _, ok := mappedIncs[inc]; !ok {
				nCnt++
			}
		}
	}
	fmt.Printf("the n-count is: %d\n", nCnt)

	fmt.Printf("*** Part2 ***\n")
	sort.Slice(mappedIncsSlice, func(i, j int) bool {
		return mappedIncsSlice[i].alg < mappedIncsSlice[j].alg
	})
	var canList []string
	for _, mi := range mappedIncsSlice {
		canList = append(canList, mi.inc)
	}
	fmt.Printf("canonical: %q\n", strings.Join(canList, ","))
}
