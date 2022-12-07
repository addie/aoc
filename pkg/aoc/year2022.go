package aoc

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func (s Solution) Year2022Day1(input string) (int, int) {
	lines := ReadFile(input)

	currentTotal := 0
	maxTotal := 0
	var allTotals []int
	for _, line := range lines {
		if line == "" {
			allTotals = append(allTotals, currentTotal)
			if currentTotal > maxTotal {
				maxTotal = currentTotal
			}
			currentTotal = 0
		}
		weight, _ := strconv.Atoi(line)
		currentTotal += weight
	}
	sort.Ints(allTotals)
	top3 := allTotals[len(allTotals)-1] + allTotals[len(allTotals)-2] + allTotals[len(allTotals)-3]
	return maxTotal, top3
}

func (s Solution) Year2022Day2(input string) (int, int) {
	lines := ReadFile(input)

	// Opp A for Rock, B for Paper, and C for Scissors
	// Me X for Rock, Y for Paper, and Z for Scissors
	// Score 6 win, 3 tie, 0 loss
	// Score 1 for Rock, 2 for Paper, and 3 for Scissors
	part1Total := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		oppMove, myMove := moves[0], moves[1]
		switch myMove {
		case "X":
			part1Total += 1
		case "Y":
			part1Total += 2
		case "Z":
			part1Total += 3
		}
		if myMove == "X" && oppMove == "A" || myMove == "Y" && oppMove == "B" || myMove == "Z" && oppMove == "C" {
			part1Total += 3
		} else if myMove == "X" && oppMove == "C" || myMove == "Y" && oppMove == "A" || myMove == "Z" && oppMove == "B" {
			part1Total += 6
		}
	}

	// Part 2
	// X means you need to lose
	// Y means you need to draw
	// Z means you need to win
	// Opp A for Rock, B for Paper, and C for Scissors
	// Score 6 win, 3 tie, 0 loss
	// Score 1 for Rock, 2 for Paper, and 3 for Scissors
	part2Total := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		oppMove, outcome := moves[0], moves[1]
		switch outcome {
		case "X": // lose
			switch oppMove {
			case "A": // rock
				part2Total += 3
			case "B": // paper
				part2Total += 1
			case "C": // scissors
				part2Total += 2
			}
		case "Y": // draw
			switch oppMove {
			case "A": // rock
				part2Total += 4
			case "B": // paper
				part2Total += 5
			case "C": // scissors
				part2Total += 6
			}
		case "Z": // win
			switch oppMove {
			case "A": // rock
				part2Total += 8
			case "B": // paper
				part2Total += 9
			case "C": // scissors
				part2Total += 7
			}
		}
	}

	return part1Total, part2Total
}

func (s Solution) Year2022Day3(input string) (int, int) {
	lines := ReadFile(input)
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	const lowerFactor int32 = 96
	const upperFactor int32 = 38
	var res1 int32 = 0
	for _, line := range lines {
		set := make(map[int32]bool)
		first, second := line[:len(line)/2], line[len(line)/2:]
		for _, char := range first {
			set[char] = true
		}
		for _, char := range second {
			if _, ok := set[char]; ok {
				factor := lowerFactor
				if char < 'a' {
					factor = upperFactor
				}
				res1 += char - factor
				break
			}
		}
	}

	var res2 int32 = 0
	for i := 0; i < len(lines); i += 3 {
		set := make(map[int32]bool)
		first, second, third := lines[i], lines[i+1], lines[i+2]
		for _, char := range first {
			set[char] = true
		}
		for _, char := range second {
			factor := lowerFactor
			if char < 'a' {
				factor = upperFactor
			}
			if _, ok := set[char]; ok {
				for _, currChar := range third {
					if currChar == char {
						res2 += char - factor
						goto breakout
					}
				}
			}
		}
	breakout:
	}
	return int(res1), int(res2)
}

func (s Solution) Year2022Day4(input string) (int, int) {
	lines := ReadFile(input)
	type interval struct {
		s, e int
	}

	contained := func(int1 interval, int2 interval) bool {
		return int1.s <= int2.s && int1.e >= int2.e
	}

	numContained := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		inter := strings.Split(sections[0], "-")
		start1, err := strconv.Atoi(inter[0])
		check(err)
		end1, err := strconv.Atoi(inter[1])
		check(err)
		inter = strings.Split(sections[1], "-")
		start2, err := strconv.Atoi(inter[0])
		check(err)
		end2, err := strconv.Atoi(inter[1])
		check(err)

		int1, int2 := interval{s: start1, e: end1}, interval{s: start2, e: end2}
		if contained(int1, int2) || contained(int2, int1) {
			numContained += 1
		}
	}

	overlap := func(int1 interval, int2 interval) bool {
		return int1.s <= int2.s && int1.e >= int2.s
	}

	numOverlap := 0
	for _, line := range lines {
		sections := strings.Split(line, ",")
		inter := strings.Split(sections[0], "-")
		start1, err := strconv.Atoi(inter[0])
		check(err)
		end1, err := strconv.Atoi(inter[1])
		check(err)
		inter = strings.Split(sections[1], "-")
		start2, err := strconv.Atoi(inter[0])
		check(err)
		end2, err := strconv.Atoi(inter[1])
		check(err)

		int1, int2 := interval{s: start1, e: end1}, interval{s: start2, e: end2}
		if overlap(int1, int2) || overlap(int2, int1) {
			numOverlap += 1
		}
	}
	return numContained, numOverlap
}

func (s Solution) Year2022Day5(input string) (string, string) {
	lines := ReadFile(input)

	// STACKS
	//             [L] [M]         [M]
	//         [D] [R] [Z]         [C] [L]
	//         [C] [S] [T] [G]     [V] [M]
	// [R]     [L] [Q] [B] [B]     [D] [F]
	// [H] [B] [G] [D] [Q] [Z]     [T] [J]
	// [M] [J] [H] [M] [P] [S] [V] [L] [N]
	// [P] [C] [N] [T] [S] [F] [R] [G] [Q]
	// [Z] [P] [S] [F] [F] [T] [N] [P] [W]
	// 1   2   3   4   5   6   7   8   9
	stacks := [][]string{
		{"Z", "P", "M", "H", "R"},
		{"P", "C", "J", "B"},
		{"S", "N", "H", "G", "L", "C", "D"},
		{"F", "T", "M", "D", "Q", "S", "R", "L"},
		{"F", "S", "P", "Q", "B", "T", "Z", "M"},
		{"T", "F", "S", "Z", "B", "G"},
		{"N", "R", "V"},
		{"P", "G", "L", "T", "D", "V", "C", "M"},
		{"W", "Q", "N", "J", "F", "M", "L"},
	}
	// PROCESS MOVES
	// ex: move 7 from 3 to 9
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		count, err := strconv.Atoi(tokens[1])
		check(err)
		start, err := strconv.Atoi(tokens[3])
		check(err)
		end, err := strconv.Atoi(tokens[5])
		check(err)

		i := 0
		for i < count {
			char := stacks[start-1][len(stacks[start-1])-1]
			stacks[start-1] = stacks[start-1][:len(stacks[start-1])-1]
			stacks[end-1] = append(stacks[end-1], char)
			i++
		}
	}
	var resArr1 []string
	for _, stack := range stacks {
		char := stack[len(stack)-1]
		resArr1 = append(resArr1, char)
	}

	stacks = [][]string{
		{"Z", "P", "M", "H", "R"},
		{"P", "C", "J", "B"},
		{"S", "N", "H", "G", "L", "C", "D"},
		{"F", "T", "M", "D", "Q", "S", "R", "L"},
		{"F", "S", "P", "Q", "B", "T", "Z", "M"},
		{"T", "F", "S", "Z", "B", "G"},
		{"N", "R", "V"},
		{"P", "G", "L", "T", "D", "V", "C", "M"},
		{"W", "Q", "N", "J", "F", "M", "L"},
	}
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		count, err := strconv.Atoi(tokens[1])
		check(err)
		start, err := strconv.Atoi(tokens[3])
		check(err)
		end, err := strconv.Atoi(tokens[5])
		check(err)

		subStack := stacks[start-1][len(stacks[start-1])-count:]
		stacks[start-1] = stacks[start-1][:len(stacks[start-1])-count]
		stacks[end-1] = append(stacks[end-1], subStack...)
	}
	var resArr2 []string
	for _, stack := range stacks {
		char := stack[len(stack)-1]
		resArr2 = append(resArr2, char)
	}
	return strings.Join(resArr1, ""), strings.Join(resArr2, "")
}

func (s Solution) Year2022Day6(input string) (int, int) {
	lines := ReadFile(input)

	check := func(alphaList []int) bool {
		for _, el := range alphaList {
			if el > 1 {
				return false
			}
		}
		return true
	}
	calcValue := func(size int) int {
		for _, str := range lines {
			alphaList := make([]int, 256)
			for i := range str {
				if i < size {
					alphaList[str[i]] += 1
					continue
				}
				if check(alphaList) {
					return i
				}
				alphaList[str[i]] += 1
				alphaList[str[i-size]] -= 1
			}
		}
		return 0
	}
	return calcValue(4), calcValue(14)
}

func (s Solution) Year2022Day7(input string) (int, int) {
	lines := ReadFile(input)
	var path []string
	sizeMap := make(map[string]int)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if tokens[1] == "ls" || tokens[0] == "dir" {
			continue
		} else if tokens[1] == "cd" {
			if tokens[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, tokens[2])
			}
		} else {
			// process files
			size, err := strconv.Atoi(tokens[0])
			if err != nil {
				fmt.Printf("ignoring %s\n", line)
				continue
			}
			for i := 1; i < len(path)+1; i++ {
				sizeMap[strings.Join(path[:i], "/")] += size
			}
		}
	}
	res1 := 0
	for _, size := range sizeMap {
		if size <= 100000 {
			res1 += size
		}
	}
	const totalDiskSpace = 70000000
	const neededDiskSpace = 30000000
	freeSpace := totalDiskSpace - sizeMap["/"]
	neededSpace := neededDiskSpace - freeSpace
	candidateSize := math.MaxInt
	for _, size := range sizeMap {
		if size >= neededSpace && size < candidateSize {
			candidateSize = size
		}
	}
	return res1, candidateSize
}
