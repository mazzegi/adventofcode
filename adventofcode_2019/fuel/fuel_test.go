package fuel

import (
	"fmt"
	"testing"
)

func TestByMass(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{
			mass: 12,
			fuel: 2,
		},
		{
			mass: 14,
			fuel: 2,
		},
		{
			mass: 1969,
			fuel: 654,
		},
		{
			mass: 100756,
			fuel: 33583,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("# %02d", i), func(t *testing.T) {
			f := ByMass(test.mass)
			if f != test.fuel {
				t.Fatalf("want %d, have %d", test.fuel, f)
			}
		})
	}
}

func TestByMassCorrected(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{
			mass: 12,
			fuel: 2,
		},
		{
			mass: 14,
			fuel: 2,
		},
		{
			mass: 1969,
			fuel: 966,
		},
		{
			mass: 100756,
			fuel: 50346,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("# %02d", i), func(t *testing.T) {
			f := FuelMass(test.mass)
			if f != test.fuel {
				t.Fatalf("want %d, have %d", test.fuel, f)
			}
		})
	}
}
