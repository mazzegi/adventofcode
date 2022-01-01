package day_06

import (
	"fmt"
	"testing"
)

const inputTest = `
eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
`

func TestErrorCorrectedMessage(t *testing.T) {

	tests := []struct {
		in  string
		msg string
	}{
		{
			in:  inputTest,
			msg: "easter",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := ErrorCorrectedMessage(test.in)
			if res != test.msg {
				t.Fatalf("ecm: expect %q, got %q", test.msg, res)
			}
		})
	}
}

func TestErrorCorrectedMessage2(t *testing.T) {

	tests := []struct {
		in  string
		msg string
	}{
		{
			in:  inputTest,
			msg: "advent",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := ErrorCorrectedMessage2(test.in)
			if res != test.msg {
				t.Fatalf("ecm: expect %q, got %q", test.msg, res)
			}
		})
	}
}
