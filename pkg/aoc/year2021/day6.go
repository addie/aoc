package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day6Part1() int {
	data, err := get6()
	if err != nil {
		log.Fatal(err)
	}
	const totalDays = 80
	return lanternFishCount(totalDays, data)
}

func Day6Part2() int {
	data, err := get6()
	if err != nil {
		log.Fatal(err)
	}
	const totalDays = 256
	return lanternFishCount(totalDays, data)
}

func lanternFishCount(days int, input []int) int {
	const resetFish = 6
	const newFish = 8
	counter := counter(input)
	for day := 0; day < days; day++ {
		fishMap := make(map[int]int) // tracks how many fish have n numbers to give birth
		for numDays, count := range counter {
			if numDays == 0 {
				fishMap[resetFish] += count
				fishMap[newFish] += count
			} else {
				fishMap[numDays-1] += count
			}
		}
		counter = fishMap
	}
	return sum(counter)
}

func sum(counter map[int]int) int {
	count := 0
	for _, val := range counter {
		count += val
	}
	return count
}

func counter(input []int) map[int]int {
	c := make(map[int]int)
	for _, in := range input {
		if _, ok := c[in]; ok {
			c[in]++
		} else {
			c[in] = 1
		}
	}
	return c
}

func get6() ([]int, error) {
	file, err := os.Open("data/day6")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		parsedRow := strings.Split(row, ",")
		for _, n := range parsedRow {
			num, _ := strconv.Atoi(n)
			input = append(input, num)
		}
	}
	return input, nil
}
