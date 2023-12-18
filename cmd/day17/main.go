package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Direction int

const (
	DirectionNo Direction = iota
	DirectionRight
	DirectionDown
	DirectionUp
	DirectionLeft
)

type Node struct {
	heatLoss int
	x        int
	y        int
	dir      Direction
	step     int
}

func (n *Node) String() string {
	return fmt.Sprintf("{ x: %d, y: %d, hl: %d, dir: %d, step: %d}", n.x, n.y, n.heatLoss, n.dir, n.step)
}

func parseInput(input string) [][]int {
	res := [][]int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := []int{}
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		res = append(res, row)
	}

	return res
}

func printGrid[T any](grid [][]T) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Printf("%v, ", grid[row][col])
		}
		fmt.Printf("\n")
	}
}

func createNextNode(grid [][]int, node *Node, dir Direction) *Node {
	if dir == DirectionNo {
		return nil
	}

	nx := node.x
	ny := node.y
	step := 1
	if node.dir == dir {
		step = node.step + 1
	}

	switch dir {
	case DirectionRight:
		nx++
	case DirectionLeft:
		nx--
	case DirectionDown:
		ny++
	case DirectionUp:
		ny--
	}

	if nx < 0 || nx >= len(grid[0]) || ny < 0 || ny >= len(grid) {
		return nil
	}

	return &Node{
		heatLoss: node.heatLoss + grid[ny][nx],
		x:        nx,
		y:        ny,
		dir:      dir,
		step:     step,
	}
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func popNode(nodes *PriorityQueue) *Node {
	return heap.Pop(nodes).(*Node)
}

func addNode(nodes *PriorityQueue, node *Node) {
	if node == nil {
		return
	}

	heap.Push(nodes, node)
}

func findBestPath(grid [][]int) int {
	visited := map[string]bool{}
	unvisited := PriorityQueue{
		{
			heatLoss: 0,
			x:        0,
			y:        0,
			dir:      DirectionNo,
			step:     0,
		},
	}
	heap.Init(&unvisited)

	for len(unvisited) > 0 {
		node := popNode(&unvisited)

		if node.x == len(grid[0])-1 && node.y == len(grid)-1 {
			return node.heatLoss
		}

		key := fmt.Sprintf("%d,%d,%d,%d", node.x, node.y, node.dir, node.step)
		if visited[key] {
			continue
		}
		visited[key] = true

		if node.step < 3 && node.dir != DirectionNo {
			addNode(&unvisited, createNextNode(grid, node, node.dir))
		}

		switch node.dir {
		case DirectionLeft, DirectionRight:
			addNode(&unvisited, createNextNode(grid, node, DirectionUp))
			addNode(&unvisited, createNextNode(grid, node, DirectionDown))
		case DirectionUp, DirectionDown:
			addNode(&unvisited, createNextNode(grid, node, DirectionLeft))
			addNode(&unvisited, createNextNode(grid, node, DirectionRight))
		case DirectionNo:
			addNode(&unvisited, createNextNode(grid, node, DirectionUp))
			addNode(&unvisited, createNextNode(grid, node, DirectionDown))
			addNode(&unvisited, createNextNode(grid, node, DirectionLeft))
			addNode(&unvisited, createNextNode(grid, node, DirectionRight))
		}
	}

	return 0
}

func findBestPath2(grid [][]int) int {
	visited := map[string]bool{}
	unvisited := PriorityQueue{
		{
			heatLoss: 0,
			x:        0,
			y:        0,
			dir:      DirectionNo,
			step:     0,
		},
	}
	heap.Init(&unvisited)

	for len(unvisited) > 0 {
		node := popNode(&unvisited)

		if node.x == len(grid[0])-1 && node.y == len(grid)-1 && node.step >= 4 {
			return node.heatLoss
		}

		key := fmt.Sprintf("%d,%d,%d,%d", node.x, node.y, node.dir, node.step)
		if visited[key] {
			continue
		}
		visited[key] = true

		if node.step < 10 && node.dir != DirectionNo {
			addNode(&unvisited, createNextNode(grid, node, node.dir))
		}

		if node.step >= 4 || node.dir == DirectionNo {
			switch node.dir {
			case DirectionLeft, DirectionRight:
				addNode(&unvisited, createNextNode(grid, node, DirectionUp))
				addNode(&unvisited, createNextNode(grid, node, DirectionDown))
			case DirectionUp, DirectionDown:
				addNode(&unvisited, createNextNode(grid, node, DirectionLeft))
				addNode(&unvisited, createNextNode(grid, node, DirectionRight))
			case DirectionNo:
				addNode(&unvisited, createNextNode(grid, node, DirectionUp))
				addNode(&unvisited, createNextNode(grid, node, DirectionDown))
				addNode(&unvisited, createNextNode(grid, node, DirectionLeft))
				addNode(&unvisited, createNextNode(grid, node, DirectionRight))
			}
		}
	}

	return 0
}

func solvePuzzle01() {
	input := getInput()
	grid := parseInput(input)

	bestPath := findBestPath(grid)
	fmt.Printf("Least heat lost: %d\n", bestPath)
}

func solvePuzzle02() {
	input := getInput()
	grid := parseInput(input)

	bestPath := findBestPath2(grid)
	fmt.Printf("Least heat lost: %d\n", bestPath)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
