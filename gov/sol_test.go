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
		{
			priorities: []int{2, 1, 3, 2},
			location:   2,
			expect:     1,
		},
		{
			priorities: []int{1, 1, 9, 1, 1, 1},
			location:   0,
			expect:     5,
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

			locationMap[cursor+1] = locationMap[cursor]
			locationMap[cursor] = len(priorities)

			for c := 0; c < len(locationMap); c++ {
				if c > cursor+1 {
					locationMap[c] = locationMap[c-1] + 1
				}

				if c < cursor {
					locationMap[c] = locationMap[c+1] - 1
				}
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
