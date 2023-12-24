package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Hailstone struct {
	x  float64
	y  float64
	z  float64
	vx float64
	vy float64
	vz float64
	a  float64
	b  float64
	c  float64
}

func NewHailstone(x, y, z, vx, vy, vz float64) Hailstone {
	return Hailstone{
		x:  x,
		y:  y,
		z:  z,
		vx: vx,
		vy: vy,
		vz: vz,
		a:  vy,
		b:  -vx,
		c:  vy*x - vx*y,
	}
}

func parseInput(input string) []Hailstone {
	res := []Hailstone{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.Replace(line, " @ ", ", ", 1)
		data := strings.Split(line, ", ")
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])
		z, _ := strconv.Atoi(data[2])
		vx, _ := strconv.Atoi(data[3])
		vy, _ := strconv.Atoi(data[4])
		vz, _ := strconv.Atoi(data[5])

		res = append(res, NewHailstone(float64(x), float64(y), float64(z), float64(vx), float64(vy), float64(vz)))
	}

	return res
}

func countIntersections(hailstones []Hailstone) int {
	res := 0

	testStart := float64(200000000000000)
	testEnd := float64(400000000000000)

	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			a1 := hailstones[i].a
			b1 := hailstones[i].b
			c1 := hailstones[i].c
			a2 := hailstones[j].a
			b2 := hailstones[j].b
			c2 := hailstones[j].c
			if a1*b2 == b1*a2 {
				continue
			}

			x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)

			if x < testStart || x > testEnd ||
				y < testStart || y > testEnd {
				continue
			}

			if (x-hailstones[i].x)*hailstones[i].vx < 0 || (y-hailstones[i].y)*hailstones[i].vy < 0 ||
				(x-hailstones[j].x)*hailstones[j].vx < 0 || (y-hailstones[j].y)*hailstones[j].vy < 0 {
				continue
			}

			res++
		}
	}

	return res
}

func solvePuzzle01() {
	input := getInput()
	hailstones := parseInput(input)

	count := countIntersections(hailstones)
	fmt.Printf("Number of intersections: %d\n", count)
}

func main() {
	solvePuzzle01()
}
