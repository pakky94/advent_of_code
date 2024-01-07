package day13

import (
	"os"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day13/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	m := parseMaps(content)
	tot := 0
	for _, x := range m {
		r := findSmudgedReflection(x)
		if !r.isVertical {
			tot += 100 * (r.pos + 1)
		} else {
			tot += r.pos + 1
		}
	}
	println("p1", tot)
}

func p1(content string) {
	m := parseMaps(content)
	tot := 0
	for _, x := range m {
		r := findReflection(x)
		if !r.isVertical {
			tot += 100 * (r.pos + 1)
		} else {
			tot += r.pos + 1
		}
	}
	println("p1", tot)
}

func findSmudgedReflection(m Map) Reflection {
	smudged := false
Row:
	for y := 0; y < len(m.pos)-1; y++ {
		smudged = false
		for y2 := y; y2 >= 0; y2-- {
			for x := 0; x < len(m.pos[0]); x++ {
				if (y + y + 1 - y2) >= len(m.pos) {
					if smudged {
						return Reflection{
							isVertical: false,
							pos:        y,
						}
					} else {
						continue Row
					}
				}
				if m.pos[y2][x] != m.pos[y+y+1-y2][x] {
					if smudged {
						continue Row
					} else {
						smudged = true
					}
				}
			}
		}
		if smudged {
			return Reflection{
				isVertical: false,
				pos:        y,
			}
		}
	}

Col:
	for x := 0; x < len(m.pos[0])-1; x++ {
		smudged = false
		for x2 := x; x2 >= 0; x2-- {
			for y := 0; y < len(m.pos); y++ {
				if (x + x + 1 - x2) >= len(m.pos[0]) {
					if smudged {
						return Reflection{
							isVertical: true,
							pos:        x,
						}
					} else {
						continue Col
					}
				}
				if m.pos[y][x2] != m.pos[y][x+x+1-x2] {
					if smudged {
						continue Col
					} else {
						smudged = true
					}
				}
			}
		}
		if smudged {
			return Reflection{
				isVertical: true,
				pos:        x,
			}
		}
	}

	return Reflection{pos: -1}
}

func findReflection(m Map) Reflection {
Row:
	for y := 0; y < len(m.pos)-1; y++ {
		for y2 := y; y2 >= 0; y2-- {
			for x := 0; x < len(m.pos[0]); x++ {
				if (y + y + 1 - y2) >= len(m.pos) {
					return Reflection{
						isVertical: false,
						pos:        y,
					}
				}
				if m.pos[y2][x] != m.pos[y+y+1-y2][x] {
					continue Row
				}
			}
		}
		return Reflection{
			isVertical: false,
			pos:        y,
		}
	}

Col:
	for x := 0; x < len(m.pos[0])-1; x++ {
		for x2 := x; x2 >= 0; x2-- {
			for y := 0; y < len(m.pos); y++ {
				if (x + x + 1 - x2) >= len(m.pos[0]) {
					return Reflection{
						isVertical: true,
						pos:        x,
					}
				}
				if m.pos[y][x2] != m.pos[y][x+x+1-x2] {
					continue Col
				}
			}
		}
		return Reflection{
			isVertical: true,
			pos:        x,
		}
	}

	return Reflection{pos: -1}
}

func parseMaps(content string) []Map {
	res := make([]Map, 0)
	lines := make([][]byte, 0)
	for _, l := range strings.Split(content, "\n") {
		if len(l) == 0 {
			if len(lines) > 0 {
				res = append(res, Map{
					pos: lines,
				})
			}
			lines = make([][]byte, 0)
		} else {
			x := ([]byte)(l)
			lines = append(lines, x)
		}
	}
	return res
}

type Map struct {
	pos [][]byte
}

type Reflection struct {
	isVertical bool
	pos        int
}
