package aoc

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type Solution struct {
	year, day int
}

func (s Solution) Execute() (any, any) {
	inputs := []reflect.Value{reflect.ValueOf(s.dataFilename())}
	res := reflect.ValueOf(s).MethodByName(s.methodName()).Call(inputs)
	return res[0].Interface(), res[1].Interface()
}

func ReadFile(filename string) []string {
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

func (s Solution) dataFilename() string {
	return fmt.Sprintf("data/year%dday%d.txt", s.year, s.day)
}

func (s Solution) methodName() string {
	return fmt.Sprintf("Year%dDay%d", s.year, s.day)
}
