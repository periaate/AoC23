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

var lookup = map[string]int{
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

	// 1 long
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var Example = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}
var ExampleTwo = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func GetInput() []string {
	input, err := util.ReadLines("one/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return input
}

func Solve(input []string) (sum int) {
	for _, problem := range input {
		// Initialize a 0 length array with 16 length of memory for mutationless appending.
		n := make([]int, 0, 16)
		for runeI := 0; runeI <= len(problem); runeI++ {
			l := len(problem) - runeI
			if l >= 5 {
				str := problem[runeI : runeI+5]
				if v, ok := lookup[str]; ok {
					n = append(n, v)
					continue
				}
			}
			if l >= 4 {
				str := problem[runeI : runeI+4]
				if v, ok := lookup[str]; ok {
					n = append(n, v)
					continue
				}
			}
			if l >= 3 {
				str := problem[runeI : runeI+3]
				if v, ok := lookup[str]; ok {
					n = append(n, v)
					continue
				}
			}

			if l >= 1 {
				str := string(problem[runeI])
				if v, ok := lookup[str]; ok {
					n = append(n, v)
				}
			}
		}
		sum += n[0]*10 + n[len(n)-1]
	}

	return sum
}
