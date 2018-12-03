package main

import (
	"fmt"

	"github.com/jonfriesen/adventofcode2018/util"
)

func main() {
	util.LoadInput("input", func(s string) { fmt.Println("P>", s) })
}
