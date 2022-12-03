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
	result := 0
	for _, g := range *d {
		result += calc_priority(g)
	}
	fmt.Println(result)
}

func calc_priority(g group) int {
	for _, c := range g[0] {
		if strings.ContainsRune(g[1], c) {
			if strings.ContainsRune(g[2], c) {
				return rune_priority(c)
			}
		}
	}
	return 0
}

func rune_priority(c rune) int {
	// int('a') = 97
	// int('A') = 65
	p := int(c)
	if p >= 97 {
		p = p - 96
	} else {
		p = p - 38
	}
	return p
}

type data []group
type group []string

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data
	var g group
	c := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		if c == 3 {
			c = 0
			d = append(d, g)
			g = nil
		}
		g = append(g, s.Text())
		c += 1
	}
	d = append(d, g)
	return &d, s.Err()
}
