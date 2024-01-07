package day18

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day18/input")
	content := string(f)
	_ = content

	//p1(content)
	p2(testContent)
}

func p1(content string) {
	instrs := parseInstructions(content)
	m := buildPerimeter(instrs)

	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for p := range m {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}

	count := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, exists := m[Pos{x, y}]
			if exists {
				count++
			} else {
				crossings := 0
				botCross, topCross := false, false
				for i := x - 1; i >= minX; i-- {
					p, e := m[Pos{i, y}]
					if !e {
						botCross, topCross = false, false
						continue
					}

					switch p.vertDir {
					case MIDDLE:
						crossings++
					case BOTTOM:
						if topCross {
							crossings++
							topCross = false
						} else {
							botCross = !botCross
						}
					case TOP:
						if botCross {
							crossings++
							botCross = false
						} else {
							topCross = !topCross
						}
					}
				}
				if crossings%2 == 1 {
					count++
				}
			}
		}
	}

	println("p1", count)
}

func buildPerimeter(instrs []Instr) map[Pos]MapPoint {
	m := map[Pos]MapPoint{}
	x, y := 0, 0

	for i := 0; i < len(instrs); i++ {
		//lastInstr := instrs[(i+len(instrs)-1)%len(instrs)]
		nextInstr := instrs[(i+1)%len(instrs)]

		for j := 0; j < instrs[i].len; j++ {
			x, y = movePos(x, y, instrs[i])
			p := Pos{x, y}
			vertDir := NONE

			if j == instrs[i].len-1 {
				if instrs[i].dir == LEFT || instrs[i].dir == RIGHT {
					if nextInstr.dir == UP {
						vertDir = TOP
					} else if nextInstr.dir == DOWN {
						vertDir = BOTTOM
					}
					//} else if instrs[i].dir == RIGHT {
					//	if nextInstr.dir == UP {
					//		vertDir = BOTTOM
					//	} else if nextInstr.dir == DOWN {
					//		vertDir = TOP
					//	}
				}
			}

			if instrs[i].dir == UP {
				if j == instrs[i].len-1 {
					vertDir = BOTTOM
				} else {
					vertDir = MIDDLE
				}
			}

			if instrs[i].dir == DOWN {
				if j == instrs[i].len-1 {
					vertDir = TOP
				} else {
					vertDir = MIDDLE
				}
			}

			existing, e := m[p]
			if e && existing.vertDir != NONE {
				vertDir = existing.vertDir
			}

			m[p] = MapPoint{vertDir}
		}
	}

	return m
}

func movePos(x int, y int, instr Instr) (int, int) {
	o := dirToOffset[instr.dir]
	return x + o.x, y + o.y
}

var dirToOffset = map[uint8]Pos{
	UP:    {0, -1},
	DOWN:  {0, 1},
	RIGHT: {1, 0},
	LEFT:  {-1, 0},
}

type MapPoint struct {
	vertDir uint8
}

type Pos struct {
	x int
	y int
}

func parseInstructions(content string) []Instr {
	var instrs []Instr
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			continue
		}
		x := strings.Split(line, " ")

		var dir uint8
		switch x[0] {
		case "R":
			dir = RIGHT
		case "L":
			dir = LEFT
		case "U":
			dir = UP
		case "D":
			dir = DOWN
		}

		n, _ := strconv.ParseInt(x[1], 10, 64)

		instrs = append(instrs, Instr{
			dir: dir,
			len: int(n),
		})
	}
	return instrs
}

const (
	BOTTOM = uint8(iota)
	MIDDLE
	TOP
	HORIZ
	NONE
)

type Instr struct {
	dir uint8
	len int
}

const (
	RIGHT = uint8(iota)
	DOWN
	LEFT
	UP
)

var testContent = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`
