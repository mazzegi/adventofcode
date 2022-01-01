package dinner

func Permutations(sl []string) <-chan []string {
	c := make(chan []string)
	go func() {
		defer close(c)
		if len(sl) == 0 {
			return
		}
		if len(sl) == 1 {
			c <- sl
			return
		}
		s0 := sl[0]
		for sub := range Permutations(sl[1:]) {
			var subPerms [][]string
			perm := append([]string{s0}, sub...)
			subPerms = append(subPerms, perm)
			for i := 0; i < len(sub); i++ {
				perm := append([]string{}, sub[:i+1]...)
				perm = append(perm, s0)
				perm = append(perm, sub[i+1:]...)
				subPerms = append(subPerms, perm)
			}
			for _, sp := range subPerms {
				c <- sp
			}
		}

	}()
	return c
}
