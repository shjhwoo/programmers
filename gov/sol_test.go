package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	priorities []int
	location   int
	expect     int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	priorities: []int{2, 1, 3, 2},
		// 	location:   2,
		// 	expect:     1,
		// },
		// {
		// 	priorities: []int{1, 1, 9, 1, 1, 1},
		// 	location:   0,
		// 	expect:     5,
		// },
		{
			priorities: []int{1, 2, 3, 4, 5, 6}, //  5 1 2 3 4
			location:   4,
			expect:     2,
		},
	}

	for _, test := range tests {
		ans := solution(test.priorities, test.location)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(priorities []int, location int) int {
	var locationMap = make(map[int]int)

	for idx, _ := range priorities {
		locationMap[idx] = idx + 1
	}

	var nonZeroCnt = len(priorities)
	var cursor int

	for nonZeroCnt > 0 && cursor < len(priorities) {
		if isCusrorHighest(priorities) {
			priorities[cursor] = 0
			nonZeroCnt--
		} else {

			priorities = append(priorities[1:], priorities[0])

			locationMap[cursor+1] = 1             //맨 앞으로 가기
			locationMap[cursor] = len(priorities) //맨 뒤로 가기

			var boundary = len(priorities)

			//순서 떙기기..
			var i int = 2
			for {
				locationMap[cursor+i] = locationMap[cursor+i-1] + 1 //어디까지??

				i++
			}

		}

		cursor++
	}

	return locationMap[location]
}

func isCusrorHighest(priorities []int) bool {
	num := priorities[0]

	for _, compareNum := range priorities[1:] {
		if compareNum > num {
			return false
		}
	}

	return true
}
