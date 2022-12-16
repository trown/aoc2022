package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Point represents a point in a 2D plane
type Point struct {
	x, y int
}

// manhattanDistance returns the Manhattan distance between two points
func manhattanDistance(p1, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

// abs returns the absolute value of an integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {

	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	min_x := 0
	max_x := 0
	for _, i := range *d {
		min := findMinX(i, 2000000)
		if min < min_x {
			min_x = min
		}

		max := findMaxX(i, 2000000)
		if max > max_x {
			max_x = max
		}
	}
	fmt.Println(abs(min_x) + abs(max_x))
}

func findMinX(s Sensor, row int) int {
	d := manhattanDistance(s.location, s.closest_beacon)
	for i := (s.location.x - d); i < s.location.x+d; i++ {
		if manhattanDistance(s.location, Point{i, row}) <= d {
			return i
		}
	}
	return 0
}

func findMaxX(s Sensor, row int) int {
	d := manhattanDistance(s.location, s.closest_beacon)
	for i := (s.location.x + d); i > s.location.x-d; i-- {
		if manhattanDistance(s.location, Point{i, row}) <= d {
			return i
		}
	}
	return 0
}

type data []Sensor
type Sensor struct {
	location       Point
	closest_beacon Point
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
		line := strings.Split(s.Text(), ",")
		sensor_x, _ := strconv.Atoi(strings.Split(line[0], "=")[1])
		sensor_y, _ := strconv.Atoi(strings.Split(strings.Split(line[1], ":")[0], "=")[1])
		beacon_x, _ := strconv.Atoi(strings.Split(line[1], "=")[2])
		beacon_y, _ := strconv.Atoi(strings.Split(line[2], "=")[1])
		sensor := Sensor{location: Point{sensor_x, sensor_y}, closest_beacon: Point{beacon_x, beacon_y}}
		d = append(d, sensor)
	}
	return &d, s.Err()
}
