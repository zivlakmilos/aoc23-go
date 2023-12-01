package main

import (
	"fmt"
	"strings"
)

// const input string = `1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet`

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func parseCalibrationValue(line string) int {
	l := 0
	r := len(line) - 1

	for !isDigit(line[l]) {
		l++
	}

	for !isDigit(line[r]) {
		r--
	}

	res := int(line[l] - '0')
	res *= 10
	res += int(line[r] - '0')

	return res
}

func main() {
	total := 0

	input := getInput()

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		total += parseCalibrationValue(line)
	}

	fmt.Printf("Sum of all calibration values is: %d\n", total)
}
