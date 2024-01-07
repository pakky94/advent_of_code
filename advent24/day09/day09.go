package day09

import (
	"os"
	"strconv"
	"strings"
)

func Run() {
	f, _ := os.ReadFile("day09/input")
	content := string(f)
	_ = content

	p1(content)
	p2(content)
}

func p1(content string) {
	tot := 0
	for _, line := range strings.Split(content, "\n") {
		nums := parseNums(line)
		tot += extrapolate(nums)
	}
	println("p1", tot)
}

func p2(content string) {
	tot := 0
	for _, line := range strings.Split(content, "\n") {
		nums := parseNums(line)
		tot += extrapolateBackwards(nums)
	}
	println("p2", tot)
}

func extrapolateBackwards(nums []int) int {
	if isZeros(nums) {
		return 0
	}

	return nums[0] - extrapolateBackwards(diff(nums))
}

func extrapolate(nums []int) int {
	if isZeros(nums) {
		return 0
	}

	return nums[len(nums)-1] + extrapolate(diff(nums))
}

func diff(nums []int) []int {
	diff := make([]int, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		diff[i-1] = nums[i] - nums[i-1]
	}
	return diff
}

func isZeros(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func parseNums(input string) []int {
	res := make([]int, 0)
	for _, x := range strings.Split(input, " ") {
		if len(strings.Trim(x, " ")) == 0 {
			continue
		}
		n, _ := strconv.ParseInt(x, 10, 64)
		res = append(res, int(n))
	}
	return res
}

var testContent = ``
