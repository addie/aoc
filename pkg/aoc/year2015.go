package aoc

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (s Solution[T]) Year2015Day1(input string) (int, int) {
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

func (s Solution[T]) Year2015Day2(input string) (int, int) {
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

func (s Solution[T]) Year2015Day3(path string) (int, int) {
	data := ReadFileToString(path)
	res1 := year2015Day3Part1(data)
	res2 := year2015Day3Part2(data)
	return res1, res2
}

func year2015Day3Part1(insts string) int {
	r, c := 0, 0
	set := map[coord]bool{coord{r: r, c: c}: true}
	for _, inst := range []byte(insts) {
		r, c = move(inst, r, c)
		set[coord{r: r, c: c}] = true
	}
	res := 0
	for _ = range set {
		res++
	}
	return res
}

func year2015Day3Part2(insts string) int {
	r, c := 0, 0
	set := map[coord]bool{coord{r: r, c: c}: true}
	for i := 0; i < len(insts); i += 2 {
		r, c = move(insts[i], r, c)
		set[coord{r: r, c: c}] = true
	}
	r, c = 0, 0
	for i := 1; i < len(insts); i += 2 {
		r, c = move(insts[i], r, c)
		set[coord{r: r, c: c}] = true
	}
	res := 0
	for _ = range set {
		res++
	}
	return res
}

func move(inst byte, r, c int) (int, int) {
	switch inst {
	case '^':
		r--
	case 'v':
		r++
	case '<':
		c--
	case '>':
		c++
	}
	return r, c
}

func (s Solution[T]) Year2015Day4(path string) (int, int) {
	data := ReadFileToString(path)
	data = strings.TrimSuffix(data, "\n")
	p1, p2 := 0, 0
	for i, sz := range []int{5, 6} {
		count := 0
		for {
			countStr := strconv.Itoa(count)
			prefix := fmt.Sprintf("%x", md5.Sum([]byte(data+countStr)))[:sz]
			if prefix == strings.Repeat("0", sz) {
				break
			}
			count++
		}
		if i == 0 {
			p1 = count
		} else {
			p2 = count
		}
	}
	return p1, p2
}

func (s Solution[T]) Year2015Day5(path string) (int, int) {
	lines := ReadFile(path)
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	bannedChars := map[string]bool{"ab": true, "cd": true, "pq": true, "xy": true}
	p1, p2 := 0, 0
	for _, line := range lines {
		L := len(line)
		vowelCt := 0
		containsDupe := false
		bannedChar := false
		for i := 0; i < L-1; i++ {
			if bannedChars[line[i:i+2]] {
				bannedChar = true
				break
			}
			if line[i] == line[i+1] {
				containsDupe = true
			}
		}
		if bannedChar {
			continue
		}
		for _, char := range line {
			if vowels[char] {
				vowelCt++
			}
		}
		if vowelCt < 3 || !containsDupe {
			continue
		}
		p1++
	}
	/* It contains a pair of any two letters that
	appears at least twice in the string without
	overlapping, like xyxy (xy) or aabcdefgaa (aa),
	but not like aaa (aa, but it overlaps).
	It contains at least one letter which
	repeats with exactly one letter between them,
	like xyx, abcdefeghi (efe), or even aaa.*/
	for _, line := range lines {
		L := len(line)
		containsPair := false
		containsRepeat := false
		for i := 0; i < L-2; i++ {
			if line[i] == line[i+2] && line[i] != line[i+1] {
				containsRepeat = true
				break
			}
		}
		if !containsRepeat {
			continue
		}
		i := 0
		for i < L-1 {
			if line[i] == line[i+1] {
				if i == L-2 && line[i-1] != line[i] {
					containsPair = true
					break
				} else if i < L-2 {
					if i > 0 && line[i+1] != line[i+2] && line[i] == line[i-1] {
						i++
						continue
					}
					if i == 0 && line[i+1] != line[i+2] {
						containsPair = true
						break
					}
				} else if i < L-3 && line[i+1] == line[i+2] && line[i+2] == line[i+3] {
					containsPair = true
					break
				}
			}
			i++
		}
		if containsPair {
			p2++
		}
	}
	return p1, p2
}
