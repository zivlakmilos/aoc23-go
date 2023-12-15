package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Lense struct {
	id       string
	focalLen int
}

func calcHash(str string) int {
	hash := 0

	for _, r := range str {
		ch := byte(r)

		hash += int(ch)
		hash *= 17
		hash %= 256
	}

	return hash
}

func removeLense(boxes [][]Lense, idx int, id string) {
	for i := 0; i < len(boxes[idx]); i++ {
		if boxes[idx][i].id == id {
			boxes[idx] = append(boxes[idx][:i], boxes[idx][i+1:]...)
		}
	}
}

func addLense(boxes [][]Lense, idx int, lense Lense) {
	for i := 0; i < len(boxes[idx]); i++ {
		if boxes[idx][i].id == lense.id {
			boxes[idx][i].focalLen = lense.focalLen
			return
		}
	}

	boxes[idx] = append(boxes[idx], lense)
}

func calcBoxesHash(boxes [][]Lense) int {
	res := 0

	for i := range boxes {
		for j := range boxes[i] {
			res += (i + 1) * (j + 1) * boxes[i][j].focalLen
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	data := strings.Split(input, ",")

	total := 0
	for _, str := range data {
		total += calcHash(str)
	}

	fmt.Printf("Sum of hashes: %d\n", total)
}

func solvePuzzle02() {
	input := getInput()
	data := strings.Split(input, ",")

	boxes := make([][]Lense, 256)

	for _, str := range data {
		if str[len(str)-1] == '-' {
			idx := calcHash(str[:len(str)-1])
			removeLense(boxes, idx, str[:len(str)-1])
		} else {
			params := strings.Split(str, "=")
			idx := calcHash(params[0])
			focalLen, _ := strconv.Atoi(params[1])
			lense := Lense{
				id:       params[0],
				focalLen: focalLen,
			}
			addLense(boxes, idx, lense)
		}
	}

	totalFocusPower := calcBoxesHash(boxes)
	fmt.Printf("Total focus power: %d\n", totalFocusPower)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
