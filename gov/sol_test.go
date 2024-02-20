package main_test

import (
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	array    []int
	commands [][]int
	expect   []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			array:    []int{1, 5, 2, 6, 3, 7, 4},
			commands: [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}},
			expect:   []int{5, 6, 3},
		},
	}

	for _, test := range tests {
		ans := solution(test.array, test.commands)
		t.Log(ans, "계산값")
		assert.True(t, slices.Equal(test.expect, ans))
	}
}

func solution(array []int, commands [][]int) []int {

	//자른다
	var answer []int
	for _, command := range commands {
		start := command[0] - 1
		end := command[1]

		var sliced []int
		part := array[start:end] //새로운 곳에 복사해야함
		sliced = append(sliced, part...)

		//정렬 - 오름차순
		sort.Ints(sliced)
		answer = append(answer, sliced[command[2]-1])
	}

	return answer
}
