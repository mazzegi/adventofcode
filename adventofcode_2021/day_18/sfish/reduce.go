package sfish

import "math"

func Reduce(p *Pair) {
	action := true
	for action {
		if action, _, _, _ = explode(p, 0); action {
			continue
		}
		if action = split(p); action {
			continue
		}
		action = false
	}
}

/*
	in:   "[[[[[9,8],1],2],3],4]",
	out:  "[[[[0,9],2],3],4]",

	in:   "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
	out:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
*/
func explode(p *Pair, level int) (act bool, apply bool, left, right Regular) {
	if level >= 4 {
		//explode candidate?
		r1, ok1 := p.Left.(Regular)
		r2, ok2 := p.Right.(Regular)
		if ok1 && ok2 {
			return true, true, r1, r2
		}
	}

	switch l := p.Left.(type) {
	case *Pair:
		if act, apl, left, right := explode(l, level+1); act {
			if apl {
				p.Left = Regular(0)
			}
			//addLeft(p.Right, right)
			switch rr := p.Right.(type) {
			case Regular:
				p.Right = Regular(rr + right)
				right = 0
			case *Pair:
				added := addLeft(rr, right)
				if added {
					right = 0
				}
			}

			// if rr, ok := p.Right.(Regular); ok {
			// 	p.Right = Regular(rr + right)
			// 	right = 0
			// }
			return true, false, left, right
		}
	}
	switch r := p.Right.(type) {
	case *Pair:
		if act, apl, left, right := explode(r, level+1); act {
			if apl {
				p.Right = Regular(0)
			}

			switch rr := p.Left.(type) {
			case Regular:
				p.Left = Regular(rr + left)
				left = 0
			case *Pair:
				added := addRight(rr, left)
				if added {
					left = 0
				}
			}

			// if lr, ok := p.Left.(Regular); ok {
			// 	p.Left = Regular(lr + left)
			// 	left = 0
			// }
			return true, false, left, right
		}
	}

	return false, false, 0, 0
}

func addLeft(p *Pair, add Regular) bool {
	switch rr := p.Left.(type) {
	case Regular:
		p.Left = Regular(rr + add)
		return true
	case *Pair:
		return addLeft(rr, add)
	default:
		panic("örgs")
	}
}

func addRight(p *Pair, add Regular) bool {
	switch rr := p.Right.(type) {
	case Regular:
		p.Right = Regular(rr + add)
		return true
	case *Pair:
		return addRight(rr, add)
	default:
		panic("örgs")
	}
}

func split(p *Pair) bool {
	switch l := p.Left.(type) {
	case *Pair:
		if act := split(l); act {
			return true
		}
	case Regular:
		if l >= 10 {
			ip := &Pair{
				Left:  Regular(math.Floor(float64(l) / 2.0)),
				Right: Regular(math.Ceil(float64(l) / 2.0)),
			}
			p.Left = ip
			return true
		}
	}

	switch r := p.Right.(type) {
	case *Pair:
		if act := split(r); act {
			return true
		}
	case Regular:
		if r >= 10 {
			ip := &Pair{
				Left:  Regular(math.Floor(float64(r) / 2.0)),
				Right: Regular(math.Ceil(float64(r) / 2.0)),
			}
			p.Right = ip
			return true
		}
	}

	return false
}
