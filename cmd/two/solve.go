package main

import (
	"aoc/lib/util"
	"aoc/two"
	"fmt"
)

func main() {
	res := two.Solve(util.ReadLines(two.Fp))
	fmt.Println(res)
}
