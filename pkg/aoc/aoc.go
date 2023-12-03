package aoc

import (
	"log"
	"strconv"
)

func Run(y, d string) (any, any) {
	year := Must(strconv.Atoi(y))
	day := Must(strconv.Atoi(d))
	s := Solution[int]{
		year: year,
		day:  day,
	}
	s.saveData()
	return s.Execute()
}

func Must[T any](i T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
