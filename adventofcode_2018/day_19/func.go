package day_19

func programFunc(r0Init int) int {
	r0, r1, r2, r3, r4 := r0Init, 0, 0, 0, 0
	dump := func() {
		log("%d, %d, %d, %d, %d", r0, r1, r2, r3, r4)
	}
	_ = dump

	// r2 = (r2 + 2) * (r2 + 2) * 19 * 11
	// r4 = (r4+1)*22 + 19
	// r2 = r2 + r4
	// if r0 == 1 {
	// 	r4 = 10550400
	// 	r2 = r2 + r4
	// 	r0 = 0
	// }
	// r3 = 1

	r4 = 41
	r2 = 877 // 836 + 41 // 2 * 2 * 19 * 11 + 41
	//r2 = r2 + r4
	if r0 == 1 {
		r2 = 10551277
		r0 = 0

		// primes are 1, 11, 959207, 10551277
		return 1 + 10551277 + 11 + 959207
	} else if r0 == 0 {
		// primes are 1, 877
		return 1 + 877
	}
	r3 = 1

	for r3 = 1; r3 <= r2; r3++ {
		for r1 = 1; r1 <= r2; r1++ {
			if r1*r3 == r2 {
				r0 += r3
			}
		}
	}
	return r0

	// loop1_loop2:
	// 	for {
	// 		r1 = 1
	// 	loop1_loop3:
	// 		for {
	// 			r4 = r1 * r3
	// 			if r2 == r4 {
	// 				r0 = r0 + r3
	// 			}

	// 			r1 = r1 + 1
	// 			if r1 <= r2 {
	// 				continue loop1_loop3
	// 			}
	// 			r3 = r3 + 1
	// 			if r3 <= r2 {
	// 				continue loop1_loop2
	// 			}
	// 			return r0
	// 		}
	// 	}
	//return r0
}

// func programFunc(r0Init int) int {
// 	r0, r1, r2, r3, r4 := r0Init, 0, 0, 0, 0
// 	dump := func() {
// 		log("%d, %d, %d, %d, %d", r0, r1, r2, r3, r4)
// 	}
// 	_ = dump

// 	r2 = (r2 + 2) * (r2 + 2) * 19 * 11
// 	r4 = (r4+1)*22 + 19
// 	r2 = r2 + r4
// 	//r4 = (27*28 + 29) * 30 * 14 * 32
// 	if r0 == 1 {
// 		r4 = 10550400
// 		r2 = r2 + r4
// 		r0 = 0
// 	}

// 	r3 = 1
// 	//dump()
// loop1_loop2:
// 	for {
// 		r1 = 1 // l2
// 	loop1_loop3:
// 		for {
// 			r4 = r1 * r3 // l3
// 			//r4 = eq(r4, r2) // l4
// 			if r2 == r4 { // == 1
// 				r0 = r0 + r3 // l7
// 			}
// 			r1 = r1 + 1 // l8
// 			//r4 = gt(r1, r2) //l9
// 			if r1 <= r2 {
// 				continue loop1_loop3
// 			}
// 			r3 = r3 + 1 // l12
// 			if r3 <= r2 {
// 				continue loop1_loop2
// 			}
// 			return r0
// 		}
// 	}
// }
