package day11

import (
	"os"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day11/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	galaxies := parse(content, 999_999)

	tot := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			tot += distance(galaxies[i], galaxies[j])
		}
	}
	println("p2", tot)
}

func p1(content string) {
	galaxies := parse(content, 1)

	tot := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			tot += distance(galaxies[i], galaxies[j])
		}
	}
	println("p1", tot)
}

func distance(a Galaxy, b Galaxy) int {
	return max(a.x, b.x) - min(a.x, b.x) + max(a.y, b.y) - min(a.y, b.y)
}

func parse(content string, expansionFactor int) []Galaxy {
	lines := make([]string, 0)
	for _, x := range strings.Split(content, "\n") {
		if len(x) != 0 {
			lines = append(lines, x)
		}
	}

	emptyRows := make([]bool, len(lines))
	emptyCols := make([]bool, len(lines[0]))

	for i := 0; i < len(lines); i++ {
		if len(strings.Replace(lines[i], ".", "", -1)) == 0 {
			emptyRows[i] = true
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		emptyCols[i] = true
		for y := 0; y < len(lines); y++ {
			if lines[y][i] != '.' {
				emptyCols[i] = false
				break
			}
		}
	}

	galaxies := make([]Galaxy, 0)
	expX, expY := 0, 0
	for y := 0; y < len(lines); y++ {
		if emptyRows[y] {
			expY += expansionFactor
		}
		expX = 0
		for x := 0; x < len(lines[0]); x++ {
			if emptyCols[x] {
				expX += expansionFactor
			}

			if lines[y][x] == '#' {
				galaxies = append(galaxies, Galaxy{
					x: x + expX,
					y: y + expY,
				})
			}
		}
	}

	return galaxies
}

type Galaxy struct {
	x int
	y int
}

var testContent = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`
