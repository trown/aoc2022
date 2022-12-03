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
	for _, r := range *d {
		result += calc_priority(r)
	}
	fmt.Println(result)
}

func calc_priority(r rucksack) int {
	for _, c := range r.first {
		if strings.ContainsRune(r.second, c) {
			return rune_priority(c)
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

type data []rucksack
type rucksack struct{
	first string
	second string
}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data
	var r rucksack

	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		r.first = text[0:len(text)/2]
		r.second = text[len(text)/2:]
		d = append(d, r)
	}
	return &d, s.Err()
}
