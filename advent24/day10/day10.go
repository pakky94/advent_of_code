package day10

import (
	"os"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day10/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	m := parseMap(content)
	workMap := make([][]bool, len(m.lines))
	for i := 0; i < len(workMap); i++ {
		workMap[i] = make([]bool, len(m.lines[0]))
	}

	currX := m.startX
	currY := m.startY
	count := 0
	dX, dY := 0, 1

	for count == 0 || currX != m.startX || currY != m.startY {
		workMap[currY][currX] = true
		currX += dX
		currY += dY
		pipe := m.lines[currY][currX]
		dX, dY = calcOffset(pipe, dX, dY)
		count++
	}

	startPipe := calcStartPipe(m)
	count = 0
	for y := 0; y < len(workMap); y++ {
		for x := 0; x < len(workMap[0]); x++ {
			if isInner(x, y, m, workMap, startPipe) {
				count += 1
			}
		}
	}

	println("p2", count)
}

func calcStartPipe(m Map) uint8 {
	for _, pipe := range []uint8{'-', '7', 'F', 'J', 'L', '|'} {
		matches := 0

		for _, offset := range offsetsMap {
			tx, ty := calcOffset(pipe, offset.x, offset.y)
			if tx == 0 && ty == 0 {
				continue
			}

			target := m.lines[m.startY+ty][m.startX+tx]
			tx, ty = calcOffset(target, tx, ty)
			if tx == 0 && ty == 0 {
				continue
			}

			matches += 1
		}

		if matches == 2 {
			return pipe
		}
	}
	return ' '
}

func isInner(x int, y int, m Map, workMap [][]bool, startPipe uint8) bool {
	if workMap[y][x] {
		return false
	}

	crossings := 0
	upCross, downCross := 0, 0
	cX := x

	for cX > 0 {
		cX--
		isPipe := workMap[y][cX]
		pipe := m.lines[y][cX]

		if !isPipe {
			upCross, downCross = 0, 0
			continue
		}

		if pipe == 'S' {
			pipe = startPipe
		}

		switch pipe {
		case '|':
			crossings++
			upCross, downCross = 0, 0
		case 'F':
			if upCross == 1 {
				crossings++
			}
			upCross, downCross = 0, 0
		case '7':
			downCross = 1
			upCross = 0
		case 'J':
			upCross = 1
			downCross = 0
		case 'L':
			if downCross == 1 {
				crossings++
			}
			upCross, downCross = 0, 0
		default:
			if pipe != '-' {
				upCross, downCross = 0, 0
			}
		}
	}

	isInner := crossings%2 == 1
	//if isInner {
	//    println("x", x+1, "y", y+1)
	//}

	return isInner
}

func p1(content string) {
	m := parseMap(content)
	currX := m.startX
	currY := m.startY
	count := 0
	dX, dY := 0, 1

	for count == 0 || currX != m.startX || currY != m.startY {
		currX += dX
		currY += dY
		pipe := m.lines[currY][currX]
		dX, dY = calcOffset(pipe, dX, dY)
		count++
	}

	println("p1", count/2)
}

var pipeMap = map[struct {
	uint8
	Offset
}]Offset{
	{'F', Offset{x: 0, y: -1}}: {1, 0},
	{'F', Offset{x: -1, y: 0}}: {0, 1},
	{'-', Offset{x: -1, y: 0}}: {-1, 0},
	{'-', Offset{x: 1, y: 0}}:  {1, 0},
	{'|', Offset{x: 0, y: 1}}:  {0, 1},
	{'|', Offset{x: 0, y: -1}}: {0, -1},
	{'L', Offset{x: 0, y: 1}}:  {1, 0},
	{'L', Offset{x: -1, y: 0}}: {0, -1},
	{'7', Offset{x: 1, y: 0}}:  {0, 1},
	{'7', Offset{x: 0, y: -1}}: {-1, 0},
	{'J', Offset{x: 1, y: 0}}:  {0, -1},
	{'J', Offset{x: 0, y: 1}}:  {-1, 0},
}

var offsetsMap = []Offset{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func calcOffset(c uint8, fromX int, fromY int) (int, int) {
	k := struct {
		uint8
		Offset
	}{c, Offset{
		x: fromX,
		y: fromY,
	}}
	x, ok := pipeMap[k]
	if ok {
		return x.x, x.y
	}

	return 0, 0
}

func parseMap(content string) Map {
	lines := strings.Split(content, "\n")
	out := make([]string, 0)
	for y := 0; y < len(lines); y++ {
		if len(lines[y]) == 0 {
			continue
		}
		out = append(out, lines[y])
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'S' {
				return Map{
					lines:  out,
					startX: x,
					startY: y,
				}
			}
		}
	}
	return Map{
		lines:  lines,
		startX: -1,
		startY: -1,
	}
}

type Offset struct {
	x int
	y int
}

type Map struct {
	lines  []string
	startX int
	startY int
}

var testContent = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`
