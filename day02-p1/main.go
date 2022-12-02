package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move uint8

const (
	Rock Move = iota
	Paper
	Scissors
)

func main() {
	g, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	for _, r := range *g {
		move_score := get_move_score(r.my_move)
		win_score := get_win_score(r.opp_move, r.my_move)
		result += move_score + win_score
	}
	fmt.Println(result)
}

func get_move_score(m Move) int {
	switch m {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

func get_win_score(o Move, m Move) int {
	switch o {
	case Rock:
		switch m {
		case Rock:
			return 3
		case Paper:
			return 6
		case Scissors:
			return 0
		}
	case Paper:
		switch m {
		case Rock:
			return 0
		case Paper:
			return 3
		case Scissors:
			return 6
		}
	case Scissors:
		switch m {
		case Rock:
			return 6
		case Paper:
			return 0
		case Scissors:
			return 3
		}
	}
	return 0
}

type round struct {
	opp_move Move
	my_move  Move
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
			r.my_move = Rock
		case "Y":
			r.my_move = Paper
		case "Z":
			r.my_move = Scissors
		}

		g = append(g, r)
	}
	return &g, s.Err()
}
