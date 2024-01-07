package day12

import (
	"os"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day12/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	gears := parseUnfolded(content)
	println("p2", totalPerm(gears))
}

func p1(content string) {
	gears := parse(content)
	println("p1", totalPerm(gears))
}

func totalPerm(gears []Gears) int {
	tot := 0
	for _, g := range gears {
		mem := map[GearKey]int{}
		tot += permutations(g, mem)
	}
	return tot
}

func parseUnfolded(content string) []Gears {
	t := parse(content)
	res := make([]Gears, len(t))
	maxD := 0
	for i := 0; i < len(t); i++ {
		res[i] = unfold(t[i])
		if len(res[i].damaged) > maxD {
			maxD = len(res[i].damaged)
		}
	}
	return res
}

func unfold(gears Gears) Gears {
	l := len(gears.damaged)
	damaged := make([]byte, l*5)
	for i := 0; i < 5; i++ {
		for j := 0; j < len(gears.damaged); j++ {
			damaged[i*l+j] = gears.damaged[j]
		}
	}

	sections := make([]string, 5)
	for i := 0; i < 5; i++ {
		sections[i] = gears.pattern
	}

	return Gears{
		pattern: strings.Join(sections, "?"),
		damaged: damaged,
	}
}

type GearKey struct {
	pattern string
	damaged string
}

func toKey(gears Gears) GearKey {
	return GearKey{
		pattern: gears.pattern,
		damaged: string(gears.damaged),
	}
}

func permutations(gears Gears, mem map[GearKey]int) int {
	gearsKey := toKey(gears)
	memRes, ok := mem[gearsKey]
	if ok {
		return memRes
	}

	r := isValid(gears)

	if !r.mayBeValid {
		return saveAndReturn(gearsKey, 0, mem)
	}

	if strings.Index(gears.pattern, "?") == -1 {
		if r.completed {
			return saveAndReturn(gearsKey, 1, mem)
		} else {
			return saveAndReturn(gearsKey, 0, mem)
		}
	}

	s1 := strings.Replace(r.substr, "?", ".", 1)
	s2 := strings.Replace(r.substr, "?", "#", 1)
	memRes = permutations(Gears{
		pattern: s1,
		damaged: r.remaining,
	}, mem) + permutations(Gears{
		pattern: s2,
		damaged: r.remaining,
	}, mem)

	return saveAndReturn(gearsKey, memRes, mem)
}

func saveAndReturn(key GearKey, val int, mem map[GearKey]int) int {
	mem[key] = val
	return val
}

func isValid(gears Gears) CheckResult {
	idx := 0
	found := make([]byte, 0)
	curr := byte(0)
Loop:
	for i := 0; i < len(gears.pattern); i++ {
		switch gears.pattern[i] {
		case '#':
			curr++
		case '.':
			idx = i
			if curr != 0 {
				found = append(found, curr)
				curr = 0
			}
		case '?':
			curr = 0
			break Loop
		}
	}

	if curr != 0 {
		found = append(found, curr)
	}

	if len(found) > len(gears.damaged) {
		return CheckResult{false, false, "", nil}
	}

	for i := 0; i < len(found); i++ {
		if found[i] != gears.damaged[i] {
			return CheckResult{false, false, "", nil}
		}
	}

	completed := len(found) == len(gears.damaged) &&
		strings.Index(gears.pattern, "?") == -1
	substr := ""
	remaining := []byte(nil)
	if !completed {
		substr = gears.pattern[idx:]
		remaining = gears.damaged[len(found):]
	}

	return CheckResult{
		mayBeValid: true,
		completed:  completed,
		substr:     substr,
		remaining:  remaining,
	}
}

type CheckResult struct {
	mayBeValid bool
	completed  bool
	substr     string
	remaining  []byte
}

func parse(content string) []Gears {
	lines := make([]Gears, 0)
	for _, x := range strings.Split(content, "\n") {
		if len(x) == 0 {
			continue
		}

		parts := strings.Split(x, " ")
		damaged := make([]byte, 0)
		for _, n := range strings.Split(parts[1], ",") {
			t, _ := strconv.ParseInt(n, 10, 64)
			damaged = append(damaged, byte(t))
		}
		lines = append(lines, Gears{
			pattern: parts[0],
			damaged: damaged,
		})
	}
	return lines
}

type Gears struct {
	pattern string
	damaged []byte
}
