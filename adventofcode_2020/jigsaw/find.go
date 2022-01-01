package jigsaw

import "fmt"

func FindTilesWithNoMatches(tiles []Tile) {
	for _, t := range tiles {
		notMatches := notMatchingSides(t, tiles)
		if len(notMatches) > 0 {
			fmt.Printf("tile %d doesn't match all sides. missing %v\n", t.id, notMatches)
		}
	}
}

func notMatchingSides(tile Tile, tiles []Tile) []MatchPosition {
	var cnt int
	tileMatches := map[MatchPosition]bool{}
	matchingTiles := map[int]Tile{}
	for _, trans := range AllTransforms() {
		ttrans := tile.Transformed(trans)
		for _, ot := range tiles {
			if ot.id == ttrans.id {
				continue
			}
			for _, trans := range AllTransforms() {
				ottrans := ot.Transformed(trans)
				mps := ttrans.Match(ottrans, []MatchPosition{})
				for _, mp := range mps {
					tileMatches[mp] = true
				}
				cnt += len(mps)
				if len(mps) > 0 {
					matchingTiles[ot.id] = ot
				}
			}
		}
	}
	var notMatches []MatchPosition
	for _, mp := range AllValidMatchPositions() {
		if _, ok := tileMatches[mp]; !ok {
			notMatches = append(notMatches, mp)
		}
	}
	fmt.Printf("tile %d matches with %d other tiles\n", tile.id, len(matchingTiles))
	return notMatches
}

func FindEdgeTiles(ts []Tile) []Tile {
	var edgeTiles []Tile
	for _, t := range ts {
		if matchingTileCnt(t, ts) < 4 {
			edgeTiles = append(edgeTiles, t)
		}
	}
	return edgeTiles
}

func matchingTileCnt(tile Tile, tiles []Tile) int {
	matchingTiles := map[int]Tile{}
	for _, trans := range AllTransforms() {
		ttrans := tile.Transformed(trans)
		for _, ot := range tiles {
			if ot.id == ttrans.id {
				continue
			}
			for _, trans := range AllTransforms() {
				ottrans := ot.Transformed(trans)
				mps := ttrans.Match(ottrans, []MatchPosition{})
				if len(mps) > 0 {
					matchingTiles[ot.id] = ot
				}
			}
		}
	}
	return len(matchingTiles)
}
