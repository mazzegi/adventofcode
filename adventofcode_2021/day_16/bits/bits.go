package bits

import "github.com/pkg/errors"

var hex2Bin = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func ParseHex(hex string) (*Message, error) {
	var bin []bool
	for _, r := range hex {
		bstr, ok := hex2Bin[r]
		if !ok {
			return nil, errors.Errorf("invalid hex letter %q", string(r))
		}
		for _, br := range bstr {
			switch br {
			case '0':
				bin = append(bin, false)
			default:
				bin = append(bin, true)
			}
		}
	}
	if len(bin) < 6 {
		return nil, errors.Errorf("to less bits to decode message (%d)", len(bin))
	}

	pkt, err := NewParser(bin).Parse()
	if err != nil {
		return nil, errors.Wrap(err, "parse")
	}
	return &Message{
		root: pkt,
	}, nil
}

type Message struct {
	root Packet
}

func (msg *Message) SumOfVersion() int {
	return msg.root.SumOfVersions()
}

func (msg *Message) Eval() (int, error) {
	return msg.root.Eval()
}
