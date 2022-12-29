package day_25

import (
	"math"

	"github.com/mazzegi/adventofcode/slices"
)

func powN(b int, e int) int {
	if e < 0 {
		return 0
	}
	agg := 1
	for i := 0; i < e; i++ {
		agg *= b
	}
	return agg
}

var digits = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

func dig(n int) rune {
	switch {
	case n == 0:
		return '0'
	case n == 1:
		return '1'
	case n == -1:
		return '-'
	case n == 2:
		return '2'
	case n == -2:
		return '='
	default:
		fatal("cannot dig %d ", n)
	}
	return ' '
}

func DecodeSNAFU(s string) int {
	var sum int
	rs := slices.Reverse([]rune(s))
	for i, r := range rs {
		place := powN(5, i)
		fac, ok := digits[r]
		if !ok {
			fatal("invalid digit %s", string(r))
		}
		sum += fac * place
	}
	return sum
}

var ln5 = math.Log(5)

func pow5(n int) int {
	return powN(5, n)
}

type place struct {
	exp int
	fac rune
}

func invSNAFU(s string) string {
	var si string
	for _, r := range s {
		n := digits[r]
		var ri rune
		switch n {
		case 0:
			ri = '0'
		case 1:
			ri = '-'
		case 2:
			ri = '='
		case -1:
			ri = '1'
		case -2:
			ri = '2'
		}
		si += string(ri)
	}
	return si
}

func EncodeSNAFU(ov int) string {
	//encode in 5 system
	v := ov
	var enc5 []int
	n := int(math.Floor(float64(math.Log(float64(ov))) / ln5))
	for i := n; i >= 0; i-- {
		fac := int(math.Floor(float64(v) / float64(pow5(i))))
		enc5 = append(enc5, fac)
		v -= fac * pow5(i)
	}
	enc5 = slices.Reverse(enc5)

	//now adjust
	var encS []int
	overflow := false
	for _, e5 := range enc5 {
		if overflow {
			e5 = e5 + 1
			overflow = false
		}
		if e5 <= 2 {
			encS = append(encS, e5)
			continue
		}
		pullDown := e5 - 5
		encS = append(encS, pullDown)

		//this overflows the next by 1
		overflow = true
	}
	if overflow {
		encS = append(encS, 1)
	}

	//stringify
	encS = slices.Reverse(encS)
	var s string
	for _, es := range encS {
		s += string(dig(es))
	}

	return s
}

// func EncodeSNAFU(ov int) string {
// 	var pls []place
// 	v := ov
// 	for {
// 		neg := false
// 		tt := v
// 		if v < 0 {
// 			tt = -v
// 			neg = true
// 		}

// 		n := int(math.Floor(float64(math.Log(float64(tt))) / ln5))
// 		p5n := pow5(n)
// 		if tt == p5n {
// 			pls = append(pls, place{n, dig(1, neg)})
// 			break
// 		}
// 		if tt == 2*p5n {
// 			pls = append(pls, place{n, dig(2, neg)})
// 			break
// 		}
// 		if 2*p5n < tt {
// 			// check if we compensate this with the next one
// 			left := tt - 2*p5n
// 			if 2*pow5(n-1) >= left {
// 				//yes
// 				pls = append(pls, place{n, dig(2, neg)})
// 				tt -= 2 * p5n
// 			} else {
// 				//no
// 				pls = append(pls, place{n + 1, dig(1, neg)})
// 				tt -= pow5(n + 1)
// 				pls = append(pls, place{n, dig(2, !neg)})
// 				tt += 2 * p5n
// 			}
// 		} else {
// 			left := tt - p5n
// 			if 2*pow5(n-1) >= left {
// 				pls = append(pls, place{n, dig(1, neg)})
// 				tt -= p5n
// 			} else {
// 				pls = append(pls, place{n, dig(2, neg)})
// 				tt -= 2 * p5n
// 			}
// 		}
// 		if !neg {
// 			v = tt
// 		} else {
// 			v = -tt
// 		}
// 		if v == 0 {
// 			break
// 		}
// 	}

// 	var senc string
// 	for i, p := range pls {
// 		senc += string(p.fac)
// 		if i > 0 {
// 			for e := p.exp; e < pls[i-1].exp-1; e++ {
// 				senc += "0"
// 			}
// 		}
// 	}

// 	check := DecodeSNAFU(senc)
// 	_ = check

// 	return senc
// }
