package day14

import (
	"os"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day14/input")
	content := string(f)
	_ = content

	p1(content)
	//p2(testContent)
	p2(content)
}

func p2(content string) {
	m := parseMap(content)
	var seenIter = map[string]int{}
	var loads = make([]int, 0)
	start, end := 0, 0

	for i := 0; true; i++ {
		m = m.rollNorth().rollWest().rollSouth().rollEast()
		//println(i, "\t", m.load())
		//printMap(m)
		mKey := m.toKey()
		last, ok := seenIter[mKey]
		if ok {
			//println("found: ", last, "-", i)
			start = last
			end = i
			break
		} else {
			seenIter[mKey] = i
			loads = append(loads, m.load())
		}
	}

	load := extrapolate(start, end, 999_999_999, loads)

	println("p2", load)
}

func printMap(m Map) {
	for _, l := range m.rocks {
		s := make([]byte, len(l))
		for i, x := range l {
			switch x {
			case Empty:
				s[i] = '.'
			case Square:
				s[i] = '#'
			case Round:
				s[i] = 'O'
			}
		}
		println(string(s))
	}
}

func extrapolate(start int, end int, target int, loads []int) int {
	period := end - start
	idx := (target - start) % period
	return loads[idx+start]
}

func p1(content string) {
	println("p1", parseMap(content).rollNorth().load())
}

func (m Map) load() int {
	height := len(m.rocks)
	tot := 0

	for y := 0; y < len(m.rocks); y++ {
		for x := 0; x < len(m.rocks[0]); x++ {
			if m.rocks[y][x] == Round {
				tot += height - y
			}
		}
	}

	return tot
}

func (m Map) rollNorth() Map {
	return m.roll(0, 1, 0, -1)
}

func (m Map) rollWest() Map {
	return m.roll(1, 0, -1, 0)
}

func (m Map) rollSouth() Map {
	return m.revRoll(0, -1, 0, 1)
}

func (m Map) rollEast() Map {
	return m.revRoll(-1, 0, 1, 0)
}

func (m Map) toKey() string {
	l := make([]string, len(m.rocks))
	for i := 0; i < len(l); i++ {
		l[i] = string(m.rocks[i])
	}
	return strings.Join(l, "")
}

func (m Map) revRoll(sX int, sY int, dX int, dY int) Map {
	changed := true
	for changed {
		changed = false
		for y := len(m.rocks) + sY - 1; y >= 0; y-- {
			for x := len(m.rocks[0]) + sX - 1; x >= 0; x-- {
				if m.rocks[y][x] == Round &&
					m.rocks[y+dY][x+dX] == Empty {
					m.rocks[y][x] = Empty
					m.rocks[y+dY][x+dX] = Round
					changed = true
				}
			}
		}
	}
	return m
}

func (m Map) roll(sX int, sY int, dX int, dY int) Map {
	changed := true
	for changed {
		changed = false
		for y := sY; y < len(m.rocks); y++ {
			for x := sX; x < len(m.rocks[0]); x++ {
				if m.rocks[y][x] == Round &&
					m.rocks[y+dY][x+dX] == Empty {
					m.rocks[y][x] = Empty
					m.rocks[y+dY][x+dX] = Round
					changed = true
				}
			}
		}
	}
	return m
}

func parseMap(content string) Map {
	lines := make([][]byte, 0)
	for _, l := range strings.Split(content, "\n") {
		if len(l) == 0 {
			continue
		}
		x := make([]byte, len(l))
		for i, b := range l {
			switch b {
			case '.':
				x[i] = Empty
			case '#':
				x[i] = Square
			case 'O':
				x[i] = Round
			}
		}
		lines = append(lines, x)
	}
	return Map{
		rocks: lines,
	}
}

type Map struct {
	rocks [][]byte
}

var testContent = `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

const (
	Empty byte = iota
	Square
	Round
)
