package aoc

import "strconv"

func Year2021Day1(input string) (int, int) {
	lines := ReadFile(input)
	numIncreases1 := 0
	for i := range lines {
		if i == 0 {
			continue
		}
		if Must(strconv.Atoi(lines[i])) > Must(strconv.Atoi(lines[i-1])) {
			numIncreases1++
		}
	}

	numIncreases2 := 0
	var window []int
	runningSum := 0
	for _, line := range lines {
		lintInt := Must(strconv.Atoi(line))
		if len(window) < 3 {
			window = append(window, lintInt)
			runningSum += lintInt
			continue
		}
		nextSum := runningSum - window[0] + lintInt
		window = append(window[1:], lintInt)
		if nextSum > runningSum {
			numIncreases2++
		}
		runningSum = nextSum
	}
	return numIncreases1, numIncreases2
}
