package coord

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Point struct {
	id   string
	x, y int
}

func (p Point) Distance(op Point) int {
	return absInt(p.x-op.x) + absInt(p.y-op.y)
}

func ParsePoint(id string, s string) (Point, error) {
	p := Point{
		id: id,
	}
	_, err := fmt.Sscanf(s, "%d, %d", &p.x, &p.y)
	return p, err
}

type Grid struct {
	points   []Point
	min, max Point
}

func ParseGrid(r io.Reader) (*Grid, error) {
	g := &Grid{}
	scanner := bufio.NewScanner(r)
	id := 0
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		p, err := ParsePoint(fmt.Sprintf("%02d", id), l)
		if err != nil {
			return nil, err
		}
		g.points = append(g.points, p)
		id++
	}
	if len(g.points) < 2 {
		return nil, errors.Errorf("too less points")
	}
	g.min, g.max = g.points[0], g.points[0]
	for i := 1; i < len(g.points); i++ {
		p := g.points[i]
		if p.x < g.min.x {
			g.min.x = p.x
		}
		if p.x > g.max.x {
			g.max.x = p.x
		}
		if p.y < g.min.y {
			g.min.y = p.y
		}
		if p.y > g.max.y {
			g.max.y = p.y
		}
	}
	w, h := g.max.x-g.min.x, g.max.y-g.min.y
	g.min.x -= w
	g.max.x += w
	g.min.y -= h
	g.max.y += h

	return g, nil
}

type DPoint struct {
	Point
	Closest string
}

type DPoints map[Point]DPoint

func (g *Grid) CalcDistances() {
	edgePoints := map[string]bool{}
	areas := map[string]int{}
	dpoints := DPoints{}
	for x := g.min.x; x <= g.max.x; x++ {
		for y := g.min.y; y <= g.max.y; y++ {
			dp := Point{
				x: x,
				y: y,
			}
			closestPt := g.points[0]
			closestDist := closestPt.Distance(dp)
			hitMin := 1
			for i := 1; i < len(g.points); i++ {
				pt := g.points[i]
				dist := pt.Distance(dp)
				if dist < closestDist {
					closestPt = pt
					closestDist = dist
					hitMin = 1
				} else if dist == closestDist {
					hitMin++
				}
			}
			if hitMin > 1 {
				dpoints[dp] = DPoint{
					Point:   dp,
					Closest: ".",
				}
			} else {
				dpoints[dp] = DPoint{
					Point:   dp,
					Closest: closestPt.id,
				}
				areas[closestPt.id]++
			}

			if x == g.min.x || x == g.max.x || y == g.min.y || y == g.max.y {
				edgePoints[closestPt.id] = true
			}
		}
	}
	var biggest string
	var biggestSize int
	for area, cnt := range areas {
		if _, contains := edgePoints[area]; contains {
			fmt.Printf("%q => edge\n", area)
		} else {
			fmt.Printf("%q => %d\n", area, cnt)
			if cnt > biggestSize {
				biggest = area
				biggestSize = cnt
			}
		}
	}
	fmt.Printf("the biggest us %q with size %d\n", biggest, biggestSize)
}

func (g *Grid) CalcSafest(safeDist int) {
	dpoints := DPoints{}
	for x := g.min.x; x <= g.max.x; x++ {
		for y := g.min.y; y <= g.max.y; y++ {
			dp := Point{
				x: x,
				y: y,
			}

			sum := 0
			for i := 0; i < len(g.points); i++ {
				pt := g.points[i]
				dist := pt.Distance(dp)
				sum += dist
			}
			if sum < safeDist {
				dpoints[dp] = DPoint{
					Point:   dp,
					Closest: "#",
				}
			}
		}
	}
	fmt.Printf("the safe-area size if %d\n", len(dpoints))
}
