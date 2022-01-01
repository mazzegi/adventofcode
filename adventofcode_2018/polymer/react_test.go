package polymer

import "testing"

func TestDestroy(t *testing.T) {
	var r1, r2 rune

	r1 = 'r'
	r2 = 'r'
	t.Logf("%s + %s => destroys = %t", string(r1), string(r2), destroys(r1, r2))

	r1 = 'r'
	r2 = 'R'
	t.Logf("%s + %s => destroys = %t", string(r1), string(r2), destroys(r1, r2))

	r1 = 'R'
	r2 = 'r'
	t.Logf("%s + %s => destroys = %t", string(r1), string(r2), destroys(r1, r2))

	r1 = 'R'
	r2 = 'R'
	t.Logf("%s + %s => destroys = %t", string(r1), string(r2), destroys(r1, r2))
}

func TestReactOnce(t *testing.T) {
	var p string
	var r string
	var c bool

	p = "abcdDefg"
	r, c = reactOnce(p)
	t.Logf("%q => %q, %t", p, r, c)

	p = "aAbcdefg"
	r, c = reactOnce(p)
	t.Logf("%q => %q, %t", p, r, c)

	p = "abcdefGg"
	r, c = reactOnce(p)
	t.Logf("%q => %q, %t", p, r, c)

	p = "abcdefh"
	r, c = reactOnce(p)
	t.Logf("%q => %q, %t", p, r, c)

}
