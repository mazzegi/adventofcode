package dinner

import "testing"

func TestPerm(t *testing.T) {
	//sl := []string{"1", "2", "3"}
	sl := []string{"1", "2", "3", "4"}
	for perm := range Permutations(sl) {
		t.Logf("%v", perm)
	}
	t.Logf("done")
}
