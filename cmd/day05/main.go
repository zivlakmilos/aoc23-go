package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
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

func parseSeedsRange(line string) []Range {
	res := []Range{}

	seeds := strings.Split(line, " ")
	current := -1

	for _, seed := range seeds {
		seed = strings.TrimSpace(seed)
		if seed == "" {
			continue
		}

		num, _ := strconv.Atoi(seed)
		if current < 0 {
			current = num
		} else {
			res = append(res, Range{
				start: current,
				end:   current + num,
			})
			current = -1
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

func pop[T any](arr *[]T) T {
	l := len(*arr)
	cur := (*arr)[l-1]
	*arr = (*arr)[:l-1]
	return cur
}

func mapSeedsRange(seeds *[]Range, out *[]Range, block string) {
	lines := strings.Split(block, "\n")

	for len(*seeds) > 0 {
		if len(*seeds) == 0 {
			break
		}

		seed := pop(seeds)
		found := false

		for idx, line := range lines {
			if idx == 0 {
				continue
			}

			data := strings.Split(line, " ")
			dest, _ := strconv.Atoi(data[0])
			src, _ := strconv.Atoi(data[1])
			count, _ := strconv.Atoi(data[2])

			os := max(seed.start, src)
			oe := min(seed.end, src+count)

			if os < oe {
				*out = append(*out, Range{
					start: os - src + dest,
					end:   oe - src + dest,
				})
				if os > seed.start {
					(*seeds) = append((*seeds), Range{
						start: seed.start,
						end:   os,
					})
				}
				if oe < seed.end {
					(*seeds) = append((*seeds), Range{
						start: oe,
						end:   seed.end,
					})
				}

				found = true
				break
			}
		}

		if !found {
			*out = append(*out, Range{
				start: seed.start,
				end:   seed.end,
			})
		}
	}
}

func findMinRange(arr []Range) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0].start

	for _, val := range arr {
		if val.start < res {
			res = val.start
		}
	}

	return res
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

func solvePuzzle02(input string) {
	blocks := strings.Split(input, "\n\n")

	seeds := parseSeedsRange(strings.Split(blocks[0], ":")[1])
	newSeeds := []Range{}

	for idx, block := range blocks {
		if idx == 0 {
			continue
		}

		mapSeedsRange(&seeds, &newSeeds, block)
		seeds = newSeeds
		newSeeds = []Range{}

	}

	min := findMinRange(seeds)
	fmt.Printf("Lowest location number: %d\n", min)
}

func main() {
	input := getInput()
	solvePuzzle01(input)
	solvePuzzle02(input)
}
