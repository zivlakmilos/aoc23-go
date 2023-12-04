package main

import (
	"fmt"
	"strings"
)

func solvePuzzle1(input string) {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		points := 0

		card := strings.Split(line, ":")
		data := strings.Split(card[1], "|")

		correct := strings.Split(data[0], " ")
		numbers := strings.Split(data[1], " ")

		for idx := range correct {
			correct[idx] = strings.TrimSpace(correct[idx])
		}

		for _, num := range numbers {
			num = strings.TrimSpace(num)
			if num == "" {
				continue
			}

			for _, cor := range correct {
				if num == cor {
					if points > 0 {
						points *= 2
					} else {
						points = 1
					}
					break
				}
			}
		}

		total += points
	}

	fmt.Printf("Total points: %d\n", total)
}

func solvePuzzle2(input string) {
	lines := strings.Split(input, "\n")

	total := len(lines)

	copies := make([]int, len(lines))

	for idx, line := range lines {
		points := 0

		card := strings.Split(line, ":")
		data := strings.Split(card[1], "|")

		correct := strings.Split(data[0], " ")
		numbers := strings.Split(data[1], " ")

		for idx := range correct {
			correct[idx] = strings.TrimSpace(correct[idx])
		}

		for _, num := range numbers {
			num = strings.TrimSpace(num)
			if num == "" {
				continue
			}

			for _, cor := range correct {
				if num == cor {
					points++
					break
				}
			}
		}

		for i := 0; i < points && i < len(copies); i++ {
			copies[idx+i+1] += copies[idx] + 1
		}
	}

	for _, copy := range copies {
		total += copy
	}

	fmt.Printf("Total scratchcards: %d\n", total)
}

func main() {
	input := getInput()
	solvePuzzle1(input)
	solvePuzzle2(input)
}
