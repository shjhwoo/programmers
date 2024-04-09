package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	ingredient []int
	expect     int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			ingredient: []int{2, 1, 1, 2, 3, 1, 2, 3, 1},
			expect:     2,
		},
		// {
		// 	ingredient: []int{1, 3, 2, 1, 2, 1, 3, 1, 2},
		// 	expect:     0,
		// },
		// {
		// 	ingredient: []int{1, 2, 3, 1, 1, 2, 3, 1, 1, 2, 3, 1},
		// 	expect:     3,
		// },
		// {
		// 	ingredient: []int{1, 2, 1, 2, 1, 2, 1, 2, 3, 1, 3, 1, 3, 1, 3, 1},
		// 	expect:     4,
		// },
		// {
		// 	ingredient: []int{1, 2, 1, 2, 3, 1, 3, 1, 2, 3, 1, 1},
		// 	expect:     2,
		// },
		//들어온 순서대로 1231이 완성되면 햄버거가 완성되어야하는데 전체에서 따지면 순서대로가아닌 초기상태에서 만든후 빠집니다.
	}

	for _, test := range tests {
		ans := solution(test.ingredient)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(ingredient []int) int {
	var answer int

	var index = 3

	for index < len(ingredient) {
		if ingredient[index] == 2 {
			index = index + 2
			continue
		}

		if ingredient[index] == 3 {
			index++
			continue
		}

		if ingredient[index] == 1 {
			if ingredient[index-1] < 3 {
				index = index + 3
				continue
			}

			if isAbleToWrap(ingredient[index-3 : index]) {
				answer++
				ingredient = append(ingredient[:index-3], ingredient[index+1:]...)
				index = index - 3
			} else {
				index++
			}
		}

	}

	return answer
}

func isAbleToWrap(sl []int) bool {
	return sl[0] == 1 && sl[1] == 2 && sl[2] == 3
}
