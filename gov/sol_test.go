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
		// {
		// 	words:  []string{"go", "gone", "guild"},
		// 	expect: 7,
		// },
		// {
		// 	words:  []string{"abc", "def", "ghi", "jklm"},
		// 	expect: 4,
		// },
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
	//트라이 자료구조를 만들어야 한다.

	trie := TrieNode{}
	for _, word := range words {
		trie.Insert(word)
	}

	b, _ := json.MarshalIndent(trie, "", "	")

	fmt.Println(string(b), "만든 trie 확인하기")

	var wordMap = make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}

	var answer int
	for _, word := range words {

		fmt.Println("개수::", trie.Count(word, wordMap))

		answer += trie.Count(word, wordMap)
	}

	return answer
}

type TrieNode struct {
	Value   string
	NextMap map[string]*TrieNode
}

func (tn *TrieNode) Insert(word string) {
	currentNode := tn
	charList := strings.Split(word, "")

	for _, char := range charList {

		if currentNode.NextMap == nil {
			currentNode.NextMap = make(map[string]*TrieNode)
		}

		if _, ok := currentNode.NextMap[char]; !ok {
			//이때만 넣어준다..
			currentNode.NextMap[char] = &TrieNode{
				Value:   currentNode.Value + char,
				NextMap: make(map[string]*TrieNode),
			}
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

func (tn *TrieNode) Count(word string, wordMap map[string]bool) int {
	currentNode := tn
	charList := strings.Split(word, "")

	var cnt int

	for _, char := range charList {
		if _, ok := currentNode.NextMap[char]; !ok {
			break
		}

		cnt++
		currentNode = currentNode.NextMap[char]
		if len(currentNode.NextMap) == 1 {
			if wordMap[currentNode.Value] && currentNode.Value != word {
				cnt++
				break
			}

			if wordMap[currentNode.Value] {
				break
			}

			//카운트 문제..!! -- 마지막 케이스.
		}
	}

	return cnt
}
