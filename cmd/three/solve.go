package main

import (
	"aoc/lib/util"
	"aoc/three"
	"fmt"
)

func main() {
	input := util.ReadLines(three.Ex)
	res1 := three.Solve(input)

	input = util.ReadLines(three.Fp)
	res2 := three.Solve(input)
	fmt.Println(res1, res2)
}
