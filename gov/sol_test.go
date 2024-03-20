package main_test

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	k      []int
	expect []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			k:      []int{1, 2, 3, 4, 5},
			expect: []int{1},
		},
		{
			k:      []int{1, 3, 2, 4, 2},
			expect: []int{1, 2, 3},
		},
		{
			k:      []int{1, 1, 3, 4, 5, 2},
			expect: []int{1},
		},
	}

	for _, test := range tests {
		ans := solution(test.k)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(answers []int) []int {
	var scoreMap = map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}

	for i := 1; i <= 3; i++ {
		for index, answer := range answers {
			supojaAns := getAnswerOfSupoja(index, i)
			if answer == supojaAns {
				scoreMap[i]++
			}
		}
	}

	var maxScore int
	for _, score := range scoreMap {
		if score >= maxScore {
			maxScore = score
		}
	}

	var answer []int
	for pno, score := range scoreMap {
		if score == maxScore {
			answer = append(answer, pno)
		}
	}

	sort.Slice(answer, func(i, j int) bool {
		return answer[i] < answer[j]
	})

	return answer
}

func getAnswerOfSupoja(index int, supojaNum int) int {
	switch supojaNum {
	case 1:
		//인덱스에 1을 더하고, 5로 나눈 나머지 구하기. 0이면 5를 준다
		rem := (index + 1) % 5
		if rem == 0 {
			return 5
		} else {
			return rem
		}
	case 2:
		if index%2 == 0 {
			return 2
		} else {
			var tag int
			if index < 8 {
				tag = index
			} else {
				tag = index % 8
			}

			switch tag {
			case 1:
				return 1
			case 3:
				return 3
			case 5:
				return 4
			case 7:
				return 5
			}
		}
	case 3:
		//인덱스가 10보다 작다면
		//홀수인덱스는 자기 바로 앞의 인덱스 숫자와 같아.
		var tag int
		if index < 10 {
			if index%2 == 1 {
				tag = index - 1
			} else {
				tag = index
			}

		} else {
			tag = index % 10
			if tag%2 == 1 {
				tag = tag - 1
			}
		}
		switch tag {
		case 0:
			return 3
		case 2:
			return 1
		case 4:
			return 2
		case 6:
			return 4
		case 8:
			return 5
		}
	}

	return 0
}
