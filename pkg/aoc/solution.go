package aoc

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Solution[T any] struct {
	year, day int
}

func (s Solution[T]) Execute(demo bool) (any, any) {
	input := s.dataFilename(demo)
	inputs := []reflect.Value{reflect.ValueOf(input)}
	res := reflect.ValueOf(s).MethodByName(s.methodName()).Call(inputs)
	return res[0].Interface(), res[1].Interface()
}

func (s Solution[T]) dataFilename(demo bool) string {
	if demo {
		return fmt.Sprintf("data/year%dday%ddemo.txt", s.year, s.day)
	}
	return fmt.Sprintf("data/year%dday%d.txt", s.year, s.day)
}

func (s Solution[T]) methodName() string {
	return fmt.Sprintf("Year%dDay%d", s.year, s.day)
}

func (s Solution[T]) saveData() {
	filename := fmt.Sprintf(Filename, s.year, s.day)

	file, err := os.Open(filename)
	if errors.Is(err, os.ErrExist) {
		file.Close()
		return
	}

	err = Get(s.day, s.year, filename)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (s Solution[T]) post(level, res int) {
	err := Post(s.day, level, res)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFileToLines(filename string) []string {
	str := strings.TrimSuffix(string(Must(os.ReadFile(filename))), "\n")
	return strings.Split(str, "\n")
}

func ReadFileToString(filename string) string {
	return strings.TrimSuffix(string(Must(os.ReadFile(filename))), "\n")
}

func ReadFileToIntGrid(filename string) [][]int {
	str := strings.TrimSuffix(string(Must(os.ReadFile(filename))), "\n")
	lines := strings.Split(str, "\n")
	grid := make([][]int, len(lines))
	for r := range lines {
		for _, v := range lines[r] {
			grid[r] = append(grid[r], Must(strconv.Atoi(string(v))))
		}
	}
	return grid
}

func ReadFileToByteGrid(filename string) [][]byte {
	str := strings.TrimSuffix(string(Must(os.ReadFile(filename))), "\n")
	lines := strings.Split(str, "\n")
	grid := make([][]byte, len(lines))
	for r := range lines {
		for _, v := range lines[r] {
			grid[r] = append(grid[r], byte(v))
		}
	}
	return grid
}
