package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputs string
	expect string
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			inputs: "abcde",
			expect: "c",
		},
		{
			inputs: "qwer",
			expect: "we",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, solution(test.inputs))
	}
}

func solution(s string) string {
	//짝수인 경우  0123 => 12  01234567 => 34
	if len(s)%2 == 0 {
		start := len(s)/2 - 1
		end := len(s) / 2
		return s[start : end+1]
	}
	//홀수인 경우
	mid := len(s) / 2
	return s[mid : mid+1]
}
