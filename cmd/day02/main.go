package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	totalRed   int = 12
	totalGreen int = 13
	totalBlue  int = 14
)

func parseGameId(game string) int {
	data := strings.Split(game, " ")
	val, _ := strconv.Atoi(data[1])
	return val
}

func isGameValid(game string) bool {
	sets := strings.Split(game, "; ")

	for _, set := range sets {
		counts := map[string]int{}

		items := strings.Split(set, ", ")
		for _, item := range items {
			data := strings.Split(item, " ")

			val, _ := strconv.Atoi(data[0])
			counts[data[1]] += val
		}

		if counts["red"] > totalRed ||
			counts["green"] > totalGreen ||
			counts["blue"] > totalBlue {
			return false
		}
	}

	return true
}

func calcPowerOfGame(game string) int {
	sets := strings.Split(game, "; ")

	counts := map[string]int{}

	for _, set := range sets {

		items := strings.Split(set, ", ")
		for _, item := range items {
			data := strings.Split(item, " ")

			val, _ := strconv.Atoi(data[0])

			if val > counts[data[1]] {
				counts[data[1]] = val
			}
		}
	}

	return counts["red"] * counts["green"] * counts["blue"]
}

func main() {
	input := getInput()
	lines := strings.Split(input, "\n")

	total := 0
	totalPowers := 0

	for _, line := range lines {
		cols := strings.Split(line, ": ")

		if isGameValid(cols[1]) {
			id := parseGameId(cols[0])
			total += id
		}

		power := calcPowerOfGame(cols[1])
		totalPowers += power
	}

	fmt.Printf("Sum of valid game ids is: %d\n", total)
	fmt.Printf("Sum of game powers is: %d\n", totalPowers)
}
