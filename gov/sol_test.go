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
		{
			ingredient: []int{1, 3, 2, 1, 2, 1, 3, 1, 2},
			expect:     0,
		},
		{
			ingredient: []int{1, 2, 3, 1, 1, 2, 3, 1, 1, 2, 3, 1},
			expect:     3,
		},
		{
			ingredient: []int{1, 2, 1, 2, 1, 2, 1, 2, 3, 1, 3, 1, 3, 1, 3, 1},
			expect:     4,
		},
		{
			ingredient: []int{1, 2, 1, 2, 3, 1, 3, 1, 2, 3, 1, 1},
			expect:     2,
		},
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

	var idx int

	for idx < len(ingredient)-3 {
		if isAbleToWrap(ingredient[idx : idx+4]) {
			answer++
			ingredient = append(ingredient[:idx], ingredient[idx+4:]...)
			idx = 0 //처음부터 다시 본다.
		} else {
			idx++
		}
	}

	return answer
}

func isAbleToWrap(sl []int) bool {
	return sl[0] == 1 && sl[1] == 2 && sl[2] == 3 && sl[3] == 1
}

// ////왜 더 빠를까? 조건식이 단순해서? --스택 이용 (쌓아 올린다.)
// func isBugger(ingredient []int) bool {
// 	return ingredient[0] == 1 && ingredient[1] == 2 && ingredient[2] == 3 && ingredient[3] == 1
// }

// func solution(ingredient []int) int {
// 	stack := make([]int, 0, len(ingredient))
// 	answer := 0

// 	for index := range ingredient {
// 		stack = append(stack, ingredient[index]) //무조건 끝에서부터 보기 때문에 위 솔루션보다 훨씬 빠름
// 		if len(stack) >= 4 && isBugger(stack[len(stack)-4:]) {
// 			answer++
// 			stack = stack[:len(stack)-4]
// 		}
// 	}

// 	return answer
// }
