package aoc

import (
	"sort"
	"strconv"
	"strings"
)

func (s Solution) Year2015Day1(input string) (int, int) {
	lines := ReadFile(input)
	m := map[string]int{"(": 1, ")": -1}
	res1, res2 := 0, 0
	for i, char := range lines[0] {
		res1 += m[string(char)]
		if res2 == 0 && res1 < 0 {
			res2 = i + 1
		}
	}
	return res1, res2
}

func (s Solution) Year2015Day2(input string) (int, int) {
	lines := ReadFile(input)
	res1, res2 := 0, 0

	for _, line := range lines {
		dims := strings.Split(line, "x")
		var i []int
		i = append(i, Must(strconv.Atoi(dims[0])), Must(strconv.Atoi(dims[1])), Must(strconv.Atoi(dims[2])))
		sort.Ints(i)
		res1 += 2*i[0]*i[1] + 2*i[1]*i[2] + 2*i[2]*i[0] + i[0]*i[1]
		res2 += 2*i[0] + 2*i[1] + i[0]*i[1]*i[2]
	}
	return res1, res2
}
