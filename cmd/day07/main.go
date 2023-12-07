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

func getCardValue(card byte) int {
	return cardValues[card]
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

func sortHands(hands []Hand) {
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			cmp := compareHands(hands[i], hands[j])
			if cmp < 0 {
				tmp := hands[i]
				hands[i] = hands[j]
				hands[j] = tmp
			}
		}
	}
}

func calcTotalWinnings(hands []Hand) int {
	res := 0

	for idx, hand := range hands {
		fmt.Printf("%v\n", hand)
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

func solvePuzzle01() {
	hands := []Hand{}

	input := getInput()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		hand := parseLine(line)
		hands = append(hands, hand)
	}

	// sortHands(hands)
	fmt.Printf("%v\n", hands)

	slices.SortFunc(hands, func(a, b Hand) int {
		return compareHands(a, b)
	})

	total := calcTotalWinnings(hands)
	fmt.Printf("Total winnings: %d\n", total)
}

func main() {
	solvePuzzle01()
}
