package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	bc := readListToCollection(f)

	fmt.Println("Checksum", bc.calculateChecksum())
	fmt.Println("Matching ID String", bc.findSingleDiffBox())
}

type BoxID string

type BoxCollection struct {
	values []BoxID
}

type BoxMetaCount struct {
	hasTwo   bool
	hasThree bool
}

func (bc *BoxCollection) calculateChecksum() int {
	three := 0
	two := 0

	for _, v := range bc.values {
		bmc := v.countDuplicates()
		if bmc.hasTwo {
			two++
		}
		if bmc.hasThree {
			three++
		}
	}

	return two * three
}

func (b *BoxID) countDuplicates() BoxMetaCount {
	d := BoxMetaCount{}
	m := make(map[rune]int)
	bs := string(*b)
	for _, l := range bs {
		_, e := m[l]
		if !e {
			m[l] = 1
		} else {
			m[l]++
		}
	}

	for _, v := range m {
		if v == 2 {
			d.hasTwo = true
		}
		if v == 3 {
			d.hasThree = true
		}
	}

	return d
}

func (bc *BoxCollection) findSingleDiffBox() string {
	for _, b1 := range bc.values {
		for _, b2 := range bc.values {
			// skip identicals
			if b1 == b2 {
				continue
			}
			m, i := b1.checkDiffChars(&b2)
			if m {
				bs := string(b1)
				return bs[:i] + bs[i+1:]
			}
		}
	}
	return "ERROR"
}

func (b *BoxID) checkDiffChars(c *BoxID) (bool, int) {
	bs := string(*b)
	cs := string(*c)

	dupeAlreadyFound := false
	diffIndex := 0
	for i, v := range bs {
		if v != rune(cs[i]) {
			if dupeAlreadyFound {
				return false, 0
			} else {
				dupeAlreadyFound = true
				diffIndex = i
			}
		}
	}

	return true, diffIndex
}

func readListToCollection(r io.Reader) BoxCollection {
	f := []BoxID{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		f = append(f, BoxID(s.Text()))
	}

	return BoxCollection{values: f}
}
