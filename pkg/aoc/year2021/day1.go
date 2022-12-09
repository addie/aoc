package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Day1Part1() int {
	data, err := get1()
	if err != nil {
		log.Fatal(err)
	}
	numIncreases := 0
	for i := range data {
		if i == 0 {
			continue
		}
		if data[i] > data[i-1] {
			numIncreases++
		}
	}
	return numIncreases
}

func Day1Part2() int {
	data, err := get1()
	if err != nil {
		log.Fatal(err)
	}
	numIncreases := 0
	var window []int
	runningSum := 0
	for i := range data {
		if len(window) < 3 {
			window = append(window, data[i])
			runningSum += data[i]
			continue
		}
		nextSum := runningSum - window[0] + data[i]
		window = append(window[1:], data[i])
		if nextSum > runningSum {
			numIncreases++
		}
		runningSum = nextSum
	}
	return numIncreases
}

func get1() ([]int, error) {
	file, err := os.Open("data/day1")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digit, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, digit)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		log.Fatal("read 0 bytes from file")
	}
	return res, nil
}
