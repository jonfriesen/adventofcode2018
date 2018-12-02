package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const startingFrequency = 0

func main() {

	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fc := createFreqCollFromReader(f)

	fmt.Println("End frequency is: ", fc.findFrequency())
	fmt.Println("First duplicate frequency is: ", fc.findDuplicateFrequency())

}

type FrequencyValue string

type FrequencyCollection struct {
	starting int32
	values   []FrequencyValue
}

func (fc *FrequencyCollection) findFrequency() int32 {
	s := fc.starting
	for _, v := range fc.values {
		s += v.asInt32()
	}

	return s
}

func (fc *FrequencyCollection) findDuplicateFrequency() int32 {
	m := make(map[int32]bool)

	count := 0
	s := fc.starting
	for {
		count++
		for _, v := range fc.values {
			s += v.asInt32()

			v := m[s]
			if v == true {
				fmt.Println("Found on ", count, "iterations")
				return s
			}

			m[s] = true
		}
	}
}

func (iv *FrequencyValue) asInt32() int32 {
	v := string(*iv)
	n, _ := strconv.ParseInt(v, 10, 32)
	return int32(n)
}

func createFreqCollFromReader(r io.Reader) FrequencyCollection {
	f := []FrequencyValue{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		f = append(f, FrequencyValue(s.Text()))
	}

	return FrequencyCollection{starting: startingFrequency, values: f}
}
