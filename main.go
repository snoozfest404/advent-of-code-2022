package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/snoozfest404/advent-of-code-2022/day01"
)

func main() {
	dataPath := filepath.Join("data", "day01.txt")
	dataBytes, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}

	input := string(dataBytes)
	solution1 := day01.Part1(input)
	solution2 := day01.Part2(input)

	fmt.Printf("Part 1 solution: %v\n", solution1)
	fmt.Printf("Part 2 solution: %v\n", solution2)
}
