package bitmask

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type Mask struct {
	raw string
}

func (m Mask) String() string {
	return m.raw
}

type Command interface {
	String() string
}

type ChangeMask struct {
	mask Mask
}

func (c ChangeMask) String() string {
	return fmt.Sprintf("change-mask: %q", c.mask)
}

func ParseChangeMask(s string) (ChangeMask, error) {
	var mr string
	_, err := fmt.Sscanf(s, "mask = %s", &mr)
	if err != nil {
		return ChangeMask{}, err
	}
	return ChangeMask{
		mask: Mask{raw: mr},
	}, nil
}

type Poke struct {
	addr uint64
	val  uint64
}

func ParsePoke(s string) (Poke, error) {
	var addr, val int
	_, err := fmt.Sscanf(s, "mem[%d] = %d", &addr, &val)
	if err != nil {
		return Poke{}, err
	}
	if addr < 0 || val < 0 {
		return Poke{}, errors.Errorf("neg. values are forbidden")
	}
	return Poke{
		addr: uint64(addr),
		val:  uint64(val),
	}, nil
}

func (c Poke) String() string {
	return fmt.Sprintf("poke: addr=%d val=%d", c.addr, c.val)
}

func ParseCommands(r io.Reader) ([]Command, error) {
	var cs []Command
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " \r\n\t")
		if l == "" {
			continue
		}
		switch {
		case strings.HasPrefix(l, "mask"):
			c, err := ParseChangeMask(l)
			if err != nil {
				return nil, err
			}
			cs = append(cs, c)
		case strings.HasPrefix(l, "mem"):
			c, err := ParsePoke(l)
			if err != nil {
				return nil, err
			}
			cs = append(cs, c)
		}
	}
	return cs, nil
}
