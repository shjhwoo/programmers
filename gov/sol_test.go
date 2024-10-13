package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	gems   []string
	expect []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			gems:   []string{"DIA", "RUBY", "RUBY", "DIA", "DIA", "EMERALD", "SAPPHIRE", "DIA"},
			expect: []int{3, 7},
		},
		// {
		// 	gems:   []string{"AA", "AB", "AC", "AA", "AC"},
		// 	expect: []int{1, 3},
		// },
		// {
		// 	gems:   []string{"XYZ", "XYZ", "XYZ"},
		// 	expect: []int{1, 1},
		// },
		// {
		// 	gems:   []string{"ZZZ", "YYY", "NNNN", "YYY", "BBB"},
		// 	expect: []int{1, 5},
		// },
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, solution(test.gems))
	}
}

//진열된 모든 종류의 보석을
//적어도 1개 이상 포함하는
//가장 짧은 구간을 찾아서 구매

func solution(gems []string) []int {
	gems = append([]string{""}, gems...) //인덱스 헷갈리지 마라공..

	gemMap := make(map[string]bool)
	for _, gem := range gems {
		if gem == "" {
			continue
		}
		gemMap[gem] = true
	}

	startIdx := 1
	endIdx := startIdx + len(gemMap) - 1

	var candidates [][]int

	for startIdx <= endIdx && endIdx < len(gems) {
		shoppingList := gems[startIdx:endIdx]

		if shopListHasAllKind(shoppingList, gemMap) {
			candidates = append(candidates, []int{startIdx, endIdx})
			startIdx++
			endIdx = startIdx + len(gemMap) - 1
			//배열도 바꿔줘야 한다

		} else {
			//인덱스를 이동하고 배열도 바꿔줘야 한다
		}
	}

	return []int{}
}

func shopListHasAllKind(shoppingList []string, gemMap map[string]bool) bool {
	shopMap := make(map[string]bool)

	for _, shopGem := range shoppingList {
		shopMap[shopGem] = true
	}

	for gem := range gemMap {
		if !shopMap[gem] {
			return false
		}
	}

	return true
}
