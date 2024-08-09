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

// 포인트는 인덱스를 같이 저장하는 것..!!
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
	var answer int = 0 //실행순서임.

	var priLo [][]int
	for idx, p := range priorities {
		priLo = append(priLo, []int{p, idx})
	}

	for len(priLo) > 0 {
		if highest(priLo[0][0], priLo[1:]) {
			outlo := priLo[0][1]
			priLo = priLo[1:]
			answer++
			if outlo == location {
				break
			}
		} else {
			priLo = append(priLo[1:], priLo[0])
		}
	}

	return answer
}

func highest(process int, priLo [][]int) bool {
	for _, item := range priLo {
		if process < item[0] {
			return false
		}
	}

	return true
}
