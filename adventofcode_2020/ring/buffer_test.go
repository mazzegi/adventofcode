package ring

import "testing"

func TestAppend(t *testing.T) {
	b := NewBuffer(nil)

	b.Append(1)
	t.Logf("%s", b.String())
	b.Append(2)
	t.Logf("%s", b.String())
	b.Append(3)
	t.Logf("%s", b.String())
	b.Append(4)
	t.Logf("%s", b.String())
}
