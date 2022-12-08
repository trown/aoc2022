package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	visible := 0
	for i := 1; i <= len(d)-2; i = i + 1 {
		for j := 1; j <= len(d)-2; j = j + 1 {
			t := Tuple{i: i, j: j}
			if is_visible(t, d) {
				visible += 1
			}
		}
	}
	fmt.Println(len(d)*4 - 4 + visible)
}

func is_visible(t Tuple, d [][]int) bool {
	row := d[t.i]
	left := row[0:t.j]
	right := row[t.j+1:]
	if d[t.i][t.j] > max(left) {
		return true
	}
	if d[t.i][t.j] > max(right) {
		return true
	}

	col := get_col(t.j, d)
	up := col[0:t.i]
	down := col[t.i+1:]
	if d[t.i][t.j] > max(up) {
		return true
	}
	if d[t.i][t.j] > max(down) {
		return true
	}
	return false
}

type Tuple struct {
	i int
	j int
}

func get_col(j int, d [][]int) []int {
	var col []int
	for _, row := range d {
		col = append(col, row[j])
	}
	return col
}

func max(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func readInput(path string) ([][]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d [][]int

	s := bufio.NewScanner(f)
	for s.Scan() {
		var row []int
		for _, r := range s.Text() {
			val, _ := strconv.Atoi(string(r))
			row = append(row, val)
		}
		d = append(d, row)

	}
	return d, s.Err()
}
