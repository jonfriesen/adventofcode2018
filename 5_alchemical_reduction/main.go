package main

import (
	"fmt"
	"strings"

	"github.com/jonfriesen/adventofcode2018/util"
)

func main() {
	var polymer string
	util.LoadInputFromPath("input", func(line string) {
		polymer = line
	})

	processedPolymer := processPolymer(polymer)

	fmt.Println(processedPolymer)
	fmt.Println("Polymer Count:", len(processedPolymer))

	count, unit := findMostInstrusiveUnit(polymer)
	fmt.Println("Most intrusive unit:", string(unit), "with a count of", count)
}

func processPolymer(p string) string {
	s1 := 0
	s2 := 1

	for s2 < len(p) {
		if unitsReact(rune(p[s1]), rune(p[s2])) {
			p = removeRangeFromString(p, s1, s2)
			if s1 > 0 {
				s1--
				s2--
			}
		} else {
			s1++
			s2++
		}
	}

	return p
}

func removeRangeFromString(s string, st, e int) string {
	r := []rune(s)
	return string(append(r[:st], r[e+1:]...))
}

func unitsReact(a, b rune) bool {
	return runesAreSameChar(a, b) && ((a >= 97 && b < 97) || (b >= 97 && a < 97))
}

func runesAreSameChar(a, b rune) bool {
	return toUpper(a) == toUpper(b)
}

func toUpper(r rune) rune {
	if r >= 97 {
		r -= 32
	}
	return r
}

func toLower(r rune) rune {
	if r < 97 {
		r += 32
	}
	return r
}

func removeUnit(s string, l rune) string {
	s = strings.Replace(s, string(toUpper(l)), "", -1)
	s = strings.Replace(s, string(toLower(l)), "", -1)

	return s
}

func findMostInstrusiveUnit(p string) (int, int) {

	mostInstrusiveUnit := -1
	processedPolymerLength := -1
	for i := 65; i < 91; i++ {
		fmt.Println("Generating for ", string(rune(i)))
		pl := len(processPolymer(removeUnit(p, rune(i))))
		if mostInstrusiveUnit == -1 {
			mostInstrusiveUnit = i
			processedPolymerLength = pl
		}
		if pl < processedPolymerLength {
			processedPolymerLength = pl
			mostInstrusiveUnit = i
		}
	}

	return processedPolymerLength, mostInstrusiveUnit
}
