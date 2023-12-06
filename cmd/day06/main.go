package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(out *[]int, line string) {
	data := strings.Split(strings.Split(line, ":")[1], " ")
	for _, el := range data {
		el = strings.TrimSpace(el)
		if el == "" {
			continue
		}

		num, _ := strconv.Atoi(el)
		*out = append(*out, num)
	}
}

func parseTimes(line string) []int {
	res := []int{}
	parseLine(&res, line)
	return res
}

func parseDistances(line string) []int {
	res := []int{}
	parseLine(&res, line)
	return res
}

func calcDistance(totalTime int, time int) int {
	return (totalTime - time) * time
}

func parseLineKerning(line string) int {
	data := strings.Split(strings.Split(line, ":")[1], " ")
	str := ""
	for _, el := range data {
		el = strings.TrimSpace(el)
		if el == "" {
			continue
		}
		str += el
	}

	res, _ := strconv.Atoi(str)
	return res
}

func parseTimesKerning(line string) int {
	return parseLineKerning(line)
}

func parseDistancesKerning(line string) int {
	return parseLineKerning(line)
}

func solvePuzzle01() {
	input := getInput()

	lines := strings.Split(input, "\n")
	times := parseTimes(lines[0])
	distances := parseTimes(lines[1])

	total := 1

	for idx := range times {
		count := 0

		time := times[idx]
		maxDistance := distances[idx]

		for t := 0; t <= time; t++ {
			distance := calcDistance(time, t)
			if distance > maxDistance {
				count++
			}
		}

		if count > 0 {
			total *= count
		}
	}

	fmt.Printf("Total number of ways to beat record: %d\n", total)
}

func solvePuzzle02() {
	input := getInput()

	lines := strings.Split(input, "\n")

	totalTime := parseTimesKerning(lines[0])
	maxDistance := parseTimesKerning(lines[1])

	total := 0

	for t := 0; t <= totalTime; t++ {
		distance := calcDistance(totalTime, t)
		if distance > maxDistance {
			total++
		}
	}

	fmt.Printf("Total number of ways to beat record: %d\n", total)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
