package main

import (
	"fmt"
	"strings"
)

func calcHash(str string) int {
	hash := 0

	for _, r := range str {
		ch := byte(r)

		hash += int(ch)
		hash *= 17
		hash %= 256
	}

	return hash
}

func solvePuzzle01() {
	input := getInput()
	data := strings.Split(input, ",")

	total := 0
	for _, str := range data {
		total += calcHash(str)
	}

	fmt.Printf("Sum of hashes: %d\n", total)
}

func main() {
	solvePuzzle01()
}
