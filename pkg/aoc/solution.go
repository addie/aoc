package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
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
	if !errors.Is(err, os.ErrNotExist) {
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

func ReadFile(filename string) []string {
	var lines []string
	readFile := Must(os.Open(filename))
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
func ReadFileToString(filename string) string {
	return string(Must(os.ReadFile(filename)))
}

func in[T comparable](val T, container []T) bool {
	for _, v := range container {
		if v == val {
			return true
		}
	}
	return false
}

func pop[T any](alist *[]T) T {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func popleft[T any](alist *[]T) T {
	lv := (*alist)[0]
	*alist = (*alist)[1:]
	return lv
}

func printGrid[T any](grid [][]T) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
