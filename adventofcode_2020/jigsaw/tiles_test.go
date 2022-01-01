package jigsaw

import "testing"

func TestTile(t *testing.T) {
	var rawTile []string

	rawTile = []string{
		"Tile 2311:",
		"..##.#..#.",
		"##..#.....",
		"#...##..#.",
		"####.#...#",
		"##.##.###.",
		"##...#.###",
		".#.#.#..##",
		"..#....#..",
		"###...#.#.",
		"..###..###",
	}

	tile, err := ParseTile(rawTile, 10, 10)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	t.Logf("\n%s", tile.String())

	rottile := tile.Rotated(Clockwise, 3)
	t.Logf("\nrot-one:\n%s", rottile.String())

	mytile := tile.FlippedColumns()
	t.Logf("\nmirror-y:\n%s", mytile.String())

	mxtile := tile.FlippedRows()
	t.Logf("\nmirror-x:\n%s", mxtile.String())
}

func TestMatchEdgePos(t *testing.T) {
	dim := 12
	for i := 1; i < 4*(dim-1); i++ {
		t.Logf("%d => %s", i, edgeMatchPosition(dim, i))
	}
}
