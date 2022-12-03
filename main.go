package main

import (
	"aoc/cmd/aoc"
	"aoc/data"
	"fmt"
	"log"
)

type version struct {
	day  int
	part int
}

func main() {
	aoc.Execute()
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
