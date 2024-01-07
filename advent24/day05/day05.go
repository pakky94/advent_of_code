package day05

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day05/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p1(content string) {
	maps := parseMaps(content)
	currMin := math.MaxInt64

	for _, seed := range maps.seeds {
		pos := recMapValue(seed, maps.maps)
		if pos < currMin {
			currMin = pos
		}
	}

	println("p1", currMin)
}

func p2(content string) {
	maps := parseMaps(content)
	currMin := math.MaxInt64

	seedRanges := make([]CurrRange, 0)
	for i := 0; i < len(maps.seeds); i += 2 {
		seedRanges = append(seedRanges, CurrRange{
			start: maps.seeds[i],
			end:   maps.seeds[i] + maps.seeds[i+1],
		})
	}

	res := recMapRange(seedRanges, maps.maps)

	for _, r := range res {
		if currMin > int(r.start) {
			currMin = int(r.start)
		}
	}

	println("p1", currMin)
}

func recMapRange(curr []CurrRange, ranges [][]Range) []CurrRange {
	result := curr
	for _, r := range ranges {
		temp := make([]CurrRange, 0)
		for _, c := range result {
			temp = append(temp, mapRanges(c, r)...)
		}
		result = temp
	}
	return result
}

func parseMaps(content string) Maps {
	seeds := make([]int64, 0)
	maps := make([][]Range, 0)
	mapId := 0

	lines := strings.Split(content, "\n")
	seedsSection := strings.Split(lines[0], " ")
	for i := 1; i < len(seedsSection); i++ {
		t, _ := strconv.ParseInt(seedsSection[i], 10, 64)
		seeds = append(seeds, t)
	}

	for i := 3; i < len(lines); i++ {
		if len(strings.Trim(lines[i], " ")) == 0 {
			i += 1
			mapId += 1
			continue
		}

		if len(maps) <= mapId {
			maps = append(maps, make([]Range, 0))
		}
		maps[mapId] = append(maps[mapId], parseRange(lines[i]))
	}

	return Maps{
		seeds,
		maps,
	}
}

type Maps struct {
	seeds []int64
	maps  [][]Range
}

type CurrRange struct {
	start int64
	end   int64
}

type Range struct {
	sourceStart int64
	destStart   int64
	length      int64
}

func mapRanges(curr CurrRange, ranges []Range) []CurrRange {
	result := make([]CurrRange, 0)

	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i].sourceStart < ranges[j].sourceStart
	})

	p := curr.start

	for p < curr.end {
		inside, matchingRange := getMatchingRange(p, ranges)

		if inside {
			mappedStart := p - matchingRange.sourceStart + matchingRange.destStart
			length := min(
				matchingRange.length-p+matchingRange.sourceStart,
				curr.end-p,
			)
			result = append(result, CurrRange{start: mappedStart, end: mappedStart + length})
			p += length
		} else {
			length := min(
				matchingRange.sourceStart-p,
				curr.end-p,
			)
			result = append(result, CurrRange{start: p, end: p + length})
			p += length
		}
	}

	return result
}

func getMatchingRange(p int64, ranges []Range) (bool, Range) {
	for _, r := range ranges {
		if r.sourceStart <= p && r.length > (p-r.sourceStart) {
			return true, r
		}

		if r.sourceStart > p {
			return false, r
		}
	}

	return false, Range{
		sourceStart: math.MaxInt64,
	}
}

func parseRange(line string) Range {
	numStrs := strings.Split(line, " ")

	destStart, _ := strconv.ParseInt(numStrs[0], 10, 64)
	sourceStart, _ := strconv.ParseInt(numStrs[1], 10, 64)
	length, _ := strconv.ParseInt(numStrs[2], 10, 64)

	return Range{
		sourceStart,
		destStart,
		length,
	}
}

func recMapValue(val int64, ranges [][]Range) int {
	curr := val
	for i := 0; i < len(ranges); i++ {
		for j := 0; j < len(ranges[i]); j++ {
			isMapped, mappedVal := mapValue(curr, ranges[i][j])

			if isMapped {
				//println("mapping: ", val, mappedVal)
				curr = mappedVal
				break
			}
		}
		//println("mapping: ", val, val)
	}
	return int(curr)
}

func mapValue(val int64, r Range) (bool, int64) {
	delta := val - r.sourceStart
	if delta >= 0 && delta < r.length {
		return true, r.destStart + delta
	}

	return false, -1
}

var testContent = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
