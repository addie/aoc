package aoc

import (
	"fmt"
	"math"
)

const day17Filename = "data/day17"

type box struct {
	x1, x2, y1, y2 int
}

func Day17() (int, int) {
	// "target area: x=169..206, y=-108..-68"
	const maxInit = 1000
	const minInit = -1000
	t := box{x1: 169, x2: 206, y1: -108, y2: -68}
	res := math.MinInt
	valid := 0
	for i := minInit; i < maxInit; i++ {
		for j := minInit; j < maxInit; j++ {
			found := false
			maxY := math.MinInt
			x, y := 0, 0
			dx, dy := i, j
			for k := 0; k < maxInit; k++ {
				x += dx
				y += dy
				maxY = int(math.Max(float64(maxY), float64(y)))
				if dx > 0 {
					dx--
				} else if dx < 0 {
					dx++
				}
				dy--
				if x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2 {
					found = true
				}
			}
			if found {
				valid++
				if maxY > res {
					res = maxY
					fmt.Printf("(%d, %d) %d\n", dx, dy, res)
				}
			}

		}
	}
	return res, valid
}
