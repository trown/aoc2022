package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move uint8
type WLD uint8

const (
	Rock Move = iota
	Paper
	Scissors
)

const (
	Win WLD = iota
	Lose
	Draw
)

func main() {
	g, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	for _, r := range *g {
		win_score := get_win_score(r.opp_move, r.result)
		result += win_score
	}
	fmt.Println(result)
}

func get_win_score(o Move, m WLD) int {
	switch o {
	case Rock:
		switch m {
		case Win:
			return 8
		case Lose:
			return 3
		case Draw:
			return 4
		}
	case Paper:
		switch m {
		case Win:
			return 9
		case Lose:
			return 1
		case Draw:
			return 5
		}
	case Scissors:
		switch m {
		case Win:
			return 7
		case Lose:
			return 2
		case Draw:
			return 6
		}
	}
	return 0
}

type round struct {
	opp_move Move
	result   WLD
}

type guide []round

func readInput(path string) (*guide, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var g guide
	var r round

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := strings.Split(s.Text(), " ")
		opp := t[0]
		me := t[1]

		switch opp {
		case "A":
			r.opp_move = Rock
		case "B":
			r.opp_move = Paper
		case "C":
			r.opp_move = Scissors
		}

		switch me {
		case "X":
			r.result = Lose
		case "Y":
			r.result = Draw
		case "Z":
			r.result = Win
		}

		g = append(g, r)
	}
	return &g, s.Err()
}
