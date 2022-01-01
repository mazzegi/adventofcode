package ints

func Diffs(in []int) []int {
	if len(in) < 2 {
		return []int{}
	}

	diffs := make([]int, len(in)-1)
	for i := 1; i < len(in); i++ {
		diffs[i-1] = in[i] - in[i-1]
	}
	return diffs
}

func Hist(in []int) map[int]int {
	h := map[int]int{}
	for _, n := range in {
		h[n]++
	}
	return h
}

func Max(in []int) int {
	if len(in) == 0 {
		return 0
	}
	m := in[0]
	for _, n := range in {
		if n > m {
			m = n
		}
	}
	return m
}

func Min(in []int) int {
	if len(in) == 0 {
		return 0
	}
	m := in[0]
	for _, n := range in {
		if n < m {
			m = n
		}
	}
	return m
}

func HasDuplicate(in []int) bool {
	m := map[int]int{}
	for _, n := range in {
		m[n]++
		if m[n] >= 2 {
			return true
		}
	}
	return false
}

func Contains(in []int, other []int) bool {
	containOne := func(n int) bool {
		for _, v := range in {
			if v == n {
				return true
			}
		}
		return false
	}
	for _, v := range other {
		if !containOne(v) {
			return false
		}
	}
	return true
}
