package intutil

func Min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}
	return n2
}

func Max(n1, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}
