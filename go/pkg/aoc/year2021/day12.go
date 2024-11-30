package aoc

import (
	"bufio"
	"os"
	"strings"
)

const day12Filename = "data/day12"

type stateProblem12 struct {
	current      string
	smallSet     map[string]bool
	visitedTwice string
}

func Day12Part1() int {
	file, _ := os.Open(day12Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := buildGraph(scanner)

	paths := 0
	queue := []stateProblem12{{
		current:      "start",
		smallSet:     map[string]bool{"start": true},
		visitedTwice: "",
	}}
	for len(queue) > 0 {
		all := queue[0]
		queue = queue[1:]
		current, smallSet := all.current, all.smallSet
		if current == "end" {
			paths++
			continue
		}
		for _, neighbor := range graph[current] {
			if _, ok := smallSet[neighbor]; !ok {
				smallSetCopy := make(map[string]bool)
				for k, v := range smallSet {
					smallSetCopy[k] = v
				}
				if strings.ToLower(neighbor) == neighbor {
					smallSetCopy[neighbor] = true
				}
				queue = append(queue, stateProblem12{
					current:  neighbor,
					smallSet: smallSetCopy,
				})
			}
		}
	}
	return paths
}

func Day12Part2() int {
	file, _ := os.Open(day12Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := buildGraph(scanner)

	paths := 0
	queue := []stateProblem12{{
		current:      "start",
		smallSet:     map[string]bool{"start": true},
		visitedTwice: "",
	}}
	for len(queue) > 0 {
		all := queue[0]
		queue = queue[1:]
		current, smallSet, visitedTwice := all.current, all.smallSet, all.visitedTwice
		if current == "end" {
			paths++
			continue
		}
		for _, neighbor := range graph[current] {
			if _, ok := smallSet[neighbor]; !ok {
				smallSetCopy := make(map[string]bool)
				for k, v := range smallSet {
					smallSetCopy[k] = v
				}
				if strings.ToLower(neighbor) == neighbor {
					smallSetCopy[neighbor] = true
				}
				queue = append(queue, stateProblem12{
					current:      neighbor,
					smallSet:     smallSetCopy,
					visitedTwice: visitedTwice,
				})
			} else if _, ok := smallSet[neighbor]; ok && visitedTwice == "" &&
				neighbor != "start" && neighbor != "end" {

				queue = append(queue, stateProblem12{
					current:      neighbor,
					smallSet:     smallSet,
					visitedTwice: neighbor,
				})
			}
		}
	}
	return paths
}

func buildGraph(scanner *bufio.Scanner) map[string][]string {
	graph := make(map[string][]string)
	for scanner.Scan() {
		strRow := strings.Split(scanner.Text(), "-")
		if _, ok := graph[strRow[0]]; !ok {
			graph[strRow[0]] = []string{strRow[1]}
		} else {
			graph[strRow[0]] = append(graph[strRow[0]], strRow[1])
		}
		if _, ok := graph[strRow[1]]; !ok {
			graph[strRow[1]] = []string{strRow[0]}
		} else {
			graph[strRow[1]] = append(graph[strRow[1]], strRow[0])
		}
	}
	return graph
}
