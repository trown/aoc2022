package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i, w := range slidingWindow(4, *d) {
		var seen string
		for _, r := range w {
			if !strings.ContainsRune(seen, r) {
				seen += string(r)
			}
		}
		if len(seen) == 4 {
			fmt.Println(i + 4)
			break
		} else {
			seen = ""
		}

	}
}

// https://github.com/golang/go/wiki/SliceTricks#sliding-window
func slidingWindow(size int, input []rune) [][]rune {
	// returns the input slice as the first element
	if len(input) <= size {
		return [][]rune{input}
	}

	// allocate slice at the precise size we need
	r := make([][]rune, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}

type data []rune

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data

	s := bufio.NewScanner(f)
	for s.Scan() {
		for _, i := range s.Text() {
			d = append(d, i)
		}
	}
	return &d, s.Err()
}
