package pkg

import (
	"math"
	"slices"
)

func ArrayContains(arr1 []string, arr2 []string) bool {
	for _, s := range arr2 {
		if !slices.Contains(arr1, s) {
			return false
		}
	}
	return true
}

func LCMM(numbers []int) int {
	lcm := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcm = LCM(lcm, numbers[i])
	}
	return lcm
}

func LCM(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(GCD(a, b)))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}
