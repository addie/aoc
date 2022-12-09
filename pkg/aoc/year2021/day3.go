package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	totalStrings = 1000
	binLength    = 12
)

func Day3Part1() int {
	data, err := get3()
	if err != nil {
		log.Fatal(err)
	}

	oneCount := make(map[int]int)
	initCount(oneCount)
	for _, dat := range data {
		for i, digit := range dat {
			if string(digit) == "1" {
				oneCount[i] += 1
			}
		}
	}
	mostCommon := ""
	leastCommon := ""
	for i := 0; i < binLength; i++ {
		if oneCount[i] > totalStrings/2 {
			mostCommon += "1"
			leastCommon += "0"
		} else {
			mostCommon += "0"
			leastCommon += "1"
		}
	}
	gamma, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(leastCommon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(gamma) * int(epsilon)
}

func Day3Part2() int {
	data, err := get3()
	if err != nil {
		log.Fatal(err)
	}
	oxygen := buildOxygenNum(data)
	co2 := buildCO2Num(data)
	return oxygen * co2
}

func buildOxygenNum(currentList []string) int {
	oneCount := 0
	var oxygen int64
	for i := 0; i < binLength; i++ {
		for _, binStr := range currentList {
			if string(binStr[i]) == "1" {
				oneCount += 1
			}
		}
		keepOnes := oneCount >= len(currentList)-oneCount
		oneCount = 0
		currentList = cull(i, currentList, keepOnes, 1)
		if len(currentList) == 1 {
			var err error
			oxygen, err = strconv.ParseInt(currentList[0], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			return int(oxygen)
		}
	}
	return -1
}

func buildCO2Num(currentList []string) int {
	zeroCount := 0
	var co2 int64
	for i := 0; i < binLength; i++ {
		for _, binStr := range currentList {
			if string(binStr[i]) == "0" {
				zeroCount += 1
			}
		}
		keepZeros := zeroCount <= len(currentList)-zeroCount
		zeroCount = 0
		currentList = cull(i, currentList, keepZeros, 0)
		if len(currentList) == 1 {
			var err error
			co2, err = strconv.ParseInt(currentList[0], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			return int(co2)
		}
	}
	return -1
}

func cull(i int, currentList []string, invariant bool, digitToKeep int) []string {
	var newList []string
	for _, binStr := range currentList {
		if invariant && string(binStr[i]) == strconv.Itoa(digitToKeep) {
			newList = append(newList, binStr)
		} else if !invariant && string(binStr[i]) == strconv.Itoa(1-digitToKeep) {
			newList = append(newList, binStr)
		}
	}
	return newList
}

func initCount(count map[int]int) {
	for i := 1; i < 13; i++ {
		count[i] = 0
	}
}

func get3() ([]string, error) {
	file, err := os.Open("data/day3")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(res) == 0 {
		log.Fatal("read 0 bytes from file")
	}
	return res, nil
}
