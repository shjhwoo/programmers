package main_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	words  []string
	expect int
}

// 포인트는 인덱스를 같이 저장하는 것..!!
func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			words:  []string{"go", "gone", "guild"},
			expect: 7,
		},
		{
			words:  []string{"abc", "def", "ghi", "jklm"},
			expect: 4,
		},
		{
			words:  []string{"word", "war", "warrior", "world"},
			expect: 15,
		},
	}

	for _, test := range tests {
		ans := solution(test.words)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(words []string) int {
	trie := TrieNode{}
	for _, word := range words {
		trie.Insert(word)
	}

	b, _ := json.MarshalIndent(trie, "", "	")
	fmt.Println(string(b), "만든 trie 확인하기")

	var answer int
	for _, word := range words {
		charSlice := strings.Split(word, "")
		curNode := trie
		var cnt int
		for _, char := range charSlice {
			if nextNode, ok := curNode.NextMap[char]; ok {
				cnt++
				curNode = *nextNode
				if curNode.Count == 1 {
					break
				}
			}
		}

		answer += cnt
	}

	return answer
}

type TrieNode struct {
	Value   string
	Count   int
	NextMap map[string]*TrieNode
}

func (tn *TrieNode) Insert(word string) {
	currentNode := tn
	charList := strings.Split(word, "")

	for _, char := range charList {

		if currentNode.NextMap == nil {
			currentNode.NextMap = make(map[string]*TrieNode)
		}

		foundNode, ok := currentNode.NextMap[char]
		if !ok {
			//이때만 넣어준다..
			currentNode.NextMap[char] = &TrieNode{
				Value:   currentNode.Value + char,
				Count:   1,
				NextMap: make(map[string]*TrieNode),
			}
		} else {
			foundNode.Count++
		}

		currentNode = currentNode.NextMap[char]
	}
}

func (tn *TrieNode) Has(word string) bool {
	currentNode := tn
	charList := strings.Split(word, "")

	for _, char := range charList {
		if _, ok := currentNode.NextMap[char]; !ok {
			return false
		}

		currentNode = currentNode.NextMap[char]
	}

	return true
}
