package util

import (
	"bufio"
	"io"
	"os"
)

type InputFunc func(line string)

func LoadInputFromPath(file string, fn InputFunc) {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	LoadInput(f, fn)
}

func LoadInput(r io.Reader, fn InputFunc) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fn(s.Text())
	}
}
