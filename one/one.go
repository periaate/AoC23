package one

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

const Fp = "one/input.txt"

// Solve is my best attempt at optimizing this algorithm.
// Single core performance on my machine gets to ~80ns per item
// Testing with various multithreading approaches, best I've gotten is ~6.2ns per item locally on 24 threads.
// Your mileage may vary.
func Solve(input []string) (sum int) {
	for _, problem := range input {
		// 0 length, 16 size, allowing for appending without resizing.
		// Using uint8 here improves performance by ~8%, likely due to cache locality.
		// For whatever reason any size other than 16 seems to lead to degraded performance, why?
		n := make([]uint8, 0, 16)
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
					runeI += 1 // Skipping only 1, as the `e` may be used by `eight`.
					continue
				case "two":
					n = append(n, 2)
					runeI += 1 // Skipping only 1, as the `o` may be used by `one`.
					continue
				case "six":
					n = append(n, 6)
					runeI += 2 // Skipping 2, as none start with `x`. Other skips use same logic.
					continue
				}
			}

			if l >= 4 {
				str := problem[runeI : runeI+4]
				switch str {
				case "zero":
					n = append(n, 0)
					runeI += 2
					continue
				case "four":
					n = append(n, 4)
					runeI += 3
					continue
				case "five":
					n = append(n, 5)
					runeI += 2
					continue
				case "nine":
					n = append(n, 9)
					runeI += 2
					continue
				}
			}

			if l >= 5 {
				str := problem[runeI : runeI+5]
				switch str {
				case "three":
					n = append(n, 3)
					runeI += 3
					continue
				case "seven":
					n = append(n, 7)
					runeI += 3
					continue
				case "eight":
					n = append(n, 8)
					runeI += 3
					continue
				}
			}
		}

		sum += int(n[0]*10 + n[len(n)-1])
	}

	return sum
}
