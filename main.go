package main

import (
	"aoc2021/aoc"
	"aoc2021/data"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type version struct {
	day  int
	part int
}

func main() {
	problemSet := os.Args[1:][0]
	dayStr := strings.Split(problemSet, ".")[0]
	day, _ := strconv.Atoi(dayStr)

	saveData(day)
	res := 0
	switch problemSet {
	case "1.1":
		res = aoc.Day1Part1()
		fmt.Printf("Answer %d\n", res)
	case "1.2":
		res = aoc.Day1Part2()
		fmt.Printf("Answer %d\n", res)
	case "2.1":
		res = aoc.Day2Part1()
		fmt.Printf("Answer %d\n", res)
	case "2.2":
		res = aoc.Day2Part2()
		fmt.Printf("Answer %d\n", res)
	case "3.1":
		res = aoc.Day3Part1()
		fmt.Printf("Answer %d\n", res)
	case "3.2":
		res = aoc.Day3Part2()
		fmt.Printf("Answer %d\n", res)
	case "4.1":
		res = aoc.Day4Part1()
		fmt.Printf("Answer %d\n", res)
	case "4.2":
		res = aoc.Day4Part2()
		fmt.Printf("Answer %d\n", res)
	case "5.1":
		res = aoc.Day5Part1()
		fmt.Printf("Answer %d\n", res)
	case "5.2":
		res = aoc.Day5Part2()
		fmt.Printf("Answer %d\n", res)
	case "6.1":
		res = aoc.Day6Part1()
		fmt.Printf("Answer %d\n", res)
	case "6.2":
		res = aoc.Day6Part2()
		fmt.Printf("Answer %d\n", res)
	case "7.1":
		res = aoc.Day7Part1()
		fmt.Printf("Answer %d\n", res)
	case "7.2":
		res = aoc.Day7Part2()
		fmt.Printf("Answer %d\n", res)
	case "8.1":
		res = aoc.Day8Part1()
		fmt.Printf("Answer %d\n", res)
	case "8.2":
		res = aoc.Day8Part2()
		fmt.Printf("Answer %d\n", res)
	case "9.1":
		res = aoc.Day9Part1()
		fmt.Printf("Answer %d\n", res)
	case "9.2":
		res = aoc.Day9Part2()
		fmt.Printf("Answer %d\n", res)
	case "10.1":
		res = aoc.Day10Part1()
		fmt.Printf("Answer %d\n", res)
	case "10.2":
		res = aoc.Day10Part2()
		fmt.Printf("Answer %d\n", res)
	case "11.1":
		res = aoc.Day11Part1()
		fmt.Printf("Answer %d\n", res)
	case "11.2":
		res = aoc.Day11Part2()
		fmt.Printf("Answer %d\n", res)
	case "12.1":
		res = aoc.Day12Part1()
		fmt.Printf("Answer %d\n", res)
	case "12.2":
		res = aoc.Day12Part2()
		fmt.Printf("Answer %d\n", res)
	case "13", "13.1", "13.2":
		res = aoc.Day13()
		fmt.Printf("Answer %d\n", res)
	case "14", "14.1", "14.2":
		res = aoc.Day14()
		fmt.Printf("Answer %d\n", res)
	case "15", "15.1", "15.2":
		res = aoc.Day15()
		fmt.Printf("Answer %d\n", res)
	case "17", "17.1", "17.2":
		p1, p2 := aoc.Day17()
		fmt.Printf("%d\n%d\n", p1, p2)
	}
}

func saveData(day int) string {
	filename := fmt.Sprintf(data.Filename, day)
	err := data.Get(day, filename)
	if err != nil {
		log.Fatal(err)
	}
	return filename
}

func post(day, level, res int) {
	err := data.Post(day, level, res)
	if err != nil {
		log.Fatal(err)
	}
}
