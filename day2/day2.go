package day2

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

type Game struct {
	ID     int
	rounds []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func Star1() {
	scan := pkg.ReadFile("./day2/input")
	idsSum := 0
	i := 1

	for scan.Scan() {
		line := scan.Text()
		game := lineToGame(i, line)
		if game.isValid() {
			idsSum += game.ID
		}
		i++
	}

	fmt.Println(idsSum)
}

func (g Game) isValid() bool {
	valid := true

	for _, r := range g.rounds {
		if r.red > maxRed || r.green > maxGreen || r.blue > maxBlue {
			valid = false
		}
	}

	return valid
}

func lineToGame(id int, line string) Game {
	g := Game{
		ID:     id,
		rounds: make([]Round, 0),
	}

	lineRounds := strings.Split(strings.Split(line, ":")[1], ";")

	for _, lineRound := range lineRounds {
		g.rounds = append(g.rounds, lineToRound(lineRound))
	}

	return g
}

func lineToRound(lineRound string) Round {
	var round Round

	colors := strings.Split(lineRound, ",")
	for _, c := range colors {
		split := strings.Split(strings.Trim(c, " "), " ")
		value := split[0]
		color := split[1]

		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		switch color {
		case "red":
			round.red = valueAsInt
		case "green":
			round.green = valueAsInt
		case "blue":
			round.blue = valueAsInt
		}
	}

	return round
}

func Star2() {
	scan := pkg.ReadFile("./day2/input")
	powersSum := 0
	i := 1

	for scan.Scan() {
		line := scan.Text()
		game := lineToGame(i, line)
		powersSum += game.getPower()
		i++
	}

	fmt.Println(powersSum)
}

func (g Game) getPower() int {
	minRed, minGreen, minBlue := 0, 0, 0
	for _, r := range g.rounds {
		if r.red > minRed {
			minRed = r.red
		}
		if r.green > minGreen {
			minGreen = r.green
		}
		if r.blue > minBlue {
			minBlue = r.blue
		}
	}
	return minRed * minGreen * minBlue
}
