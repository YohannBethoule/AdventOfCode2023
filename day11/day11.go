package day11

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"math"
	"slices"
)

type Coord struct {
	x int
	y int
}

const expansion = 999999

func Star1() {
	universe := parseInput()
	expandedUniverse := expandUniverse(universe)
	galaxies := getGalaxiesCoordinates(expandedUniverse)

	sum := 0.0

	for i, galaxy := range galaxies {
		for _, nextGalaxy := range galaxies[i+1:] {
			sum += math.Abs(float64(nextGalaxy.x-galaxy.x)) + math.Abs(float64(nextGalaxy.y-galaxy.y))
		}
	}

	fmt.Println(sum)
}

func Star2() {
	universe := parseInput()
	galaxies := getGalaxiesCoordinates(universe)
	expandedGalaxies := expandUniverseWithGalaxies(universe, galaxies)

	sum := 0.0

	for i, galaxy := range expandedGalaxies {
		for _, nextGalaxy := range expandedGalaxies[i+1:] {
			sum += math.Abs(float64(nextGalaxy.x-galaxy.x)) + math.Abs(float64(nextGalaxy.y-galaxy.y))
		}
	}

	fmt.Println(sum)
}

func parseInput() [][]int {
	scan := pkg.ReadFile("./day11/input")
	universe := make([][]int, 0)
	for scan.Scan() {
		textLine := scan.Text()
		line := make([]int, 0)
		for _, char := range textLine {
			if string(char) == "." {
				line = append(line, 0)
			} else if string(char) == "#" {
				line = append(line, 1)
			}
		}
		universe = append(universe, line)
	}
	return universe
}

func expandUniverseWithGalaxies(universe [][]int, galaxies []Coord) []Coord {
	expandedGalaxies := make([]Coord, len(galaxies))

	// expand lines
	for i := 0; i < len(universe); i++ {
		if slices.Max(universe[i]) == 0 {
			for k, galaxy := range galaxies {
				if galaxy.x > i {
					expandedGalaxies[k].x++
				}
			}
		}
	}

	// expand columns
	for j := 0; j < len(universe[0]); j++ {
		shouldExpand := true
		for i := 0; i < len(universe); i++ {
			if universe[i][j] == 1 {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			for k, galaxy := range galaxies {
				if galaxy.y > j {
					expandedGalaxies[k].y++
				}
			}
		}
	}

	for i, expand := range expandedGalaxies {
		galaxies[i].x += expand.x * expansion
		galaxies[i].y += expand.y * expansion
	}

	return galaxies
}

func expandUniverse(universe [][]int) [][]int {
	expandedUniverse := make([][]int, 0)

	// expand lines
	for i := 0; i < len(universe); i++ {
		if slices.Max(universe[i]) == 0 {
			expandedUniverse = append(expandedUniverse, universe[i])
		}
		expandedUniverse = append(expandedUniverse, universe[i])
	}

	// expand columns
	for j := 0; j < len(expandedUniverse[0]); j++ {
		shouldExpand := true
		for i := 0; i < len(expandedUniverse); i++ {
			if expandedUniverse[i][j] == 1 {
				shouldExpand = false
				break
			}
		}
		if shouldExpand {
			for i := 0; i < len(expandedUniverse); i++ {
				rest := expandedUniverse[i][j:]
				expandedUniverse[i] = append(expandedUniverse[i][0:j], expandedUniverse[i][j])
				expandedUniverse[i] = append(expandedUniverse[i], rest...)
			}
			j++
		}
	}

	return expandedUniverse
}

func getGalaxiesCoordinates(universe [][]int) []Coord {
	galaxies := make([]Coord, 0)

	for i, line := range universe {
		for j, isGalaxy := range line {
			if isGalaxy == 1 {
				galaxies = append(galaxies, Coord{
					i,
					j,
				})
			}
		}
	}

	return galaxies
}
