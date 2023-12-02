package two

import (
	"aoc/lib/arrays"
	"aoc/lib/util"
)

const (
	CorrectPartOne = 2204
	CorrectPartTwo = 71036
)

const Fp = "two/input.txt"

/*
To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).
*/

var Example = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

/*
In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
*/

/*
Part 2
- Game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
- Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
- Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
- Game 4 required at least 14 red, 3 green, and 15 blue cubes.
- Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.
*/

func SolvePartTwo(input []string) (sum int) {
	for _, item := range input {
		cubesInGame := arrays.SplitWithAll(item, []string{":", ";", ","})[1:]
		/* Logic goes:
		item := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		cubesInGame := arrays.SplitWithAll(item, []string{":", ";", ","})[1:]
		cubesInGame == ["3 blue", "4 red", "2 green", "6 blue", "2 green"]
		*/
		localMax := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, cube := range cubesInGame {
			item := arrays.SplitWithAll(cube, []string{" "})

			color := item[1]
			num := util.ToInt(item[0])

			if num > localMax[color] {
				localMax[color] = num
			}
		}

		sum += localMax["red"] * localMax["green"] * localMax["blue"]
	}

	return sum
}

func SolvePartOne(input []string) (sum int) {
	var isPossible = map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, item := range input {
		game := arrays.SplitWithAll(item, []string{":", ";", ","})

		// "Game 123..."[5:] == "123..."
		index := util.ToInt(game[0][5:])

		cubes := game[1:]
		gamePossible := true

		for _, cube := range cubes {
			item := arrays.SplitWithAll(cube, []string{" "})

			color := item[1]
			num := util.ToInt(item[0])

			if v, ok := isPossible[color]; ok && v < num {
				gamePossible = false
			}
		}

		if gamePossible {
			sum += index
		}
	}

	return sum
}
