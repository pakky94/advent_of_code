package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func day02() {
	d02p1()
	d02p2()
}

func d02p1() {
	f, err := os.Open("input02")
	check(err)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	res := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game := parseCubes(scanner.Text())
		if isGameValid(game) {
			res += game.id
		}
	}

	println("p1 ", res)
}

func d02p2() {
	f, err := os.Open("input02")
	check(err)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	res := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		game := parseCubes(scanner.Text())
		res += cubePower(game)
	}

	println("p1 ", res)
}

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int
	blue  int
	green int
}

func parseCubes(line string) Game {
	s := strings.Split(line, ":")

	i := strings.Split(s[0], " ")
	id, _ := strconv.ParseInt(i[1], 10, 64)

	sets := make([]Set, 0)

	for _, set := range strings.Split(s[1], ";") {
		sets = append(sets, parseSet(set))
	}

	return Game{
		int(id),
		sets,
	}
}

func parseSet(input string) Set {
	set := Set{
		red:   0,
		blue:  0,
		green: 0,
	}

	for _, cubes := range strings.Split(strings.Trim(input, " "), ",") {
		x := strings.Split(strings.Trim(cubes, " "), " ")
		color := strings.Trim(x[1], " ")
		n, _ := strconv.ParseInt(strings.Trim(x[0], " "), 10, 64)
		//println(cubes, "'", x[0], "'", x[1], "color: '", color, "'", "n:", n)
		switch color {
		case "red":
			set.red += int(n)
		case "blue":
			set.blue += int(n)
		case "green":
			set.green += int(n)
		}
	}

	//println(input, " - red", set.red, "blue", set.blue, "green", set.green)

	return set
}

func isGameValid(game Game) bool {
	for _, set := range game.sets {
		if set.green > 13 || set.red > 12 || set.blue > 14 {
			return false
		}
	}

	return true
}

func cubePower(game Game) int {
	r, g, b := 0, 0, 0

	for _, set := range game.sets {
		r = max(r, set.red)
		g = max(g, set.green)
		b = max(b, set.blue)
	}

	return r * g * b
}
