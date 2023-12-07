package aoc

import (
	"strconv"
)

type coord struct{ r, c int }

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

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func toStr[T Number](i T) string {
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
