package bits

import "github.com/pkg/errors"

func NewParser(bin []bool) *Parser {
	return &Parser{
		bin: bin,
		pos: 0,
	}
}

type Parser struct {
	bin []bool
	pos int
}

func (p *Parser) Parse() (Packet, error) {
	return p.parsePacket()
}

func (p *Parser) nextSlice(cnt int) ([]bool, error) {
	if p.pos+cnt-1 >= len(p.bin) {
		return nil, errors.Errorf("EOF")
	}
	defer func() {
		p.pos += cnt
	}()
	return p.bin[p.pos : p.pos+cnt], nil
}

func (p *Parser) next() (bool, error) {
	bs, err := p.nextSlice(1)
	if err != nil {
		return false, err
	}
	return bs[0], nil
}

func (p *Parser) parsePacket() (Packet, error) {
	//take first 3 => version
	//take next  3 => typeID
	bs, err := p.nextSlice(3)
	if err != nil {
		return nil, err
	}
	ver := bitsToNumber(bs)
	bs, err = p.nextSlice(3)
	if err != nil {
		return nil, err
	}
	typeID := bitsToNumber(bs)
	ph := PacketHeader{
		Version: uint8(ver),
		TypeID:  uint8(typeID),
	}

	// defer func() {
	// 	p.eatZeros()
	// }()
	switch ph.TypeID {
	case TypeIDLiteralValue:
		return p.parseLiteralValue(ph)
	default:
		return p.parseOperator(ph)
	}

}

func (p *Parser) eatZeros() {
	for {
		b, err := p.next()
		if err != nil {
			return
		}
		if b {
			return
		}
	}
}

// 1011 1111 1000 1010 00
// 10111 11110 00101 000

func (p *Parser) parseLiteralValue(head PacketHeader) (*LiteralValue, error) {
	var nbs []bool
	for {
		g, err := p.nextSlice(5)
		if err != nil {
			return nil, err
		}
		nbs = append(nbs, g[1:]...)
		if !g[0] {
			break
		}
	}
	num := bitsToNumber(nbs)
	return &LiteralValue{
		Header: head,
		Value:  num,
	}, nil
}

func (p *Parser) parseOperator(head PacketHeader) (*Operator, error) {
	lengthTypeID, err := p.next()
	if err != nil {
		return nil, err
	}

	if !lengthTypeID {
		// total length in bits
		nbs, err := p.nextSlice(15)
		if err != nil {
			return nil, err
		}
		totalLengthBits := bitsToNumber(nbs)
		return p.parseOperatorTotalLength(head, totalLengthBits)
	} else {
		// number of sub packets
		nbs, err := p.nextSlice(11)
		if err != nil {
			return nil, err
		}
		numberOfSubPackets := bitsToNumber(nbs)
		return p.parseOperatorNumPacktes(head, numberOfSubPackets)
	}
}

func (p *Parser) parseOperatorTotalLength(head PacketHeader, length int) (*Operator, error) {
	op := &Operator{
		Header: head,
	}
	curr := p.pos
	for {
		pkt, err := p.parsePacket()
		if err != nil {
			return nil, err
		}
		op.Packets = append(op.Packets, pkt)
		if p.pos-curr >= length {
			return op, nil
		}
	}
}

func (p *Parser) parseOperatorNumPacktes(head PacketHeader, num int) (*Operator, error) {
	op := &Operator{
		Header: head,
	}
	for i := 0; i < num; i++ {
		pkt, err := p.parsePacket()
		if err != nil {
			return nil, err
		}
		op.Packets = append(op.Packets, pkt)
	}
	return op, nil
}
