package day4

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	ID              int
	WinningsNumbers []int
	YourNumbers     []int
}

func Star1() {
	scan := pkg.ReadFile("./day4/input")
	id := 1
	sum := 0.0

	for scan.Scan() {
		line := scan.Text()
		card := parseCardFromString(id, line)
		sum += card.GetScore()
		id++
	}

	fmt.Println(sum)
}

func (c Card) GetScore() float64 {
	winningNumbersCount := c.getNumberOfWinningNumbers()
	if winningNumbersCount == 0 {
		return 0
	}
	return math.Pow(2, float64(winningNumbersCount-1))
}

func (c Card) getNumberOfWinningNumbers() int {
	counter := 0

	for _, n := range c.YourNumbers {
		if slices.Contains(c.WinningsNumbers, n) {
			counter++
		}
	}

	return counter
}

func parseCardFromString(id int, s string) Card {
	c := Card{
		ID:              id,
		WinningsNumbers: make([]int, 0),
		YourNumbers:     make([]int, 0),
	}

	numbersStrings := strings.Split(strings.Split(s, ":")[1], "|")

	c.WinningsNumbers = parseIntsFromString(numbersStrings[0])
	c.YourNumbers = parseIntsFromString(numbersStrings[1])
	return c
}

func parseIntsFromString(s string) []int {
	s = strings.Trim(s, " ")
	stringNumbers := strings.Split(s, " ")
	numbers := make([]int, 0)
	for _, stringNumber := range stringNumbers {
		n, err := strconv.Atoi(stringNumber)
		if err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func Star2() {
	scan := pkg.ReadFile("./day4/input")
	id := 1
	cards := make([]Card, 0)

	for scan.Scan() {
		line := scan.Text()
		card := parseCardFromString(id, line)
		cards = append(cards, card)
		id++
	}

	cardsCounters := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		cardsCounters[i] = 1
	}

	for i, c := range cards {
		for rep := 0; rep < cardsCounters[i]; rep++ {
			winCount := c.getNumberOfWinningNumbers()
			for j := i + 1; j <= i+winCount; j++ {
				cardsCounters[j]++
			}
		}
	}

	sum := 0
	for _, count := range cardsCounters {
		sum += count
	}

	fmt.Println(sum)
}
