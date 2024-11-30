package aoc

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Run(d, filename string, demo bool) (any, any) {
	name := strings.Split(filename, ".")[0]
	reg, _ := regexp.Compile(`year(\d+)`)
	y := reg.FindStringSubmatch(name)[1]
	year := Must(strconv.Atoi(y))
	day := Must(strconv.Atoi(d))
	s := Solution[int]{
		year: year,
		day:  day,
	}
	s.saveData()
	return s.Execute(demo)
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
