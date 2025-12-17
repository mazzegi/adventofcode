package day_08

import (
	"fmt"
	"sort"
	"time"

	"github.com/mazzegi/adventofcode/errutil"
	"github.com/mazzegi/adventofcode/readutil"
	"github.com/mazzegi/adventofcode/vector"
)

func log(pattern string, args ...interface{}) {
	fmt.Printf(pattern+"\n", args...)
}

func Part1() {
	t0 := time.Now()
	res, err := part1MainFunc(input, 1000)
	errutil.ExitOnErr(err)
	log("part1: result = %d (%s)", res, time.Since(t0))
}

func Part2() {
	t0 := time.Now()
	res, err := part2MainFunc(input)
	errutil.ExitOnErr(err)
	log("part2: result = %d (%s)", res, time.Since(t0))
}

func parseV3D(s string) (vector.Vector3D, error) {
	var v vector.Vector3D
	_, err := fmt.Sscanf(s, "%f,%f,%f", &v.X, &v.Y, &v.Z)
	if err != nil {
		return vector.Vector3D{}, fmt.Errorf("scan: %w", err)
	}
	return v, nil
}

type JBox struct {
	ID          int
	Position    vector.Vector3D
	ConnectedTo []*JBox
}

type Cluster struct {
	JBoxes []*JBox
}

func (c *Cluster) Contains(jb *JBox) bool {
	for _, exjb := range c.JBoxes {
		if exjb.ID == jb.ID {
			return true
		}
	}
	return false
}

func (jb *JBox) isDirectlyConnectedTo(otherJB *JBox) bool {
	if jb.ID == otherJB.ID {
		return true
	}
	for _, conJB := range jb.ConnectedTo {
		if conJB.ID == otherJB.ID {
			return true
		}
	}
	return false
}

func formatV3D(v vector.Vector3D) string {
	return fmt.Sprintf("(%.0f, %.0f, %.0f)", v.X, v.Y, v.Z)
}

func part1MainFunc(in string, numPairs int) (int, error) {
	lines := readutil.ReadLines(in)
	var jBoxes []*JBox
	log("parse input ...")
	for i, line := range lines {
		v, err := parseV3D(line)
		if err != nil {
			return 0, fmt.Errorf("parse_v3d %q: %w", line, err)
		}
		jBoxes = append(jBoxes, &JBox{
			ID:       i,
			Position: v,
		})
	}

	makeShortestConnection := func() error {
		var (
			minDist        float64
			minJB1, minJB2 *JBox
		)

		for i1, jb1 := range jBoxes {
			_ = i1
			for i2, jb2 := range jBoxes {
				_ = i2
				if jb1.ID == jb2.ID {
					continue
				}
				if jb1.isDirectlyConnectedTo(jb2) {
					continue
				}
				dist := jb1.Position.DistTo(jb2.Position)
				if minJB1 == nil || dist < minDist {
					minDist = dist
					minJB1 = jb1
					minJB2 = jb2
				}
			}
		}
		if minJB1 == nil {
			// no connection where made
			return fmt.Errorf("no connection possible")
		}
		minJB1.ConnectedTo = append(minJB1.ConnectedTo, minJB2)
		minJB2.ConnectedTo = append(minJB2.ConnectedTo, minJB1)
		log("connected: %s <-> %s", formatV3D(minJB1.Position), formatV3D(minJB2.Position))
		return nil
	}

	log("make %d connections", numPairs)
	for i := range numPairs {
		err := makeShortestConnection()
		if err != nil {
			return 0, fmt.Errorf("make_shortest_connection")
		}
		log("make connection %d", i+1)
	}

	// make cluster
	jBoxMap := map[int]*JBox{}
	for _, jb := range jBoxes {
		jBoxMap[jb.ID] = jb
	}
	takeOne := func() (*JBox, bool) {
		for _, jb := range jBoxMap {
			return jb, true
		}
		return nil, false
	}

	var fillClusterRec func(c *Cluster, jb *JBox)
	fillClusterRec = func(c *Cluster, jb *JBox) {
		if c.Contains(jb) {
			return
		}
		c.JBoxes = append(c.JBoxes, jb)
		delete(jBoxMap, jb.ID)

		for _, cjb := range jb.ConnectedTo {
			fillClusterRec(c, cjb)
		}
	}

	log("fill clusters")
	// loop until map is empty
	var clusters []*Cluster
	for {
		jb, ok := takeOne()
		if !ok {
			break
		}
		//
		cluster := &Cluster{}
		fillClusterRec(cluster, jb)
		clusters = append(clusters, cluster)
	}
	//
	log("fill %d clusters done", len(clusters))

	sort.Slice(clusters, func(i, j int) bool {
		return len(clusters[i].JBoxes) > len(clusters[j].JBoxes)
	})

	val := len(clusters[0].JBoxes) *
		len(clusters[1].JBoxes) *
		len(clusters[2].JBoxes)

	return val, nil
}

func part2MainFunc(in string) (int, error) {
	return 0, nil
}
