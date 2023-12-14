package main

import (
	"fmt"
	"strings"
)

func parseInput(lines []string) [][]byte {
	res := [][]byte{}

	for _, line := range lines {
		l := []byte{}
		for _, ch := range line {
			l = append(l, byte(ch))
		}
		res = append(res, l)
	}

	return res
}

func tiltNorth(input [][]byte) {
	for col := range input[0] {
		emptyQueue := []int{}
		for row := range input {
			switch input[row][col] {
			case 'O':
				if len(emptyQueue) > 0 {
					input[row][col] = '.'
					input[emptyQueue[0]][col] = 'O'

					emptyQueue = emptyQueue[1:]
					if input[row][col] == '.' {
						emptyQueue = append(emptyQueue, row)
					}
				}
			case '#':
				emptyQueue = []int{}
			case '.':
				emptyQueue = append(emptyQueue, row)
			}
		}
	}
}

func calcLoad(input [][]byte) int {
	res := 0

	for row := range input {
		load := 0
		for col := range input {
			if input[row][col] == 'O' {
				load++
			}
		}
		load *= len(input) - row
		res += load
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	lines := strings.Split(input, "\n")
	data := parseInput(lines)

	tiltNorth(data)
	totalLoad := calcLoad(data)

	fmt.Printf("Total load: %d\n", totalLoad)
}

func main() {
	solvePuzzle01()
}
