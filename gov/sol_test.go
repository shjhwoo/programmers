package main_test

import (
	"fmt"
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
		// {
		// 	words:  []string{"abc", "def", "ghi", "jklm"},
		// 	expect: 4,
		// },
		// {
		// 	words:  []string{"word", "war", "warrior", "world"},
		// 	expect: 15,
		// },
	}

	for _, test := range tests {
		ans := solution(test.words)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

type TrieNode struct {
	CurrentChar string
	PointerMap  map[string]*TrieNode
}

func (tn *TrieNode) Add(word string) {
	found, node := tn.Find(word)
	if !found {

		if node.PointerMap == nil {
			node.PointerMap = make(map[string]*TrieNode)
		}
		//이때만 추가.
		node.PointerMap[word[:len(word)-1]] = &TrieNode{
			CurrentChar: word[len(word)-1:],
		}
	}
}

// 분기점을 찾아주는..(찾은 경우 찾은 노드, 못 찾은 경우 부모 노드라도 주기?)
func (tn *TrieNode) Find(word string) (bool, *TrieNode) {
	var found bool
	var foundNode *TrieNode

	if tn.CurrentChar == word {
		found = true
		foundNode = tn
		return found, foundNode
	} else {
		foundNode = tn //못 찾은 경우 부모.
	}

	subword := word[:len(word)-1]

	for pointer, nextNode := range tn.PointerMap {
		if subword == nextNode.CurrentChar {
			//해당 노드로 내려가서 계속 찾아줌..
			newSub := subword + pointer
			found, foundNode = nextNode.Find(newSub)
			break
		}
	}

	return found, foundNode
}

func solution(words []string) int {
	//트라이 자료구조를 만들어야 한다.
	trie := TrieNode{}
	for _, word := range words {
		trie.Add(word)
	}

	fmt.Println(trie, "trie 구조 만든거")

	return 0
}
