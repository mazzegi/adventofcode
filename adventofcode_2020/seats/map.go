package seats

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type State rune

const (
	Floor    State = '.'
	Empty    State = 'L'
	Occupied State = '#'
	Unknown  State = '?'
)

type Row struct {
	States     []State
	NextStates []State
}

type Map struct {
	Rows           []*Row
	NumRows        int
	NumSeatsPerRow int
}

func ParseRow(s string) (*Row, error) {
	row := &Row{}
	for _, r := range s {
		switch State(r) {
		case Floor, Empty, Occupied:
			row.States = append(row.States, State(r))
			if State(r) == Floor {
				row.NextStates = append(row.NextStates, Floor)
			} else {
				row.NextStates = append(row.NextStates, State(r))
			}
		default:
			return nil, errors.Errorf("invalid state: %q", string(r))
		}
	}
	return row, nil
}

func ParseMap(r io.Reader) (*Map, error) {
	m := &Map{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		row, err := ParseRow(l)
		if err != nil {
			return nil, errors.Wrapf(err, "parse-row %q", l)
		}
		if len(m.Rows) > 0 && len(row.States) != len(m.Rows[0].States) {
			return nil, errors.Errorf("different row-size: %q", l)
		}
		m.Rows = append(m.Rows, row)
	}
	if len(m.Rows) == 0 {
		return nil, errors.Errorf("map is empty")
	}
	m.NumRows = len(m.Rows)
	m.NumSeatsPerRow = len(m.Rows[0].States)
	return m, nil
}

func (row *Row) String() string {
	rs := make([]rune, len(row.States))
	for i := 0; i < len(row.States); i++ {
		rs[i] = rune(row.States[i])
	}
	return string(rs)
}

func (m *Map) String() string {
	var sl []string
	for _, row := range m.Rows {
		sl = append(sl, row.String())
	}
	return strings.Join(sl, "\n")
}

func (row *Row) Next() {
	for i := 0; i < len(row.States); i++ {
		if row.NextStates[i] == Floor {
			continue
		}
		row.States[i] = row.NextStates[i]
		row.NextStates[i] = Unknown
	}
}

func (m *Map) AdjacentOccupiedSeats(rowIdx, seatIdx int) int {
	var cnt int
	for ri := rowIdx - 1; ri <= rowIdx+1; ri++ {
		if ri < 0 || ri >= m.NumRows {
			continue
		}
		for si := seatIdx - 1; si <= seatIdx+1; si++ {
			if si < 0 || si >= m.NumSeatsPerRow {
				continue
			}
			if ri == rowIdx && si == seatIdx {
				continue
			}
			state := m.Rows[ri].States[si]
			if state == Occupied {
				cnt++
			}
		}
	}
	return cnt
}

func (m *Map) NextByAdj() int {
	var changes int
	for ri, row := range m.Rows {
		for si, state := range row.States {
			if state == Floor {
				continue
			}
			numAO := m.AdjacentOccupiedSeats(ri, si)
			switch {
			case state == Empty && numAO == 0:
				row.NextStates[si] = Occupied
				if row.States[si] != Occupied {
					changes++
				}
			case state == Occupied && numAO >= 4:
				row.NextStates[si] = Empty
				if row.States[si] != Empty {
					changes++
				}
			default:
				row.NextStates[si] = state
			}
		}
	}
	for _, row := range m.Rows {
		row.Next()
	}
	return changes
}

type dir struct {
	dx, dy int
}

func (d dir) addTo(ri, si *int) {
	*ri += d.dy
	*si += d.dx
}

var dirs = []dir{
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
}

func (m *Map) SeenOccupiedSeats(rowIdx, seatIdx int) int {

	validPos := func(ri, si int) bool {
		if ri < 0 || ri >= m.NumRows {
			return false
		}
		if si < 0 || si >= m.NumSeatsPerRow {
			return false
		}
		return true
	}

	var cnt int
	for _, d := range dirs {
		ri, si := rowIdx, seatIdx
		d.addTo(&ri, &si)
		for validPos(ri, si) {
			if m.Rows[ri].States[si] == Occupied {
				cnt++
				break
			}
			if m.Rows[ri].States[si] == Empty {
				break
			}
			d.addTo(&ri, &si)
		}
	}

	return cnt
}

func (m *Map) NextBySeen() int {
	var changes int
	for ri, row := range m.Rows {
		for si, state := range row.States {
			if state == Floor {
				continue
			}
			numAO := m.SeenOccupiedSeats(ri, si)
			switch {
			case state == Empty && numAO == 0:
				row.NextStates[si] = Occupied
				if row.States[si] != Occupied {
					changes++
				}
			case state == Occupied && numAO >= 5:
				row.NextStates[si] = Empty
				if row.States[si] != Empty {
					changes++
				}
			default:
				row.NextStates[si] = state
			}
		}
	}
	for _, row := range m.Rows {
		row.Next()
	}
	return changes
}

func (row *Row) OccupiedSeats() int {
	var cnt int
	for _, s := range row.States {
		if s == Occupied {
			cnt++
		}
	}
	return cnt
}

func (m *Map) OccupiedSeats() int {
	var cnt int
	for _, row := range m.Rows {
		cnt += row.OccupiedSeats()
	}
	return cnt
}
