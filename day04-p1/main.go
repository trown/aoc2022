package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	for _, i := range *d {
		if contains_range(i.elf1, i.elf2) || contains_range(i.elf2, i.elf1) {
			result += 1
		}
	}
	fmt.Println(result)
}

func contains_element(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func contains_range(s []int, r []int) bool {
	for _, a := range s {
		if !contains_element(r, a) {
			return false
		}

	}
	return true
}

type data []elfPair
type elfPair struct {
	elf1 []int
	elf2 []int
}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data
	var ep elfPair

	s := bufio.NewScanner(f)
	for s.Scan() {
		elves := strings.Split(s.Text(), ",")
		ep.elf1 = ElfToRangeInt(elves[0])
		ep.elf2 = ElfToRangeInt(elves[1])
		d = append(d, ep)
	}
	return &d, s.Err()
}

func ElfToRangeInt(e string) []int {
	nums := strings.Split(e, "-")
	begin, _ := strconv.Atoi(nums[0])
	end, _ := strconv.Atoi(nums[1])
	return RangeInt(begin, end)
}

func RangeInt(begin int, end int) (sequence []int) {
	count := end - begin + 1

	sequence = make([]int, count)
	for i := 0; i < count; i, begin = i+1, begin+1 {
		sequence[i] = begin
	}
	return sequence
}
