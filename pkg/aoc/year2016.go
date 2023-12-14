package aoc

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type dir int

const (
	north = iota
	east
	south
	west
)

func (s Solution[T]) Year2016Day1(input string) (int, int) {
	instructions := strings.Split(ReadFileToString(input), ", ")
	res1, res2 := 0, 0
	divisor := big.NewInt(4)
	for _, p2 := range []bool{false, true} {
		if !p2 {
			continue
		}
		facing := int64(0)
		r, c := 0, 0
		visited := map[coord]bool{coord{0, 0}: true}
		for _, instruction := range instructions {
			dividend := new(big.Int)
			if instruction[0] == 'R' {
				dividend = big.NewInt(facing + 1)
			} else {
				dividend = big.NewInt(facing - 1)
			}
			facing = new(big.Int).Mod(dividend, divisor).Int64()
			val := Must(strconv.Atoi(strings.TrimSpace(instruction[1:])))
			for step := 0; step < val; step++ {
				switch facing {
				case north:
					r--
				case east:
					c++
				case south:
					r++
				case west:
					c--
				}
				if p2 {
					cell := coord{r: r, c: c}
					if visited[cell] {
						goto end
					}
					visited[cell] = true
				}
			}
		}
	end:
		if p2 {
			res2 = int(math.Abs(float64(r)) + math.Abs(float64(c)))
		} else {
			res1 = int(math.Abs(float64(r)) + math.Abs(float64(c)))
		}
	}
	return res1, res2
}

func (s Solution[T]) Year2016Day2(input string) (int, int) {
	data := ReadFileToString(input)
	res1, res2 := 0, 0
	var code []string
	inst := strings.Split(data, "\n")
	for _, p2 := range []bool{false, true} {
		grid := [][]byte{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		}
		if p2 {
			grid = [][]byte{
				{'.', '.', '1', '.', '.'},
				{'.', '2', '3', '4', '.'},
				{'5', '6', '7', '8', '9'},
				{'.', 'A', 'B', 'C', '.'},
				{'.', '.', 'D', '.', '.'},
			}
		}
		curr := coord{r: 1, c: 1}
		if p2 {
			curr = coord{r: 2, c: 0}
		}
		d := map[rune]coord{'U': {-1, 0}, 'R': {0, 1}, 'D': {1, 0}, 'L': {0, -1}}
		for _, c := range inst {
			if c == "" {
				continue
			}
			for _, m := range c {
				n := coord{curr.r + d[m].r, curr.c + d[m].c}
				if n.r < 0 || n.r >= len(grid) || n.c < 0 || n.c >= len(grid[0]) || grid[n.r][n.c] == '.' {
					continue
				}
				curr = n
				fmt.Println(string(grid[curr.r][curr.c]))
			}
			code = append(code, string(grid[curr.r][curr.c]))
		}
		fmt.Println(strings.Join(code, ""))
	}
	return res1, res2
}
