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
		// 0 length, 16 size, allowing for appending without resizing
		n := make([]int, 0, 16)
		for runeI := 0; runeI <= len(problem); runeI++ {
			l := len(problem) - runeI
			if l != 0 {
				r := problem[runeI]
				switch r {
				case '0':
					n = append(n, 0)
					continue
				case '1':
					n = append(n, 1)
					continue
				case '2':
					n = append(n, 2)
					continue
				case '3':
					n = append(n, 3)
					continue
				case '4':
					n = append(n, 4)
					continue
				case '5':
					n = append(n, 5)
					continue
				case '6':
					n = append(n, 6)
					continue
				case '7':
					n = append(n, 7)
					continue
				case '8':
					n = append(n, 8)
					continue
				case '9':
					n = append(n, 9)
					continue
				}
			}

			if l >= 3 {
				str := problem[runeI : runeI+3]
				switch str {
				case "one":
					n = append(n, 1)
					continue
				case "two":
					n = append(n, 2)
					continue
				case "six":
					n = append(n, 6)
					continue
				}
			}

			if l >= 4 {
				str := problem[runeI : runeI+4]
				switch str {
				case "zero":
					n = append(n, 0)
					continue
				case "four":
					n = append(n, 4)
					continue
				case "five":
					n = append(n, 5)
					continue
				case "nine":
					n = append(n, 9)
					continue
				}
			}

			if l >= 5 {
				str := problem[runeI : runeI+5]
				switch str {
				case "three":
					n = append(n, 3)
					continue
				case "seven":
					n = append(n, 7)
					continue
				case "eight":
					n = append(n, 8)
					continue
				}
			}
		}
		sum += n[0]*10 + n[len(n)-1]
	}

	return sum
}
