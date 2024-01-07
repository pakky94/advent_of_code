package day07

import (
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day07/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	hands := parseHands(content, jollyCardToValue)
	sort.Slice(hands, func(i int, j int) bool {
		return cardCmpIsLess(hands[i], hands[j], jollyHandType)
	})

	tot := 0
	for i := 0; i < len(hands); i++ {
		tot += (i + 1) * hands[i].bid
	}

	println("p2", tot)
}

func p1(content string) {
	hands := parseHands(content, cardToValue)
	sort.Slice(hands, func(i int, j int) bool {
		return cardCmpIsLess(hands[i], hands[j], handType)
	})

	tot := 0
	for i := 0; i < len(hands); i++ {
		tot += (i + 1) * hands[i].bid
	}

	println("p1", tot)
}

func parseHands(content string, cardToVal func(card int32) int) []Hand {
	hands := make([]Hand, 0)

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if len(strings.Trim(line, " ")) == 0 {
			continue
		}

		x := strings.Split(line, " ")
		hand := Hand{
			source: x[0],
			cards:  []int{},
			bid:    0,
		}
		for _, card := range x[0] {
			hand.cards = append(hand.cards, cardToVal(card))
		}
		n, _ := strconv.ParseInt(x[1], 10, 64)
		hand.bid = int(n)

		hands = append(hands, hand)
	}

	return hands
}

func cardCmpIsLess(a Hand, b Hand, handType func(hand Hand) []int) bool {
	aT := handType(a)
	bT := handType(b)

	for i := 0; i < len(aT); i++ {
		if aT[i] == bT[i] {
			continue
		}

		return aT[i] < bT[i]
	}

	for i := 0; i < 5; i++ {
		if a.cards[i] == b.cards[i] {
			continue
		}

		return a.cards[i] < b.cards[i]
	}

	return true
}

func jollyHandType(hand Hand) []int {
	cards := append([]int{}, hand.cards...)
	slices.Sort(cards)
	groups := make([]int, 0)

	jollys := 0
	for i := 0; i < 5; i++ {
		if cards[i] == 1 {
			jollys++
		}
	}

	if jollys == 5 {
		return []int{5}
	}

	count := 1
	curr := cards[jollys]
	for i := 1 + jollys; i < 5; i++ {
		if curr != cards[i] {
			groups = append(groups, count)
			curr = cards[i]
			count = 1
		} else {
			count += 1
		}
	}
	groups = append(groups, count)

	slices.Sort(groups)
	slices.Reverse(groups)

	groups[0] += jollys

	return groups
}

func jollyCardToValue(card int32) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		n, _ := strconv.ParseInt(string(card), 10, 64)
		return int(n)
	}
}

func handType(hand Hand) []int {
	cards := append([]int{}, hand.cards...)
	slices.Sort(cards)
	groups := make([]int, 0)

	count := 1
	curr := cards[0]
	for i := 1; i < 5; i++ {
		if curr != cards[i] {
			groups = append(groups, count)
			curr = cards[i]
			count = 1
		} else {
			count += 1
		}
	}
	groups = append(groups, count)

	slices.Sort(groups)
	slices.Reverse(groups)

	return groups
}

func cardToValue(card int32) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		n, _ := strconv.ParseInt(string(card), 10, 64)
		return int(n)
	}
}

type Hand struct {
	source string
	cards  []int
	bid    int
}

var testContent = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
