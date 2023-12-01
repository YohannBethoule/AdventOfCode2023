package pkg

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}
