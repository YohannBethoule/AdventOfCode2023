package day6

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"strconv"
	"unicode"
)

type Race struct {
	Duration int
	Distance int
}

func Star1() {
	scan := pkg.ReadFile("./day6/input")

	var races []Race

	scan.Scan()
	timesString := scan.Text()
	times := parseIntsFromString(timesString)

	scan.Scan()
	distancesString := scan.Text()
	distances := parseIntsFromString(distancesString)

	if len(times) != len(distances) {
		fmt.Println("Times and distances not the same length")
	}

	result := 1

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			Duration: times[i],
			Distance: distances[i],
		})
		result = result * races[i].getOutstandingMoveCount()
	}

	fmt.Println(result)
}

func (r Race) getOutstandingMoveCount() int {
	count := 0
	for i := 0; i < r.Duration; i++ {
		if calculateTravelDistance(i, r.Duration) > r.Distance {
			count++
		}
	}
	return count
}

func calculateTravelDistance(btnPushDuration int, totalDuration int) int {
	return btnPushDuration * (totalDuration - btnPushDuration)
}

func parseIntsFromString(s string) []int {
	numbers := make([]int, 0)
	var nbString string
	isReadingNb := false
	for _, c := range s {
		if unicode.IsDigit(c) {
			nbString = nbString + string(c)
			isReadingNb = true
		} else {
			if isReadingNb {
				nb, err := strconv.Atoi(nbString)
				if err == nil {
					numbers = append(numbers, nb)
				}
				nbString = ""
			}
			isReadingNb = false
		}
	}
	nb, err := strconv.Atoi(nbString)
	if err == nil {
		numbers = append(numbers, nb)
	}

	return numbers
}

func Star2() {
	scan := pkg.ReadFile("./day6/input2")

	scan.Scan()
	time, _ := strconv.Atoi(scan.Text())
	scan.Scan()
	distance, _ := strconv.Atoi(scan.Text())

	race := Race{
		Duration: time,
		Distance: distance,
	}

	result := race.getOutstandingMoveCount()
	fmt.Println(result)
}
