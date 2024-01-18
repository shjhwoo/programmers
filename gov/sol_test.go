package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input [2]int
	gcd   int
	lcm   int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			input: [2]int{3, 12},
			gcd:   3,
			lcm:   12,
		},
		{
			input: [2]int{2, 5},
			gcd:   1,
			lcm:   10,
		},
	}

	for _, test := range tests {
		ans := solution(test.input[0], test.input[1])
		assert.Equal(t, test.gcd, ans[0])
		assert.Equal(t, test.lcm, ans[1])
	}
}

func solution(n int, m int) []int {
	return []int{getGreatestCommonDividor(n, m), getLeastCommonMultiple(n, m)}
}

func getGreatestCommonDividor(n, m int) int {
	var bigNum int
	var smallNum int

	if n < m {
		bigNum = m
		smallNum = n
	} else {
		bigNum = n
		smallNum = m
	}

	if bigNum%smallNum == 0 {
		return smallNum
	}

	return getGreatestCommonDividor(smallNum, bigNum%smallNum)
}

func getLeastCommonMultiple(n, m int) int {
	gcd := getGreatestCommonDividor(n, m)
	gcd_left := n / gcd
	gcd_right := m / gcd

	return gcd * gcd_left * gcd_right
}
