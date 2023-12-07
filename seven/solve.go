package main

import (
	"aoc/lib/arrays"
	"aoc/lib/util"
	"fmt"
	"sort"
)

const Ex = "seven/example.txt"
const Fp = "seven/input.txt"

func main() {
	input := util.ReadLines(Fp)
	res := Solve(input)
	fmt.Println(res)
	fmt.Println(Solve(util.ReadLines(Ex)))
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

var Strength = map[uint8]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type hand struct {
	Cards string
	Type  int
	Value int
}

func Solve(input []string) (res int) {
	hands := []hand{}
	for _, h := range input {
		s := arrays.SplitWithAll(h, []string{" "})
		hand := hand{
			Type:  score(s[0]),
			Cards: s[0],
			Value: util.ToInt(s[1]),
		}
		hands = append(hands, hand)
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			for k := 0; k < len(hands[i].Cards); k++ {
				a := Strength[hands[i].Cards[k]]
				b := Strength[hands[j].Cards[k]]
				if a != b {
					return a < b
				}
			}
		}
		return hands[i].Type < hands[j].Type
	})

	for i, hand := range hands {
		res += (i + 1) * hand.Value
	}
	return res
}

var scoreMap = map[[2]int]int{
	{5, 0}: FiveKind,
	{4, 1}: FourKind,
	{3, 2}: FullHouse,
	{3, 1}: ThreeKind,
	{2, 2}: TwoPair,
	{2, 1}: OnePair,
	{1, 1}: HighCard,
}

func score(str string) int {
	var jCount int
	labels := make([]int, 'T'+1)

	for _, c := range str {
		if c != 'J' {
			labels[c]++
		} else {
			jCount++
		}
	}

	sort.SliceStable(labels, func(i, j int) bool { return labels[i] > labels[j] })
	return scoreMap[[2]int{labels[0] + jCount, labels[1]}]
}
