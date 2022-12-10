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
	result, cycle, register := 0, 1, 1
	for _, i := range *d {
		if i.name == "addx" {
			result += do_step(cycle, register)
			cycle += 1
			result += do_step(cycle, register)
			cycle += 1
			register += i.value
		} else {
			result += do_step(cycle, register)
			cycle += 1
		}
	}
	fmt.Println(result)
}

func do_step(cycle int, register int) int {
	if (cycle-20)%40 == 0 {
		return register * cycle
	}
	return 0
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
