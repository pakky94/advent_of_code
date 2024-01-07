package day15

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day15/input")
	content := string(f)
	_ = content

	//p1(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)
	p1(content)
	//p2(`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`)
	p2(content)
}

func p2(content string) {
	input := strings.Trim(content, " \n\r")
	lenses := make([][]Lens, 256)
	for i := 0; i < 256; i++ {
		lenses[i] = make([]Lens, 0)
	}

	for _, s := range strings.Split(input, ",") {
		instr := parseInstr(s)
		box := calcHash(instr.label)
		if instr.remove {
			lenses[box] = removeLens(lenses[box], instr.label)
		} else {
			lenses[box] = addLens(lenses[box], instr.label, instr.focal)
		}
	}
	tot := 0

	for i := 0; i < 256; i++ {
		for j, n := range lenses[i] {
			tot += (i + 1) * (j + 1) * n.focal
		}
	}

	println("p2", tot)
}

func addLens(lenses []Lens, label string, focal int) []Lens {
	for i := 0; i < len(lenses); i++ {
		if lenses[i].label == label {
			lenses[i].focal = focal
			return lenses
		}
	}
	return append(lenses, Lens{
		label: label,
		focal: focal,
	})
}

func removeLens(lenses []Lens, label string) []Lens {
	lenses = slices.DeleteFunc(lenses, func(l Lens) bool {
		return l.label == label
	})
	return lenses
}

func parseInstr(s string) Instruction {
	dash := strings.Index(s, "-")
	if dash >= 0 {
		return Instruction{
			label:  s[:dash],
			remove: true,
			focal:  0,
		}
	}
	eq := strings.Index(s, "=")
	n, err := strconv.ParseInt(s[eq+1:], 10, 64)
	if err != nil {
		panic(err)
	}
	return Instruction{
		label:  s[:eq],
		remove: false,
		focal:  int(n),
	}
}

type Instruction struct {
	label  string
	remove bool
	focal  int
}

type Lens struct {
	label string
	focal int
}

func p1(content string) {
	input := strings.Trim(content, " \n\r")
	tot := 0
	for _, s := range strings.Split(input, ",") {
		tot += calcHash(s)
	}
	println("p1", tot)
}

type Hasher struct {
	curr uint8
}

func calcHash(s string) int {
	h := Hasher{curr: 0}
	for i := 0; i < len(s); i++ {
		h.write(s[i])
	}
	return int(h.curr)
}

func (h *Hasher) write(c byte) *Hasher {
	t := uint16(h.curr)
	r := uint8(((t + uint16(c)) * 17) % 256)
	h.curr = r
	return h
}
