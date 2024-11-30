package aoc

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day8Part1() int {
	file, _ := os.Open("data/day8")
	defer file.Close()

	numTimes := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		splitRow := strings.Split(row, "|")
		_, after := splitRow[0], splitRow[1]
		afterList := strings.Split(strings.TrimSpace(after), " ")
		for _, num := range afterList {
			switch len(num) {
			case 2, 3, 4, 7:
				numTimes++
			}
		}
	}

	return numTimes
}

var letters = map[string]string{
	"abcefg":  "0",
	"cf":      "1",
	"acdeg":   "2",
	"acdfg":   "3",
	"bcdf":    "4",
	"abdfg":   "5",
	"abdefg":  "6",
	"acf":     "7",
	"abcdefg": "8",
	"abcdfg":  "9",
}

// beacf afbd bcead cgefa ecdbga efb gbfdeac ecgfbd acbdfe fb | bf efb bgecdfa egcfa

func Day8Part2() int {
	file, _ := os.Open("data/day8")
	defer file.Close()

	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		splitRow := strings.Split(row, "|")
		before := strings.Split(strings.TrimSpace(splitRow[0]), " ")
		after := strings.Split(strings.TrimSpace(splitRow[1]), " ")

		numStr := ""
		M := buildMap(before)
		for _, w := range after {
			var p []string
			for i := 0; i < len(w); i++ {
				p = append(p, string(M[w[i]]))
			}
			sort.Strings(p)
			word := strings.Join(p, "")
			numStr += letters[word]
		}
		num, _ := strconv.Atoi(numStr)
		res += num
	}
	return res
}

func buildMap(before []string) map[byte]byte {
	M := make(map[byte]byte)
	cf := deriveCF(before)
	wIsSix(before, cf, M)
	wIsSeven(before, cf, M)
	bd := deriveBD(before, cf)
	wIsZero(before, bd, M)
	eg := deriveEG(M)
	wIsNine(before, eg, M)
	return M
}

func deriveEG(M map[byte]byte) string {
	eg := ""
	for _, c := range []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
		if _, ok := M[c]; !ok {
			eg += string(c)
		}
	}
	return eg
}

func wIsNine(before []string, eg string, M map[byte]byte) {
	for _, w := range before {
		if len(w) == 6 && in(w, eg[0]) != in(w, eg[1]) {
			if in(w, eg[0]) {
				M[eg[0]] = 'g'
				M[eg[1]] = 'e'
			} else {
				M[eg[0]] = 'e'
				M[eg[1]] = 'g'
			}
		}
	}
}

func deriveCF(before []string) string {
	cf := ""
	for _, w := range before {
		if len(w) == 2 {
			cf = w
		}
	}
	return cf
}

func wIsZero(before []string, bd string, M map[byte]byte) {
	for _, w := range before {
		// w is 0
		if len(w) == 6 && in(w, bd[0]) != in(w, bd[1]) {
			if in(w, bd[0]) {
				M[bd[0]] = 'b'
				M[bd[1]] = 'd'
			} else {
				M[bd[0]] = 'd'
				M[bd[1]] = 'b'
			}
		}
	}
}

func deriveBD(before []string, cf string) string {
	bd := ""
	for _, w := range before {
		if len(w) == 4 {
			for i := 0; i < len(w); i++ {
				if !in(cf, w[i]) {
					bd += string(w[i])
				}
			}
		}
	}
	return bd
}

func wIsSeven(before []string, cf string, M map[byte]byte) {
	for _, w := range before {
		if len(w) == 3 {
			for i := 0; i < len(w); i++ {
				if !in(cf, w[i]) {
					M[w[i]] = 'a'
				}
			}
		}
	}
}

func wIsSix(before []string, cf string, M map[byte]byte) {
	for _, w := range before {
		if len(w) == 6 && in(w, cf[0]) != in(w, cf[1]) {
			if in(w, cf[0]) {
				M[cf[0]] = 'f'
				M[cf[1]] = 'c'
			} else {
				M[cf[0]] = 'c'
				M[cf[1]] = 'f'
			}
		}
	}
}

func in(w string, b byte) bool {
	for i := 0; i < len(w); i++ {
		if w[i] == b {
			return true
		}
	}
	return false
}
