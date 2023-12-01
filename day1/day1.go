package day1

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var charsToDigit = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Star1() {
	scan := pkg.ReadFile("./day1/input")

	sum := 0

	for scan.Scan() {
		line := scan.Text()
		sum += retrieveDigitsFromString(line)
	}

	fmt.Println(sum)
}

func retrieveDigitsFromString(s string) int {
	var firstDigit string
	var lastDigit string

	for _, char := range s {
		if unicode.IsDigit(char) {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		}
	}

	numberAsString := firstDigit + lastDigit
	number, err := strconv.Atoi(numberAsString)

	if err != nil {
		return 0
	}

	return number
}

func Star2() {
	scan := pkg.ReadFile("./day1/input")

	sum := 0

	for scan.Scan() {
		line := scan.Text()
		sum += retrieveDigitsFromStringWithChars(line)
	}

	fmt.Println(sum)
}

func retrieveDigitsFromStringWithChars(s string) int {
	var firstDigit string
	var lastDigit string

	for key, value := range charsToDigit {
		strings.ReplaceAll(s, key, value)
	}

	for i, char := range s {
		if unicode.IsDigit(char) {
			if firstDigit == "" {
				firstDigit = string(char)
			}
			lastDigit = string(char)
		} else {
			var substr string
			substr = s[i:]

			for key, value := range charsToDigit {
				if strings.HasPrefix(substr, key) {
					if firstDigit == "" {
						firstDigit = value
					}
					lastDigit = value
				}
			}
		}
	}

	numberAsString := firstDigit + lastDigit
	number, err := strconv.Atoi(numberAsString)

	if err != nil {
		return 0
	}

	return number
}
