package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ConditionType int

const (
	ConditionTypeNo ConditionType = iota
	ConditionTypeGT
	ConditionTypeLT
)

type Condition struct {
	conditionType ConditionType
	parameter     string
	destination   string
	value         int
}

func (c *Condition) check(part map[string]int) bool {
	switch c.conditionType {
	case ConditionTypeNo:
		return true
	case ConditionTypeGT:
		return part[c.parameter] > c.value
	case ConditionTypeLT:
		return part[c.parameter] < c.value
	}

	return false
}

func parseCondition(txt string, isLast bool) Condition {
	if isLast {
		return Condition{
			conditionType: ConditionTypeNo,
			destination:   txt,
		}
	}

	conditionType := ConditionTypeNo

	data := strings.Split(txt, ":")
	cond := data[0]
	destination := data[1]

	separator := cond[1]
	switch separator {
	case '>':
		conditionType = ConditionTypeGT
	case '<':
		conditionType = ConditionTypeLT
	}

	arr := strings.Split(cond, string(separator))
	parameter := arr[0]
	value, _ := strconv.Atoi(arr[1])

	return Condition{
		conditionType: conditionType,
		parameter:     parameter,
		destination:   destination,
		value:         value,
	}
}

func parseConditions(block string) map[string][]Condition {
	res := map[string][]Condition{}

	lines := strings.Split(block, "\n")
	for _, line := range lines {
		data := strings.Split(line, "{")
		name := data[0]
		arr := strings.Split(data[1][:len(data[1])-1], ",")

		res[name] = []Condition{}
		for i, el := range arr {
			condition := parseCondition(el, i == len(arr)-1)
			res[name] = append(res[name], condition)
		}
	}

	return res
}

func parsePart(line string) map[string]int {
	res := map[string]int{}

	arr := strings.Split(line[1:len(line)-1], ",")
	for _, el := range arr {
		data := strings.Split(el, "=")
		res[data[0]], _ = strconv.Atoi(data[1])
	}

	return res
}

func processPart(conditions map[string][]Condition, part map[string]int, current string) string {
	for _, condition := range conditions[current] {
		if condition.check(part) {
			return condition.destination
		}
	}

	return ""
}

func isPartAccepted(conditions map[string][]Condition, part map[string]int) bool {
	current := "in"
	for current != "A" && current != "R" {
		current = processPart(conditions, part, current)
	}

	return current == "A"
}

func sumPart(part map[string]int) int {
	return part["x"] + part["m"] + part["a"] + part["s"]
}

func createRanges() map[string][2]int {
	return map[string][2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
}

func copyRanges(r map[string][2]int) map[string][2]int {
	return map[string][2]int{
		"x": {r["x"][0], r["x"][1]},
		"m": {r["m"][0], r["m"][1]},
		"a": {r["a"][0], r["a"][1]},
		"s": {r["s"][0], r["s"][1]},
	}
}

func countAcceptedRanges(ranges map[string][2]int, conditions map[string][]Condition, name string) int {
	if name == "R" {
		return 0
	}

	if name == "A" {
		product := 1
		for _, val := range ranges {
			product *= val[1] - val[0] + 1
		}
		return product
	}

	total := 0

	for _, condition := range conditions[name] {
		t := [2]int{}
		f := [2]int{}
		lo := ranges[condition.parameter][0]
		hi := ranges[condition.parameter][1]

		switch condition.conditionType {
		case ConditionTypeNo:
			total += countAcceptedRanges(ranges, conditions, condition.destination)
			continue
		case ConditionTypeLT:
			t[0] = lo
			t[1] = condition.value - 1
			f[0] = condition.value
			f[1] = hi
		case ConditionTypeGT:
			t[0] = condition.value + 1
			t[1] = hi
			f[0] = lo
			f[1] = condition.value
		}

		if t[0] <= t[1] {
			copy := copyRanges(ranges)
			copy[condition.parameter] = t
			total += countAcceptedRanges(copy, conditions, condition.destination)
		}
		if f[0] <= f[1] {
			ranges = copyRanges(ranges)
			ranges[condition.parameter] = f
		} else {
			break
		}
	}

	return total
}

func solvePuzzle01() {
	input := getInput()
	blocks := strings.Split(input, "\n\n")
	lines := strings.Split(blocks[1], "\n")

	conditions := parseConditions(blocks[0])

	total := 0
	for _, line := range lines {
		part := parsePart(line)
		if isPartAccepted(conditions, part) {
			total += sumPart(part)
		}
	}

	fmt.Printf("Total accepted parts value: %d\n", total)
}

func solvePuzzle02() {
	input := getInput()
	blocks := strings.Split(input, "\n\n")

	conditions := parseConditions(blocks[0])
	ranges := createRanges()

	total := countAcceptedRanges(ranges, conditions, "in")
	fmt.Printf("Total accepted parts combinations: %d\n", total)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
