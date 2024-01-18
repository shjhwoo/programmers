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

}
