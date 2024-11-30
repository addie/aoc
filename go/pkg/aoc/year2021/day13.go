package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const day13Filename = "data/day13"

type xy struct {
	x, y int
}

func Day13() int {
	file, _ := os.Open(day13Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make(map[xy]bool)
	foldCount := 0
	for scanner.Scan() {
		r := strings.TrimSpace(scanner.Text())
		if strings.Contains(r, "fold") {
			foldCount++
			gridCopy := make(map[xy]bool)
			instr := strings.Split(strings.Split(r, " ")[2], "=")
			n, _ := strconv.Atoi(instr[1])
			dir, num := instr[0], n
			if dir == "x" { // fold left
				for crd := range grid {
					if crd.x < num {
						gridCopy[xy{x: crd.x, y: crd.y}] = true
					} else {
						gridCopy[xy{x: num - (crd.x - num), y: crd.y}] = true
					}
				}
			} else { // fold up
				for crd := range grid {
					if crd.y < num {
						gridCopy[xy{x: crd.x, y: crd.y}] = true
					} else {
						gridCopy[xy{x: crd.x, y: num - (crd.y - num)}] = true
					}
				}
			}
			grid = gridCopy
			if foldCount == 1 {
				fmt.Printf("fold %d count %d\n", foldCount, len(gridCopy))
			}
		} else if len(r) > 0 {
			rowStr := strings.Split(r, ",")
			x, _ := strconv.Atoi(rowStr[0])
			y, _ := strconv.Atoi(rowStr[1])
			grid[xy{x: x, y: y}] = true
		}
	}
	X := getMaxX(grid)
	Y := getMaxY(grid)

	result := ""
	for y := 0; y <= Y; y++ {
		for x := 0; x <= X; x++ {
			if _, ok := grid[xy{x: x, y: y}]; ok {
				result += "*"
			} else {
				result += " "
			}
		}
		fmt.Println(result)
		result = ""
	}
	return 0
}

func getMaxX(grid map[xy]bool) int {
	m := 0
	for k := range grid {
		if k.x > m {
			m = k.x
		}
	}
	return m
}

func getMaxY(grid map[xy]bool) int {
	m := 0
	for k := range grid {
		if k.y > m {
			m = k.y
		}
	}
	return m
}
