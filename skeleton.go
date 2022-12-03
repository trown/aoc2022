package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	d, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	for _, i := range *d {
		fmt.Println(i)
	}
	fmt.Println(result)
}

type data []struct{}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data

	s := bufio.NewScanner(f)
	for s.Scan() {

	}
	return &d, s.Err()
}
