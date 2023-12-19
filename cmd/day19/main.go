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

func main() {
	solvePuzzle01()
}
