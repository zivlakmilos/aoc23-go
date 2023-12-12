package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) (string, []int) {
	data := strings.Split(line, " ")
	groups := strings.Split(data[1], ",")
	groupsNum := []int{}

	for _, group := range groups {
		num, _ := strconv.Atoi(group)
		groupsNum = append(groupsNum, num)
	}
	return data[0], groupsNum
}

func getAllCombinations(line []byte, idx int) []string {
	res := []string{}

	for i := idx; i < len(line); i++ {
		if line[i] == '?' {
			line[i] = '.'
			comb1 := getAllCombinations(line, i+1)
			line[i] = '#'
			comb2 := getAllCombinations(line, i+1)
			line[i] = '?'

			res = append(res, comb1...)
			res = append(res, comb2...)

			break
		}
	}

	if len(res) == 0 {
		res = []string{strings.Clone(string(line))}
	}

	return res
}

func isValid(line string, groups []int) bool {
	groupIdx := 0
	broken := 0

	for i := range line {
		ch := line[i]
		if ch == '#' {
			broken++
		} else if broken > 0 {
			if groupIdx >= len(groups) || groups[groupIdx] != broken {
				return false
			}
			broken = 0
			groupIdx++
		}
	}

	if broken > 0 {
		if groupIdx >= len(groups) || groups[groupIdx] != broken {
			return false
		}
		groupIdx++
	}

	return groupIdx == len(groups)
}

func calcCombinationNumber(line string) int {
	res := 0

	data, groups := parseLine(line)
	combinations := getAllCombinations([]byte(data), 0)

	for _, comb := range combinations {
		if isValid(comb, groups) {
			res++
		}
	}

	return res
}

func calcCombinationNumberOptimized(line string, groups []int) int {
	if line == "" {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(groups) == 0 {
		if strings.ContainsRune(line, '#') {
			return 0
		} else {
			return 1
		}
	}

	res := 0

	if line[0] == '.' || line[0] == '?' {
		res += calcCombinationNumberOptimized(line[1:], groups)
	}

	if line[0] == '#' || line[0] == '?' {
		if groups[0] <= len(line) && !strings.ContainsRune(line[:groups[0]], '.') && (groups[0] == len(line) || line[groups[0]] != '#') {
			if groups[0] < len(line) {
				res += calcCombinationNumberOptimized(line[groups[0]+1:], groups[1:])
			} else {
				res += calcCombinationNumberOptimized("", groups[1:])
			}
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	lines := strings.Split(input, "\n")

	totalCombinations := 0
	for _, line := range lines {
		data, groups := parseLine(line)
		totalCombinations += calcCombinationNumberOptimized(data, groups)
	}

	fmt.Printf("Total arrangements: %d\n", totalCombinations)
}

func solvePuzzle02() {
	input := getInput()
	lines := strings.Split(input, "\n")

	totalCombinations := 0
	for _, line := range lines {
		data, groups := parseLine(line)
		totalCombinations += calcCombinationNumberOptimized(data, groups)
	}

	fmt.Printf("Total arrangements: %d\n", totalCombinations)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
