package nano

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func ParseBot(s string, id int) (*Bot, error) {
	b := Bot{
		ID: id,
	}
	_, err := fmt.Sscanf(s, "pos=<%d,%d,%d>, r=%d", &b.Position.X, &b.Position.Y, &b.Position.Z, &b.Radius)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

type Vector struct {
	X, Y, Z int
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d, %d, %d)", v.X, v.Y, v.Z)
}

func (v Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

func (v Vector) Len() int {
	return v.Distance(Vector{})
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func unit(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}

func (v Vector) Dir(toV Vector) Vector {
	return Vector{
		X: unit(toV.X - v.X),
		Y: unit(toV.Y - v.Y),
		Z: unit(toV.Z - v.Z),
	}
}

func (v Vector) Sub(ov Vector) Vector {
	return Vector{
		X: ov.X - v.X,
		Y: ov.Y - v.Y,
		Z: ov.Z - v.Z,
	}
}

func (v Vector) Add(ov Vector) Vector {
	return Vector{
		X: ov.X + v.X,
		Y: ov.Y + v.Y,
		Z: ov.Z + v.Z,
	}
}

func (v Vector) Mult(n int) Vector {
	return Vector{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func (v Vector) Distance(ov Vector) int {
	return absInt(v.X-ov.X) + absInt(v.Y-ov.Y) + absInt(v.Z-ov.Z)
}

type Bot struct {
	ID       int
	Position Vector
	Radius   int
}

func (b *Bot) InRange(pos Vector) bool {
	return b.Position.Distance(pos) <= b.Radius
}

type Intersection struct {
	Bots []*Bot
}

func (is *Intersection) Smallest() {
	smallest := -1
	smallestDirLen := -1
	minDirLen := -1
	var sB *Bot
	var soB *Bot
	for _, b := range is.Bots {
		for _, ob := range is.Bots {
			if b.ID == ob.ID {
				continue
			}

			// is == b.Position.Distance(ob.Position) <= b.Radius+ob.Radius
			isDist := b.Radius + ob.Radius - b.Position.Distance(ob.Position)
			if isDist < 0 {
				panic("no intersection")
			}
			dir := b.Position.Dir(ob.Position)
			if minDirLen < 0 || dir.Len() < minDirLen {
				minDirLen = dir.Len()
			}

			if smallest < 0 {
				smallest = isDist
				sB = b
				soB = ob
				smallestDirLen = dir.Len()
			} else if isDist < smallest {
				smallest = isDist
				sB = b
				soB = ob
				smallestDirLen = dir.Len()
			} else if isDist == smallest {
				if dir.Len() < smallestDirLen {
					smallest = isDist
					sB = b
					soB = ob
					smallestDirLen = dir.Len()
				}
			}
		}
	}
	fmt.Printf("smallest: %d\n", smallest)
	fmt.Printf("min-dir-len: %d\n", minDirLen)
	fmt.Printf("bot-1 (%d) %s with radius %d\n", sB.ID, sB.Position, sB.Radius)
	fmt.Printf("bot-2 (%d) %s with radius %d\n", soB.ID, soB.Position, soB.Radius)
	fmt.Printf("bots distance: %d (added radius %d)\n", sB.Position.Distance(soB.Position), sB.Radius+soB.Radius)

	if smallest > 0 {
		return
	}

	dir := sB.Position.Dir(soB.Position)
	fmt.Printf("direction 1-2: %s\n", dir)

	ip := sB.Position
	for sB.Position.Distance(ip) < sB.Radius {
		dir = ip.Dir(soB.Position)
		ip = ip.Add(dir)
	}
	fmt.Printf("intersection-point: %s\n", ip)

	////
	collected := []*Bot{sB, soB}
	validPosition := func(p Vector) bool {
		for _, cb := range collected {
			if !cb.InRange(p) {
				return false
			}
		}
		return true
	}

	notInRange := func() []*Bot {
		var nir []*Bot
		for _, b := range is.Bots {
			if !b.InRange(ip) {
				nir = append(nir, b)
			}
		}
		return nir
	}

	dirMutations := func(dir Vector) []Vector {
		var ms []Vector
		if dir.Len() == 0 || dir.Len() == 1 {
			//impossible
			return ms
		}
		if dir.Len() == 2 {
			return []Vector{dir}
		}
		ms = append(ms, dir)

		mdir := dir
		mdir.X = 0
		ms = append(ms, mdir)

		mdir = dir
		mdir.Y = 0
		ms = append(ms, mdir)

		mdir = dir
		mdir.Z = 0
		ms = append(ms, mdir)

		return ms
	}

	nir := notInRange()
	for len(nir) > 0 {
		first := nir[0]
		fmt.Printf("integrate %d\n", first.ID)
		for !first.InRange(ip) {
			dir := ip.Dir(first.Position)
			mdirs := dirMutations(dir)
			found := false
			for _, mdir := range mdirs {
				nextIp := ip.Add(mdir)
				if validPosition(nextIp) {
					ip = nextIp
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("no step further ...\n")
				return
			}
		}

		var nowNir []*Bot
		for _, b := range nir {
			if b.InRange(ip) {
				fmt.Printf("integrated %d\n", b.ID)
				collected = append(collected, b)
			} else {
				nowNir = append(nowNir, b)
			}
		}
		fmt.Printf("left %d not in range\n", len(nowNir))

		nir = nowNir
	}

	////

	// //sdir := dir.Mult(sB.Radius)
	// sdir := dir
	// for sdir.Len() < sB.Radius {

	// 	sdir = sdir.Add(dir)
	// }

	// fmt.Printf("sub-vec 1-2: %s\n", sdir)
	//ip := sB.Position.Add(sdir)
	if !sB.InRange(ip) {
		fmt.Printf("oops: ip is not in range of b1\n")
		return
	}
	if !soB.InRange(ip) {
		fmt.Printf("oops: ip is not in range of b1\n")
		return
	}

	//check if ip is really in all bots range
	// var notInRange []int
	// for _, b := range is.Bots {
	// 	if !b.InRange(ip) {
	// 		notInRange = append(notInRange, b.ID)
	// 		//fmt.Printf("opps - %s is not in bots %d (r=%d) range (dist %d)\n", ip, b.ID, b.Radius, b.Position.Distance(ip))
	// 		//return
	// 	}
	// }
	// fmt.Printf("not-in-range (%d / %d): %v\n", len(notInRange), len(is.Bots), notInRange)

	distFromZero := Vector{X: 0, Y: 0, Z: 0}.Distance(ip)
	fmt.Printf("its dist from zero: %d\n", distFromZero)

	// iv := sB.FirstIntersectionPoint(soB)
	// fmt.Printf("vec: %s\n", iv)
	// distFromZero := Vector{X: 0, Y: 0, Z: 0}.Distance(iv)
	// fmt.Printf("its dist from zero: %d\n", distFromZero)
}

func (i *Intersection) FirstPoint() Vector {
	if len(i.Bots) == 0 {
		return Vector{X: 0, Y: 0, Z: 0}
	} else if len(i.Bots) == 1 {
		return i.Bots[0].Position
	}
	return i.Bots[0].FirstIntersectionPoint(i.Bots[1])
}

func (i *Intersection) Intersects(ob *Bot) bool {
	for _, b := range i.Bots {
		if !b.Intersects(ob) {
			return false
		}
	}
	return true
}

func (i *Intersection) Hash() string {
	var ns []int
	for _, b := range i.Bots {
		ns = append(ns, b.ID)
	}
	sort.Ints(ns)
	var s string
	for _, n := range ns {
		s += fmt.Sprintf("%04d", n)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//

func (b *Bot) Intersects(ob *Bot) bool {
	return b.Position.Distance(ob.Position) <= b.Radius+ob.Radius
}

func (b *Bot) FirstIntersectionPoint(ob *Bot) Vector {
	inRangeOfBoth := func(v Vector) bool {
		return b.InRange(v) && ob.InRange(v)
	}

	v := b.Position
	for !inRangeOfBoth(v) {
		dir := b.Position.Dir(ob.Position)
		if dir.IsZero() {
			panic("dir is zero")
		}
		v = v.Add(dir)
	}

	return v
}

func (b *Bot) Intersect(obots ...*Bot) *Intersection {
	is := &Intersection{
		Bots: []*Bot{b},
	}
	for _, ob := range obots {
		if ob.ID == b.ID {
			continue
		}
		if is.Intersects(ob) {
			is.Bots = append(is.Bots, ob)
		}
	}
	return is
}
