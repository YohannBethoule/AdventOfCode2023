package day5

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type MapRange struct {
	destinationStart int
	sourceStart      int
	rangeNb          int
}

func Star1() {
	scan := pkg.ReadFile("./day5/input")

	var seeds []int
	var mappings [][]MapRange

	i := 0
	for scan.Scan() {
		line := scan.Text()
		if i == 0 {
			seeds = readSeeds(line)
			i++
			continue
		}

		if len(line) == 0 {
			continue
		}

		if unicode.IsLetter(rune(line[0])) {
			//new map
			mappings = append(mappings, make([]MapRange, 0))
		}

		if unicode.IsDigit(rune(line[0])) {
			// process map
			mappings[len(mappings)-1] = append(mappings[len(mappings)-1], processMappingFromString(line))
		}
	}
	minLocation := math.MaxInt32
	for _, seed := range seeds {
		location := browseMappings(seed, mappings)
		if location < minLocation {
			minLocation = location
		}
	}
	fmt.Println(minLocation)
}

// browseMappings returns location for given seed
func browseMappings(seed int, mappings [][]MapRange) int {
	for len(mappings) > 1 {
		nextNumber := getMappedNumber(seed, mappings[0])
		seed = nextNumber
		mappings = mappings[1:]
	}

	return seed
}

func getMappedNumber(nb int, mapping []MapRange) int {
	for _, mapRange := range mapping {
		if isInRange(nb, mapRange) {
			index := nb - mapRange.sourceStart
			return mapRange.destinationStart + index
		}
	}
	return 0
}

func isInRange(nb int, mapRange MapRange) bool {
	return nb >= mapRange.sourceStart && nb < mapRange.sourceStart+mapRange.rangeNb
}

func processMappingFromString(s string) MapRange {
	stringNumbers := strings.Split(s, " ")
	numbers := make([]int, 0)
	for _, stringNumber := range stringNumbers {
		n, err := strconv.Atoi(stringNumber)
		if err == nil {
			numbers = append(numbers, n)
		}
	}

	return MapRange{
		destinationStart: numbers[0],
		sourceStart:      numbers[1],
		rangeNb:          numbers[2],
	}
}

func readSeeds(s string) []int {
	seedsString := strings.Split(s, ":")[1]
	return parseIntsFromString(seedsString)
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
