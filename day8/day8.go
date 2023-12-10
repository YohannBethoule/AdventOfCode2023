package day8

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"slices"
	"strings"
)

type Input struct {
	instructions  string
	startingNodes []string
	endNodes      []string
	network       map[string]map[string]string
}

func Star1() {
	input := parseInput()
	fmt.Println(input.calculateStepCount())
}

func Star2() {
	input := parseInput()
	fmt.Println(input.calculateGhostStepCount())
}

func (input Input) calculateGhostStepCount() int {
	distances := make([]int, 0)

	for _, from := range input.startingNodes {
		distances = append(distances, input.calculateStepCountForOneNode(from))
	}

	return pkg.LCMM(distances)
}

func (input Input) calculateStepCountForOneNode(from string) int {
	count := 0
	for i := 0; i <= len(input.instructions); i++ {
		if i == len(input.instructions) {
			i = -1
			continue
		}
		count++
		from = input.network[from][string(input.instructions[i])]
		if slices.Contains(input.endNodes, from) {
			return count
		}
	}
	fmt.Println("on a pas trouvé la fin")
	return -1
}

func (input Input) calculateStepCount() int {
	from := "AAA"
	count := 0
	for i := 0; i <= len(input.instructions); i++ {
		if i == len(input.instructions) {
			i = -1
			continue
		}
		count++
		from = input.network[from][string(input.instructions[i])]
		if from == "ZZZ" {
			return count
		}
	}
	fmt.Println("on a pas trouvé la fin")
	return -1
}

func parseInput() Input {
	input := Input{
		instructions:  "",
		startingNodes: make([]string, 0),
		endNodes:      make([]string, 0),
		network:       make(map[string]map[string]string),
	}

	scan := pkg.ReadFile("./day8/input")
	scan.Scan()
	input.instructions = scan.Text()
	scan.Scan()
	for scan.Scan() {
		line := scan.Text()
		from, node := parseNode(line)
		if string(from[2]) == "A" {
			input.startingNodes = append(input.startingNodes, from)
		} else if string(from[2]) == "Z" {
			input.endNodes = append(input.endNodes, from)
		}
		input.network[from] = node
	}

	return input
}

func parseNode(s string) (string, map[string]string) {
	lineAsArray := strings.Split(s, "=")
	from := strings.Trim(lineAsArray[0], " ")
	destinationsAsArray := strings.Split(cleanNodeString(lineAsArray[1]), ",")

	return from, map[string]string{
		"L": destinationsAsArray[0],
		"R": destinationsAsArray[1],
	}
}

func cleanNodeString(s string) string {
	cleaned := strings.ReplaceAll(s, "(", "")
	cleaned = strings.ReplaceAll(cleaned, ")", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	return cleaned
}
