package jigsaw

import "fmt"

func edgeMatchPosition(dim int, i int) MatchPosition {
	if i == 0 {
		return MatchNone
	}
	i -= 1
	//m := i % dim
	r := i / (dim - 1)
	switch r {
	case 0:
		return MatchRight
	case 1:
		return MatchBottom
	case 2:
		return MatchLeft
	default:
		return MatchTop
	}
}

func TileIDs(tiles []Tile) []int {
	var ids []int
	for _, tile := range tiles {
		ids = append(ids, tile.id)
	}
	return ids
}

func ArrangeEdge(dim int, tiles []Tile) {
	if 2*dim+2*(dim-2) != len(tiles) {
		panic("invalid dim or edge count")
	}
	for _, t := range tiles {
		for _, trans := range AllTransforms() {
			transt := t.Transformed(trans)

			exclude := []Tile{transt}
			edgePos := 1
			subTiles, ok := constructEdge(dim, edgePos, transt, transt, tiles, exclude)
			if ok {
				edgeTiles := []Tile{transt}
				edgeTiles = append(edgeTiles, subTiles...)
				fmt.Printf("constructed edge: %v\n", TileIDs(edgeTiles))
			}
		}
	}
}

func constructEdge(dim int, edgePos int, firstTile Tile, prevTile Tile, tiles []Tile, exclude []Tile) ([]Tile, bool) {
	mustExclude := func(t Tile) bool {
		for _, ext := range exclude {
			if ext.id == t.id {
				return true
			}
		}
		return false
	}

	for _, t := range tiles {
		if mustExclude(t) {
			continue
		}
		for _, trans := range AllTransforms() {
			transt := t.Transformed(trans)
			matchPos := edgeMatchPosition(dim, edgePos)
			if !prevTile.MatchAt(matchPos, transt) {
				continue
			}
			//match!
			if edgePos == 4*(dim-1)-1 {
				if !firstTile.MatchAt(MatchBottom, transt) {
					continue
				} else {
					return []Tile{transt}, true
				}
			}

			nex := append(exclude, transt)
			subTiles, ok := constructEdge(dim, edgePos+1, firstTile, transt, tiles, nex)
			if ok {
				edgeTiles := []Tile{transt}
				edgeTiles = append(edgeTiles, subTiles...)
				return edgeTiles, true
			}
		}
	}

	return nil, false
}
