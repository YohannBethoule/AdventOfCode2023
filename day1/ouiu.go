package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var numberMap = map[string]string{
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

func spelledNumberToNumber(text string) (string, error) {
	text = strings.ToLower(text)
	if val, ok := numberMap[text]; ok {
		return val, nil
	}
	return "", fmt.Errorf("not found")
}

func Step2(text string) int {
	fmt.Println("Day 1")
	bullshits := strings.Split(text, "\n")
	results := make([]string, 0)
	for _, bullshit := range bullshits {
		var first, second string
		fmt.Println(bullshit)
		for i := 0; i < len(bullshit); i++ {
			char := bullshit[i]
			if unicode.IsNumber(rune(char)) {
				if first == "" {
					first = string(char)
				}
				second = string(char)
				continue
			} else {
				for j := len(bullshit); j > i; j-- {
					substr := bullshit[i:j]
					if val, err := spelledNumberToNumber(substr); err == nil {
						if first == "" {
							first = val
						}
						second = val
						break
					}
				}
			}
		}

		results = append(results, first+second)
	}
	finalResult := 0
	for _, result := range results {
		fmt.Println(result)
		number, _ := strconv.Atoi(result)
		finalResult += number
	}
	return finalResult
}
