package main

import (
	"fmt"
	"strings"
)

type Node struct {
	row   int
	col   int
	steps int
}

type Stack []Node

func (s *Stack) Push(node Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() Node {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

var dirs = map[byte][]Node{
	'^': {Node{row: -1, col: 0}},
	'v': {Node{row: 1, col: 0}},
	'<': {Node{row: 0, col: -1}},
	'>': {Node{row: 0, col: 1}},
	'.': {Node{row: -1, col: 0}, Node{row: 1, col: 0}, Node{row: 0, col: -1}, Node{row: 0, col: 1}},
}

func parseInput(input string) [][]byte {
	res := [][]byte{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		res = append(res, []byte(line))
	}

	return res
}

func printGrid(grid [][]byte) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Printf("%c", grid[row][col])
		}
		fmt.Printf("\n")
	}
}

func countNeighbours(grid [][]byte, row, col int) int {
	count := 0

	if row > 0 && grid[row-1][col] != '#' {
		count++
	}
	if col > 0 && grid[row][col-1] != '#' {
		count++
	}
	if row < len(grid)-1 && grid[row+1][col] != '#' {
		count++
	}
	if col < len(grid[row])-1 && grid[row][col+1] != '#' {
		count++
	}

	return count
}

func findEmptyNode(grid [][]byte, row int) Node {
	for col := range grid[row] {
		if grid[row][col] == '.' {
			return Node{
				row: row,
				col: col,
			}
		}
	}

	return Node{}
}

func parseNodes(grid [][]byte) []Node {
	res := []Node{}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '#' {
				continue
			}
			neighbours := countNeighbours(grid, row, col)
			if neighbours > 2 {
				res = append(res, Node{
					row: row,
					col: col,
				})
			}
		}
	}

	return res
}

func createGraph(rows, cols int) [][][]Node {
	res := make([][][]Node, rows)

	for row := 0; row < rows; row++ {
		res[row] = make([][]Node, cols)
		for col := 0; col < cols; col++ {
			res[row][col] = []Node{}
		}
	}

	return res
}

func createMap[T any](rows, cols int, value T) [][]T {
	res := make([][]T, rows)

	for row := 0; row < rows; row++ {
		res[row] = make([]T, cols)
		for col := 0; col < cols; col++ {
			res[row][col] = value
		}
	}

	return res
}

func dfs(node, finish Node, graph [][][]Node, seen [][]bool) int {
	if node.row == finish.row && node.col == finish.col {
		return 0
	}

	m := 0

	seen[node.row][node.col] = true
	for _, n := range graph[node.row][node.col] {
		if !seen[n.row][n.col] {
			m = max(m, dfs(n, finish, graph, seen)+n.steps)
		}
	}
	seen[node.row][node.col] = false

	return m
}

func solvePuzzle01() {
	input := getInput()
	grid := parseInput(input)
	nodes := parseNodes(grid)

	start := findEmptyNode(grid, 0)
	finish := findEmptyNode(grid, len(grid)-1)
	nodes = append([]Node{start, finish}, nodes...)

	graph := createGraph(len(grid), len(grid[0]))

	for _, node := range nodes {
		node.steps = 0
		stack := Stack{node}
		seen := createMap(len(grid), len(grid[0]), false)
		seen[node.row][node.col] = true

		for len(stack) > 0 {
			current := stack.Pop()

			if current.steps > 0 {
				found := false
				for _, node := range nodes {
					if node.row == current.row && node.col == current.col {
						found = true
						break
					}
				}
				if found {
					graph[node.row][node.col] = append(graph[node.row][node.col], current)
					continue
				}
			}

			for _, dir := range dirs[grid[current.row][current.col]] {
				next := Node{
					row:   current.row + dir.row,
					col:   current.col + dir.col,
					steps: current.steps + 1,
				}

				if next.row >= 0 && next.row < len(grid) &&
					next.col >= 0 && next.col < len(grid) &&
					grid[next.row][next.col] != '#' && !seen[next.row][next.col] {
					stack.Push(next)
					seen[next.row][next.col] = true
				}
			}
		}
	}

	seen := createMap(len(grid), len(grid[0]), false)
	steps := dfs(start, finish, graph, seen)
	fmt.Printf("Longest hike: %d\n", steps)
}

func solvePuzzle02() {
	input := getInput()
	grid := parseInput(input)
	nodes := parseNodes(grid)

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '#' {
				grid[row][col] = '.'
			}
		}
	}

	start := findEmptyNode(grid, 0)
	finish := findEmptyNode(grid, len(grid)-1)
	nodes = append([]Node{start, finish}, nodes...)

	graph := createGraph(len(grid), len(grid[0]))

	for _, node := range nodes {
		node.steps = 0
		stack := Stack{node}
		seen := createMap(len(grid), len(grid[0]), false)
		seen[node.row][node.col] = true

		for len(stack) > 0 {
			current := stack.Pop()

			if current.steps > 0 {
				found := false
				for _, node := range nodes {
					if node.row == current.row && node.col == current.col {
						found = true
						break
					}
				}
				if found {
					graph[node.row][node.col] = append(graph[node.row][node.col], current)
					continue
				}
			}

			for _, dir := range dirs[grid[current.row][current.col]] {
				next := Node{
					row:   current.row + dir.row,
					col:   current.col + dir.col,
					steps: current.steps + 1,
				}

				if next.row >= 0 && next.row < len(grid) &&
					next.col >= 0 && next.col < len(grid) &&
					grid[next.row][next.col] != '#' && !seen[next.row][next.col] {
					stack.Push(next)
					seen[next.row][next.col] = true
				}
			}
		}
	}

	seen := createMap(len(grid), len(grid[0]), false)
	steps := dfs(start, finish, graph, seen)
	fmt.Printf("Longest hike: %d\n", steps)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
