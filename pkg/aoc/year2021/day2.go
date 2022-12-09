package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() int {
	data, err := get2()
	if err != nil {
		log.Fatal(err)
	}

	res := pos{}
	for _, fullCmd := range data {
		dir := fullCmd[0]
		num, _ := strconv.Atoi(fullCmd[1])
		switch dir {
		case "forward":
			res.horizontal += num
		case "down":
			res.depth += num
		case "up":
			res.depth -= num
		}
	}
	return res.horizontal * res.depth
}

func Day2Part2() int {
	data, err := get2()
	if err != nil {
		log.Fatal(err)
	}
	res := pos{}
	for _, fullCmd := range data {
		dir := fullCmd[0]
		num, _ := strconv.Atoi(fullCmd[1])
		switch dir {
		case "forward":
			res.horizontal += num
			res.depth += res.aim * num
		case "down":
			res.aim += num
		case "up":
			res.aim -= num
		}
	}
	return res.horizontal * res.depth
}

type pos struct {
	horizontal, depth, aim int
}

func get2() ([][]string, error) {
	file, err := os.Open("data/day2")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fullCmd := strings.Split(scanner.Text(), " ")
		dir, numStr := fullCmd[0], fullCmd[1]
		res = append(res, []string{dir, numStr})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		log.Fatal("read 0 bytes from file")
	}
	return res, nil
}
