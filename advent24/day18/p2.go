package day18

import (
	"slices"
	"strconv"
	"strings"
)

func p2(content string) {
	instrs := parseInstructions2(content)
	blocks := buildGrid(instrs)
	for _, row := range blocks {
		s := make([]byte, len(row))
		for ix, x := range row {
			switch x.t {
			case NONE:
				s[ix] = '.'
			case BOTTOM:
				s[ix] = 'B'
			case MIDDLE:
				s[ix] = 'M'
			case TOP:
				s[ix] = 'T'
			case HORIZ:
				s[ix] = 'H'
			}
		}
		println(string(s))
	}

	count := 0
	for y := 0; y < len(blocks); y++ {
		for x := 0; x < len(blocks); x++ {
			block := blocks[y][x]
			if block.t != NONE {
				count += block.size
			} else {
				crossings := 0
				botCross, topCross := false, false
				for i := x - 1; i >= 0; i-- {
					b := blocks[y][i]

					switch b.t {
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
					//case HORIZ:
					//	continue
					case NONE:
						botCross, topCross = false, false
					}
				}
				if crossings%2 == 1 {
					count += block.size
				}
			}
		}
	}

	println("p2", count)
}

type Block struct {
	t    uint8
	size int
}

func buildGrid(instrs []Instr) [][]Block {
	xs := make([]int, 0)
	ys := make([]int, 0)
	x, y := 0, 0
	lines := make([]Line, 0)

	for i := 0; i < len(instrs); i++ {
		x1, y1 := x, y
		instr := instrs[i]
		nextInstr := instrs[(i+1)%len(instrs)]
		xs = append(xs, x)
		xs = append(xs, x+1)
		ys = append(ys, y)
		ys = append(ys, y+1)
		if instr.dir == RIGHT {
			x1++
			x += instr.len
			xs = append(xs, x)
			xs = append(xs, x-1)
		} else if instr.dir == LEFT {
			x1--
			x -= instr.len
			xs = append(xs, x+1)
			xs = append(xs, x)
		} else if instr.dir == UP {
			y1--
			y -= instr.len
			ys = append(ys, y+1)
			ys = append(ys, y)
		} else if instr.dir == DOWN {
			y1++
			y += instr.len
			ys = append(ys, y)
			ys = append(ys, y-1)
		}
		lines = append(lines, Line{
			x1:      min(x1, x),
			y1:      min(y1, y),
			x2:      max(x1, x),
			y2:      max(y1, y),
			dir:     instr.dir,
			nextDir: nextInstr.dir,
		})
	}

	slices.Sort(xs)
	xs = slices.Compact(xs)
	slices.Sort(ys)
	ys = slices.Compact(ys)

	_ = lines
	blocks := make([][]Block, len(ys))
	for i := 0; i < len(ys); i++ {
		blocks[i] = make([]Block, len(xs))
	}

	for iy := 0; iy < len(ys); iy++ {
		for ix := 0; ix < len(xs); ix++ {
			matchingLine, found := findLine(lines, xs[ix], ys[iy])
			size := blockSize(xs, ys, ix, iy)
			if found {
				blocks[iy][ix] = Block{
					t:    blockType(matchingLine, xs[ix], ys[iy]),
					size: size,
				}
			} else {
				blocks[iy][ix] = Block{
					t:    NONE,
					size: size,
				}
			}
		}
	}

	return blocks
}

func blockSize(xs []int, ys []int, ix int, iy int) int {
	var w int
	if ix+1 < len(xs) {
		w = xs[ix+1] - xs[ix]
	} else {
		w = 1
	}

	var h int
	if iy+1 < len(ys) {
		h = ys[iy+1] - ys[iy]
	} else {
		h = 1
	}
	return w * h
}

func blockType(line Line, x int, y int) uint8 {
	vertDir := HORIZ

	if (x == line.x2 && line.dir == RIGHT) ||
		(x == line.x1 && line.dir == LEFT) {
		if line.nextDir == UP {
			vertDir = TOP
		} else if line.nextDir == DOWN {
			vertDir = BOTTOM
		}
	}

	if line.dir == UP {
		if y == line.y1 {
			vertDir = BOTTOM
		} else {
			vertDir = MIDDLE
		}
	}

	if line.dir == DOWN {
		if y == line.y2 {
			vertDir = TOP
		} else {
			vertDir = MIDDLE
		}
	}

	return vertDir
}

func findLine(lines []Line, x int, y int) (Line, bool) {
	f := Line{}
	count := 0
	for _, l := range lines {
		if x >= l.x1 && x <= l.x2 && y >= l.y1 && y <= l.y2 {
			//return l, true
			count++
			f = l
		}
	}
	if count > 1 {
		panic("wrong match number")
	}
	return f, count > 0
	//return Line{}, false
}

func parseInstructions2(content string) []Instr {
	var instrs []Instr
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			continue
		}
		x := strings.Split(line, " ")

		var dir uint8
		switch x[2][7:8] {
		case "0":
			dir = RIGHT
		case "2":
			dir = LEFT
		case "3":
			dir = UP
		case "1":
			dir = DOWN
		}

		n, _ := strconv.ParseInt(x[2][2:7], 16, 64)

		instrs = append(instrs, Instr{
			dir: dir,
			len: int(n),
		})
	}
	return instrs
}

type Line struct {
	x1      int
	y1      int
	x2      int
	y2      int
	dir     uint8
	nextDir uint8
}
