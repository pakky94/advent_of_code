package day17

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day17/input")
	content := string(f)
	_ = content

	p1(content)
	//	p2(`
	//111111111111
	//999999999991
	//999999999991
	//999999999991
	//999999999991`)
	p2(content)
}

func p1(content string) {
	println("p1", internalCalc(content, standardValidate, 1))
}

func p2(content string) {
	println("p2", internalCalc(content, ultraValidate, 4))
}

func ultraValidate(point MovPoint, m Map, dir uint8) (MovPoint, bool) {
	if point.len >= 10 && point.dir == dir {
		return MovPoint{}, false
	}

	if point.len <= 3 && point.dir != dir {
		return MovPoint{}, false
	}

	l := point.len + 1
	if dir != point.dir {
		l = 1
	}

	offset := dirToOffset[dir]
	p := MovPoint{
		x:   point.x + offset.x,
		y:   point.y + offset.y,
		dir: dir,
		len: l,
	}

	if p.x < 0 || p.y < 0 || p.x >= len(m.costs[0]) || p.y >= len(m.costs) {
		return MovPoint{}, false
	}

	return p, true
}

func internalCalc(content string,
	validateNextPoint func(point MovPoint, m Map, dir uint8) (MovPoint, bool),
	finishMinLen uint8,
) int {
	m := parseMap(content)

	endX := len(m.costs[0]) - 1
	endY := len(m.costs) - 1

	var mem = map[MovPoint]int{}
	var foundCount = map[MovPoint]uint8{}
	var nextPoint = map[MovPoint]MovPoint{}

	for d := uint8(0); d < 4; d++ {
		for l := finishMinLen; l < 11; l++ {
			mem[MovPoint{
				y:   endY,
				x:   endX,
				dir: d,
				len: l,
			}] = int(m.costs[endY][endX])
		}
	}

	done := false
	for !done {
		done = true
		for y := endY; y >= 0; y-- {
			for x := endX; x >= 0; x-- {
				for d := uint8(0); d < 4; d++ {
					for l := uint8(1); l < 11; l++ {
						point := MovPoint{x, y, d, l}
						old, exists := mem[point]
						c, v, found, next := pointCost(point, m, mem, validateNextPoint)
						foundC, foundE := foundCount[point]
						if v && (!exists || old > c || !foundE || found > foundC) {
							done = false
							mem[point] = c
							foundCount[point] = found
							nextPoint[point] = next
						}
					}
				}
			}
		}
		_ = mem
	}

	startCost := int(m.costs[0][0])
	minCost := math.MaxInt64
	//minStart := MovPoint{}

	for d := uint8(0); d < 4; d++ {
		//for l := uint8(1); l < 11; l++ {
		p := MovPoint{0, 0, d, 1}
		t, found := mem[p]
		//println(d, l, t, found)
		//if !found {
		//	panic("match not found")
		//}
		if found && t < minCost {
			minCost = t
			//minStart = p
		}
		//}
	}

	//currPoint := minStart
	//for currPoint.x != endX || currPoint.y != endY {
	//	currPoint = nextPoint[currPoint]
	//	println(currPoint.x, "\t", currPoint.y, "\t", mem[currPoint])
	//}

	return minCost - startCost
}

func pointCost(point MovPoint, m Map, mem map[MovPoint]int,
	validateNextPoint func(point MovPoint, m Map, dir uint8) (MovPoint, bool),
) (int, bool, uint8, MovPoint) {
	thisCost := int(m.costs[point.y][point.x])

	nFound := uint8(0)
	found := false
	next := MovPoint{}
	minC := math.MaxInt64
	for _, dir := range allowedDirs[point.dir] {
		nP, valid := validateNextPoint(point, m, dir)
		if valid {
			t, exists := mem[nP]
			if !exists {
				continue
			}
			nFound++
			if t < minC {
				next = nP
				found = true
				minC = t
			}
		}
	}

	if !found {
		return -1, false, 0, next
	}

	return thisCost + minC, true, nFound, next
}

func standardValidate(point MovPoint, m Map, dir uint8) (MovPoint, bool) {
	if point.len >= 3 && point.dir == dir {
		return MovPoint{}, false
	}

	l := point.len + 1
	if dir != point.dir {
		l = 1
	}

	offset := dirToOffset[dir]
	p := MovPoint{
		x:   point.x + offset.x,
		y:   point.y + offset.y,
		dir: dir,
		len: l,
	}

	if p.x < 0 || p.y < 0 || p.x >= len(m.costs[0]) || p.y >= len(m.costs) {
		return MovPoint{}, false
	}

	return p, true
}

type PointLog struct {
	cost    int
	success bool
}

type MovPoint struct {
	x   int
	y   int
	dir uint8
	len uint8
}

var dirToOffset = map[uint8]Offset{
	LEFT:  {-1, 0},
	RIGHT: {1, 0},
	UP:    {0, -1},
	DOWN:  {0, 1},
}

var allowedDirs = map[uint8][]uint8{
	UP:    {UP, RIGHT, LEFT},
	DOWN:  {DOWN, RIGHT, LEFT},
	LEFT:  {LEFT, DOWN, UP},
	RIGHT: {RIGHT, DOWN, UP},
}

type Offset struct {
	x int
	y int
}

const (
	UP = uint8(iota)
	DOWN
	LEFT
	RIGHT
)

func parseMap(content string) Map {
	c := make([][]uint8, 0)

	for _, l := range strings.Split(content, "\n") {
		c2 := make([]uint8, 0)

		for _, x := range l {
			n, err := strconv.ParseInt(string(x), 10, 64)
			if err != nil {
				panic(err)
			}
			c2 = append(c2, uint8(n))
		}

		if len(c2) > 0 {
			c = append(c, c2)
		}
	}

	return Map{
		costs: c,
	}
}

type Map struct {
	costs [][]uint8
}

var testContent = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`
