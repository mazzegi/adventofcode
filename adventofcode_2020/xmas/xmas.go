package xmas

type Processor struct {
	in        []int64
	blockSize int64
}

func NewProcessor(in []int64, blockSize int64) *Processor {
	return &Processor{
		in:        in,
		blockSize: blockSize,
	}
}

func (p *Processor) FindFirstNotValid() (pos, val int64, found bool) {
	for i := p.blockSize; i < int64(len(p.in)); i++ {
		block := p.in[i-p.blockSize : i]
		val := p.in[i]
		if !p.valid(block, val) {
			return i, val, true
		}
	}
	return 0, 0, false
}

func (p *Processor) valid(block []int64, val int64) bool {
	for _, n1 := range block {
		for _, n2 := range block {
			if n1+n2 == val {
				return true
			}
		}
	}
	return false
}

func (p *Processor) FindWeakness(num int64) ([]int64, bool) {
	sum := func(set []int64) int64 {
		var s int64
		for _, n := range set {
			s += n
		}
		return s
	}
	for start := 0; start < len(p.in)-1; start++ {
		for end := start + 1; end < len(p.in); end++ {
			set := p.in[start : end+1]
			//fmt.Printf("check [%d, %d] => %v\n", start, end, set)
			if sum(set) == num {
				return set, true
			}
		}
	}
	return nil, false
}
