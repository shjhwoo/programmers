package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	s      string
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	s:      "banana",
		// 	expect: 3,
		// },
		// {
		// 	s:      "abracadabra",
		// 	expect: 6,
		// },
		// {
		// 	s:      "aaabbaccccabba",
		// 	expect: 3,
		// },
		{
			s:      "zzzzzz",
			expect: 1,
		},
	}

	for _, test := range tests {
		ans := solution(test.s)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(s string) int {

	var answer int

	var xCount int
	var otherCount int

	var x rune

	runes := []rune(s)
	for index, char := range runes {
		if xCount == 0 && otherCount == 0 {
			x = char
			xCount++

			if index == len(runes)-1 {
				answer++
			}
			continue
		}

		if x != char {
			otherCount++

			if otherCount == xCount {
				answer++
				xCount = 0
				otherCount = 0
			}

			if otherCount != xCount && index == len(runes)-1 {
				answer++
				break
			}
		}

		if x == char {
			xCount++

			if index == len(runes)-1 {
				answer++
				continue
			}
		}

	}

	return answer
}
