package main_test

import (
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	stringSlice []string
	num         int
	expect      []string
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			stringSlice: []string{"sun", "bed", "car"},
			num:         1,
			expect:      []string{"car", "bed", "sun"},
		},
		{
			stringSlice: []string{"abce", "abcd", "cdx"},
			num:         2,
			expect:      []string{"abcd", "abce", "cdx"},
		},
	}

	for _, test := range tests {
		ans := solution(test.stringSlice, test.num)
		t.Log(ans, "계산값")
		assert.True(t, slices.Equal(test.expect, ans))
	}
}

var dict = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
}

func solution(strings []string, n int) []string {

	dictionary := strings
	sort.Strings(dictionary)

	dictMap := make(map[string]int)
	for idx, word := range dictionary {
		dictMap[word] = idx
	}

	sort.Slice(strings, func(i, j int) bool {
		leftword := strings[i]
		rightword := strings[j]

		leftChrIdx := dict[leftword[n:n+1]]
		rightChrIdx := dict[rightword[n:n+1]]

		if leftChrIdx != rightChrIdx {
			return leftChrIdx < rightChrIdx
		}

		//사전순으로 앞선 단어 먼저.
		leftDictIdx := dictMap[leftword]
		rightDictIdx := dictMap[rightword]
		return leftDictIdx < rightDictIdx
	})

	return strings
}

// 문자끼리 대소비교 가능..
// func solution(strings []string, n int) []string {

// 	sort.Slice(strings, func(i, j int) bool {
// 		strA := []byte(strings[i])
// 		strB := []byte(strings[j])
// 		if strA[n] == strB[n] {
// 			return strings[i] < strings[j]
// 		}
// 		return strA[n] < strB[n]
// 	})

// 	return strings
// }
