package main

import (
	"fmt"
	"strings"
)

type LR int

const (
	left LR = iota
	right
)

type Node struct {
	key    string
	coords [3]string
}

func parseLR(line string) []LR {
	res := []LR{}

	for _, ch := range line {
		lr := left
		if ch == 'R' {
			lr = right
		}
		res = append(res, lr)
	}

	return res
}

func parseLine(line string) Node {
	res := Node{}

	eq := strings.Split(line, " = ")
	res.key = eq[0]

	coords := strings.Split(eq[1], ", ")
	l := coords[0][1:]
	r := coords[1][:len(coords[1])-1]

	res.coords[left] = l
	res.coords[right] = r

	return res
}

func parseMap(block string) map[string]Node {
	res := map[string]Node{}
	lines := strings.Split(block, "\n")

	for _, line := range lines {
		node := parseLine(line)
		res[node.key] = node
	}

	return res
}

func calcGCD(a, b int) int {
	h := max(a, b)
	l := min(a, b)

	for h != 0 {
		t := h
		h = l % h
		l = t
	}

	return l
}

func calcLCM(a, b int) int {
	gcd := calcGCD(a, b)
	return (a * b) / gcd
}

func solvePuzzle01() {
	input := getInput()
	blocks := strings.Split(input, "\n\n")

	instructions := parseLR(blocks[0])
	coords := parseMap(blocks[1])
	pos := "AAA"

	found := false
	steps := 0
	for !found {
		for _, ins := range instructions {
			steps++
			pos = coords[pos].coords[ins]

			if pos == "ZZZ" {
				found = true
				break
			}
		}
	}

	fmt.Printf("Number of required steps is %d\n", steps)
}

func solvePuzzle02() {
	input := getInput()
	blocks := strings.Split(input, "\n\n")

	instructions := parseLR(blocks[0])
	coords := parseMap(blocks[1])
	pos := []string{}

	for _, coord := range coords {
		end := coord.key[2]
		if end == 'A' {
			pos = append(pos, coord.key)
		}
	}

	cycles := []int{}
	steps := 0
	for len(cycles) < len(pos) {
		for _, ins := range instructions {
			steps++

			for idx, p := range pos {
				if p == "" {
					continue
				}

				pos[idx] = coords[p].coords[ins]

				if pos[idx][2] == 'Z' {
					pos[idx] = ""
					cycles = append(cycles, steps)
				}
			}
		}
	}

	res := cycles[0]
	for i := 1; i < len(cycles); i++ {
		res = calcLCM(res, cycles[i])
	}

	fmt.Printf("Number of required steps is %d\n", res)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
