package four

import (
	"aoc/lib/arrays"
	"fmt"
)

const Ex = "four/example.txt"
const Fp = "four/input.txt"

func Solve(input []string) (sum int) {

	r := arrays.SplitWithAll(input[0], []string{" "})
	p1Length := 0

	for j, el := range r {
		if el == "|" {
			p1Length = j
			fmt.Println(j)
			break
		}
	}

	copies := make([]int, len(input))
	wins := make([]int, len(input))

	for i, line := range input {
		prod := 0
		c := 0
		r := arrays.SplitWithAll(line, []string{":", "|", " "})
		winning := r[2:p1Length]
		having := r[p1Length:]

		wMap := make(map[string]string, len(winning))
		for _, el := range winning {
			wMap[el] = el
		}

		hMap := make(map[string]string, len(having))
		for _, el := range having {
			hMap[el] = el
		}

		j := 0
		copies[i]++
		for el := range wMap {
			if res, ok := hMap[el]; ok {
				if el == res {
					prod = 1 << c
					c += 1
					wins[i]++
				}
			}
			j++
		}

		for l := copies[i]; l > 0; l-- {
			for k := wins[i]; k > 0; k-- {
				j := i + k
				copies[j] += 1
			}
		}

		sum += prod

	}

	totalCards := 0
	for _, copy := range copies {
		totalCards += copy
	}
	fmt.Println(totalCards)
	return sum
}
