package ops

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseInt(s string) (int, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	return int(n), err
}

func Eval(expr string) (value int, numChars int, err error) {
	expr = strings.ReplaceAll(expr, " ", "")
	// fmt.Printf("eval %q\n", expr)
	// defer func() {
	// 	fmt.Printf("res=%d\n", value)
	// }()

	var op rune
	wantNumber := true
	for i := 0; i < len(expr); i++ {
		r := rune(expr[i])
		//fmt.Printf("eval %d: %q\n", i, string(r))
		numChars++
		if r == ')' {
			return
		}
		if wantNumber {
			var n int
			if r == '(' {
				subExpr := expr[i+1:]
				var cnt int
				n, cnt, err = Eval(subExpr)
				if err != nil {
					return
				}
				i += cnt
				numChars += cnt
			} else {
				n, err = parseInt(string(r))
				if err != nil {
					return
				}
			}
			switch op {
			case '+':
				value += n
			case '*':
				value *= n
			default:
				value = n
			}
			wantNumber = false
		} else {
			//want op
			switch r {
			case '+', '*':
				op = r
			default:
				err = errors.Errorf("want number - got %q", string(r))
				return
			}
			wantNumber = true
		}
	}
	return
}

type Token struct {
	Op    rune
	Value int
}

func (t Token) String() string {
	return fmt.Sprintf("%s%d", string(t.Op), t.Value)
}

type Stack []Token

func (s Stack) String() string {
	var sl []string
	for _, t := range s {
		sl = append(sl, t.String())
	}
	return strings.Join(sl, " ")
}

//+1 +2 *3 +4 *5 +6
func (s Stack) Eval() int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return s[0].Value
	}

	var mstack Stack
	curr := s[0].Value
	for i := 1; i < len(s); i++ {
		t := s[i]
		if t.Op == '+' {
			curr += t.Value
		} else { // op == '*'
			mstack = append(mstack, Token{
				Op:    '*',
				Value: curr,
			})
			curr = t.Value
		}
	}
	mstack = append(mstack, Token{
		Op:    '*',
		Value: curr,
	})

	if len(mstack) == 0 {
		return 0
	} else if len(mstack) == 1 {
		return mstack[0].Value
	}
	res := mstack[0].Value
	for i := 1; i < len(mstack); i++ {
		res *= mstack[i].Value
	}
	return res
}

func EvalPrecedence(expr string) (value int, numChars int, err error) {
	expr = strings.ReplaceAll(expr, " ", "")
	// fmt.Printf("eval %q\n", expr)
	// defer func() {
	// 	fmt.Printf("res=%d\n", value)
	// }()
	var stack Stack
	var op rune = '+'
	defer func() {
		//fmt.Printf("%q => [%s]\n", expr, stack.String())
		value = stack.Eval()
	}()

	wantNumber := true
	for i := 0; i < len(expr); i++ {
		r := rune(expr[i])
		//fmt.Printf("eval %d: %q\n", i, string(r))
		numChars++
		if r == ')' {
			return
		}
		if wantNumber {
			var n int
			if r == '(' {
				subExpr := expr[i+1:]
				var cnt int
				n, cnt, err = EvalPrecedence(subExpr)
				if err != nil {
					return
				}
				i += cnt
				numChars += cnt
			} else {
				n, err = parseInt(string(r))
				if err != nil {
					return
				}
			}
			stack = append(stack, Token{
				Op:    op,
				Value: n,
			})
			wantNumber = false
		} else {
			//want op
			switch r {
			case '+', '*':
				op = r
			default:
				err = errors.Errorf("want number - got %q", string(r))
				return
			}
			wantNumber = true
		}
	}
	return
}
