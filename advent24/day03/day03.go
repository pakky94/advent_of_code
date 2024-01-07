package day03

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	p1()
	p2()
}

func p2() {
	f, _ := os.ReadFile("day03/input")
	content := string(f)

	//    content := `467..114..
	//...*......
	//..35..633.
	//......#...
	//617*......
	//.....+.58.
	//..592.....
	//......755.
	//...$.*....
	//.664.598..`

	runP2(content)
}

func runP2(input string) {
	gameMap := parseMap(input)
	tot := 0
	for _, s := range gameMap.symbols {
		tot += gearRatio(s, gameMap.nums)
	}

	println("p2", tot)
}

func gearRatio(symbol Symbol, nums []Num) int {
	if symbol.val != '*' {
		return 0
	}

	nears := make([]Num, 0)

	for _, n := range nums {
		if areNear(n, symbol) {
			nears = append(nears, n)
		}
	}

	if len(nears) == 2 {
		return nears[0].val * nears[1].val
	}

	return 0
}

func p1() {
	f, _ := os.ReadFile("day03/input")
	content := string(f)

	//    content := `467..114..
	//...*......
	//..35..633.
	//......#...
	//617*......
	//.....+.58.
	//..592.....
	//......755.
	//...$.*....
	//.664.598..`

	runP1(content)
}

func runP1(input string) {
	gameMap := parseMap(input)
	tot := 0
	for _, n := range gameMap.nums {
		if isNearSymbol(n, gameMap.symbols) {
			tot += n.val
		}
	}

	println("p1", tot)
}

func isNearSymbol(num Num, symbols []Symbol) bool {
	for _, s := range symbols {
		if areNear(num, s) {
			return true
		}
	}

	return false
}

func areNear(num Num, symbol Symbol) bool {
	for x := num.start; x <= num.end; x++ {
		if symbol.row <= num.row+1 &&
			symbol.row >= num.row-1 &&
			symbol.col <= x+1 &&
			symbol.col >= x-1 {
			return true
		}
	}

	return false
}

func parseMap(input string) Map {
	nums := make([]Num, 0)
	symbols := make([]Symbol, 0)

	numStart := -1
	currentNum := -1

	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if unicode.IsDigit(c) {
				val, _ := strconv.ParseInt(string(c), 10, 8)
				n := int(val)

				if currentNum == -1 {
					numStart = x
					currentNum = n
				} else {
					currentNum = (currentNum * 10) + n
				}
				continue
			}

			if currentNum != -1 {
				nums = append(nums, Num{
					val:   currentNum,
					row:   y,
					start: numStart,
					end:   x - 1,
				})

				numStart = -1
				currentNum = -1
			}

			if c != '.' {
				symbols = append(symbols, Symbol{
					val: c,
					row: y,
					col: x,
				})
			}
		}

		if currentNum != -1 {
			nums = append(nums, Num{
				val:   currentNum,
				row:   y,
				start: numStart,
				end:   len(line) - 1,
			})

			numStart = -1
			currentNum = -1
		}
	}

	return Map{
		nums:    nums,
		symbols: symbols,
	}
}

type Map struct {
	nums    []Num
	symbols []Symbol
}

type Num struct {
	val   int
	row   int
	start int
	end   int
}

type Symbol struct {
	val rune
	row int
	col int
}
