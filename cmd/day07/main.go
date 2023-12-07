package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type HandType int

const (
	HandTypeFiveOfKind HandType = iota
	HandTypeFourOfKind
	HandTypeFullHouse
	HandTypeThreeOfKind
	HandTypeTwoPair
	HandTypeOnePair
	HandTypeHighCard
)

type Hand struct {
	hand     string
	handType HandType
	bid      int
}

var cardValues map[byte]int = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardValues2 map[byte]int = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getCardValue(card byte) int {
	return cardValues[card]
}

func getCardValue2(card byte) int {
	return cardValues2[card]
}

func parseHandType(hand string) HandType {
	handType := HandTypeHighCard

	cards := map[rune]int{}
	for _, card := range hand {
		cards[card]++
	}

	max := 0
	for _, card := range cards {
		switch card {
		case 5:
			handType = HandTypeFiveOfKind
		case 4:
			if card > max {
				handType = HandTypeFourOfKind
			}
		case 3:
			if handType == HandTypeOnePair {
				handType = HandTypeFullHouse
			} else if card > max {
				handType = HandTypeThreeOfKind
			}
		case 2:
			if handType == HandTypeThreeOfKind {
				handType = HandTypeFullHouse
			} else if card >= max {
				if handType == HandTypeOnePair {
					handType = HandTypeTwoPair
				} else {
					handType = HandTypeOnePair
				}
			}
		}

		if card > max {
			max = card
		}
	}

	return handType
}

func parseHandType2(hand string) HandType {
	handType := HandTypeHighCard

	cards := map[rune]int{}
	for _, card := range hand {
		cards[card]++
	}

	max := 0
	for key, card := range cards {
		if key == 'J' {
			continue
		}

		switch card {
		case 5:
			handType = HandTypeFiveOfKind
		case 4:
			if card > max {
				handType = HandTypeFourOfKind
			}
		case 3:
			if handType == HandTypeOnePair {
				handType = HandTypeFullHouse
			} else if card > max {
				handType = HandTypeThreeOfKind
			}
		case 2:
			if handType == HandTypeThreeOfKind {
				handType = HandTypeFullHouse
			} else if card >= max {
				if handType == HandTypeOnePair {
					handType = HandTypeTwoPair
				} else {
					handType = HandTypeOnePair
				}
			}
		}

		if card > max {
			max = card
		}
	}

	if cards['J'] > 0 {
		switch handType {
		case HandTypeFourOfKind:
			handType = HandTypeFiveOfKind
		case HandTypeThreeOfKind:
			if cards['J'] > 1 {
				handType = HandTypeFiveOfKind
			} else {
				handType = HandTypeFourOfKind
			}
		case HandTypeTwoPair:
			handType = HandTypeFullHouse
		case HandTypeOnePair:
			if cards['J'] > 2 {
				handType = HandTypeFiveOfKind
			} else if cards['J'] > 1 {
				handType = HandTypeFourOfKind
			} else {
				handType = HandTypeThreeOfKind
			}
		case HandTypeHighCard:
			if cards['J'] > 3 {
				handType = HandTypeFiveOfKind
			} else if cards['J'] > 2 {
				handType = HandTypeFourOfKind
			} else if cards['J'] > 1 {
				handType = HandTypeThreeOfKind
			} else {
				handType = HandTypeOnePair
			}
		}
	}

	return handType
}

func compareHands(l, r Hand) int {
	if l.handType < r.handType {
		return 1
	}

	if l.handType > r.handType {
		return -1
	}

	for idx := range l.hand {
		lVal := getCardValue(l.hand[idx])
		rVal := getCardValue(r.hand[idx])

		cmp := lVal - rVal
		if cmp != 0 {
			return cmp
		}
	}

	return 0
}

func compareHands2(l, r Hand) int {
	if l.handType < r.handType {
		return 1
	}

	if l.handType > r.handType {
		return -1
	}

	for idx := range l.hand {
		lVal := getCardValue2(l.hand[idx])
		rVal := getCardValue2(r.hand[idx])

		cmp := lVal - rVal
		if cmp != 0 {
			return cmp
		}
	}

	return 0
}

func calcTotalWinnings(hands []Hand) int {
	res := 0

	for idx, hand := range hands {
		win := (idx + 1) * hand.bid
		res += win
	}

	return res
}

func parseLine(line string) Hand {
	res := Hand{}

	data := strings.Split(line, " ")
	res.hand = strings.TrimSpace(data[0])
	res.handType = parseHandType(res.hand)

	bid, _ := strconv.Atoi(strings.TrimSpace(data[1]))
	res.bid = bid

	return res
}

func parseLine2(line string) Hand {
	res := Hand{}

	data := strings.Split(line, " ")
	res.hand = strings.TrimSpace(data[0])
	res.handType = parseHandType2(res.hand)

	bid, _ := strconv.Atoi(strings.TrimSpace(data[1]))
	res.bid = bid

	return res
}

func solvePuzzle01() {
	hands := []Hand{}

	input := getInput()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hand := parseLine(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return compareHands(a, b)
	})

	total := calcTotalWinnings(hands)
	fmt.Printf("Total winnings: %d\n", total)
}

func solvePuzzle02() {
	hands := []Hand{}

	input := getInput()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hand := parseLine2(line)
		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return compareHands2(a, b)
	})

	total := calcTotalWinnings(hands)
	fmt.Printf("Total winnings: %d\n", total)
}

func main() {
	solvePuzzle01()
	solvePuzzle02()
}
