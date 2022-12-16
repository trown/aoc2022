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
	potential_points := findAllPointsOneUnitAway(d)
	for _, p := range potential_points {
		if checkPoint(d, p) {
			fmt.Println(p)
			fmt.Println(p.x*4000000 + p.y)
		}
	}

}

func findAllPointsOneUnitAway(d *data) []Point {
	var r []Point
	for _, s := range *d {
		dist := manhattanDistance(s.location, s.closest_beacon)
		//minX->maxY
		j := s.location.y
		for i := s.location.x - dist - 1; i <= s.location.x; i++ {
			r = append(r, Point{i, j})
			j++
		}
		//maxY->maxX
		i := s.location.x
		for j := s.location.y + dist + 1; j >= s.location.y; j-- {
			r = append(r, Point{i, j})
			i++
		}
		//maxX->minY
		j = s.location.y
		for i := s.location.x + dist + 1; i >= s.location.x; i-- {
			r = append(r, Point{i, j})
			j--
		}
		//minY->minX
		i = s.location.x
		for j := s.location.y - dist - 1; j <= s.location.y; j++ {
			r = append(r, Point{i, j})
			i++
		}
	}
	return r
}

func checkPoint(d *data, p Point) bool {
	r := true
	if p.y > 4000000 || p.y < 0 || p.x > 4000000 || p.x < 0 {
		return false
	}
	for _, s := range *d {
		if manhattanDistance(s.location, p) <= manhattanDistance(s.location, s.closest_beacon) {
			r = false
			break
		}
	}
	return r
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
