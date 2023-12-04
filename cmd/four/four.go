package main

import (
	"aoc/four"
	"aoc/lib/util"
	"fmt"
)

func main() {
	inputEx := util.ReadLines(four.Ex)
	outEx := four.Solve(inputEx)
	fmt.Println(outEx)

	input := util.ReadLines(four.Fp)
	out := four.Solve(input)
	fmt.Println(out)
}
