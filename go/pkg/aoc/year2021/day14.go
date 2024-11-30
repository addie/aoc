package aoc

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const day14Filename = "data/day14"

func Day14() int {
	file, _ := os.Open(day14Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s := ""
	rules := make(map[string]string)
	for scanner.Scan() {
		r := strings.TrimSpace(scanner.Text())
		if r == "" {
			continue
		} else if !strings.Contains(r, "->") {
			s = r
			continue
		}
		pi := strings.Split(r, " -> ")
		rules[pi[0]] = pi[1]
	}

	currentCounter := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		inc(currentCounter, string(s[i])+string(s[i+1]))
	}

	res := 0
	for ct := 0; ct < 41; ct++ {
		if ct == 10 || ct == 40 {
			c1 := make(map[string]int)
			for k := range currentCounter {
				inc(c1, string(k[0]), currentCounter[k])
			}
			inc(c1, string(s[len(s)-1]))
			res = max(c1) - min(c1)
			fmt.Println(res)
		}
		c2 := make(map[string]int)
		for k := range currentCounter {
			inc(c2, string(k[0])+rules[k], currentCounter[k])
			inc(c2, rules[k]+string(k[1]), currentCounter[k])
		}
		currentCounter = c2
	}
	return res
}

func min(c map[string]int) int {
	min := math.MaxInt
	for _, v := range c {
		if v < min {
			min = v
		}
	}
	return min
}

func max(c map[string]int) int {
	max := 0
	for _, v := range c {
		if v > max {
			max = v
		}
	}
	return max
}

func inc(m map[string]int, key string, optVal ...int) {
	val := 1
	if optVal != nil {
		val = optVal[0]
	}
	if _, ok := m[key]; !ok {
		m[key] = val
	} else {
		m[key] += val
	}
}
