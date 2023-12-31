package main

import (
	"fmt"
	"strings"
)

type BeamDir int

const (
	BeamDirRight BeamDir = iota
	BeamDirDown
	BeamDirLeft
	BeamDirUp
)

type Node struct {
	ch      byte
	visited [4]bool
}

var blameDirMap = map[byte][4][]BeamDir{
	'/':  {[]BeamDir{BeamDirUp}, []BeamDir{BeamDirLeft}, []BeamDir{BeamDirDown}, []BeamDir{BeamDirRight}},
	'\\': {[]BeamDir{BeamDirDown}, []BeamDir{BeamDirRight}, []BeamDir{BeamDirUp}, []BeamDir{BeamDirLeft}},
	'-':  {[]BeamDir{}, []BeamDir{BeamDirRight, BeamDirLeft}, []BeamDir{}, []BeamDir{BeamDirRight, BeamDirLeft}},
	'|':  {[]BeamDir{BeamDirDown, BeamDirUp}, []BeamDir{}, []BeamDir{BeamDirDown, BeamDirUp}, []BeamDir{}},
}

func parseInput(input string) [][]Node {
	res := [][]Node{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := []Node{}
		for _, ch := range line {
			node := Node{
				ch:      byte(ch),
				visited: [4]bool{},
			}
			row = append(row, node)
		}
		res = append(res, row)
	}

	return res
}

func calcNextPost(x, y int, dir BeamDir) (int, int) {
	switch dir {
	case BeamDirRight:
		x++
	case BeamDirDown:
		y++
	case BeamDirLeft:
		x--
	case BeamDirUp:
		y--
	}

	return x, y
}

func visit(grid [][]Node, x, y int, dir BeamDir) {
	if y < 0 || y >= len(grid) ||
		x < 0 || x >= len(grid[y]) {
		return
	}

	if grid[y][x].visited[dir] {
		return
	}

	grid[y][x].visited[dir] = true

	if blameMap, ok := blameDirMap[grid[y][x].ch]; ok {
		if len(blameMap[dir]) > 0 {
			for _, nextDir := range blameMap[dir] {
				nextX, nextY := calcNextPost(x, y, nextDir)
				visit(grid, nextX, nextY, nextDir)
			}
			return
		}
	}

	nextX, nextY := calcNextPost(x, y, dir)
	visit(grid, nextX, nextY, dir)
}

func countVisited(grid [][]Node) int {
	res := 0

	for _, row := range grid {
		for _, node := range row {
			if node.visited[0] || node.visited[1] || node.visited[2] || node.visited[3] {
				res++
			}
		}
	}

	return res
}

func cleanVisited(grid [][]Node) {
	for row := range grid {
		for col := range grid[row] {
			grid[row][col].visited[0] = false
			grid[row][col].visited[1] = false
			grid[row][col].visited[2] = false
			grid[row][col].visited[3] = false
		}
	}
}

func solvePuzzle01() {
	input := getInput()
	grid := parseInput(input)

	visit(grid, 0, 0, BeamDirRight)
	visited := countVisited(grid)

	fmt.Printf("Number of energized tiles: %d\n", visited)
}

func solvePuzzle02() {
	input := getInput()
	grid := parseInput(input)

	maxVisited := 0

	for row := range grid {
		visit(grid, 0, row, BeamDirRight)
		visited := countVisited(grid)
		if visited > maxVisited {
			maxVisited = visited
		}
		cleanVisited(grid)

		visit(grid, len(grid)-1, row, BeamDirLeft)
		visited = countVisited(grid)
		if visited > maxVisited {
			maxVisited = visited
		}
		cleanVisited(grid)
	}

	for col := range grid[0] {
		visit(grid, col, 0, BeamDirDown)
		visited := countVisited(grid)
		if visited > maxVisited {
			maxVisited = visited
		}
		cleanVisited(grid)

		visit(grid, col, len(grid[0])-1, BeamDirUp)
		visited = countVisited(grid)
		if visited > maxVisited {
			maxVisited = visited
		}
		cleanVisited(grid)
	}

	fmt.Printf("Max number of energized tiles: %d\n", maxVisited)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
