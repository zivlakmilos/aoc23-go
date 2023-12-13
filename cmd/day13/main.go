package main

import (
	"fmt"
	"strings"
)

func isColumnsMirror(block []string, a, b, smudged int) bool {
	diff := 0

	for a >= 0 && b < len(block[0]) {
		for row := range block {
			if block[row][a] != block[row][b] {
				diff++
				if diff > smudged {
					return false
				}
			}
		}
		a--
		b++
	}

	return diff == smudged
}

func countMirrorCols(block []string, smudged int) int {
	for col := 0; col < len(block[0])-1; col++ {
		if isColumnsMirror(block, col, col+1, smudged) {
			return col + 1
		}
	}

	return 0
}

func isRowMirror(block []string, a, b int, smudged int) bool {
	diff := 0

	for a >= 0 && b < len(block) {
		for col := range block[0] {
			if block[a][col] != block[b][col] {
				diff++
				if diff > smudged {
					return false
				}
			}
		}
		a--
		b++
	}

	return diff == smudged
}

func countMirrorRows(block []string, smudged int) int {
	for row := 0; row < len(block)-1; row++ {
		if isRowMirror(block, row, row+1, smudged) {
			return row + 1
		}
	}

	return 0
}

func solvePuzzle01() {
	input := getInput()
	blocks := strings.Split(input, "\n\n")

	res := 0
	for _, block := range blocks {
		blk := strings.Split(block, "\n")
		colsCount := countMirrorCols(blk, 0)
		rowsCount := countMirrorRows(blk, 0) * 100
		res += colsCount
		res += rowsCount
	}

	fmt.Printf("Reflection pattern analysis: %d\n", res)
}

func solvePuzzle02() {
	input := getInput()
	blocks := strings.Split(input, "\n\n")

	res := 0
	for _, block := range blocks {
		blk := strings.Split(block, "\n")
		colsCount := countMirrorCols(blk, 1)
		rowsCount := countMirrorRows(blk, 1) * 100
		res += colsCount
		res += rowsCount
	}

	fmt.Printf("Reflection pattern analysis: %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
