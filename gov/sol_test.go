package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	lottos   [6]int
	win_nums [6]int
	expect   []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			lottos:   [6]int{44, 1, 0, 0, 31, 25},
			win_nums: [6]int{31, 10, 45, 1, 6, 19},
			expect:   []int{3, 5},
		},
		{
			lottos:   [6]int{0, 0, 0, 0, 0, 0},
			win_nums: [6]int{38, 19, 20, 40, 15, 25},
			expect:   []int{1, 6},
		},
		{
			lottos:   [6]int{45, 4, 35, 20, 3, 9},
			win_nums: [6]int{20, 9, 3, 45, 4, 35},
			expect:   []int{1, 1},
		},
	}

	for _, test := range tests {
		ans := solution(test.lottos, test.win_nums)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(lottos [6]int, win_nums [6]int) []int {
	win_numMap := make(map[int]bool)
	for _, num := range win_nums {
		win_numMap[num] = true
	}

	var answer []int

	var baseHitCount int
	var matchedNumberMap = make(map[int]bool)
	var numverOfZero int
	for _, num := range lottos {
		if _, exist := win_numMap[num]; exist {
			baseHitCount++
			matchedNumberMap[num] = true
		}

		if num == 0 {
			numverOfZero++
		}
	}

	var unmatchedNumberMap = make(map[int]bool)
	for _, wnum := range win_nums {
		if _, exist := matchedNumberMap[wnum]; !exist {
			unmatchedNumberMap[wnum] = true
		}
	}

	//최저 순위: 로또구매자가 선택한 번호 중 0을 제외하고 당첨번호와 일치하는 개수를 센다.
	var lowestRank int
	if baseHitCount < 2 {
		lowestRank = 6
	} else {
		lowestRank = 7 - baseHitCount
	}

	var highestRank int
	//최고 순위) 0으로 되어있는 부분.
	if numverOfZero == 0 || len(unmatchedNumberMap) == 0 {
		highestRank = lowestRank
	} else {
		hitCount := (baseHitCount + numverOfZero)
		if hitCount < 2 {
			highestRank = 6
		} else {
			highestRank = 7 - hitCount
		}
	}

	answer = append(answer, highestRank)
	answer = append(answer, lowestRank)

	return answer
}
