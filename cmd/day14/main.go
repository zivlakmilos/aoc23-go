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
	for col := 0; col < len(input[0]); col++ {
		emptyQueue := []int{}
		for row := 0; row < len(input); row++ {
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

func tiltSouth(input [][]byte) {
	for col := 0; col < len(input[0]); col++ {
		emptyQueue := []int{}
		for row := len(input) - 1; row >= 0; row-- {
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

func tiltWest(input [][]byte) {
	for row := 0; row < len(input); row++ {
		emptyQueue := []int{}
		for col := 0; col < len(input[0]); col++ {
			switch input[row][col] {
			case 'O':
				if len(emptyQueue) > 0 {
					input[row][col] = '.'
					input[row][emptyQueue[0]] = 'O'

					emptyQueue = emptyQueue[1:]
					if input[row][col] == '.' {
						emptyQueue = append(emptyQueue, col)
					}
				}
			case '#':
				emptyQueue = []int{}
			case '.':
				emptyQueue = append(emptyQueue, col)
			}
		}
	}
}

func tiltEast(input [][]byte) {
	for row := 0; row < len(input); row++ {
		emptyQueue := []int{}
		for col := len(input[0]) - 1; col >= 0; col-- {
			switch input[row][col] {
			case 'O':
				if len(emptyQueue) > 0 {
					input[row][col] = '.'
					input[row][emptyQueue[0]] = 'O'

					emptyQueue = emptyQueue[1:]
					if input[row][col] == '.' {
						emptyQueue = append(emptyQueue, col)
					}
				}
			case '#':
				emptyQueue = []int{}
			case '.':
				emptyQueue = append(emptyQueue, col)
			}
		}
	}
}

func createKey(input [][]byte) string {
	var sb strings.Builder

	for _, row := range input {
		sb.Write(row)
	}

	return sb.String()
}

func tilt(input [][]byte, cycles int) {
	currentCycle := 0
	cycleStart := 0
	cache := map[string]int{}

	for {
		tiltNorth(input)
		tiltWest(input)
		tiltSouth(input)
		tiltEast(input)

		currentCycle++

		key := createKey(input)
		if idx, ok := cache[key]; ok {
			cycleStart = idx
			break
		}
		cache[key] = currentCycle
	}

	cycles = (cycles - currentCycle) % (currentCycle - cycleStart)
	for cycles > 0 {
		tiltNorth(input)
		tiltWest(input)
		tiltSouth(input)
		tiltEast(input)

		cycles--
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

func solvePuzzle02() {
	input := getInput()
	lines := strings.Split(input, "\n")
	data := parseInput(lines)

	tilt(data, 1000000000)
	totalLoad := calcLoad(data)

	fmt.Printf("Total load: %d\n", totalLoad)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
