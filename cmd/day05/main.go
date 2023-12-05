package main

import (
	"fmt"
	"strconv"
	"strings"
)

var mappingNames []string = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func parseSeeds(line string) []int {
	res := []int{}

	seeds := strings.Split(line, " ")
	for _, seed := range seeds {
		seed = strings.TrimSpace(seed)
		if seed != "" {
			num, _ := strconv.Atoi(seed)
			res = append(res, num)
		}
	}

	return res
}

func parseMap(out map[int]int, line string) {
	data := strings.Split(line, " ")
	dest, _ := strconv.Atoi(data[0])
	src, _ := strconv.Atoi(data[1])
	count, _ := strconv.Atoi(data[2])

	for i := 0; i < count; i++ {
		out[src+i] = dest + i
	}
}

func mapSeeds(seeds []int, line string, seedsCompleted []bool) {
	data := strings.Split(line, " ")
	dest, _ := strconv.Atoi(data[0])
	src, _ := strconv.Atoi(data[1])
	count, _ := strconv.Atoi(data[2])

	for idx := range seeds {
		if seedsCompleted[idx] {
			continue
		}

		if seeds[idx] >= src && seeds[idx] < src+count {
			seeds[idx] = dest + seeds[idx] - src
			seedsCompleted[idx] = true
		}
	}
}

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0]

	for _, val := range arr {
		if val < res {
			res = val
		}
	}

	return res
}

func solvePuzzle01(input string) {
	lines := strings.Split(input, "\n")
	block := false

	seeds := []int{}
	var seedsCompleted []bool

	for _, line := range lines {
		if line == "" {
			block = false
			continue
		}

		if !block {
			data := strings.Split(line, ":")
			if data[0] == "seeds" {
				seeds = parseSeeds(data[1])
			} else {
				block = true
				seedsCompleted = make([]bool, len(seeds))
			}
			continue
		}

		mapSeeds(seeds, line, seedsCompleted)
	}

	min := findMin(seeds)
	fmt.Printf("Lowest location number: %d\n", min)
}

func main() {
	input := getInput()
	solvePuzzle01(input)
}
