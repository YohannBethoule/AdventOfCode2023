package day9

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Star1() {
	historics := parseInput()
	sum := 0
	for _, historic := range historics {
		sum += extrapolateFuture(historic)
	}
	fmt.Println(sum)
}

func Star2() {
	historics := parseInput()
	sum := 0
	for _, historic := range historics {
		sum += extrapolatePast(historic)
	}
	fmt.Println(sum)
}

func parseInput() [][]int {
	scan := pkg.ReadFile("./day9/input")

	historics := make([][]int, 0)

	for scan.Scan() {
		historic := make([]int, 0)
		line := scan.Text()
		lineAsArray := strings.Split(line, " ")
		for _, s := range lineAsArray {
			nb, _ := strconv.Atoi(s)
			historic = append(historic, nb)
		}
		historics = append(historics, historic)
	}

	return historics
}

func extrapolatePast(historic []int) int {
	extras := [][]int{
		historic,
	}
	for i := 0; i < len(extras); i++ {
		if slices.Max(extras[i]) == 0 && slices.Min(extras[i]) == 0 {
			extras[i] = append(extras[i], 0)
			break
		}
		extras = append(extras, calculateNextLineForPast(extras[i]))
	}

	for i := len(extras) - 1; i > 0; i-- {
		extras[i-1] = append([]int{extras[i-1][0] + extras[i][0]}, extras[i-1]...)
	}

	return extras[0][0]
}

func extrapolateFuture(historic []int) int {
	extras := [][]int{
		historic,
	}
	for i := 0; i < len(extras); i++ {
		if slices.Max(extras[i]) == 0 && slices.Min(extras[i]) == 0 {
			extras[i] = append(extras[i], 0)
			break
		}
		extras = append(extras, calculateNextLine(extras[i]))
	}

	for i := len(extras) - 1; i > 0; i-- {
		extras[i-1] = append(extras[i-1], extras[i-1][len(extras[i-1])-1]+extras[i][len(extras[i])-1])
	}

	return extras[0][len(extras[0])-1]
}

func calculateNextLine(historic []int) []int {
	nextLine := make([]int, 0)
	for i := 0; i < len(historic)-1; i++ {
		nextLine = append(nextLine, historic[i+1]-historic[i])
	}
	return nextLine
}

func calculateNextLineForPast(historic []int) []int {
	nextLine := make([]int, 0)
	for i := len(historic) - 1; i > 0; i-- {
		nextLine = append([]int{historic[i-1] - historic[i]}, nextLine...)
	}
	return nextLine
}
