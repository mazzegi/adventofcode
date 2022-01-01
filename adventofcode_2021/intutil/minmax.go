package intutil

func MinInt(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}
	return n2
}

func MaxInt(n1, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}
