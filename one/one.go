package one

import (
	"aoc/util"
	"log"
)

/*
On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.
*/

/*
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
*/

var lookupStr = map[string]int{
	// 3 long
	"one": 1,
	"two": 2,
	"six": 6,

	// 4 long
	"zero": 0,
	"four": 4,
	"five": 5,
	"nine": 9,

	// 5 long
	"three": 3,
	"seven": 7,
	"eight": 8,
}

var lookupNum = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

var Example = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

func Solve() (sum int) {
	input, err := util.ReadLines("one/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	nums := make([][]int, 0, len(input))

	for _, problem := range input {
		current := []int{}
		for runeI := 0; runeI < len(problem); runeI++ {

			if v, ok := lookupNum[problem[runeI]]; ok {
				current = append(current, v)
				continue
			}

			switch l := runeI - len(problem); l {
			case 5:
				str := problem[runeI : runeI+5]
				if v, ok := lookupStr[str]; ok {
					current = append(current, v)
				}
				fallthrough
			case 4:
				str := problem[runeI : runeI+4]
				if v, ok := lookupStr[str]; ok {
					current = append(current, v)
				}
				fallthrough
			case 3:
				str := problem[runeI : runeI+3]
				if v, ok := lookupStr[str]; ok {
					current = append(current, v)
				}
			}
		}

		nums = append(nums, current)
	}

	for _, n := range nums {
		sum += n[0]*10 + n[len(n)-1]
	}

	return sum
}
