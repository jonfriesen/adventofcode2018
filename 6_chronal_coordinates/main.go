package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonfriesen/adventofcode2018/util"
)

// Resources:
// https://www.quora.com/Is-there-something-called-the-Manhattan-Algorithm
// https://en.wikipedia.org/wiki/Taxicab_geometry

func main() {
	coords := []Point{}
	getInput(&coords)

	canvas := createCanvas(coords)
	printCanvas(canvas)
}

type Point struct {
	x int
	y int
	l rune
}

func createCanvas(c []Point) [][]Point {
	max := Point{0, 0, 0}

	for _, v := range c {
		if v.x > max.x {
			max.x = v.x
		}
		if v.y > max.y {
			max.y = v.y
		}
	}

	m := make([][]Point, max.x+2)
	for i := range m {
		m[i] = make([]Point, max.y+2)
	}

	for _, v := range c {
		m[v.x][v.y] = v
	}

	return m

}

func getInput(in *[]Point) {
	label := 'A'

	util.LoadInputFromPath("sampleinput", func(line string) {
		v := strings.Split(line, ", ")
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		l := label
		label++
		*in = append(*in, Point{x, y, l})
	})
}

// puzzle is rendering incorrectly
func printCanvas(c [][]Point) {
	for x := range c {
		for y := range c[x] {
			if c[x][y].l == 0 {
				fmt.Print("🤠")
			}
			fmt.Print(string(c[x][y].l))
		}
		fmt.Print("\n")
	}
}
