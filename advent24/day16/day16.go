package day16

import (
	"os"
	"slices"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day16/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	m := parseMap(content)
	ill := 0

	for i := 0; i < len(m); i++ {
		t := calcIlluminated(m, RayPoint{0, i, Right})
		if t > ill {
			ill = t
		}
		t = calcIlluminated(m, RayPoint{len(m[0]) - 1, i, Left})
		if t > ill {
			ill = t
		}
	}

	for i := 0; i < len(m); i++ {
		t := calcIlluminated(m, RayPoint{i, 0, Down})
		if t > ill {
			ill = t
		}
		t = calcIlluminated(m, RayPoint{i, len(m) - 1, Up})
		if t > ill {
			ill = t
		}
	}

	println("p2", ill)
}

func p1(content string) {
	m := parseMap(content)
	println("p1", calcIlluminated(m, RayPoint{0, 0, Right}))
}

type Offset struct {
	x int
	y int
}

var dirToOffset = map[uint8]Offset{
	Left:  {-1, 0},
	Right: {1, 0},
	Up:    {0, -1},
	Down:  {0, 1},
}

var interactions = map[byte]map[uint8][]uint8{
	'\\': {
		Up:    {Left},
		Down:  {Right},
		Left:  {Up},
		Right: {Down},
	},
	'/': {
		Down:  {Left},
		Up:    {Right},
		Right: {Up},
		Left:  {Down},
	},
	'-': {
		Up:    {Left, Right},
		Down:  {Left, Right},
		Right: {Right},
		Left:  {Left},
	},
	'|': {
		Right: {Up, Down},
		Left:  {Up, Down},
		Up:    {Up},
		Down:  {Down},
	},
	'.': {
		Right: {Right},
		Left:  {Left},
		Up:    {Up},
		Down:  {Down},
	},
}

func calcIlluminated(m []string, start RayPoint) int {
	calculated := map[RayPoint]bool{}
	illuminated := map[Point]bool{}

	s := []RayPoint{start}
	for len(s) > 0 {
		p := s[len(s)-1]
		s = s[:len(s)-1]

		if p.x < 0 || p.y < 0 || p.x >= len(m[0]) || p.y >= len(m) {
			continue
		}

		_, c := calculated[p]
		if c {
			continue
		}

		calculated[p] = true
		illuminated[Point{p.x, p.y}] = true

		for _, newDir := range interactions[m[p.y][p.x]][p.dir] {
			offset := dirToOffset[newDir]
			nP := RayPoint{
				x:   p.x + offset.x,
				y:   p.y + offset.y,
				dir: newDir,
			}
			s = append(s, nP)
		}

		/*
			switch m[p.y][p.x] {
			case '\\':
				switch p.dir {
				case Up:
					s = append(s, RayPoint{p.x - 1, p.y, Left})
				case Down:
					s = append(s, RayPoint{p.x + 1, p.y, Right})
				case Left:
					s = append(s, RayPoint{p.x, p.y - 1, Up})
				case Right:
					s = append(s, RayPoint{p.x, p.y + 1, Down})
				}
			case '/':
				switch p.dir {
				case Down:
					s = append(s, RayPoint{p.x - 1, p.y, Left})
				case Up:
					s = append(s, RayPoint{p.x + 1, p.y, Right})
				case Right:
					s = append(s, RayPoint{p.x, p.y - 1, Up})
				case Left:
					s = append(s, RayPoint{p.x, p.y + 1, Down})
				}
			case '-':
				switch p.dir {
				case Down:
					s = append(s, RayPoint{p.x - 1, p.y, Left})
					s = append(s, RayPoint{p.x + 1, p.y, Right})
				case Up:
					s = append(s, RayPoint{p.x - 1, p.y, Left})
					s = append(s, RayPoint{p.x + 1, p.y, Right})
				case Right:
					s = append(s, RayPoint{p.x + 1, p.y, Right})
				case Left:
					s = append(s, RayPoint{p.x - 1, p.y, Left})
				}
			case '|':
				switch p.dir {
				case Down:
					s = append(s, RayPoint{p.x, p.y + 1, Down})
				case Up:
					s = append(s, RayPoint{p.x, p.y - 1, Up})
				case Right:
					s = append(s, RayPoint{p.x, p.y - 1, Up})
					s = append(s, RayPoint{p.x, p.y + 1, Down})
				case Left:
					s = append(s, RayPoint{p.x, p.y - 1, Up})
					s = append(s, RayPoint{p.x, p.y + 1, Down})
				}
			case '.':
				switch p.dir {
				case Down:
					s = append(s, RayPoint{p.x, p.y + 1, Down})
				case Up:
					s = append(s, RayPoint{p.x, p.y - 1, Up})
				case Right:
					s = append(s, RayPoint{p.x + 1, p.y, Right})
				case Left:
					s = append(s, RayPoint{p.x - 1, p.y, Left})
				}
			}
		*/
	}

	return len(illuminated)
}

type RayPoint struct {
	x   int
	y   int
	dir uint8
}

type Point struct {
	x int
	y int
}

const (
	Up = uint8(iota)
	Down
	Left
	Right
)

func parseMap(content string) []string {
	lines := strings.Split(content, "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool {
		return len(s) == 0
	})
	return lines
}

var testContent = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`
