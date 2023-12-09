package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	res := []int{}

	nums := strings.Split(line, " ")
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		res = append(res, n)
	}

	return res
}

func calcArrays(data []int) [][]int {
	res := [][]int{data}
	idx := 1

	isAllZeros := false
	for !isAllZeros {
		isAllZeros = true
		res = append(res, []int{})
		for i := 1; i < len(res[idx-1]); i++ {
			diff := res[idx-1][i] - res[idx-1][i-1]
			res[idx] = append(res[idx], diff)

			if diff != 0 {
				isAllZeros = false
			}
		}
		idx++
	}

	return res
}

func calcPrediction(data []int) int {
	arr := calcArrays(data)

	for idx := len(arr) - 2; idx >= 0; idx-- {
		prev := arr[idx][len(arr[idx])-1]
		diff := arr[idx+1][len(arr[idx+1])-1]
		next := prev + diff
		arr[idx] = append(arr[idx], next)
	}

	return arr[0][len(arr[0])-1]
}

func solvePuzzle01() {
	input := getInput()
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		data := parseLine(line)
		prediction := calcPrediction(data)
		total += prediction
	}

	fmt.Printf("Total extrapolated values: %d\n", total)
}

func main() {
	solvePuzzle01()
}
