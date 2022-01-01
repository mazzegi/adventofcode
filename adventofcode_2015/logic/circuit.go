package logic

import (
	"strconv"

	"github.com/mazzegi/wasa/errors"
)

type Circuit struct {
	blocks []*Block
}

func NewCircuit(blocks []BlockScheme) *Circuit {
	c := &Circuit{}
	for _, bs := range blocks {
		c.blocks = append(c.blocks, NewBlock(bs))
	}
	return c
}

func (c *Circuit) findBlockWithOutput(wire string) (*Block, bool) {
	for _, b := range c.blocks {
		if b.Scheme.Output == wire {
			return b, true
		}
	}
	return nil, false
}

func (c *Circuit) ValueOf(wire string) (uint16, error) {
	//find block with output == wire
	//fmt.Printf("find value of %q\n", wire)

	block, ok := c.findBlockWithOutput(wire)
	if !ok {
		return 0, errors.Errorf("found no block with output %q", wire)
	}
	if val, ok := block.Value(); ok {
		return val, nil
	}

	var ins []uint16
	for _, in := range block.Scheme.Input {
		num, err := strconv.ParseUint(in, 10, 16)
		if err == nil {
			ins = append(ins, uint16(num))
			continue
		}

		vin, err := c.ValueOf(in)
		if err != nil {
			return 0, err
		}
		ins = append(ins, vin)
	}
	val, err := block.Evaluate(ins...)
	if err != nil {
		return 0, err
	}
	return val, nil
}
