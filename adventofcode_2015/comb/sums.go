package comb

func FixedSumInts(sum int, len int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		if len <= 0 {
			return
		}
		if len == 1 {
			c <- []int{sum}
			return
		}
		for n1 := 0; n1 <= sum; n1++ {
			for sub := range FixedSumInts(sum-n1, len-1) {
				c <- append([]int{n1}, sub...)
			}
		}
	}()
	return c
}
