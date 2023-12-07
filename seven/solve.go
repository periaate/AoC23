package main

import (
	"aoc/lib/arrays"
	"aoc/lib/util"
	"fmt"
	"sort"
)

const Ex = "seven/example.txt"
const CorrectEx = 6440
const CorrectEx2 = 5905
const Fp = "seven/input.txt"

func main() {

	input := util.ReadLines(Fp)
	res := Solve(input)
	fmt.Println(res)
	fmt.Println(Solve(util.ReadLines(Ex)))
	fmt.Println(CorrectEx2 == Solve(util.ReadLines(Ex)))
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

var Strength = map[rune]int{
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
	for _, hand := range input {
		s := arrays.SplitWithAll(hand, []string{" "})
		hand := score(s[0])
		hand.Cards = s[0]
		hand.Value = util.ToInt(s[1])
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Type == hands[j].Type {
			k := 0
			for {
				a := rune(hands[i].Cards[k])
				b := rune(hands[j].Cards[k])
				if a != b {
					return Strength[a] < Strength[b]
				}
				k++
			}
		}
		return hands[i].Type < hands[j].Type
	})

	for i, hand := range hands {
		fmt.Println(hand)
		res += (i + 1) * hand.Value
	}
	return
}

func score(str string) (h hand) {
	var jCount int
	labels := map[rune]int{}
	for _, r := range str {
		labels[r]++
		if r == 'J' {
			jCount++
		}
	}

	if jCount == 5 || jCount == 4 {
		h.Type = FiveKind
		return h
	}
	counted := false
	sett := map[int]bool{}

	switch len(labels) {
	case 5: // Must be high card
		for k, v := range labels {
			if !sett[v] {
				sett[v] = true
				h.Type = HighCard
			}
			if k != 'J' && !counted {
				h.Type += jCount
				counted = true
			}
		}
	case 4: // Must be one pair
		for k, v := range labels {
			if !sett[v] {
				sett[v] = true
				h.Type = OnePair
			}
			if k != 'J' && !counted {
				h.Type += jCount
				counted = true
			}
		}
	case 3: // Must be two pair, three kind
		for k, v := range labels {
			if v == 2 {
				if !sett[v] {
					sett[v] = true
					h.Type = TwoPair
				}
				if k != 'J' && !counted {
					h.Type += jCount
					counted = true
				}
			}
			if v == 3 {
				if !sett[v] {
					sett[v] = true
					h.Type = ThreeKind
				}
				if k != 'J' && !counted {
					h.Type += jCount
					counted = true
				}
			}
		}
	case 2: // Must be full house, four kind
		for k, v := range labels {
			if v == 4 {
				if !sett[v] {
					sett[v] = true
					h.Type = FourKind
				}
				if k != 'J' && !counted {
					h.Type += jCount
				}
			}
			if v == 3 {
				if !sett[v] {
					sett[v] = true
					h.Type = FullHouse
				}
				if k != 'J' && !counted {
					h.Type += jCount
					counted = true
				}
			}
		}
	case 1: // Must be Five kind
		h.Type = FiveKind
	}

	return h
}
