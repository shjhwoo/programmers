package main_test

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	n      int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			n:      4,
			expect: 2,
		},
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.n))
	}
}

func solution(n int) int {
	queens := []int{}
	for i := 0; i < n; i++ {
		queens = append(queens, 0)
	}

	return findQueenPos(queens, 0)

}

func findQueenPos(queens []int, row int) int {

	count := 0

	if len(queens) == row {
		return 1
	}

	for i := 0; i < len(queens); i++ {
		queens[row] = i

		if canTakeQueenPos(queens, row) {
			count += findQueenPos(queens, row+1)
		}
	}

	return count
}

func canTakeQueenPos(queens []int, row int) bool {
	for i := 0; i < row; i++ {
		if queens[i] == queens[row] || math.Abs(float64(queens[i]-queens[row])) == math.Abs(float64(i-row)) {
			return false
		}
	}

	return true
}
