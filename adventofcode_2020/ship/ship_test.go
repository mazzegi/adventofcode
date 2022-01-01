package ship

import "testing"

func TestDir(t *testing.T) {
	d := DirEast

	var add int
	var dd Dir

	add = 90
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = 180
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = 270
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = 360
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = -90
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = -180
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = -270
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)

	add = -360
	dd = d.Added(add)
	t.Logf("%s + %d = %s", d, add, dd)
}
