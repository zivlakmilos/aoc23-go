package main

import (
	"fmt"
	"strings"
)

type ModuleType int

const (
	ModuleTypeBroadcaster ModuleType = iota
	ModuleTypeFlipFlop
	ModuleTypeConjuction
)

type Module struct {
	state       map[string]int
	connections []string
	moduleType  ModuleType
	value       int
}

func (m *Module) String() string {
	return fmt.Sprintf("{ type: %d, value: %d, connections: %s}\n", m.moduleType, m.value, m.connections)
}

type QueueData struct {
	src   string
	dst   string
	value int
}

type Queue []QueueData

func (q *Queue) Pop() QueueData {
	item := (*q)[0]
	(*q) = (*q)[1:]
	return item
}

func (q *Queue) Add(data QueueData) {
	*q = append(*q, data)
}

func parseInput(input string) map[string]*Module {
	res := map[string]*Module{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		data := strings.Split(line, " -> ")
		connections := strings.Split(data[1], ", ")

		key := data[0]
		moduleType := ModuleTypeBroadcaster
		state := map[string]int{}

		if data[0] != "broadcaster" {
			switch data[0][0] {
			case '%':
				moduleType = ModuleTypeFlipFlop
			case '&':
				moduleType = ModuleTypeConjuction
			}
			key = key[1:]
		}

		module := &Module{
			moduleType:  moduleType,
			connections: connections,
			value:       0,
			state:       state,
		}
		res[key] = module
	}

	for key, val := range res {
		for _, con := range val.connections {
			if _, ok := res[con]; !ok {
				continue
			}
			res[con].state[key] = 0
		}
	}

	return res
}

func processPush(modules map[string]*Module) (int, int) {
	totalLow := 0
	totalHigh := 0

	modules["broadcaster"].value = 0
	queue := Queue{QueueData{src: "button", dst: "broadcaster", value: 0}}

	for len(queue) > 0 {
		pulse := queue.Pop()

		if pulse.value > 0 {
			totalHigh++
		} else {
			totalLow++
		}

		module, ok := modules[pulse.dst]
		if !ok {
			continue
		}

		sends := true
		switch module.moduleType {
		case ModuleTypeFlipFlop:
			if pulse.value > 0 {
				sends = false
			} else {
				if module.value > 0 {
					module.value = 0
				} else {
					module.value = 1
				}
			}
		case ModuleTypeConjuction:
			module.state[pulse.src] = pulse.value
			module.value = 0
			for _, s := range module.state {
				if s == 0 {
					module.value = 1
					break
				}
			}
		default:
			module.value = pulse.value
		}

		if sends {
			for _, con := range module.connections {
				queue.Add(QueueData{src: pulse.dst, dst: con, value: module.value})
			}
		}
	}

	return totalLow, totalHigh
}

func getRXDestModule(modules map[string]*Module) string {
	for key, val := range modules {
		for _, con := range val.connections {
			if con == "rx" {
				return key
			}
		}
	}

	return ""
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func countUntilRX(modules map[string]*Module) int {
	feed := getRXDestModule(modules)

	cycleLenghts := map[string]int{}
	seen := map[string]int{}

	for key, val := range modules {
		for _, con := range val.connections {
			if con == feed {
				seen[key] = 0
			}
		}
	}

	presses := 0
	for {
		presses++

		modules["broadcaster"].value = 0
		queue := Queue{QueueData{src: "button", dst: "broadcaster", value: 0}}

		for len(queue) > 0 {
			pulse := queue.Pop()

			module, ok := modules[pulse.dst]
			if !ok {
				continue
			}

			if pulse.dst == feed && pulse.value == 1 {
				seen[pulse.src]++
				if _, ok := cycleLenghts[pulse.src]; !ok {
					cycleLenghts[pulse.src] = presses
				}

				seenAll := true
				for _, s := range seen {
					if s == 0 {
						seenAll = false
						break
					}
				}

				if seenAll {
					x := 1
					for _, cycleLen := range cycleLenghts {
						x *= cycleLen / gcd(x, cycleLen)
					}
					return x
				}
			}

			sends := true
			switch module.moduleType {
			case ModuleTypeFlipFlop:
				if pulse.value > 0 {
					sends = false
				} else {
					if module.value > 0 {
						module.value = 0
					} else {
						module.value = 1
					}
				}
			case ModuleTypeConjuction:
				module.state[pulse.src] = pulse.value
				module.value = 0
				for _, s := range module.state {
					if s == 0 {
						module.value = 1
						break
					}
				}
			default:
				module.value = pulse.value
			}

			if sends {
				for _, con := range module.connections {
					queue.Add(QueueData{src: pulse.dst, dst: con, value: module.value})
				}
			}
		}
	}
}

func solvePuzzle01() {
	input := getInput()

	modules := parseInput(input)

	lo := 0
	hi := 0
	for i := 0; i < 1000; i++ {
		l, h := processPush(modules)
		lo += l
		hi += h
	}

	fmt.Printf("Total pulses: %d\n", lo*hi)
}

func solvePuzzle02() {
	input := getInput()

	modules := parseInput(input)

	presses := countUntilRX(modules)
	fmt.Printf("Button presses: %d\n", presses)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
