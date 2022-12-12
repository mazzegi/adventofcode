package grid2d

func E(p1, p2 Pointf) Edge {
	return Edge{p1, p2}
}

type Edge struct {
	P1 Pointf
	P2 Pointf
}

func (e Edge) Length() float64 {
	return e.P1.DistTo(e.P2)
}
