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
	d.crate_map = [][]string{
		{"C", "F", "B", "L", "D", "P", "Z", "S"},
		{"B", "W", "H", "P", "G", "V", "N"},
		{"G", "J", "B", "W", "F"},
		{"S", "C", "W", "L", "F", "N", "J", "G"},
		{"H", "S", "M", "P", "T", "L", "J", "W"},
		{"S", "F", "G", "W", "C", "B"},
		{"W", "B", "Q", "M", "P", "T", "H"},
		{"T", "W", "S", "F"},
		{"R", "C", "N"},
	}
	fmt.Println(d.crate_map)
	for _, m := range d.move_list {
		fmt.Println(m)
		executeMove(d, m)
		fmt.Println(d.crate_map)
	}
	fmt.Println(d.crate_map)
}

func executeMove(d *data, m move) {
	var crates []string
	// Pop Front "amount" crates from the "from" colummn
	crates, d.crate_map[m.from-1] = d.crate_map[m.from-1][0:m.amount], d.crate_map[m.from-1][m.amount:]

	crates_copy := make([]string, len(crates))
	copy(crates_copy, crates)

	// Push Front "to" column
	d.crate_map[m.to-1] = append(crates_copy, d.crate_map[m.to-1]...)
}

type data struct {
	crate_map [][]string
	move_list []move
}

type move struct {
	amount int
	from   int
	to     int
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
		if strings.ContainsRune(s.Text(), '[') {
			d = addToCrateMap(s.Text(), d)
		}
		if strings.Contains(s.Text(), "move") {
			d = addToMoveList(s.Text(), d)
		}

	}
	return &d, s.Err()
}

// TODO: Gave up on this and just hardcoded the map.
func addToCrateMap(text string, d data) data {
	col1 := text[1:2]
	col2 := " "
	if len(text) > 6 {
		col2 = text[5:6]
	}
	col3 := " "
	if len(text) > 10 {
		col3 = text[9:10]
	}
	col4 := " "
	if len(text) > 14 {
		col4 = text[13:14]
	}
	col5 := " "
	if len(text) > 18 {
		col5 = text[17:18]
	}
	col6 := " "
	if len(text) > 22 {
		col6 = text[21:22]
	}
	col7 := " "
	if len(text) > 26 {
		col7 = text[25:26]
	}
	col8 := " "
	if len(text) > 30 {
		col8 = text[29:30]
	}
	col9 := " "
	if len(text) > 34 {
		col9 = text[33:34]
	}
	fmt.Println(col1, col2, col3, col4, col5, col6, col7, col8, col9)
	return d
}

func addToMoveList(text string, d data) data {
	text = strings.ReplaceAll(text, "move ", "")
	text = strings.ReplaceAll(text, "from", "")
	text = strings.ReplaceAll(text, "to", "")
	nums := strings.Split(text, " ")
	var newMove move
	newMove.amount, _ = strconv.Atoi(nums[0])
	newMove.from, _ = strconv.Atoi(nums[2])
	newMove.to, _ = strconv.Atoi(nums[4])
	d.move_list = append(d.move_list, newMove)
	return d
}
