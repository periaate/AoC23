package main

import (
	"aoc/lib/arrays"
	"aoc/lib/util"
	"fmt"
	"strings"
)

const Ex = "six/example.txt"
const Fp = "six/input.txt"

func main() { fmt.Println(Solve2(util.ReadLines(Fp))) }

func Solve(input []string) (prod int) {
	tS := arrays.SplitWithAll(input[0], []string{":", " "})[1:]
	dS := arrays.SplitWithAll(input[1], []string{":", " "})[1:]

	t := make([]int, len(tS))
	d := make([]int, len(tS))
	for i := range tS {
		t[i] = util.ToInt(tS[i])
		d[i] = util.ToInt(dS[i])
	}

	// tValues := make([]int, len(t))

	for j, time := range t {
		timeDistances := make([]int, 0)
		for i := 1; i <= time; i++ {
			cv := (time - i) * i
			if cv > d[j] {
				fmt.Println(cv)
				timeDistances = append(timeDistances, cv)
			}
		}
		fmt.Println("===	", j, len(timeDistances))
		if prod == 0 {
			prod += len(timeDistances)
			continue
		}
		prod *= len(timeDistances)
	}

	return prod
}

func Solve2(input []string) (prod int) {
	tS := arrays.SplitWithAll(input[0], []string{":", " "})[1:]
	dS := arrays.SplitWithAll(input[1], []string{":", " "})[1:]

	t := util.ToInt(strings.Join(tS, ""))
	d := util.ToInt(strings.Join(dS, ""))

	// tValues := make([]int, len(t))

	for i := 1; i <= t; i++ {
		cv := (t - i) * i
		if cv > d {
			prod++
		}
	}

	return
}
