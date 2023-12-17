package aoc

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func (s Solution[T]) Year2022Day1(input string) (int, int) {
	lines := ReadFileToLines(input)

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

func (s Solution[T]) Year2022Day2(input string) (int, int) {
	lines := ReadFileToLines(input)

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

func (s Solution[T]) Year2022Day3(input string) (int, int) {
	lines := ReadFileToLines(input)
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

func (s Solution[T]) Year2022Day4(input string) (int, int) {
	lines := ReadFileToLines(input)
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
		start1 := Must(strconv.Atoi(inter[0]))
		end1 := Must(strconv.Atoi(inter[1]))
		inter = strings.Split(sections[1], "-")
		start2 := Must(strconv.Atoi(inter[0]))
		end2 := Must(strconv.Atoi(inter[1]))

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
		start1 := Must(strconv.Atoi(inter[0]))
		end1 := Must(strconv.Atoi(inter[1]))
		inter = strings.Split(sections[1], "-")
		start2 := Must(strconv.Atoi(inter[0]))
		end2 := Must(strconv.Atoi(inter[1]))

		int1, int2 := interval{s: start1, e: end1}, interval{s: start2, e: end2}
		if overlap(int1, int2) || overlap(int2, int1) {
			numOverlap += 1
		}
	}
	return numContained, numOverlap
}

func (s Solution[T]) Year2022Day5(input string) (string, string) {
	lines := ReadFileToLines(input)

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
		count := Must(strconv.Atoi(tokens[1]))
		start := Must(strconv.Atoi(tokens[3]))
		end := Must(strconv.Atoi(tokens[5]))

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
		count := Must(strconv.Atoi(tokens[1]))
		start := Must(strconv.Atoi(tokens[3]))
		end := Must(strconv.Atoi(tokens[5]))

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

func (s Solution[T]) Year2022Day6(input string) (int, int) {
	lines := ReadFileToLines(input)

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

func (s Solution[T]) Year2022Day7(input string) (int, int) {
	// Year2022Day7 traverses the file system using the input
	// and tracks the current working directory in an array and
	// the size of each file in a map. Each time that we
	// hit a file in the traversal, we add the file and its size
	// to the map, and we add the size to each parent directory in
	// the map.
	//
	// For example if the current path is [/,a,b,abc.txt] the algorithm
	// adds the size of abc.txt to //, //a, //a/b in the map.
	var path []string
	sizeMap := make(map[string]int)
	lines := ReadFileToLines(input)
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
			size := Must(strconv.Atoi(tokens[0]))
			for i := 1; i < len(path)+1; i++ {
				filePath := strings.Join(path[:i], "/")
				fmt.Println(filePath)
				sizeMap[filePath] += size
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

func (s Solution[T]) Year2022Day8(input string) (int, int) {
	var grid [][]int
	// demo := []string{"30373", "25512", "65332", "33549", "35390"}
	lines := ReadFileToLines(input)
	for _, line := range lines {
		var row []int
		for _, c := range line {
			row = append(row, Must(strconv.Atoi(string(c))))
		}
		grid = append(grid, row)
	}
	mx := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	type tuple struct{ horiz, vert int }
	counted := make([][]bool, len(grid))
	for i := range counted {
		counted[i] = make([]bool, len(grid))
	}

	maxSoFar := make([][]tuple, len(grid))
	for i := range maxSoFar {
		maxSoFar[i] = make([]tuple, len(grid))
	}
	// from top left
	for r := range grid {
		for c := range grid {
			if r == 0 && c == 0 {
				maxSoFar[r][c].horiz = grid[r][c]
				maxSoFar[r][c].vert = grid[r][c]
			} else if r == 0 {
				maxSoFar[r][c].horiz = mx(maxSoFar[r][c-1].horiz, grid[r][c])
				maxSoFar[r][c].vert = grid[r][c]
			} else if c == 0 {
				maxSoFar[r][c].horiz = grid[r][c]
				maxSoFar[r][c].vert = mx(maxSoFar[r-1][c].vert, grid[r][c])
			} else {
				maxSoFar[r][c].horiz = mx(maxSoFar[r][c-1].horiz, grid[r][c])
				maxSoFar[r][c].vert = mx(maxSoFar[r-1][c].vert, grid[r][c])
			}
		}
	}
	visibleTrees := 0
	for r := range grid {
		for c := range grid {
			if r == 0 || c == 0 ||
				r == len(grid)-1 || c == len(grid)-1 ||
				grid[r][c] > maxSoFar[r-1][c].vert ||
				grid[r][c] > maxSoFar[r][c-1].horiz {
				counted[r][c] = true
				visibleTrees++
			}
		}
	}
	maxSoFar = make([][]tuple, len(grid))
	for i := range maxSoFar {
		maxSoFar[i] = make([]tuple, len(grid[0]))
	}
	// from bottom right
	for r := len(grid) - 1; r >= 0; r-- {
		for c := len(grid) - 1; c >= 0; c-- {
			if r == len(grid)-1 && c == len(grid)-1 {
				maxSoFar[r][c].horiz = grid[r][c]
				maxSoFar[r][c].vert = grid[r][c]
			} else if r == len(grid)-1 {
				maxSoFar[r][c].horiz = mx(maxSoFar[r][c+1].horiz, grid[r][c])
				maxSoFar[r][c].vert = grid[r][c]
			} else if c == len(grid)-1 {
				maxSoFar[r][c].horiz = grid[r][c]
				maxSoFar[r][c].vert = mx(maxSoFar[r+1][c].vert, grid[r][c])
			} else {
				maxSoFar[r][c].horiz = mx(maxSoFar[r][c+1].horiz, grid[r][c])
				maxSoFar[r][c].vert = mx(maxSoFar[r+1][c].vert, grid[r][c])
			}
		}
	}
	for r := len(grid) - 2; r > 0; r-- {
		for c := len(grid) - 2; c > 0; c-- {
			if counted[r][c] == false &&
				(grid[r][c] > maxSoFar[r+1][c].vert ||
					grid[r][c] > maxSoFar[r][c+1].horiz) {
				visibleTrees++
			}
		}
	}
	// calc scenic scores from top left
	scenicScores := make([][]int, len(grid))
	for i := range scenicScores {
		scenicScores[i] = make([]int, len(grid[0]))
	}
	calculateScore := func(grid [][]int, row, col int) int {
		s1 := 0
		r := row
		for r > 0 {
			s1++
			if grid[r-1][col] >= grid[row][col] {
				break
			}
			r--
		}
		s2 := 0
		c := col
		for c > 0 {
			s2++
			if grid[row][c-1] >= grid[row][col] {
				break
			}
			c--
		}
		s3 := 0
		r = row
		for r < len(grid)-1 {
			s3++
			if grid[r+1][col] >= grid[row][col] {
				break
			}
			r++
		}
		s4 := 0
		c = col
		for c < len(grid)-1 {
			s4++
			if grid[row][c+1] >= grid[row][col] {
				break
			}
			c++
		}
		return s1 * s2 * s3 * s4
	}
	for r := range grid {
		for c := range grid {
			if r > 0 && c > 0 && r < len(grid)-1 && c < len(grid)-1 {
				scenicScores[r][c] = calculateScore(grid, r, c)
			}
		}
	}
	maxScore := 0
	for r := range scenicScores {
		for c := range scenicScores {
			if scenicScores[r][c] > maxScore {
				maxScore = scenicScores[r][c]
			}
		}
	}
	return visibleTrees, maxScore
}

func (s Solution[T]) Year2022Day9(input string) (int, int) {
	type coord struct{ r, c int }
	visited1 := make(map[coord]bool)
	visited9 := make(map[coord]bool)
	knots := []coord{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	isTouching := func(head coord, tail coord) bool {
		return math.Abs(float64(head.c)-float64(tail.c)) < 2 && math.Abs(float64(head.r)-float64(tail.r)) < 2
	}
	lines := ReadFileToLines(input)
	for _, line := range lines {
		ins := strings.Split(line, " ")
		dir, count := ins[0], Must(strconv.Atoi(ins[1]))
		for move := 0; move < count; move++ {
			for i := range knots {
				if i == len(knots)-1 {
					continue
				}
				if i == 0 {
					switch dir {
					case "R":
						knots[i].c++
					case "L":
						knots[i].c--
					case "U":
						knots[i].r++
					case "D":
						knots[i].r--
					}
				}
				touching := isTouching(knots[i], knots[i+1])
				if !touching {
					if knots[i].r > knots[i+1].r {
						knots[i+1].r++
					} else if knots[i].r < knots[i+1].r {
						knots[i+1].r--
					}
					if knots[i].c > knots[i+1].c {
						knots[i+1].c++
					} else if knots[i].c < knots[i+1].c {
						knots[i+1].c--
					}
				}
				if i == 0 {
					visited1[knots[i+1]] = true
				} else if i == 8 {
					visited9[knots[i+1]] = true
				}
			}
		}
	}
	count1 := 0
	for k := range visited1 {
		if visited1[k] {
			count1++
		}
	}
	count9 := 0
	for k := range visited9 {
		if visited9[k] {
			count9++
		}
	}
	return count1, count9
}

func (s Solution[T]) Year2022Day10(input string) (int, [][]string) {
	lines := ReadFileToLines(input)
	type instr struct {
		cmd string
		arg int
	}
	getSignalStrength := func(clock int, X int) int {
		if clock%40-20 != 0 {
			return 0
		}
		return clock * X
	}
	var instructions []instr
	for _, line := range lines {
		s := strings.Split(line, " ")
		instructions = append(instructions, instr{cmd: s[0]})
		if len(s) > 1 {
			instructions[len(instructions)-1].arg = Must(strconv.Atoi(s[1]))
		}
	}
	drawPixel := func(screen [][]string, clock, X int) {
		row := clock / 40
		pos := clock % 40
		if X-1 == pos || X == pos || X+1 == pos {
			screen[row][pos] = "0"
		} else {
			screen[row][pos] = " "
		}
	}
	clock := 1
	ins := 0
	argComplete := false
	totalSignalStrength := 0
	X := 1
	screen := make([][]string, 6)
	for i := range screen {
		row := make([]string, 40)
		screen[i] = row
	}
	for ins < len(instructions) {
		drawPixel(screen, clock-1, X)
		curr := instructions[ins]
		if curr.cmd == "noop" {
			ins++
		} else if !argComplete {
			argComplete = !argComplete
		} else {
			argComplete = !argComplete
			X += instructions[ins].arg
			ins++
		}
		clock++
		totalSignalStrength += getSignalStrength(clock, X)
	}
	return totalSignalStrength, screen
}

func (s Solution[T]) Year2022Day11() (int, int) {
	type monkeyType struct {
		items                                   []int
		opp                                     string
		val, test, ifTrue, ifFalse, inspections int
	}

	// demo data
	demoMonkeys := []monkeyType{
		{items: []int{79, 98}, opp: "*", val: 19, test: 23, ifTrue: 2, ifFalse: 3},
		{items: []int{54, 65, 75, 74}, opp: "+", val: 6, test: 19, ifTrue: 2, ifFalse: 0},
		{items: []int{79, 60, 97}, opp: "*", val: -1, test: 13, ifTrue: 1, ifFalse: 3},
		{items: []int{74}, opp: "+", val: 3, test: 17, ifTrue: 0, ifFalse: 1},
	}
	_ = demoMonkeys // compile if unused
	// monkeys data
	monkeysData := []monkeyType{{
		items: []int{62, 92, 50, 63, 62, 93, 73, 50}, opp: "*", val: 7, test: 2, ifTrue: 7, ifFalse: 1}, {
		items: []int{51, 97, 74, 84, 99}, opp: "+", val: 3, test: 7, ifTrue: 2, ifFalse: 4}, {
		items: []int{98, 86, 62, 76, 51, 81, 95}, opp: "+", val: 4, test: 13, ifTrue: 5, ifFalse: 4}, {
		items: []int{53, 95, 50, 85, 83, 72}, opp: "+", val: 5, test: 19, ifTrue: 6, ifFalse: 0}, {
		items: []int{59, 60, 63, 71}, opp: "*", val: 5, test: 11, ifTrue: 5, ifFalse: 3}, {
		items: []int{92, 65}, opp: "*", val: -1, test: 5, ifTrue: 6, ifFalse: 3}, {
		items: []int{78}, opp: "+", val: 8, test: 3, ifTrue: 0, ifFalse: 7}, {
		items: []int{84, 93, 54}, opp: "+", val: 1, test: 17, ifTrue: 2, ifFalse: 1}}
	_ = monkeysData // compile if unused

	op := func(item int, opp string, value int) int {
		if value < 0 {
			value = item
		}
		if opp == "*" {
			return item * value
		}
		return item + value
	}
	inspect := func(m monkeyType, item int, part int) (int, int) {
		item = op(item, m.opp, m.val)
		if part == 1 {
			item = item / 3
		}
		if item%m.test == 0 {
			return item, m.ifTrue
		}
		return item, m.ifFalse
	}

	var monkeyBusiness []int

	// PART 1
	monkeys := make([]monkeyType, len(monkeysData))
	copy(monkeys, monkeysData)
	round := 0
	for round < 20 {
		for curr := range monkeys {
			for range monkeys[curr].items {
				item := monkeys[curr].items[0]
				monkeys[curr].items = monkeys[curr].items[1:]
				updatedItem, next := inspect(monkeys[curr], item, 1)
				monkeys[next].items = append(monkeys[next].items, updatedItem)
				monkeys[curr].inspections++
			}
		}
		round++
		fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", round)
		for i, m := range monkeys {
			fmt.Printf("Monkey %d: %v\n", i, m.items)
		}
		fmt.Println()
	}
	for i, m := range monkeys {
		fmt.Printf("Monkey %d: inspected items %d times\n", i, m.inspections)
	}
	fmt.Println()
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	monkeyBusiness = append(monkeyBusiness, monkeys[0].inspections*monkeys[1].inspections)

	// PART 2
	monkeys = make([]monkeyType, len(monkeysData))
	copy(monkeys, monkeysData)
	round = 0
	for round < 10000 {
		for curr := range monkeys {
			for range monkeys[curr].items {
				item := monkeys[curr].items[0]
				monkeys[curr].items = monkeys[curr].items[1:]
				updatedItem, next := inspect(monkeys[curr], item, 2)
				monkeys[next].items = append(monkeys[next].items, updatedItem)
				monkeys[curr].inspections++
			}
		}
		round++
		if round == 1 || round == 20 || round%1000 == 0 {
			fmt.Println("== After round", round, "==")
			for i, m := range monkeys {
				fmt.Printf("Monkey %d inspected items %d times\n", i, m.inspections)
			}
			fmt.Println()
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
	monkeyBusiness = append(monkeyBusiness, monkeys[0].inspections*monkeys[1].inspections)

	// for _, part := range []int{1, 2} {
	// 	M := make([]monkeyType, len(monkeysData))
	// 	copy(M, monkeysData)
	// 	round := 0
	// 	maxRound := 20
	// 	if part == 2 {
	// 		maxRound = 10000
	// 	}
	// 	for round < maxRound {
	// 		for m := range M {
	// 			for _, item := range M[m].items {
	// 				M[m].inspections++
	// 				item = op(m.items[0], m.opp, m.val)
	// 				m := M[curr]
	// 				M[curr].items = M[curr].items[1:]
	// 				M[curr].items[0] = op(m.items[0], m.opp, m.val)
	// 				if part == 1 {
	// 					M[curr].items[0] /= 3
	// 				}
	// 				next := m.ifFalse
	// 				if m.items[0]%m.test == 0 {
	// 					next = m.ifTrue
	// 				}
	// 				M[next].items = append(M[next].items, M[curr].items[0])
	// 			}
	// 		}
	// 		/*
	// 			item = op(item, m.opp, m.val)
	// 				if part == 1 {
	// 					item = item / 3
	// 				}
	// 				if item%m.test == 0 {
	// 					return item, m.ifTrue
	// 				}
	// 				return item, m.ifFalse
	// 		*/
	// 		round++
	// 		debug := false
	// 		if debug {
	// 			fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", round)
	// 			for i, m := range M {
	// 				fmt.Printf("Monkey %d: %v\n", i, m.items)
	// 			}
	// 			fmt.Println()
	// 			if round == 1 || round == 20 || round%1000 == 0 {
	// 				fmt.Println("== After round", round, "==")
	// 				for i, m := range M {
	// 					fmt.Printf("Monkey %d inspected items %d times\n", i, m.inspections)
	// 				}
	// 				fmt.Println()
	// 			}
	// 		}
	// 	}
	// 	sort.Slice(M, func(i, j int) bool {
	// 		return M[i].inspections > M[j].inspections
	// 	})
	// 	monkeyBusiness = append(monkeyBusiness, M[0].inspections*M[1].inspections)
	// }
	return monkeyBusiness[0], monkeyBusiness[1]
}

func (s Solution[T]) Year2022Day12(input string) (int, int) {
	st, en := coord{}, coord{}
	var grid [][]rune
	lines := ReadFileToLines(input)
	for r, line := range lines {
		var row []rune
		for c, col := range line {
			if col == 'S' {
				st = coord{r, c}
				row = append(row, 'a')
			} else if col == 'E' {
				en = coord{r, c}
				row = append(row, 'z')
			} else {
				row = append(row, col)
			}
		}
		grid = append(grid, row)
	}
	fmt.Println(st, en)
	// shortest path
	type vertex struct {
		coord coord
		path  []coord
	}
	visited := make(map[coord]bool)
	q := []vertex{{st, []coord{st}}}
	var path []coord
	for len(q) > 0 {
		vert := popleft(&q)
		visited[vert.coord] = true
		for _, dir := range []coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			node := coord{vert.coord.r + dir.r, vert.coord.c + dir.c}
			if node == en {
				vert.path = append(vert.path, en)
				path = vert.path
				goto end
			}
			if _, ok := visited[node]; !ok {
				if node.r >= 0 && node.r < len(grid) && node.c >= 0 && node.c < len(grid[0]) {
					q = append(q, vertex{node, append(vert.path, node)})
				}
			}
		}
	}
end:
	fmt.Println(path)
	fmt.Println(len(path))
	return 0, 0
}
