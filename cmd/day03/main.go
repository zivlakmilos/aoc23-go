package main

import (
	"fmt"
	"strings"
)

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isSymbol(ch byte) bool {
	return !isDigit(ch) && ch != '.'
}

func isGearSymbol(ch byte) bool {
	return ch == '*'
}

func checkAdjacent(lines []string, y, x int) bool {
	xDifN := 0
	xDifP := 0
	yDifN := 0
	yDifP := 0

	if y > 0 {
		yDifN -= 1
	}
	if y < len(lines)-1 {
		yDifP += 1
	}
	if x > 0 {
		xDifN -= 1
	}
	if x < len(lines)-1 {
		xDifP += 1
	}

	if isSymbol(lines[y+yDifP][x]) ||
		isSymbol(lines[y+yDifN][x]) ||
		isSymbol(lines[y][x+xDifP]) ||
		isSymbol(lines[y][x+xDifN]) ||
		isSymbol(lines[y+yDifP][x+xDifP]) ||
		isSymbol(lines[y+yDifP][x+xDifN]) ||
		isSymbol(lines[y+yDifN][x+xDifP]) ||
		isSymbol(lines[y+yDifN][x+xDifN]) {
		return true
	}

	return false
}

func getGearAdjacent(lines []string, y, x int) (string, bool) {
	xDifN := 0
	xDifP := 0
	yDifN := 0
	yDifP := 0

	if y > 0 {
		yDifN -= 1
	}
	if y < len(lines)-1 {
		yDifP += 1
	}
	if x > 0 {
		xDifN -= 1
	}
	if x < len(lines)-1 {
		xDifP += 1
	}

	if isGearSymbol(lines[y+yDifP][x]) {
		return fmt.Sprintf("%dx%d", y+yDifP, x), true
	}
	if isGearSymbol(lines[y+yDifN][x]) {
		return fmt.Sprintf("%dx%d", y+yDifN, x), true
	}
	if isGearSymbol(lines[y][x+xDifP]) {
		return fmt.Sprintf("%dx%d", y, x+xDifP), true
	}
	if isGearSymbol(lines[y][x+xDifN]) {
		return fmt.Sprintf("%dx%d", y, x+xDifN), true
	}
	if isGearSymbol(lines[y+yDifP][x+xDifP]) {
		return fmt.Sprintf("%dx%d", y+yDifP, x+xDifP), true
	}
	if isGearSymbol(lines[y+yDifP][x+xDifN]) {
		return fmt.Sprintf("%dx%d", y+yDifP, x+xDifN), true
	}
	if isGearSymbol(lines[y+yDifN][x+xDifP]) {
		return fmt.Sprintf("%dx%d", y+yDifN, x+xDifP), true
	}
	if isGearSymbol(lines[y+yDifN][x+xDifN]) {
		return fmt.Sprintf("%dx%d", y+yDifN, x+xDifN), true
	}

	return "", false
}

func getTotalGearRatio(lines []string) int {
	total := 0

	currentNumber := 0
	isGear := false
	gearKey := ""

	gearsMap := map[string]int{}

	for y := range lines {
		for x := range lines[y] {
			if isDigit(lines[y][x]) {
				currentNumber *= 10
				currentNumber += int(lines[y][x] - '0')

				if !isGear {
					if key, ok := getGearAdjacent(lines, y, x); ok {
						isGear = true
						gearKey = key
					}
				}
			} else {
				if isGear {
					if gearsMap[gearKey] > 0 {
						total += currentNumber * gearsMap[gearKey]
						gearsMap[gearKey] = 0
					} else {
						gearsMap[gearKey] = currentNumber
					}
				}
				isGear = false
				currentNumber = 0
			}
		}
	}

	return total
}

func main() {
	input := getInput()
	lines := strings.Split(input, "\n")

	total := 0
	currentNumber := 0
	isAdjacent := false

	for y := range lines {
		for x := range lines[y] {
			if isDigit(lines[y][x]) {
				currentNumber *= 10
				currentNumber += int(lines[y][x] - '0')

				if !isAdjacent {
					isAdjacent = checkAdjacent(lines, y, x)
				}
			} else {
				if isAdjacent {
					total += currentNumber
				}
				isAdjacent = false
				currentNumber = 0
			}
		}
	}

	totalGearRatios := getTotalGearRatio(lines)

	fmt.Printf("Sum of part numbers is %d\n", total)
	fmt.Printf("Sum of gear ratios is %d\n", totalGearRatios)
}
