package comb

/*
1 2 3 4
1 2
1 3
1 4
2 3
2 4
3 4
*/

func Parts(in []int, size int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		if size < 0 || size > len(in) {
			return
		}
		if size == 0 {
			c <- []int{}
			return
		}
		if size == len(in) {
			c <- in
			return
		}

		for i := 0; i < len(in); i++ {
			inSub := append([]int{}, in[:i]...)
			for subPart := range Parts(inSub, size-1) {
				part := []int{in[i]}
				part = append(part, subPart...)
				c <- part
			}
		}
	}()
	return c
}
