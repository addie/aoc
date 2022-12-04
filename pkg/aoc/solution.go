package aoc

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Solution struct {
	year, day int
}

func (s Solution) dataFilename() string {
	return "data/day" + strconv.Itoa(s.day) + ".txt"
}

func (s Solution) Execute() (int, int) {
	name := fmt.Sprintf("Year%dDay%d", s.year, s.day)
	inputs := []reflect.Value{reflect.ValueOf(s.dataFilename())}
	res := reflect.ValueOf(s).MethodByName(name).Call(inputs)
	return res[0].Interface().(int), res[1].Interface().(int)
}

func (s Solution) Year2022Day1(input string) (int, int) {
	strArr := readFile(input)

	currentTotal := 0
	maxTotal := 0
	var allTotals []int
	for _, line := range strArr {
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
	strArr := readFile(input)

	// Opp A for Rock, B for Paper, and C for Scissors
	// Me X for Rock, Y for Paper, and Z for Scissors
	// Score 6 win, 3 tie, 0 loss
	// Score 1 for Rock, 2 for Paper, and 3 for Scissors
	part1Total := 0
	for _, line := range strArr {
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
	for _, line := range strArr {
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
	strArr := readFile(input)
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	const lowerFactor = int32(96)
	const upperFactor = int32(38)
	res1 := int32(0)
	for _, line := range strArr {
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

	res2 := int32(0)
	for i := 0; i < len(strArr); i += 3 {
		set := make(map[int32]bool)
		first, second, third := strArr[i], strArr[i+1], strArr[i+2]
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
	strArr := readFile(input)
	type interval struct {
		s, e int
	}

	contained := func(int1 interval, int2 interval) bool {
		return int1.s <= int2.s && int1.e >= int2.e
	}

	numContained := 0
	for _, line := range strArr {
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
	for _, line := range strArr {
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

func readFile(filename string) []string {
	var lines []string
	readFile, err := os.Open(filename)
	check(err)
	defer func(readFile *os.File) {
		err := readFile.Close()
		check(err)
	}(readFile)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
