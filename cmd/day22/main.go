package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type CoordAxis int

const (
	CoordAxisX CoordAxis = iota
	CoordAxisY
	CoordAxisZ
)

type Brick struct {
	coord1 [3]int
	coord2 [3]int
}

func parseInput(input string) []Brick {
	res := []Brick{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.Replace(line, "~", ",", 1)
		data := strings.Split(line, ",")

		x1, _ := strconv.Atoi(data[0])
		y1, _ := strconv.Atoi(data[1])
		z1, _ := strconv.Atoi(data[2])
		x2, _ := strconv.Atoi(data[3])
		y2, _ := strconv.Atoi(data[4])
		z2, _ := strconv.Atoi(data[5])

		brick := Brick{
			coord1: [3]int{x1, y1, z1},
			coord2: [3]int{x2, y2, z2},
		}

		res = append(res, brick)
	}

	return res
}

func isOverlaps(a, b Brick) bool {
	return max(a.coord1[CoordAxisX], b.coord1[CoordAxisX]) <= min(a.coord2[CoordAxisX], b.coord2[CoordAxisX]) &&
		max(a.coord1[CoordAxisY], b.coord1[CoordAxisY]) <= min(a.coord2[CoordAxisY], b.coord2[CoordAxisY])
}

func createMap(count int) [][]int {
	res := make([][]int, count)

	for i := 0; i < count; i++ {
		res[i] = []int{}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	bricks := parseInput(input)

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].coord1[CoordAxisZ] < bricks[j].coord1[CoordAxisZ]
	})

	for idx, brick := range bricks {
		maxZ := 1
		for _, check := range bricks[:idx] {
			if isOverlaps(brick, check) {
				maxZ = max(maxZ, check.coord2[CoordAxisZ]+1)
			}
		}
		bricks[idx].coord2[CoordAxisZ] -= bricks[idx].coord1[CoordAxisZ] - maxZ
		bricks[idx].coord1[CoordAxisZ] = maxZ
	}
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].coord1[CoordAxisZ] < bricks[j].coord1[CoordAxisZ]
	})

	kSupportsV := createMap(len(bricks))
	vSupportsK := createMap(len(bricks))

	for i, upper := range bricks {
		for j, lower := range bricks[:i] {
			if isOverlaps(lower, upper) && upper.coord1[CoordAxisZ] == lower.coord2[CoordAxisZ]+1 {
				kSupportsV[j] = append(kSupportsV[j], i)
				vSupportsK[i] = append(vSupportsK[i], j)
			}
		}
	}

	total := 0
	for i := range bricks {
		canRemove := true
		for _, j := range kSupportsV[i] {
			if len(vSupportsK[j]) < 2 {
				canRemove = false
				break
			}
		}
		if canRemove {
			total++
		}
	}

	fmt.Printf("Number of bricks that could be desintegrated: %d\n", total)
}

func main() {
	solvePuzzle01()
}
