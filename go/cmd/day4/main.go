package main

import (
	"aoc-2024-go/day4"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE of %v:\n\t%v <path to input txt file>\n", os.Args[0], os.Args[0])
		return
	}
	var input = os.Args[1]

	// part one
	var resultOne = day4.SolvePartOne(input)
	fmt.Printf("[Part 1] Number of found XMAS words: %d\n", resultOne)

	// part two
	var resultTwo = day4.SolvePartTwo(input)
	fmt.Printf("[Part 2] Number of found X-MAS: %d\n", resultTwo)
}
