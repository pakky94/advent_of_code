package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func day01() {
	d01p1()
	d01p2()
}

func d01p1() {
	f, err := os.Open("input01")
	check(err)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	scanner := bufio.NewScanner(f)
	tot := 0
	for scanner.Scan() {
		tot += code(scanner.Text())
	}

	println("part 1: ", tot)
}

func d01p2() {
	f, err := os.Open("input01")
	check(err)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	scanner := bufio.NewScanner(f)
	tot := 0
	for scanner.Scan() {
		tot += code2(scanner.Text())
	}
	println("part 2: ", tot)
}

func code(input string) int {
	first, last := 0, 0

	for _, c := range input {
		if unicode.IsDigit(c) {
			x, _ := strconv.ParseInt(string(c), 10, 8)
			first = int(x)
			break
		}
	}

	for _, c := range input {
		if unicode.IsDigit(c) {
			x, _ := strconv.ParseInt(string(c), 10, 8)
			last = int(x)
		}
	}

	return first*10 + last
}

var nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
}

func code2(input string) int {
	matches := make(map[int]int)

	for s, val := range nums {
		index := strings.Index(input, s)
		if index >= 0 {
			matches[index] = val
		}

		index = strings.LastIndex(input, s)
		if index >= 0 {
			matches[index] = val
		}
	}

	first, last := 0, 0
	minI, maxI := 999999999999, -1

	for idx, val := range matches {
		if idx < minI {
			first = val
			minI = idx
		}

		if idx > maxI {
			last = val
			maxI = idx
		}
	}

	res := first*10 + last
	return res
}
