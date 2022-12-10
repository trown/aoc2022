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
	cycle, pixel, register := 1, 0, 1
	for _, i := range *d {
		if i.name == "addx" {
			cycle, pixel = do_step(cycle, pixel, register)
			cycle, pixel = do_step(cycle, pixel, register)
			register += i.value
		} else {
			cycle, pixel = do_step(cycle, pixel, register)
		}
	}
}

func do_step(cycle int, pixel int, register int) (int, int) {
	fmt.Print(get_pixel(pixel, register))
	pixel += 1
	if cycle%40 == 0 {
		fmt.Println()
		pixel = 0
	}
	cycle += 1
	return cycle, pixel
}

func get_pixel(p int, r int) string {
	if p >= r-1 && p <= r+1 {
		return "#"
	} else {
		return "."
	}
}

type data []instruction

type instruction struct {
	name  string
	value int
}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data

	s := bufio.NewScanner(f)
	for s.Scan() {
		var i instruction
		line := strings.Split(s.Text(), " ")
		if line[0] == "noop" {
			i.name = "noop"
			i.value = 0
		} else {
			i.name = "addx"
			i.value, _ = strconv.Atoi(line[1])
		}
		d = append(d, i)

	}
	return &d, s.Err()
}
