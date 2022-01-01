package jigsaw

import (
	"fmt"
	"math"
	"strings"

	"github.com/pkg/errors"
)

func copyInts(ns []int) []int {
	cns := make([]int, len(ns))
	copy(cns, ns)
	return cns
}

func flippedInts(ns []int) []int {
	fns := make([]int, len(ns))
	for i := 0; i < len(ns); i++ {
		fns[i] = ns[len(ns)-1-i]
	}
	return fns
}

func intsEqual(ns1, ns2 []int) bool {
	if len(ns1) != len(ns2) {
		return false
	}
	for i := 0; i < len(ns1); i++ {
		if ns1[i] != ns2[i] {
			return false
		}
	}
	return true
}

func stringIntsInner(ns []int) string {
	var s string
	for i := 1; i < len(ns)-1; i++ {
		n := ns[i]
		switch n {
		case 0:
			s += "."
		case 1:
			s += "#"
		default:
			s += "?"
		}
	}
	return s
}

func stringInts(ns []int) string {
	var s string
	for _, n := range ns {
		switch n {
		case 0:
			s += "."
		case 1:
			s += "#"
		default:
			s += "?"
		}
	}
	return s
}

func ParseTile(sl []string, dimRows int, dimCols int) (Tile, error) {
	if len(sl) != dimRows+1 {
		return Tile{}, errors.Errorf("invalid number of lines %d (want %d)", len(sl), dimRows+1)
	}
	var id int
	_, err := fmt.Sscanf(sl[0], "Tile %d:", &id)
	if err != nil {
		return Tile{}, errors.Wrapf(err, "scan id %q", sl[0])
	}
	t := Tile{id: id}
	for i := 1; i < len(sl); i++ {
		var row []int
		for _, r := range sl[i] {
			switch r {
			case '#':
				row = append(row, 1)
			case '.':
				row = append(row, 0)
			default:
				return Tile{}, errors.Errorf("invalid rune %q", string(r))
			}
		}
		if len(row) != dimCols {
			return Tile{}, errors.Errorf("invalid row-size %d (want %d)", len(row), dimCols)
		}
		t.rows = append(t.rows, row)
	}
	if len(t.rows) != dimRows {
		return Tile{}, errors.Errorf("invalid grid-size %d (want %d)", len(t.rows), dimRows)
	}
	return t, nil
}

func (t Tile) String() string {
	var sl []string
	for _, row := range t.rows {
		sl = append(sl, stringInts(row))
	}
	return fmt.Sprintf("Tile %d:\n%s", t.id, strings.Join(sl, "\n"))
}

type Tile struct {
	id   int
	rows [][]int //y,x (row, col)
}

func (t Tile) ID() int {
	return t.id
}

func (t Tile) Col(ci int) []int {
	cs := make([]int, len(t.rows))
	for r := 0; r < len(t.rows); r++ {
		cs[r] = t.rows[r][ci]
	}
	return cs
}

func (t Tile) Row(ri int) []int {
	return t.rows[ri]
}

func (t Tile) FirstRow() []int {
	return t.rows[0]
}

func (t Tile) LastRow() []int {
	return t.rows[len(t.rows)-1]
}

func (t Tile) FirstCol() []int {
	return t.Col(0)
}

func (t Tile) LastCol() []int {
	return t.Col(len(t.FirstRow()) - 1)
}

func (t Tile) FlippedRows() Tile {
	nt := Tile{
		id:   t.id,
		rows: make([][]int, len(t.rows)),
	}
	for r := 0; r < len(t.rows); r++ {
		nt.rows[r] = copyInts(t.rows[len(t.rows)-1-r])
	}
	return nt
}

func (t Tile) FlippedColumns() Tile {
	nt := Tile{
		id:   t.id,
		rows: make([][]int, len(t.rows)),
	}
	for r := 0; r < len(t.rows); r++ {
		nt.rows[r] = flippedInts(t.rows[r])
	}
	return nt
}

type RotateDir string

const (
	Clockwise        RotateDir = "clockwise"
	CounterClockwise RotateDir = "counter-clockwise"
)

func (t Tile) Rotated(dir RotateDir, rcnt int) Tile {
	rcnt = rcnt % 4
	if rcnt == 0 {
		return t
	}
	switch dir {
	case Clockwise:
		rt := t
		for i := 0; i < rcnt; i++ {
			rt = rt.rotatedOneClockwise()
		}
		return rt
	case CounterClockwise:
		return t.Rotated(Clockwise, 4-rcnt)
	default:
		panic("invalid rotate dir")
	}
}

func (t Tile) rotatedOneClockwise() Tile {
	nt := Tile{
		id:   t.id,
		rows: make([][]int, len(t.rows)),
	}
	for r := 0; r < len(t.rows); r++ {
		nt.rows[r] = flippedInts(t.Col(r))
	}
	return nt
}

//
type MatchPosition string

const (
	MatchNone   MatchPosition = "none"
	MatchLeft   MatchPosition = "left"
	MatchTop    MatchPosition = "top"
	MatchRight  MatchPosition = "right"
	MatchBottom MatchPosition = "bottom"
)

func AllValidMatchPositions() []MatchPosition {
	return []MatchPosition{
		MatchLeft, MatchRight, MatchTop, MatchBottom,
	}
}

type Transform string

const (
	TransformNone      Transform = "none"
	TransformRotate90  Transform = "rotate-90"
	TransformRotate180 Transform = "rotate-180"
	TransformRotate270 Transform = "rotate-270"
	TransformFlipRows  Transform = "flip-rows"
	TransformFlipCols  Transform = "flip-cols"

	TransformFlipRowsRotate90  Transform = "flip-rows-rotate-90"
	TransformFlipRowsRotate180 Transform = "flip-rows-rotate-180"
	TransformFlipRowsRotate270 Transform = "flip-rows-rotate-270"

	TransformFlipColsRotate90  Transform = "flip-cols-rotate-90"
	TransformFlipColsRotate180 Transform = "flip-cols-rotate-180"
	TransformFlipColsRotate270 Transform = "flip-cols-rotate-270"
)

func AllTransforms() []Transform {
	return []Transform{
		TransformNone,
		TransformRotate90,
		TransformRotate180,
		TransformRotate270,
		TransformFlipRows,
		TransformFlipCols,
		TransformFlipRowsRotate90,
		TransformFlipRowsRotate180,
		TransformFlipRowsRotate270,
		TransformFlipColsRotate90,
		TransformFlipColsRotate180,
		TransformFlipColsRotate270,
	}
}

func (t Tile) Transformed(trans Transform) Tile {
	switch trans {
	case TransformRotate90:
		return t.Rotated(Clockwise, 1)
	case TransformRotate180:
		return t.Rotated(Clockwise, 2)
	case TransformRotate270:
		return t.Rotated(Clockwise, 3)
	case TransformFlipRows:
		return t.FlippedRows()
	case TransformFlipCols:
		return t.FlippedColumns()

	case TransformFlipRowsRotate90:
		return t.FlippedRows().Rotated(Clockwise, 1)
	case TransformFlipRowsRotate180:
		return t.FlippedRows().Rotated(Clockwise, 2)
	case TransformFlipRowsRotate270:
		return t.FlippedRows().Rotated(Clockwise, 3)
	case TransformFlipColsRotate90:
		return t.FlippedColumns().Rotated(Clockwise, 1)
	case TransformFlipColsRotate180:
		return t.FlippedColumns().Rotated(Clockwise, 2)
	case TransformFlipColsRotate270:
		return t.FlippedColumns().Rotated(Clockwise, 3)

	default: //TransformNone
		return t
	}
}

func (t Tile) Match(ot Tile, exclude []MatchPosition) []MatchPosition {
	mustExclude := func(mpos MatchPosition) bool {
		for _, exmpos := range exclude {
			if exmpos == mpos {
				return true
			}
		}
		return false
	}

	var mposs []MatchPosition
	if !mustExclude(MatchRight) && intsEqual(t.LastCol(), ot.FirstCol()) {
		mposs = append(mposs, MatchRight)
	}
	if !mustExclude(MatchLeft) && intsEqual(t.FirstCol(), ot.LastCol()) {
		mposs = append(mposs, MatchLeft)
	}
	if !mustExclude(MatchTop) && intsEqual(t.FirstRow(), ot.LastRow()) {
		mposs = append(mposs, MatchTop)
	}
	if !mustExclude(MatchBottom) && intsEqual(t.LastRow(), ot.FirstRow()) {
		mposs = append(mposs, MatchBottom)
	}
	return mposs
}

type LinkedTile struct {
	Tile                     Tile
	Left, Top, Right, Bottom *LinkedTile
}

func (lt *LinkedTile) TryLink(t Tile) (*LinkedTile, bool) {
	exclude := []MatchPosition{}
	if lt.Left != nil {
		exclude = append(exclude, MatchLeft)
	}
	if lt.Right != nil {
		exclude = append(exclude, MatchRight)
	}
	if lt.Top != nil {
		exclude = append(exclude, MatchTop)
	}
	if lt.Bottom != nil {
		exclude = append(exclude, MatchBottom)
	}

	for _, trans := range AllTransforms() {
		transTile := t.Transformed(trans)
		mposs := lt.Tile.Match(transTile, exclude)
		if len(mposs) > 0 {
			nlt := &LinkedTile{
				Tile: transTile,
			}
			mpos := mposs[0]
			switch mpos {
			case MatchLeft:
				lt.Left = nlt
				nlt.Right = lt
			case MatchRight:
				lt.Right = nlt
				nlt.Left = lt
			case MatchTop:
				lt.Top = nlt
				nlt.Bottom = lt
			case MatchBottom:
				lt.Bottom = nlt
				nlt.Top = lt
			}
			fmt.Printf("linked %d to %d at %s (trans=%s) (poss=%d)\n", transTile.id, lt.Tile.id, mpos, trans, len(mposs))
			return nlt, true
		}
	}
	return nil, false
}

func AlignTiles(ts []Tile) ([][]*Tile, error) {
	squareDim := int(math.Sqrt(float64(len(ts))))

	for i, t := range ts {
		for _, trans := range AllTransforms() {
			transt := t.Transformed(trans)
			other := map[int]Tile{}
			for _, ot := range ts {
				if ot.id != transt.id {
					other[ot.id] = ot
				}
			}

			fmt.Printf("*** (%d / %d) find-square (dim %d) with %d (trans = %q) at top-left ***\n", i, len(ts), squareDim, transt.id, trans)
			sq, ok := findSquare(squareDim, transt, other)
			if ok {
				return sq, nil
			}
		}
	}

	return nil, errors.Errorf("found no square")
}

func (t *Tile) MatchBottom(ot Tile) bool {
	if t == nil {
		return true
	}
	return intsEqual(t.LastRow(), ot.FirstRow())
}

func (t *Tile) MatchTop(ot Tile) bool {
	if t == nil {
		return true
	}
	return intsEqual(t.FirstRow(), ot.LastRow())
}

func (t *Tile) MatchLeft(ot Tile) bool {
	if t == nil {
		return true
	}
	return intsEqual(t.FirstCol(), ot.LastCol())
}

func (t *Tile) MatchRight(ot Tile) bool {
	if t == nil {
		return true
	}
	return intsEqual(t.LastCol(), ot.FirstCol())
}

func (t Tile) MatchAt(pos MatchPosition, ot Tile) bool {
	switch pos {
	case MatchLeft:
		return intsEqual(t.FirstCol(), ot.LastCol())
	case MatchRight:
		return intsEqual(t.LastCol(), ot.FirstCol())
	case MatchTop:
		return intsEqual(t.FirstRow(), ot.LastRow())
	case MatchBottom:
		return intsEqual(t.LastRow(), ot.FirstRow())
	default:
		return false
	}
}

func findSquare(dim int, startTile Tile, tiles map[int]Tile) ([][]*Tile, bool) {
	square := make([][]*Tile, dim)
	for r := 0; r < dim; r++ {
		square[r] = make([]*Tile, dim)
	}

	match := func(tile Tile, row, col int) bool {
		//top
		if row > 0 {
			topTile := square[row-1][col]
			if !topTile.MatchBottom(tile) {
				return false
			}
		}
		//bottom
		if row < dim-1 {
			btmTile := square[row+1][col]
			if !btmTile.MatchTop(tile) {
				return false
			}
		}
		//left
		if col > 0 {
			leftTile := square[row][col-1]
			if !leftTile.MatchRight(tile) {
				return false
			}
		}
		//right
		if col < dim-1 {
			rightTile := square[row][col+1]
			if !rightTile.MatchLeft(tile) {
				return false
			}
		}

		return true
	}

	square[0][0] = &startTile

	for row := 0; row < dim; row++ {
	range_cols:
		for col := 0; col < dim; col++ {
			if col == 0 && row == 0 {
				continue
			}
			//try to position tile
			for _, tile := range tiles {
				for _, trans := range AllTransforms() {
					transt := tile.Transformed(trans)
					if match(transt, row, col) {
						square[row][col] = &transt
						delete(tiles, transt.id)
						continue range_cols
					}
				}
			}
		}
	}
	if len(tiles) > 0 {
		fmt.Printf("*** cannot complete square (%d tiles left) *** \n\n", len(tiles))
		return nil, false
	}
	fmt.Printf("*** completed square ***\n")
	for row := 0; row < dim; row++ {
		for col := 0; col < dim; col++ {
			fmt.Printf("%d ", square[row][col].id)
		}
		fmt.Printf("\n")
	}

	magic := square[0][0].id *
		square[0][dim-1].id *
		square[dim-1][0].id *
		square[dim-1][dim-1].id
	fmt.Printf("magic: %d\n", magic)

	//dump square
	fmt.Printf("*** image (without border) ***\n")
	tileRows := len(startTile.rows)
	for _, srow := range square {
		for tr := 1; tr < tileRows-1; tr++ {
			for _, scol := range srow {
				fmt.Printf("%s", stringIntsInner(scol.Row(tr)))
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("*** image-end ***")

	return square, true
}

func AlignTilesNaiv(ts []Tile) ([][]Tile, error) {
	if len(ts) < 2 {
		return nil, errors.Errorf("too less tiles")
	}
	//starting with one
	lt := &LinkedTile{
		Tile: ts[0],
	}
	linked := []*LinkedTile{lt}

	unlinked := map[int]Tile{}
	for i := 1; i < len(ts); i++ {
		unlinked[ts[i].id] = ts[i]
	}

	linkOne := func() error {
		for _, ul := range unlinked {
			for _, lt := range linked {
				nlt, ok := lt.TryLink(ul)
				if !ok {
					continue
				}
				delete(unlinked, ul.id)
				linked = append(linked, nlt)
				return nil
			}
		}
		return errors.Errorf("cannot link any tile anymore")
	}

	for len(unlinked) > 0 {
		err := linkOne()
		if err != nil {
			return nil, errors.Wrap(err, "link-one")
		}
	}

	//find left, top empty
	var tl *LinkedTile
	for _, lt := range linked {
		if lt.Left == nil && lt.Top == nil {
			tl = lt
			break
		}
	}
	if tl == nil {
		return nil, errors.Errorf("found no left top tile")
	}
	rt := tl
	for {
		ct := rt
		for {
			fmt.Printf("%d ", ct.Tile.id)
			if ct.Right != nil {
				ct = ct.Right
			} else {
				fmt.Printf("\n")
				break
			}
		}

		if rt.Bottom != nil {
			rt = rt.Bottom
		} else {
			break
		}
	}

	return nil, nil
}
