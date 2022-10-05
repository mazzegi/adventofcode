package intutil

func Between(n int, low, high int) bool {
	if low > high {
		high, low = low, high
	}
	return n >= low && n <= high
}
