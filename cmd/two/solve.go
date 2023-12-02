package main

import (
	"aoc/lib/util"
	"aoc/two"
	"fmt"
)

func main() {
	res := two.SolvePartTwo(util.ReadLines(two.Fp))
	fmt.Println(res)
	res = two.SolvePartOne(util.ReadLines(two.Fp))
	fmt.Println(res)
}
