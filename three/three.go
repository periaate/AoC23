package three

import (
	"aoc/lib/util"
	"fmt"
)

const Fp = "three/input.txt"
const Ex = "three/example.txt"

type check struct {
	num       []rune
	starI     [][2]int
	hasSymbol bool
	hasStar   bool
}

var partTwo map[int]map[int][]int

func Solve(input []string) (sum int) {
	partTwo = make(map[int]map[int][]int, len(input))

	for i := range input {
		partTwo[i] = make(map[int][]int)
		for j := len(input[0]); j > 0; j-- {
			partTwo[i][j] = make([]int, 0)
		}
	}

	c := check{[]rune{}, make([][2]int, 0), false, false}

	checkState := func() {
		if len(c.num) != 0 {
			if c.hasSymbol {
				num := util.ToInt(string(c.num))
				sum += num

				if c.hasStar {
					for _, i := range c.starI {
						partTwo[i[0]][i[1]] = append(partTwo[i[0]][i[1]], num)
					}
				}
			}

			c = check{[]rune{}, make([][2]int, 0), false, false}
		}
	}

	for h, item := range input {
		for w, r := range item {
			if util.IsDigit(r) {
				hasSymbol, hasStar, starIndexes := neighborHasSymbols(input, h, w)
				if hasSymbol {
					c.hasSymbol = true
				}
				if hasStar {
					if c.hasStar {
						for _, v := range c.starI {
							if v[0] != starIndexes[0] && v[1] != starIndexes[1] {
								c.starI = append(c.starI, starIndexes)
							}
						}
					} else {
						c.starI = append(c.starI, starIndexes)
						c.hasStar = true
					}
				}
				c.num = append(c.num, r)
			} else {
				checkState()
			}
		}

		checkState()
	}

	checkState()

	product := 0
	for _, star := range partTwo {
		for _, list := range star {

			if len(list) != 0 {
				fmt.Println(list)
			}
			if len(list) == 2 {
				fmt.Println(list[0], list[1])
				product += list[0] * list[1]
			}
		}
	}

	fmt.Println(product)

	return
}

func neighborHasSymbols(s []string, h, w int) (has bool, hasStar bool, star [2]int) {
	match := map[byte]bool{
		'=': true,
		'*': true,
		'#': true,
		'$': true,
		'@': true,
		'&': true,
		'/': true,
		'-': true,
		'+': true,
		'%': true,
	}

	sh := len(s)    // number of lines
	sw := len(s[0]) // width of each line

	var (
		N = (h - 1)
		S = (h + 1)
		W = (w - 1)
		E = (w + 1)
	)

	if N < 0 {
		N = 0
	}
	if W < 0 {
		W = 0
	}
	if S >= sh {
		S = sh - 1
	}
	if E >= sw {
		E = sw - 1
	}

	// fmt.Printf("	%v	\n", N)
	// fmt.Printf("%v		%v\n", W, E)
	// fmt.Printf("	%v	\n", S)

	ll := []bool{
		match[s[N][w]],
		match[s[N][E]],
		match[s[h][E]],
		match[s[S][E]],
		match[s[S][w]],
		match[s[S][W]],
		match[s[h][W]],
		match[s[N][W]],
	}

	indexes := [][2]int{
		{N, w},
		{N, E},
		{h, E},
		{S, E},
		{S, w},
		{S, W},
		{h, W},
		{N, W},
	}
	star = [2]int{-1, -1}

	for _, i := range indexes {
		r := s[i[0]][i[1]]
		if rune(r) == '*' {
			// partTwo[i[0]][i[1]] = make([]int, 0)
			star = i
		}
	}

	for _, l := range ll {
		if l {
			return true, star[0] != -1, star
		}
	}

	return false, false, star
}
