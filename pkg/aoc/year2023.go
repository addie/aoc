package aoc

import (
	"cmp"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/mowshon/iterium"
)

func (s Solution[T]) Year2023Day1(path string) (int, int) {
	data := ReadFile(path)
	res1 := year2023Day1Part1(data)
	res2 := year2023Day1Part2(data)
	return res1, res2
}

func year2023Day1Part1(data []string) int {
	res := 0
	for _, line := range data {
		left := ""
		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				left = string(line[i])
				break
			}
		}
		right := ""
		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				right = string(line[i])
				break
			}
		}
		intVal, _ := strconv.Atoi(left + right)
		res += intVal
	}
	return res
}

func year2023Day1Part2(data []string) int {
	var numMap = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	res := 0
	for _, line := range data {
		digits := findDigits(numMap, line)
		curr := digits[0]*10 + digits[len(digits)-1]
		res += curr
	}
	return res
}

func findDigits(numMap map[string]int, line string) []int {
	var digits []int
	for i, char := range line {
		if unicode.IsDigit(char) {
			val, _ := strconv.Atoi(string(char))
			digits = append(digits, val)
			continue
		}
		for k := range numMap {
			if strings.HasPrefix(line[i:], k) {
				digits = append(digits, numMap[k])
			}
		}
	}
	return digits
}

func (s Solution[T]) Year2023Day2(path string) (int, int) {
	data := ReadFile(path)
	res1 := year2023Day2Part1(data)
	res2 := year2023Day2Part2(data)
	return res1, res2
}

func year2023Day2Part1(data []string) int {
	type Color int

	const red Color = 12
	const green Color = 13
	const blue Color = 14

	var colorMap = map[string]Color{
		"red":   red,
		"green": green,
		"blue":  blue,
	}

	validGames := make([]bool, 0, len(data))
	for _, game := range data {
		gameValid := true
		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		for _, rd := range rounds {
			rd = strings.TrimSpace(rd)
			rolls := strings.Split(rd, ", ")
			for _, roll := range rolls {
				r := strings.Split(roll, " ")
				n, color := r[0], r[1]
				num, _ := strconv.Atoi(n)
				if num > int(colorMap[color]) {
					gameValid = false
					break
				}
			}
		}
		validGames = append(validGames, gameValid)
	}
	res := 0
	for i, valid := range validGames {
		if valid {
			res += i + 1
		}
	}
	return res
}

func year2023Day2Part2(data []string) int {
	powers := make([]int, 0, len(data))
	for _, game := range data {
		minBlue, minGreen, minRed := 0, 0, 0
		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		for _, rd := range rounds {
			rd = strings.TrimSpace(rd)
			rolls := strings.Split(rd, ", ")
			for _, roll := range rolls {
				r := strings.Split(roll, " ")
				n, color := r[0], r[1]
				num, _ := strconv.Atoi(n)
				switch color {
				case "red":
					minRed = max(minRed, num)
				case "green":
					minGreen = max(minGreen, num)
				case "blue":
					minBlue = max(minBlue, num)
				}
			}
		}
		powers = append(powers, minRed*minBlue*minGreen)
	}
	res := 0
	for _, power := range powers {
		res += power
	}
	return res
}

func (s Solution[T]) Year2023Day3(path string) (int, int) {
	data := ReadFile(path)
	res1 := year2023Day3Part1(data)
	res2 := year2023Day3Part2(data)
	return res1, res2
}

func isDigit(s string) bool { return strings.ContainsAny(s, "0123456789") }

func year2023Day3Part1(data []string) int {
	grid := make([][]string, len(data))
	for i := range data {
		grid[i] = strings.Split(data[i], "")
	}
	var partNumbers []int
	isPrevDigit := false
	isPartNo := false
	currDigit := ""
	checkPartNumber := func(grid [][]string, r, c int) bool {
		directions := []coord{{r: 0, c: 1}, {r: 1, c: 0}, {r: 1, c: 1}, {r: -1, c: 0}, {r: 0, c: -1}, {r: -1, c: -1}, {r: -1, c: 1}, {r: 1, c: -1}}
		for _, dir := range directions {
			nc := coord{r: r + dir.r, c: c + dir.c}
			if nc.r < 0 || nc.r >= len(grid) || nc.c < 0 || nc.c >= len(grid[0]) {
				continue
			}
			if grid[nc.r][nc.c] != "." && !isDigit(grid[nc.r][nc.c]) {
				return true
			}
		}
		return false
	}
	nextIsDigit := func(grid [][]string, r, c int) bool {
		if c+1 < len(grid[0]) && isDigit(grid[r][c+1]) {
			return true
		}
		return false
	}
	for r := range grid {
		for c := range grid[r] {
			cell := grid[r][c]
			if isDigit(cell) {
				currDigit += cell
			} else {
				currDigit = ""
				isPrevDigit = false
				isPartNo = false
				continue
			}
			if !nextIsDigit(grid, r, c) {
				if checkPartNumber(grid, r, c) {
					isPartNo = true
				}
				if isPartNo {
					partNumbers = append(partNumbers, toInt(currDigit))
					currDigit = ""
					continue
				}
			}
			if !isPrevDigit {
				isPrevDigit = true
				currDigit = cell
				if checkPartNumber(grid, r, c) {
					isPartNo = true
				}
				continue
			} else {
				if checkPartNumber(grid, r, c) {
					isPartNo = true
				}
				continue
			}
		}
		currDigit = ""
		isPrevDigit = false
		isPartNo = false
	}
	sumOfPartNos := 0
	for _, partNo := range partNumbers {
		sumOfPartNos += partNo
	}
	return sumOfPartNos
}

func year2023Day3Part2(data []string) int {
	grid := make([][]string, len(data))
	for i := range data {
		grid[i] = strings.Split(data[i], "")
	}
	R, C := len(grid), len(grid[0])
	nums := make(map[coord][]int)
	for r := range grid {
		// set containing positions of '*' (gears)
		// collect all the gears surrounding each number group
		gears := make(map[coord]struct{})
		n := 0
		for c := range grid[0] {
			if c < C && isDigit(grid[r][c]) {
				n = n*10 + Must(strconv.Atoi(grid[r][c]))
				for _, rd := range []int{-1, 0, 1} {
					for _, cd := range []int{-1, 0, 1} {
						nr, nc := r+rd, c+cd
						if nr >= 0 && nr < R && nc >= 0 && nc < C {
							newRC := grid[nr][nc]
							if newRC == "*" {
								gears[coord{r: nr, c: nc}] = struct{}{}
							}
						}
					}
				}
			} else if n > 0 { // every time you pass a number, add the gears to the map
				for gear := range gears {
					nums[gear] = append(nums[gear], n)
				}
				n = 0
				gears = make(map[coord]struct{})
			}
		}
	}
	res := 0
	for _, v := range nums {
		if len(v) == 2 {
			res += v[0] * v[1]
		}
	}
	return res
}

func (s Solution[T]) Year2023Day4(path string) (int, int) {
	data := ReadFile(path)
	res1 := year2023Day4Part1(data)
	res2 := year2023Day4Part2(data)
	return res1, res2
}

// Card 1: 41 48 83 86 17 | 83 86 6 31 17 9 48 53
func year2023Day4Part1(data []string) int {
	allPoints := 0
	for _, card := range data {
		mine := make(map[string]struct{})
		d := strings.Split(strings.Split(card, ":")[1], "|")
		card := strings.Fields(strings.TrimSpace(d[0]))
		mineList := strings.Fields(strings.TrimSpace(d[1]))
		for _, m := range mineList {
			mine[m] = struct{}{}
		}
		pow := 0
		for _, n := range card {
			if _, ok := mine[n]; ok {
				pow++
			}
		}
		curr := math.Pow(2, float64(pow)-1)
		allPoints += int(curr)
	}
	return allPoints
}

func year2023Day4Part2(data []string) int {
	calculateMatches := func(matchCount map[int]int, cardNo int, card []string, mine map[string]struct{}) int {
		matches := 0
		for _, n := range card {
			if _, ok := mine[n]; ok {
				matches++
			}
		}
		matchCount[cardNo] = matches
		return matches
	}
	addCards := func(m map[int]int, count map[int]int, cardNo int) {
		numOfCards := count[cardNo]
		for numOfCards > 0 {
			cardNo++
			m[cardNo]++
			numOfCards--
		}
	}
	summarize := func(matchCount map[int]int, cardCount int) map[int]int {
		m := make(map[int]int)
		for i := 1; i < cardCount; i++ {
			addCards(m, matchCount, i)
			for j := 0; j < m[i]; j++ {
				addCards(m, matchCount, i)
			}
		}
		return m
	}

	matchCount := make(map[int]int)
	for i, card := range data {
		mine := make(map[string]struct{})
		d := strings.Split(strings.Split(card, ":")[1], "|")
		card := strings.Fields(strings.TrimSpace(d[0]))
		mineList := strings.Fields(strings.TrimSpace(d[1]))
		for _, m := range mineList {
			mine[m] = struct{}{}
		}
		calculateMatches(matchCount, i+1, card, mine)
	}
	m := summarize(matchCount, len(data))
	res := 0
	for _, v := range m {
		res += v
	}
	return res + len(data)
}

func (s Solution[T]) Year2023Day5(path string) (int, int) {
	data := ReadFileToString(path)
	res1 := year2023Day5Part1(data)
	res2 := year2023Day5Part2(data)
	return res1, res2
}

/*********************/
/* Helpers for Day 5 */
/*********************/

func minLists[T cmp.Ordered](lists [][]T, key ...func(t []T) T) []T {
	k := func(t []T) T {
		return t[0]
	}
	if key != nil {
		k = key[0]
	}
	minList := lists[0]
	for i := 1; i < len(lists); i++ {
		if k(lists[i]) < k(minList) {
			minList = lists[i]
		}
	}
	return minList
}

type fs struct {
	tuples [][]int
}

// fs constructor
func newFS(s string) fs {
	f := fs{}
	lines := strings.Split(s, "\n")[1:]
	for _, line := range lines {
		var grouping []int
		for _, x := range strings.Fields(line) {
			grouping = append(grouping, toInt(x))
		}
		f.tuples = append(f.tuples, grouping)
	}
	return f
}

func (f fs) applyOne(x int) int {
	for _, tuple := range f.tuples {
		dst, src, sz := tuple[0], tuple[1], tuple[2]
		if src <= x && x < src+sz {
			return x + dst - src
		}
	}
	return x
}
func (f fs) applyRange(R [][]int) [][]int {
	var A [][]int
	for _, t := range f.tuples {
		dest, src, sz := t[0], t[1], t[2]
		srcEnd := src + sz
		var NR [][]int
		for len(R) > 0 {
			// [start                                     end)
			//          [src       srcEnd]
			// [BEFORE ][MIDDLE          ][AFTER             )
			a := R[len(R)-1]
			R = R[:len(R)-1]
			start, end := a[0], a[1]
			// (src,sz) might cut (start,end)
			before := []int{start, min(end, src)}
			middle := []int{max(start, src), min(srcEnd, end)}
			after := []int{max(srcEnd, start), end}
			if before[1] > before[0] {
				NR = append(NR, before)
			}
			if middle[1] > middle[0] {
				A = append(A, []int{middle[0] - src + dest, middle[1] - src + dest})
			}
			if after[1] > after[0] {
				NR = append(NR, after)
			}
		}
		R = NR
	}
	return append(A, R...)
}

func year2023Day5Part1(data string) int {
	parts := strings.Split(data, "\n\n")
	seedStr, others := parts[0], parts[1:]
	var seed []int
	for _, x := range strings.Fields(strings.Split(seedStr, ":")[1]) {
		seed = append(seed, toInt(x))
	}
	var ff []fs
	for _, other := range others {
		ff = append(ff, newFS(strings.TrimSpace(other)))
	}
	var resList []int
	for _, x := range seed {
		for _, f := range ff {
			x = f.applyOne(x)
		}
		resList = append(resList, x)
	}
	return slices.Min(resList)
}

func year2023Day5Part2(data string) int {
	parts := strings.Split(data, "\n\n")
	seedStr, others := parts[0], parts[1:]
	var seed []int
	for _, x := range strings.Fields(strings.Split(seedStr, ":")[1]) {
		seed = append(seed, toInt(x))
	}
	var ff []fs
	for _, other := range others {
		ff = append(ff, newFS(strings.TrimSpace(other)))
	}
	pairs := zip(alt(seed), alt(seed[1:]))
	var resList []int
	for _, t := range pairs {
		st, sz := t[0], t[1]
		// inclusive on the left, exclusive on the right
		// e.g. [1,3) = [1,2]
		// length of [a,b) = b-a
		// [a,b) + [b,c) = [a,c)
		R := [][]int{{st, st + sz}}
		for _, f := range ff {
			R = f.applyRange(R)
		}
		resList = append(resList, minLists(R)[0])
	}
	return slices.Min(resList)
}

// Year2023Day6 using a Simple linear search solution.
// One optimization we can do is use binary search to find the two points
// where the travel time breaks the record and then assume all times in between
// should be counted
func (s Solution[T]) Year2023Day6(path string) (int, int) {
	data := ReadFileToString(path)
	res1 := year2023Day6Part1(data)
	res2 := year2023Day6Part2(data)
	return res1, res2
}

func year2023Day6Part1(data string) int {
	lines := strings.Split(data, "\n")
	times := strings.Fields(strings.Split(lines[0], ":")[1])
	dists := strings.Fields(strings.Split(lines[1], ":")[1])
	res := 1
	for i, t := range times {
		time := toInt(t)
		c := 0
		for vel := 1; vel < time; vel++ {
			dist := vel * (time - vel)
			if dist > toInt(dists[i]) {
				c++
			}
		}
		if c > 0 {
			res *= c
		}
	}
	return res
}

func year2023Day6Part2(data string) int {
	lines := strings.Split(data, "\n")
	time := toInt(strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), ""))
	record := toInt(strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), ""))
	res := 1
	c := 0
	for vel := 1; vel < time; vel++ {
		dist := vel * (time - vel)
		if dist > record {
			c++
		}
	}
	if c > 0 {
		res *= c
	}
	return res
}

func (s Solution[T]) Year2023Day7(path string) (int, int) {
	data := ReadFileToString(path)
	res1 := year2023Day7Part1(data)
	res2 := year2023Day7Part2(data)
	return res1, res2
}

// Hands
// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2
// 7 Five of a kind, where all five cards have the same label: AAAAA
// 6 Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 5 Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 4 Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 3 Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 2 One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 1 High card, where all cards' labels are distinct: 23456
func year2023Day7Part1(data string) int {
	type hb struct {
		hand string
		bid  int
	}
	var hands []hb
	handsBids := strings.Split(data, "\n")
	for _, handsBid := range handsBids {
		if handsBid != "" {
			f := strings.Fields(handsBid)
			hands = append(hands, hb{hand: f[0], bid: toInt(f[1])})
		}
	}
	getPoints := func(a hb) int {
		tc := make(map[int32]int)
		for _, c := range a.hand {
			tc[c]++
		}
		c := make([]int, 5)
		for _, v := range tc {
			c[5-v]++
		}
		if c[0] > 0 {
			return 7
		} // five of a kind
		if c[1] > 0 {
			return 6
		} // four of a kind
		if c[2] > 0 && c[3] > 0 {
			return 5
		} // full house
		if c[2] > 0 {
			return 4
		} // three of a kind
		if c[1] > 1 {
			return 3
		} // two pair
		if c[1] > 0 {
			return 2
		} // one pair
		return 1
	}
	slices.SortFunc(hands, func(a, b hb) int {
		cardRanks := map[string]int{
			"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9,
			"8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
		}
		//
		pointsA := getPoints(a)
		pointsB := getPoints(b)
		if pointsA < pointsB {
			return -1
		}
		if pointsA > pointsB {
			return 1
		}
		for i := 0; i < 5; i++ {
			if cardRanks[string(a.hand[i])] < cardRanks[string(b.hand[i])] {
				return -1
			}
			if cardRanks[string(a.hand[i])] > cardRanks[string(b.hand[i])] {
				return 1
			}
		}
		return 0
	})
	res := 0
	for i, h := range hands {
		rank := i + 1
		res += h.bid * rank
	}
	return res
}

func year2023Day7Part2(data string) int {
	type hb struct {
		hand string
		bid  int
	}
	var hands []hb
	handsBids := strings.Split(data, "\n")
	for _, handsBid := range handsBids {
		if handsBid != "" {
			f := strings.Fields(handsBid)
			hands = append(hands, hb{hand: f[0], bid: toInt(f[1])})
		}
	}
	getPoints := func(tc map[int32]int) int {
		c := make([]int, 5)
		for _, v := range tc {
			c[5-v]++
		}
		if c[0] > 0 {
			return 7
		} // five of a kind
		if c[1] > 0 {
			return 6
		} // four of a kind
		if c[2] > 0 && c[3] > 0 {
			return 5
		} // full house
		if c[2] > 0 {
			return 4
		} // three of a kind
		if c[1] > 1 {
			return 3
		} // two pair
		if c[1] > 0 {
			return 2
		} // one pair
		return 1
	}
	getRank := func(a hb) int {
		tc := make(map[int32]int)
		for _, c := range a.hand {
			tc[c]++
		}
		jc := tc['J']
		delete(tc, 'J')
		if jc == 0 {
			return getPoints(tc)
		}
		combinations := iterium.CombinationsWithReplacement([]int32{'A', 'K', 'Q', 'T', '9', '8',
			'7', '6', '5', '4', '3', '2'}, jc)
		rank := 0
		for _, c := range Must(combinations.Slice()) {
			newTC := make(map[int32]int)
			maps.Copy(newTC, tc)
			for i := 0; i < jc; i++ {
				newTC[c[i]]++
			}
			r := getPoints(newTC)
			rank = max(rank, r)
		}
		return rank
	}
	slices.SortFunc(hands, func(a, b hb) int {
		rankA := getRank(a)
		rankB := getRank(b)
		if rankA < rankB {
			return -1
		}
		if rankA > rankB {
			return 1
		}
		cardRanks := map[string]int{
			"A": 14, "K": 13, "Q": 12, "T": 11, "9": 10, "8": 9,
			"7": 8, "6": 7, "5": 6, "4": 5, "3": 4, "2": 3, "J": 2,
		}
		for i := 0; i < 5; i++ {
			if cardRanks[string(a.hand[i])] < cardRanks[string(b.hand[i])] {
				return -1
			}
			if cardRanks[string(a.hand[i])] > cardRanks[string(b.hand[i])] {
				return 1
			}
		}
		return 0
	})
	res := 0
	for i, h := range hands {
		rank := i + 1
		res += h.bid * rank
	}
	return res
}
