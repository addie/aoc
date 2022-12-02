/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package new

func RuneGrid(size int) [][]rune {
	grid := make([][]rune, size)
	for r := range grid {
		grid[r] = make([]rune, size)
	}
	return grid
}

func IntGrid(size int) [][]int {
	grid := make([][]int, size)
	for r := range grid {
		grid[r] = make([]int, size)
	}
	return grid
}
