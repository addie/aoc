package aoc

import (
	"fmt"
	"strconv"
)

type coord struct{ r, c int }

var (
	Up    = coord{r: -1}
	Down  = coord{1, 0}
	Left  = coord{0, -1}
	Right = coord{0, 1}
)

func Counter(input []int) map[int]int {
	c := make(map[int]int)
	for _, in := range input {
		c[in]++
	}
	return c
}

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[rune]*TrieNode)}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{Root: NewTrieNode()}
}

// Insert inserts a word into the trie
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

// InsertReversed inserts a word into the trie in reverse
func (t *Trie) InsertReversed(word string) {
	node := t.Root
	for i := len(word) - 1; i >= 0; i-- {
		char := rune(word[i])
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

// Search returns true if the word is in the trie
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, char := range word {
		if _, ok := node.Children[char]; !ok {
			return false
		}
		node = node.Children[char]
	}
	return node.IsEnd
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, char := range prefix {
		if _, ok := node.Children[char]; !ok {
			return false
		}
		node = node.Children[char]
	}
	return true
}

func Reversed(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func toInt(s string) int {
	return Must(strconv.Atoi(s))
}

type Integer interface {
	SignedInteger | UnsignedInteger
}
type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}
type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func toStr[T Integer](i T) string {
	return strconv.Itoa(int(i))
}

// zip groups elements in two slices by index number
// into a list of tuples so s1 = [1, 2, 3] and
// s2 = [4, 5, 6] returns [(1, 4), (2, 5), (3, 6)]
func zip[T any](s1, s2 []T) [][2]T {
	zippedLength := min(len(s1), len(s2))
	zipped := make([][2]T, zippedLength)
	for i := 0; i < zippedLength; i++ {
		zipped[i] = [2]T{s1[i], s2[i]}
	}
	return zipped
}

// alt extracts every other value from a slice into a new slice
func alt[T any](orig []T) []T {
	var evenVals []T
	for i, value := range orig {
		if i%2 == 0 {
			evenVals = append(evenVals, value)
		}
	}
	return evenVals
}

func mapKeys[K comparable, V any](m map[K]V) []K {
	var s []K
	for k := range m {
		s = append(s, k)
	}
	return s
}

func in[T comparable](val T, container []T) bool {
	for _, v := range container {
		if v == val {
			return true
		}
	}
	return false
}

func pop[T any](alist []T) T {
	f := len(alist)
	rv := (alist)[f-1]
	alist = (alist)[:f-1]
	return rv
}

func popleft[T any](alist *[]T) T {
	lv := (*alist)[0]
	*alist = (*alist)[1:]
	return lv
}

func printStringGrid[T string | byte](grid [][]T) {
	for r := range grid {
		for c := range grid[r] {
			fmt.Printf("%s ", string(grid[r][c]))
		}
		fmt.Println()
	}
}

func printIntGrid[T Integer](grid [][]T) {
	for r := range grid {
		for c := range grid[r] {
			fmt.Printf("%d ", grid[r][c])
		}
		fmt.Println()
	}
}

func inBounds[T any](grid [][]T, n coord) bool {
	return n.r >= 0 && n.r < len(grid) && n.c >= 0 && n.c < len(grid[0])
}
