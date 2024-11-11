package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	strs   []string
	t      string
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	strs:   []string{"ba", "na", "n", "a"},
		// 	t:      "banana",
		// 	expect: 3,
		// },
		{
			strs:   []string{"app", "ap", "p", "l", "e", "ple", "pp"},
			t:      "apple",
			expect: 2,
		},
		// {
		// 	strs:   []string{"ba", "an", "nan", "ban", "n"},
		// 	t:      "banana",
		// 	expect: -1,
		// },
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.strs, test.t))
	}
}

func solution(strs []string, t string) int {
	var preMap = make(map[string]bool)

	for _, pre := range strs {
		preMap[pre] = true
	}

	var memoMap []int

	for j := 0; j < len(t)+1; j++ {
		memoMap = append(memoMap, 0)
	}

	//제일 작은 문제부터 정의한다.
	for i := 1; i < len(t); i++ {
		//제일 작은 문제는 i == 1일때.
		subWord := t[0:i] //단어 ba 

		for j := i-1; j >= 0; j-- {}
		s2 := t[0:i-j] //이전의 작은 문제 b 
		if 

		if preMap[subWord] {
			//기억해둔다.
			memoMap[i] = 1 //?
		} else {
			//없는 경우. -1
			memoMap[i] = -1 //?
		}
	}

	return 0
}
