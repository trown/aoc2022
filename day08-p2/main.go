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
	max_score := 0
	for i := 1; i <= len(d)-2; i = i + 1 {
		for j := 1; j <= len(d)-2; j = j + 1 {
			t := Tuple{i: i, j: j}
			score := calc_score(t, d)
			if score > max_score {
				max_score = score
			}
		}
	}
	fmt.Println(max_score)
}

func calc_score(t Tuple, d [][]int) int {
	row := d[t.i]
	left := row[0:t.j]
	right := row[t.j+1:]

	col := get_col(t.j, d)
	up := col[0:t.i]
	down := col[t.i+1:]

	return row_score(reverse(left), d[t.i][t.j]) *
		row_score(right, d[t.i][t.j]) *
		row_score(reverse(up), d[t.i][t.j]) *
		row_score(down, d[t.i][t.j])
}

func reverse(n []int) []int {
	numbers := make([]int, len(n))
	copy(numbers, n)
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func row_score(row []int, i int) int {
	score := 0
	for _, r := range row {
		if i <= r {
			score += 1
			break
		}
		score += 1
	}
	return score
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
