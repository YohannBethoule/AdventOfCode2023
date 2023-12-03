package day3

import (
	"AdventOfCode2023/pkg"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func Star1() {
	scan := pkg.ReadFile("./day3/input")

	i, sum := 0, 0
	var schema [140][140]rune

	for scan.Scan() {
		line := scan.Text()
		var l [140]rune
		for i, r := range line {
			l[i] = r
		}
		schema[i] = l
		i++
	}

	for i = 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if isDigit(schema[i][j]) {
				isPart := false

				// Etape 1 : trouver le nombre, son début et sa fin
				strNb := string(schema[i][j])
				start := j
				for j < 139 && isDigit(schema[i][j+1]) {
					strNb += string(schema[i][j+1])
					j++
				}
				end := j
				nb, err := strconv.Atoi(strNb)
				if err != nil {
					fmt.Printf("j'ai fait de la merde : %s\n", string(strNb))
				}

				// Etape 2 : déterminer si le nombre est une partie du moteur
				if start > 0 && !isEmpty(schema[i][start-1]) {
					isPart = true
				}
				if end < 139 && !isEmpty(schema[i][end+1]) {
					isPart = true
				}
				if i > 0 {
					prevStart := int(math.Max(0, float64(start-1)))
					prevEnd := int(math.Min(139, float64(end+1)))
					for _, r := range schema[i-1][prevStart : prevEnd+1] {
						if !isDigit(r) && !isEmpty(r) {
							isPart = true
						}
					}
				}
				if i < 139 {
					nextStart := int(math.Max(0, float64(start-1)))
					nextEnd := int(math.Min(139, float64(end+1)))
					for _, r := range schema[i+1][nextStart : nextEnd+1] {
						if !isDigit(r) && !isEmpty(r) {
							isPart = true
						}
					}
				}

				// Etape 3 : si c'est une partie, on ajoute à la somme
				if isPart {
					sum += nb
				}
			}
		}
	}

	fmt.Println(sum)
}

// J'ai honte
func Star2() {
	scan := pkg.ReadFile("./day3/input")

	i, sum := 0, 0
	var schema [140][140]rune

	for scan.Scan() {
		line := scan.Text()
		var l [140]rune
		for i, r := range line {
			l[i] = r
		}
		schema[i] = l
		i++
	}

	for i = 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if string(schema[i][j]) == "*" {
				parts := make([]int, 0)

				if j > 0 && isDigit(schema[i][j-1]) {
					nbAsStr := string(schema[i][j-1])
					k := j - 1
					for k > 0 && isDigit(schema[i][k-1]) {
						nbAsStr = string(schema[i][k-1]) + nbAsStr
						k--
					}
					nb, err := strconv.Atoi(nbAsStr)
					if err != nil {
						fmt.Printf("j'ai fait de la merde 1: %s\n", nbAsStr)
					}
					parts = append(parts, nb)
				}
				if j < 139 && isDigit(schema[i][j+1]) {
					nbAsStr := string(schema[i][j+1])
					k := j + 1
					for k < 139 && isDigit(schema[i][k+1]) {
						nbAsStr += string(schema[i][k+1])
						k++
					}
					nb, err := strconv.Atoi(nbAsStr)
					if err != nil {
						fmt.Printf("j'ai fait de la merde 2: %s\n", nbAsStr)
					}
					parts = append(parts, nb)
				}
				if i > 0 {
					for l := 0; l < 140; l++ {
						if isDigit(schema[i-1][l]) {
							strNb := string(schema[i-1][l])
							start := l
							for l < 139 && isDigit(schema[i-1][l+1]) {
								strNb += string(schema[i-1][l+1])
								l++
							}
							end := l
							if (start <= j+1 && start >= j-1) || (end <= j+1 && end >= j-1) || (start <= j-1 && end >= j+1) {
								nb, err := strconv.Atoi(strNb)
								if err != nil {
									fmt.Printf("j'ai fait de la merde 3: %s\n", string(strNb))
								}
								parts = append(parts, nb)
							}
						}
					}
				}
				if i < 139 {
					for l := 0; l < 140; l++ {
						if isDigit(schema[i+1][l]) {
							strNb := string(schema[i+1][l])
							start := l
							for l < 139 && isDigit(schema[i+1][l+1]) {
								strNb += string(schema[i+1][l+1])
								l++
							}
							end := l
							if (start <= j+1 && start >= j-1) || (end <= j+1 && end >= j-1) || (start <= j-1 && end >= j+1) {
								nb, err := strconv.Atoi(strNb)
								if err != nil {
									fmt.Printf("j'ai fait de la merde 4: %s\n", string(strNb))
								}
								parts = append(parts, nb)
							}
						}
					}
				}

				// Si c'est un rouage, calculer le gear power et ajoutez à la somme
				if len(parts) == 2 {
					sum += parts[0] * parts[1]
				} else {
					fmt.Println("ici c pas bon")
				}
			}
		}
	}

	fmt.Println(sum)
}

func isEmpty(r rune) bool {
	return string(r) == "."
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}
