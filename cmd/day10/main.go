package main

import (
	"fmt"
	"strings"
)

type Node struct {
	conns    [][2]int
	distance int
	ch       byte
	row      int
	col      int
}

var connsMap = map[byte][][2]int{
	'|': {{-1, 0}, {1, 0}},
	'-': {{0, -1}, {0, 1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{0, -1}, {-1, 0}},
	'7': {{0, -1}, {1, 0}},
	'F': {{0, 1}, {1, 0}},
}

func parseNode(ch byte, row, col int) Node {
	conns := [][2]int{}

	tmp, ok := connsMap[ch]
	if ok {
		conns = tmp
	}

	return Node{
		conns:    conns,
		distance: 0,
		ch:       ch,
		row:      row,
		col:      col,
	}
}

func parseInput(input string) [][]Node {
	res := [][]Node{}

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		row := []Node{}
		for j, ch := range line {
			node := parseNode(byte(ch), i, j)
			row = append(row, node)
		}
		res = append(res, row)
	}

	return res
}

func findStartNode(nodes [][]Node) *Node {
	for row := range nodes {
		for col := range nodes[row] {
			if nodes[row][col].ch == 'S' {
				nodes[row][col].conns = [][2]int{
					{-1, 0},
					{0, -1},
					{0, 1},
					{1, 0},
				}
				return &nodes[row][col]
			}
		}
	}

	return nil
}

func findConnectedNode(nodes [][]Node, con [2]int, node *Node) *Node {
	row := node.row + con[0]
	col := node.col + con[1]

	if row < 0 || row > len(nodes)-1 {
		return nil
	}

	if col < 0 || col > len(nodes[row])-1 {
		return nil
	}

	return &nodes[row][col]
}

func isFullyConnected(a, b *Node) bool {
	isOneConnected := false

	for _, con := range a.conns {
		if a.row+con[0] == b.row && a.col+con[1] == b.col {
			isOneConnected = true
			break
		}
	}

	if !isOneConnected {
		return false
	}

	for _, con := range b.conns {
		if b.row+con[0] == a.row && b.col+con[1] == a.col {
			return true
		}
	}

	return false
}

func solvePuzzle01() {
	input := getInput()
	nodes := parseInput(input)
	startNode := findStartNode(nodes)

	nodeQueue := []*Node{startNode}

	for len(nodeQueue) > 0 {
		currentNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		for _, con := range currentNode.conns {
			node := findConnectedNode(nodes, con, currentNode)
			if node == nil {
				continue
			}
			if node.row == currentNode.row && node.col == currentNode.col {
				continue
			}
			if !isFullyConnected(currentNode, node) {
				continue
			}

			distance := currentNode.distance + 1
			if node.distance == 0 || node.distance > distance {
				node.distance = distance
				nodeQueue = append(nodeQueue, node)
			}
		}
	}

	maxDistance := 0
	for _, row := range nodes {
		for _, node := range row {
			if node.distance > maxDistance {
				maxDistance = node.distance
			}
		}
	}

	fmt.Printf("Minimal steps to farthest distance: %d\n", maxDistance)
}

func main() {
	solvePuzzle01()
}
