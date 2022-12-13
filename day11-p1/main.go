package main

import (
	"fmt"
)

func main() {
	m := readInput("input.txt")
	for i := 0; i < 20; i++ {
		do_round(*m)
	}
	for _, m := range *m {
		fmt.Println(m.inspected)
	}
}

func take_turn(i int, monkeys data) data {
	for _, item := range monkeys[i].items {
		// operation
		switch monkeys[i].op.op {
		case "plus":
			item = item + monkeys[i].op.value
		case "times":
			item = item * monkeys[i].op.value
		case "exp":
			item = item * item
		}
		// divide by 3
		item = item / 3
		// test
		if item%monkeys[i].check.divisor == 0 {
			monkeys[monkeys[i].check.pass].items = append(monkeys[monkeys[i].check.pass].items, item)
		} else {
			monkeys[monkeys[i].check.fail].items = append(monkeys[monkeys[i].check.fail].items, item)
		}
		monkeys[i].inspected += 1
	}
	monkeys[i].items = nil
	return monkeys
}

func do_round(monkeys data) data {
	for i := 0; i < len(monkeys); i++ {
		//fmt.Println(monkeys[i])
		monkeys = take_turn(i, monkeys)
	}
	return monkeys
}

type data []monkey
type monkey struct {
	items     []int
	op        operation
	check     test
	inspected int
}
type operation struct {
	op    string
	value int
}
type test struct {
	divisor int
	pass    int
	fail    int
}

func readInput(path string) *data {
	var m data
	m = append(m, monkey{
		items: []int{52, 60, 85, 69, 75, 75},
		op: operation{
			op:    "times",
			value: 17,
		},
		check: test{
			divisor: 13,
			pass:    6,
			fail:    7,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{96, 82, 61, 99, 82, 84, 85},
		op: operation{
			op:    "plus",
			value: 8,
		},
		check: test{
			divisor: 7,
			pass:    0,
			fail:    7,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{95, 79},
		op: operation{
			op:    "plus",
			value: 6,
		},
		check: test{
			divisor: 19,
			pass:    5,
			fail:    3,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{88, 50, 82, 65, 77},
		op: operation{
			op:    "times",
			value: 19,
		},
		check: test{
			divisor: 2,
			pass:    4,
			fail:    1,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{66, 90, 59, 90, 87, 63, 53, 88},
		op: operation{
			op:    "plus",
			value: 7,
		},
		check: test{
			divisor: 5,
			pass:    1,
			fail:    0,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{92, 75, 62},
		op: operation{
			op:    "exp",
			value: 2,
		},
		check: test{
			divisor: 3,
			pass:    3,
			fail:    4,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{94, 86, 76, 67},
		op: operation{
			op:    "plus",
			value: 1,
		},
		check: test{
			divisor: 11,
			pass:    5,
			fail:    2,
		},
		inspected: 0,
	})
	m = append(m, monkey{
		items: []int{57},
		op: operation{
			op:    "plus",
			value: 2,
		},
		check: test{
			divisor: 17,
			pass:    6,
			fail:    2,
		},
		inspected: 0,
	})
	return &m
}
