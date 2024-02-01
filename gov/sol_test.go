package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  []int
	expect int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			input:  []int{-2, 3, 0, 2, -5},
			expect: 2,
		},
		{
			input:  []int{-3, -2, -1, 0, 1, 2, 3},
			expect: 5,
		},
		{
			input:  []int{-1, 1, -1, 1},
			expect: 0,
		},
	}

	for _, test := range tests {
		ans := solution(test.input)
		assert.Equal(t, test.expect, ans)
	}
}

func solution(number []int) int {
	var answer int

	/*
		0번째 + 1번째 + 2번째
		0번째 + 1번째 + 3번째
		0번째 + 1번째 + ...
		0번째 + 1번째 + 마지막 원소.

		0번째 + 2번째 + 3번째
		...

	*/

	for i := 0; i < len(number); i++ {
		firstPrson := number[i]

		for j := i + 1; j < len(number); j++ {
			secondPrson := number[j]

			for k := j + 1; k < len(number); k++ {
				thirdPrson := number[k]

				if firstPrson+secondPrson+thirdPrson == 0 {
					answer++
				}
			}
		}
	}

	return answer
}
