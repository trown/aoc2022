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
	rope := make([]coordinate, 2)
	for r := range rope {
		rope[r] = coordinate{X: 0, Y: 0}
	}
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%#v", rope[len(rope)-1])] = true
	for _, i := range *d {
		//fmt.Println(i)
		for j := 0; j < i.Distance; j++ {
			//fmt.Println(rope)
			rope = process_rope(i.Direction, rope)
			//fmt.Println(rope)
			visited[fmt.Sprintf("%#v", rope[len(rope)-1])] = true
		}
		//fmt.Println(rope)
	}
	fmt.Println(len(visited))
	//fmt.Println(visited)
}

func process_rope(d string, rope []coordinate) []coordinate {
	rope[0] = move_head(d, rope[0])
	for i := 0; i < len(rope)-1; i++ {
		//fmt.Println("processing rope:", rope)
		rope[i+1] = move_tail(rope[i], rope[i+1])
		//fmt.Println("processed rope:", rope)
	}
	return rope
}
func move_tail(head coordinate, tail coordinate) coordinate {
	if head.X-tail.X > 1 {
		tail = coordinate{X: head.X - 1, Y: head.Y}
	}
	if head.X-tail.X < -1 {
		tail = coordinate{X: head.X + 1, Y: head.Y}
	}
	if head.Y-tail.Y > 1 {
		tail = coordinate{X: head.X, Y: head.Y - 1}
	}
	if head.Y-tail.Y < -1 {
		tail = coordinate{X: head.X, Y: head.Y + 1}
	}
	return tail
}
func move_head(d string, head coordinate) coordinate {
	switch d {
	case "R":
		head = coordinate{X: head.X + 1, Y: head.Y}
	case "L":
		head = coordinate{X: head.X - 1, Y: head.Y}
	case "U":
		head = coordinate{X: head.X, Y: head.Y + 1}
	case "D":
		head = coordinate{X: head.X, Y: head.Y - 1}
	}
	return head
}

type data []instruction
type coordinate struct {
	X int
	Y int
}
type instruction struct {
	Direction string
	Distance  int
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
		dir_dist := strings.Split(s.Text(), " ")
		dist, _ := strconv.Atoi(dir_dist[1])
		d = append(d, instruction{Direction: dir_dist[0], Distance: dist})
	}
	return &d, s.Err()
}
