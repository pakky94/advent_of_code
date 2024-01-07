package day08

import (
	"os"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day08/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p2(content string) {
	parsed := parseMap(content)

	currents := getStartNodes(parsed)
	paths := len(currents)
	loops := make([]int, paths)

	for i := 0; i < paths; i++ {
		count := 0

		for !strings.HasSuffix(currents[i], "Z") {
			dir := parsed.directions[count%parsed.length]
			node := findNode(parsed, currents[i])
			if dir == 'L' {
				currents[i] = node.left
			} else {
				currents[i] = node.right
			}

			count++
		}

		loops[i] = count
	}

	res := 1

	for i := 0; i < paths; i++ {
		res = lcm(res, loops[i])
	}

	println("p2", res)
}

func lcm(a int, b int) int {
	return a * (b / gcd(a, b))
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a > b {
		return gcd(a%b, b)
	}

	return gcd(b%a, a)
}

func checkEndCondition(nodes []string) bool {
	for _, n := range nodes {
		if !strings.HasSuffix(n, "Z") {
			return false
		}
	}

	return true
}

func getStartNodes(parsed Map) []string {
	res := make([]string, 0)

	for _, n := range parsed.nodes {
		if strings.HasSuffix(n.position, "A") {
			res = append(res, n.position)
		}
	}

	return res
}

func p1(content string) {
	parsed := parseMap(content)

	count := 0
	current := "AAA"
	for strings.Compare(current, "ZZZ") != 0 {
		dir := parsed.directions[count%parsed.length]
		node := findNode(parsed, current)
		if dir == 'L' {
			current = node.left
		} else {
			current = node.right
		}
		count++
	}

	println("p1", count)
}

func findNode(parsed Map, current string) Node {
	for _, n := range parsed.nodes {
		if strings.Compare(n.position, current) == 0 {
			return n
		}
	}

	return Node{}
}

func parseMap(content string) Map {
	lines := strings.Split(content, "\n")
	nodes := make([]Node, 0)

	for i := 2; i < len(lines); i++ {
		s := strings.Split(lines[i], "=")
		if len(s) == 1 {
			continue
		}

		dest := strings.Split(s[1], ",")
		nodes = append(nodes, Node{
			position: strings.Trim(s[0], " ()"),
			left:     strings.Trim(dest[0], " ()"),
			right:    strings.Trim(dest[1], " ()"),
		})
	}

	return Map{
		directions: lines[0],
		length:     len(lines[0]),
		nodes:      nodes,
	}
}

type Map struct {
	directions string
	length     int
	nodes      []Node
}

type Node struct {
	position string
	left     string
	right    string
}

var testContent2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

var testContent = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
