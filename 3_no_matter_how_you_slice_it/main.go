package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jonfriesen/adventofcode2018/util"
)

func main() {
	fcs := []FabricClaim{}
	util.LoadInputFromPath("input", createFabricClaimFunc(&fcs))

	f := NewFabric(1000, 1000)

	for _, v := range fcs {
		f.addClaim(&v)
	}

	fmt.Println("Amount of overlapping inches: ", f.sumOverlaps())

	for _, v := range fcs {
		if f.checkClaim(&v) {
			fmt.Printf("ID# %v has no overlaps\n", v.id)
		}
	}
}

type FabricClaim struct {
	id     int
	coordX int
	coordY int
	width  int
	height int
}

type Fabric struct {
	grid [][]int
	row  int
	col  int
}

func NewFabric(row, col int) *Fabric {

	// Not a hug fan of this method...
	// might want to move this to a
	// single array where row * col
	// is the size.
	g := make([][]int, row)
	for i := range g {
		g[i] = make([]int, col)
	}

	return &Fabric{
		grid: g,
		row:  row,
		col:  col,
	}
}

func (f *Fabric) addClaim(c *FabricClaim) {
	for y := c.coordY; y < c.coordY+c.height; y++ {
		for x := c.coordX; x < c.coordX+c.width; x++ {
			f.grid[x][y]++
		}
	}
}

func (f *Fabric) checkClaim(c *FabricClaim) bool {
	for y := c.coordY; y < c.coordY+c.height; y++ {
		for x := c.coordX; x < c.coordX+c.width; x++ {
			if f.grid[x][y] > 1 {
				return false
			}
		}
	}
	return true
}

func (f *Fabric) sumOverlaps() int {
	count := 0
	for _, vx := range f.grid {
		for _, vy := range vx {
			if vy > 1 {
				count++
			}
		}
	}

	return count
}

func createFabricClaimFunc(fcs *[]FabricClaim) util.InputFunc {
	return func(line string) {
		re := regexp.MustCompile(`(?m)\d+`)

		in := re.FindAllString(line, -1)

		fc := FabricClaim{}
		fc.id, _ = strconv.Atoi(in[0])
		fc.coordX, _ = strconv.Atoi(in[1])
		fc.coordY, _ = strconv.Atoi(in[2])
		fc.width, _ = strconv.Atoi(in[3])
		fc.height, _ = strconv.Atoi(in[4])

		*fcs = append(*fcs, fc)
	}
}
