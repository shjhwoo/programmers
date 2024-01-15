package main_test

import (
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	arr     []int
	divisor int
	expect  []int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			arr:     []int{5, 9, 7, 10},
			divisor: 5,
			expect:  []int{5, 10},
		},
		{
			arr:     []int{2, 36, 1, 3},
			divisor: 1,
			expect:  []int{1, 2, 3, 36},
		},
		{
			arr:     []int{3, 2, 6},
			divisor: 10,
			expect:  []int{-1},
		},
	}

	for _, test := range tests {
		ans := solution(test.arr, test.divisor)
		if !assert.True(t, slices.Equal(test.expect, ans)) {
			t.Log(test.arr, ans)
		}
	}
}

func solution(arr []int, divisor int) []int {
	if divisor == 1 {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})

		return arr
	}

	var answer []int

	for _, num := range arr {
		if num >= divisor && num%divisor == 0 {
			answer = append(answer, num)
		}
	}

	if len(answer) == 0 {
		return []int{-1}
	}

	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})

	return answer
}
