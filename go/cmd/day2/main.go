package main

import (
	"aoc-2024-go/day2"
	"fmt"
	"os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("USAGE of %v:\n\t%v <path to input txt file>\n", os.Args[0], os.Args[0])
        return
    }

    var input = os.Args[1]

    resultPartOne := day2.SolvePartOne(input)
    fmt.Printf("[Part 1] The total of valid reports equal: %d\n", resultPartOne)

    resultPartTwo := day2.SolvePartTwo(input)
    fmt.Printf("[Part 2] The total of valid reports equal: %d\n", resultPartTwo)
}
