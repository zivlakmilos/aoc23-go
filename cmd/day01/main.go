package main

import (
	"fmt"
	"strings"
)

var digitWords = [...]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func isDigitWord(line string, idx int) bool {
	found := 1
	pos := 0

	var ignore uint16 = 0

	for found > 0 && idx < len(line) {
		found = 0
		for i, word := range digitWords {
			if ignore&(1<<i) > 0 || pos >= len(word) {
				continue
			}

			ch := line[idx]
			if ch == word[pos] {
				if pos == len(word)-1 {
					return true
				}
				found++
			} else {
				ignore |= (1 << i)
			}
		}

		idx++
		pos++
	}

	return false
}

func parseDigitWord(line string, idx int) byte {
	pos := 0

	var ignore uint16 = 0

	for idx < len(line) {
		for i, word := range digitWords {
			if ignore&(1<<i) > 0 || pos >= len(word) {
				continue
			}

			if word[pos] == line[idx] {
				if pos == len(word)-1 {
					return byte(i + '1')
				}
			} else {
				ignore |= (1 << i)
			}
		}

		pos++
		idx++
	}

	return 0
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func parseCalibrationValue(line string) int {
	l := 0
	r := len(line) - 1

	var ch1, ch2 byte

	for !isDigit(line[l]) && !isDigitWord(line, l) {
		l++
	}

	if isDigit(line[l]) {
		ch1 = line[l]
	} else {
		ch1 = parseDigitWord(line, l)
	}

	for !isDigit(line[r]) && !isDigitWord(line, r) {
		r--
	}

	if isDigit(line[r]) {
		ch2 = line[r]
	} else {
		ch2 = parseDigitWord(line, r)
	}

	res := int(ch1 - '0')
	res *= 10
	res += int(ch2 - '0')

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
