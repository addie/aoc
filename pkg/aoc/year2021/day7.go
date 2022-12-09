package aoc

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day7Part1() int {
	data, err := get7()
	if err != nil {
		log.Fatal(err)
	}
	max := getMax(data)
	minCost := math.MaxInt
	for i := 0; i <= max; i++ {
		cost := calcCostSimple(data, i)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func Day7Part2() int {
	data, err := get7()
	if err != nil {
		log.Fatal(err)
	}
	max := getMax(data)
	minCost := math.MaxInt
	for i := 0; i <= max; i++ {
		cost := calcCostComplex(data, i)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func calcCostSimple(input []int, idx int) int {
	totalCost := 0
	for _, in := range input {
		cost := math.Abs(float64(in - idx))
		totalCost += int(cost)
	}
	return totalCost
}

func calcCostComplex(input []int, idx int) int {
	totalCost := 0
	for _, in := range input {
		dist := int(math.Abs(float64(in - idx)))
		cost := dist * (dist + 1) / 2
		totalCost += int(cost)
	}
	return totalCost
}

func getMax(input []int) int {
	max := 0
	for _, in := range input {
		if in > max {
			max = in
		}
	}
	return max
}

func get7() ([]int, error) {
	file, err := os.Open("data/day7")
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
