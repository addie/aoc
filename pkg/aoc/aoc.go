package aoc

import (
	"aoc/data"
	"fmt"
	"log"
	"strconv"
)

func Run(y, d string) (any, any) {
	year := Must(strconv.Atoi(y))
	day := Must(strconv.Atoi(d))
	s := Solution{
		year: year,
		day:  day,
	}
	return s.Execute()
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

func Must(i int, err error) int {
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
