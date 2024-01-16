package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  [2]int
	expect int64
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			input:  [2]int{3, 5},
			expect: 12,
		},
		{
			input:  [2]int{3, 3},
			expect: 3,
		},
		{
			input:  [2]int{5, 3},
			expect: 12,
		},
	}

	for _, test := range tests {
		ans := solution(test.input[0], test.input[1])
		if !assert.Equal(t, test.expect, ans) {
			t.Log(test.input, "에서 실패~")
		}
	}
}

func solution(a int, b int) int64 {
	if a == b {
		return int64(a)
	}

	var numberOfInt int64
	if a < b {
		numberOfInt = int64(b - a + 1)
	} else {
		numberOfInt = int64(a - b + 1)
	}

	subSum := int64(a + b)
	if numberOfInt%2 == 0 {
		return subSum * (numberOfInt / 2)
	}

	midNum := int64((a + b) / 2)
	return subSum*(numberOfInt/2) + midNum

	/*
		//a,b 경계를 합해서 나열한 숫자의 개수가 짝수일 때 합은
		-2 -1 0 1 2 3 4 5 6 7

		//홀수일 때 합은
		-2 -1 0 1    2    3 4 5 6

	*/
}
