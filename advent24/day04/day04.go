package day04

import (
	"os"
	"strconv"
	"strings"
)

var testContent = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func Run() {
	f, _ := os.ReadFile("day04/input")
	content := string(f)

	p1(content)
	p2(content)
}

func p2(content string) {
	cards := parseCards(content)
	counts := make([]int, len(cards))

	for i := 0; i < len(cards); i++ {
		counts[i] = 1
	}

	for i := 0; i < len(cards); i++ {
		wins := winningCount(cards[i])

		for x := 0; x < wins; x++ {
			counts[i+x+1] += counts[i]
		}
	}

	tot := 0
	for _, n := range counts {
		tot += n
	}
	println("p2", tot)
}

func p1(content string) {
	cards := parseCards(content)
	tot := 0
	for _, card := range cards {
		tot += value(card)
	}
	println("p1", tot)
}

func winningCount(card Card) int {
	val := 0

	for _, actual := range card.actual {
		for _, winning := range card.winning {
			if actual == winning {
				val += 1
			}
		}
	}

	return val
}

func value(card Card) int {
	val := 0

	for _, actual := range card.actual {
		for _, winning := range card.winning {
			if actual == winning {
				if val == 0 {
					val = 1
				} else {
					val = val * 2
				}
			}
		}
	}

	return val
}

func parseCards(input string) []Card {
	cards := make([]Card, 0)

	for _, line := range strings.Split(input, "\n") {
		if len(line) != 0 {
			cards = append(cards, parseCard(line))
		}
	}

	return cards
}

func parseCard(line string) Card {
	s := strings.Split(line, ":")

	i := strings.Split(s[0], " ")
	id := 0
	for _, t := range i[1:] {
		tId, _ := strconv.ParseInt(t, 10, 64)
		id = int(tId)
	}

	winning := make([]int, 0)
	actual := make([]int, 0)

	nums := strings.Split(s[1], "|")

	for _, nS := range strings.Split(strings.Trim(nums[0], " "), " ") {
		n, _ := strconv.ParseInt(strings.Trim(nS, " "), 10, 64)
		if n != 0 {
			winning = append(winning, int(n))
		}
	}

	for _, nS := range strings.Split(strings.Trim(nums[1], " "), " ") {
		n, _ := strconv.ParseInt(strings.Trim(nS, " "), 10, 64)
		if n != 0 {
			actual = append(actual, int(n))
		}
	}

	return Card{
		id:      id,
		winning: winning,
		actual:  actual,
	}
}

type Card struct {
	id      int
	winning []int
	actual  []int
}
