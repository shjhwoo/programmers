package main_test

import (
	"strconv"
	"strings"
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
			ingredient: []int{1, 1, 1, 2, 1, 2, 3, 1, 2, 3, 1},
			expect:     1,
		},
	}

	for _, test := range tests {
		ans := solution(test.ingredient)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(ingredient []int) int {
	var ingredientString string
	for _, v := range ingredient {
		ingredientString += strconv.Itoa(v)
	}

	return wrapHamburger(ingredientString)
}

func wrapHamburger(ingredientString string) int {
	leftingredientSlice := strings.Split(ingredientString, "1231")
	if len(leftingredientSlice) == 1 {
		return 0
	}

	var answer = (len(leftingredientSlice) - 1) + wrapHamburger(strings.Join(leftingredientSlice, ""))

	return answer
}
