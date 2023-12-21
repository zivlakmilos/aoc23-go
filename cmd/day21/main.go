package main

import (
	"fmt"
	"strings"
)

type Point struct {
	row int
	col int

	remainingSteps int
}

var dirs = []Point{
	{row: 1, col: 0},
	{row: -1, col: 0},
	{row: 0, col: 1},
	{row: 0, col: -1},
}

type Queue []Point

func (q *Queue) Pop() Point {
	res := (*q)[0]

	(*q) = (*q)[1:]

	return res
}

func (q *Queue) Add(data Point) {
	*q = append(*q, data)
}

func parseInput(input string) [][]byte {
	res := [][]byte{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		res = append(res, []byte(line))
	}

	return res
}

func findStart(grid [][]byte) Point {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				return Point{
					row: row,
					col: col,
				}
			}
		}
	}

	return Point{}
}

func createGrid[T any](row, col int, value T) [][]T {
	res := make([][]T, row)

	for r := 0; r < row; r++ {
		res[r] = make([]T, col)
		for c := 0; c < col; c++ {
			res[r][c] = value
		}
	}

	return res
}

func isValidPoint(point Point, grid [][]byte, seen [][]bool) bool {
	if point.row < 0 || point.row >= len(grid) || point.col < 0 || point.col >= len(grid) {
		return false
	}

	if grid[point.row][point.col] == '#' || seen[point.row][point.col] {
		return false
	}

	return true
}

func countAns(ans [][]bool) int {
	res := 0

	for row := range ans {
		for col := range ans[row] {
			if ans[row][col] {
				res++
			}
		}
	}

	return res
}

func countReached(grid [][]byte, startPoint Point) int {
	queue := Queue{startPoint}
	seen := createGrid(len(grid), len(grid[0]), false)
	ans := createGrid(len(grid), len(grid[0]), false)
	for len(queue) > 0 {
		point := queue.Pop()

		if point.remainingSteps%2 == 0 {
			ans[point.row][point.col] = true
		}
		if point.remainingSteps == 0 {
			continue
		}

		for _, dir := range dirs {
			nextPoint := Point{
				row:            point.row + dir.row,
				col:            point.col + dir.col,
				remainingSteps: point.remainingSteps - 1,
			}
			if !isValidPoint(nextPoint, grid, seen) {
				continue
			}
			seen[nextPoint.row][nextPoint.col] = true
			queue.Add(nextPoint)
		}
	}

	steps := countAns(ans)
	return steps
}

func solvePuzzle01() {
	input := getInput()
	grid := parseInput(input)

	startPoint := findStart(grid)
	startPoint.remainingSteps = 64

	steps := countReached(grid, startPoint)

	fmt.Printf("Reached plots: %d\n", steps)
}

func solvePuzzle02() {
	input := getInput()
	grid := parseInput(input)

	steps := 26501365
	size := len(grid)

	startPoint := findStart(grid)
	startPoint.remainingSteps = steps

	gridWidth := steps/size - 1

	sr := startPoint.row
	sc := startPoint.col

	odd := (gridWidth/2*2 + 1) * (gridWidth/2*2 + 1)
	even := ((gridWidth + 1) / 2 * 2) * ((gridWidth + 1) / 2 * 2)

	oddPoints := countReached(grid, Point{row: sr, col: sc, remainingSteps: size*2 + 1})
	evenPoints := countReached(grid, Point{row: sr, col: sc, remainingSteps: size * 2})

	cornerT := countReached(grid, Point{row: size - 1, col: sr, remainingSteps: size - 1})
	cornerR := countReached(grid, Point{row: sr, col: 0, remainingSteps: size - 1})
	cornerB := countReached(grid, Point{row: 0, col: sc, remainingSteps: size - 1})
	cornerL := countReached(grid, Point{row: sr, col: size - 1, remainingSteps: size - 1})

	smallTR := countReached(grid, Point{row: size - 1, col: 0, remainingSteps: size/2 - 1})
	smallTL := countReached(grid, Point{row: size - 1, col: size - 1, remainingSteps: size/2 - 1})
	smallBR := countReached(grid, Point{row: 0, col: 0, remainingSteps: size/2 - 1})
	smallBL := countReached(grid, Point{row: 0, col: size - 1, remainingSteps: size/2 - 1})

	largeTR := countReached(grid, Point{row: size - 1, col: 0, remainingSteps: size*3/2 - 1})
	largeTL := countReached(grid, Point{row: size - 1, col: size - 1, remainingSteps: size*3/2 - 1})
	largeBR := countReached(grid, Point{row: 0, col: 0, remainingSteps: size*3/2 - 1})
	largeBL := countReached(grid, Point{row: 0, col: size - 1, remainingSteps: size*3/2 - 1})

	totalSteps := odd*oddPoints +
		even*evenPoints +
		cornerT + cornerR + cornerB + cornerL +
		(gridWidth+1)*(smallTR+smallTL+smallBR+smallBL) +
		gridWidth*(largeTR+largeTL+largeBR+largeBL)

	fmt.Printf("Reached plots: %d\n", totalSteps)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
