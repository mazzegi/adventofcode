package pki

import "testing"

func TestCalc(t *testing.T) {
	var subj int

	subj = 7
	loopSize1 := 8
	pk1 := Calc(subj, loopSize1)
	t.Logf("%d, %d => %d", subj, loopSize1, pk1)

	subj = 7
	loopSize2 := 11
	pk2 := Calc(subj, loopSize2)
	t.Logf("%d, %d => %d", subj, loopSize2, pk2)

	ek1 := Calc(pk2, loopSize1)
	t.Logf("ek1: %d", ek1)
	ek2 := Calc(pk1, loopSize2)
	t.Logf("ek2: %d", ek2)
}

func TestFindLoopSize(t *testing.T) {
	subj := 7
	pk1 := 5764801
	pk2 := 17807724

	ls1, ok := FindLoopSizeForSubj(subj, pk1)
	if !ok {
		t.Fatalf("not-found ls1")
	}
	t.Logf("ls1: %d", ls1)

	ls2, ok := FindLoopSizeForSubj(subj, pk2)
	if !ok {
		t.Fatalf("not-found ls2")
	}
	t.Logf("ls2: %d", ls2)
}

func TestFindLoopSize2(t *testing.T) {
	subj := 7
	pk1 := 9033205
	pk2 := 9281649

	ls1, ok := FindLoopSizeForSubj(subj, pk1)
	if !ok {
		t.Fatalf("not-found ls1")
	}
	t.Logf("ls1: %d", ls1)

	ls2, ok := FindLoopSizeForSubj(subj, pk2)
	if !ok {
		t.Fatalf("not-found ls2")
	}
	t.Logf("ls2: %d", ls2)
}

func TestFindLoopSizeAndSubj(t *testing.T) {
	// pk1 := 5764801
	// pk2 := 17807724
	pk1 := 9033205
	pk2 := 9281649

	ls1, sub1, ok := FindLoopSizeAndSubj(pk1)
	if !ok {
		t.Fatalf("not-found pk1")
	}
	t.Logf("ls1: %d, sub1: %d", ls1, sub1)

	ls2, sub2, ok := FindLoopSizeAndSubj(pk2)
	if !ok {
		t.Fatalf("not-found pk2")
	}
	t.Logf("ls2: %d, sub2: %d", ls2, sub2)
}
