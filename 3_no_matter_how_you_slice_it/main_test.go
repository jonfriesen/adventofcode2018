package main

import (
	"testing"
)

func TestFabric_addClaim(t *testing.T) {
	f := NewFabric(3, 3)

	// Single box
	// 1 1 0
	// 1 1 0
	// 0 0 0

	b1 := FabricClaim{
		id:     1,
		coordX: 0,
		coordY: 0,
		width:  2,
		height: 2,
	}

	f.addClaim(&b1)

	b1Expected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 0}}
	for x := range f.grid {
		for y := range f.grid[x] {
			if f.grid[x][y] != b1Expected[x][y] {
				t.Errorf("Add claim failed, expected sequence didn't exist, wanted %v got %v", b1Expected, f.grid)
			}
		}
	}

	// Two boxes
	// Overlap @ 2,2
	// 1 1 0
	// 1 2 1
	// 0 1 1

	b2 := FabricClaim{
		id:     2,
		coordX: 1,
		coordY: 1,
		width:  2,
		height: 2,
	}

	f.addClaim(&b2)

	b1b2Expected := [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}}
	for x := range f.grid {
		for y := range f.grid[x] {
			if f.grid[x][y] != b1b2Expected[x][y] {
				t.Errorf("Add claim failed, expected sequence didn't exist, wanted %v got %v", b1Expected, f.grid)
			}
		}
	}

}

func TestFabric_sumOverlaps(t *testing.T) {
	f := NewFabric(3, 3)

	b1 := FabricClaim{
		id:     1,
		coordX: 0,
		coordY: 0,
		width:  2,
		height: 2,
	}

	f.addClaim(&b1)

	b2 := FabricClaim{
		id:     2,
		coordX: 1,
		coordY: 1,
		width:  2,
		height: 2,
	}

	f.addClaim(&b2)

	if f.sumOverlaps() != 1 {
		t.Error("Did not find the expected 1 sum")
	}
}
