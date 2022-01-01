package bits

import "github.com/pkg/errors"

const (
	TypeIDLiteralValue uint8 = 4
	TypeIDOpSum        uint8 = 0
	TypeIDOpProduct    uint8 = 1
	TypeIDOpMin        uint8 = 2
	TypeIDOpMax        uint8 = 3
	TypeIDOpGreater    uint8 = 5
	TypeIDOpLess       uint8 = 6
	TypeIDOpEqual      uint8 = 7
)

type Packet interface {
	SumOfVersions() int
	Eval() (int, error)
}

type PacketHeader struct {
	Version uint8
	TypeID  uint8
}

//
type LiteralValue struct {
	Header PacketHeader
	Value  int
}

func (v *LiteralValue) SumOfVersions() int {
	return int(v.Header.Version)
}

func (v *LiteralValue) Eval() (int, error) {
	return v.Value, nil
}

//
type Operator struct {
	Header  PacketHeader
	Packets []Packet
}

func (op *Operator) SumOfVersions() int {
	var sum int
	for _, sp := range op.Packets {
		sum += sp.SumOfVersions()
	}
	return sum + int(op.Header.Version)
}

func sum(vs []int) int {
	var sum int
	for _, v := range vs {
		sum += v
	}
	return sum
}

func product(vs []int) int {
	var prd int
	first := true
	for _, v := range vs {
		if first {
			prd = v
			first = false
			continue
		}
		prd *= v
	}
	return prd
}

func min(vs []int) int {
	var mv int
	first := true
	for _, v := range vs {
		if first {
			mv = v
			first = false
			continue
		}
		if v < mv {
			mv = v
		}
	}
	return mv
}

func max(vs []int) int {
	var mv int
	first := true
	for _, v := range vs {
		if first {
			mv = v
			first = false
			continue
		}
		if v > mv {
			mv = v
		}
	}
	return mv
}

func greater(vs []int) (int, error) {
	if len(vs) != 2 {
		return 0, errors.Errorf("greater with non-2-len values")
	}
	if vs[0] > vs[1] {
		return 1, nil
	}
	return 0, nil
}

func less(vs []int) (int, error) {
	if len(vs) != 2 {
		return 0, errors.Errorf("less with non-2-len values")
	}
	if vs[0] < vs[1] {
		return 1, nil
	}
	return 0, nil
}

func equal(vs []int) (int, error) {
	if len(vs) != 2 {
		return 0, errors.Errorf("equal with non-2-len values")
	}
	if vs[0] == vs[1] {
		return 1, nil
	}
	return 0, nil
}

func (op *Operator) Eval() (int, error) {
	var subValues []int
	for _, p := range op.Packets {
		v, err := p.Eval()
		if err != nil {
			return 0, err
		}
		subValues = append(subValues, v)
	}

	switch op.Header.TypeID {
	case TypeIDOpSum:
		return sum(subValues), nil
	case TypeIDOpProduct:
		return product(subValues), nil
	case TypeIDOpMin:
		return min(subValues), nil
	case TypeIDOpMax:
		return max(subValues), nil
	case TypeIDOpGreater:
		return greater(subValues)
	case TypeIDOpLess:
		return less(subValues)
	case TypeIDOpEqual:
		return equal(subValues)
	default:
		return 0, errors.Errorf("invalid typeID %d", op.Header.TypeID)
	}
}
