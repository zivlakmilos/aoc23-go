package main

import (
	"fmt"
	"strings"
)

func isColumnsMirror(block []string, a, b int) bool {
	for a >= 0 && b < len(block[0]) {
		for row := range block {
			if block[row][a] != block[row][b] {
				return false
			}
		}
		a--
		b++
	}

	return true
}

func countMirrorCols(block []string) int {
	for col := 0; col < len(block[0])-1; col++ {
		if block[0][col] == block[0][col+1] {
			if isColumnsMirror(block, col, col+1) {
				return col + 1
			}
		}
	}

	return 0
}

func isRowMirror(block []string, a, b int) bool {
	for a >= 0 && b < len(block) {
		for col := range block[0] {
			if block[a][col] != block[b][col] {
				return false
			}
		}
		a--
		b++
	}

	return true
}

func countMirrorRows(block []string) int {
	for row := 0; row < len(block)-1; row++ {
		if block[row][0] == block[row+1][0] {
			if isRowMirror(block, row, row+1) {
				return row + 1
			}
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
		colsCount := countMirrorCols(blk)
		rowsCount := countMirrorRows(blk) * 100
		res += colsCount
		res += rowsCount
	}

	fmt.Printf("Reflection pattern analysis: %d\n", res)
}

func main() {
	solvePuzzle01()
}
