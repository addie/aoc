/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package print

import "fmt"

func GridInt(grid [][]int) {
	for r := range grid {
		fmt.Print("[")
		for c := range grid[r] {
			fmt.Print(grid[r][c])
			fmt.Print(" ")
		}
		fmt.Println("]")
	}
	fmt.Println()
}

func GridBool(grid [][]bool) {
	for r := range grid {
		fmt.Print("[")
		for c := range grid[r] {
			if !grid[r][c] {
				fmt.Print(".")
			} else {
				fmt.Print("T")
			}
			fmt.Print(" ")
		}
		fmt.Println("]")
	}
	fmt.Println()
}

func GridString(grid [][]string) {
	for r := range grid {
		fmt.Print("[")
		for c := range grid[r] {
			fmt.Print(grid[r][c])
			fmt.Print(" ")
		}
		fmt.Println("]")
	}
	fmt.Println()
}
