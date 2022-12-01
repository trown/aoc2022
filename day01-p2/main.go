package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	elfTroop, err := readInput("input.txt")
	var totals []int
	for _, e := range *elfTroop {
		totals = append(totals, sum(e))
	}
	sort.Ints(totals)
	fmt.Println(sum(totals[len(totals)-3:]))
	if err != nil {
		log.Fatal(err)
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

type elf []int
type elves []elf

func readInput(path string) (*elves, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var e elf
	var elfTroop elves

	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() != "" {
			i, err := strconv.Atoi(s.Text())
			if err != nil {
				return nil, fmt.Errorf("could not convert line to int. line: %s, err: %v", s.Text(), err)
			}
			e = append(e, i)
		} else {
			elfTroop = append(elfTroop, e)
			e = nil
		}
	}
	elfTroop = append(elfTroop, e)
	return &elfTroop, s.Err()
}
