package day_06

import (
	"testing"

	"github.com/mazzegi/adventofcode/adventofcode_2021/testutil"
)

func TestFishCount(t *testing.T) {
	res80, err := FishCount(inputTest, 80)
	testutil.CheckUnexpectedError(t, err)
	exp80 := 5934
	if exp80 != res80 {
		t.Fatalf("want %d, have %d", exp80, res80)
	}
}

func TestFishCountSharded(t *testing.T) {
	res80, err := FishCountSharded(inputTest, 80)
	testutil.CheckUnexpectedError(t, err)
	exp80 := 5934
	if exp80 != res80 {
		t.Fatalf("want %d, have %d", exp80, res80)
	}
}

func TestFishCount256(t *testing.T) {
	res256, err := FishCountLookup256(inputTest)
	testutil.CheckUnexpectedError(t, err)
	exp256 := 26984457539
	if exp256 != res256 {
		t.Fatalf("want %d, have %d", exp256, res256)
	}
}

func TestFishCountCalced(t *testing.T) {
	res80, err := FishCountCalced(inputTest, 18)
	testutil.CheckUnexpectedError(t, err)
	exp80 := 26
	if exp80 != res80 {
		t.Fatalf("want %d, have %d", exp80, res80)
	}
}
