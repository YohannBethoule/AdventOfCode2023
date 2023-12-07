package day7

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"A": 15,
	"K": 14,
	"Q": 13,
	"J": 12,
	"T": 11,
	"9": 10,
	"8": 9,
	"7": 8,
	"6": 7,
	"5": 6,
	"4": 5,
	"3": 4,
	"2": 3,
}

type Round struct {
	Hand    string
	HandMap map[string]int
	Bid     int
}

func Star1() {
	scan := pkg.ReadFile("./day7/input")

	var rounds []Round

	for scan.Scan() {
		line := scan.Text()
		arr := strings.Split(line, " ")
		bid, _ := strconv.Atoi(arr[1])
		rounds = append(rounds, Round{
			Hand:    arr[0],
			HandMap: parseHand(arr[0]),
			Bid:     bid,
		})
	}

	sortedRounds := sortRounds(rounds)

	sum := 0
	for i, round := range sortedRounds {
		sum += (i + 1) * round.Bid
	}

	fmt.Println(sum)
}

func sortRounds(rounds []Round) []Round {
	slices.SortFunc(rounds, compareHands)
	return rounds
}

func compareHands(hand1 Round, hand2 Round) int {
	value1 := hand1.getHandValue()
	value2 := hand2.getHandValue()

	if value1 == value2 {
		for i := 0; i < 5; i++ {
			cmpChar := cardValues[string(hand1.Hand[i])] - cardValues[string(hand2.Hand[i])]
			if cmpChar != 0 {
				return cmpChar
			}
		}
	}

	return value1 - value2
}

func (r Round) getHandValue() int {
	if r.isFiveOfAKind() {
		return 6
	}
	if r.isFourOfAKind() {
		return 5
	}
	if r.isFull() {
		return 4
	}
	if r.isThreeOfAKind() {
		return 3
	}
	if r.isDoublePair() {
		return 2
	}
	if r.isPair() {
		return 1
	}
	if r.isHighCard() {
		return 0
	}
	fmt.Println("Y A UN SOUCIS")
	return -1
}

func (r Round) isFiveOfAKind() bool {
	return len(r.HandMap) == 1
}

func (r Round) isFourOfAKind() bool {
	for _, card := range r.HandMap {
		if card == 4 {
			return true
		}
	}
	return false
}

func (r Round) isFull() bool {
	return len(r.HandMap) == 2
}

func (r Round) isThreeOfAKind() bool {
	for _, card := range r.HandMap {
		if card == 3 {
			return true
		}
	}
	return false
}

func (r Round) isDoublePair() bool {
	return len(r.HandMap) == 3
}

func (r Round) isPair() bool {
	return len(r.HandMap) == 4
}

func (r Round) isHighCard() bool {
	return len(r.HandMap) == 5
}

func parseHand(hand string) map[string]int {
	handMap := make(map[string]int, 0)
	for _, c := range hand {
		_, found := handMap[string(c)]
		if found {
			handMap[string(c)]++
		} else {
			handMap[string(c)] = 1
		}
	}
	return handMap
}
