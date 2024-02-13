package main_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputS string
	expect []int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			inputS: "banana",
			expect: []int{-1, -1, -1, 2, 2, 2},
		},
		{
			inputS: "foobar",
			expect: []int{-1, -1, 1, -1, -1, -1},
		},
	}

	for _, test := range tests {
		ans := solution(test.inputS)

		t.Log(ans, "계산값")

		assert.True(t, slices.Equal(test.expect, ans))
	}
}

func solution(s string) []int {
	var answer []int

	if len(s) == 1 {
		return []int{-1}
	}

	for idx, ch := range strings.Split(s, "") {
		if idx == 0 {
			answer = append(answer, -1)
		} else {
			prevChs := strings.Split(s[0:idx], "")
			idxOfClosestCh := findIdxOfClosestPrevCh(prevChs, ch)

			if idxOfClosestCh == -1 {
				answer = append(answer, -1)
				continue
			}

			a := idx - idxOfClosestCh
			answer = append(answer, a)
		}
	}

	return answer
}

func findIdxOfClosestPrevCh(prevChs []string, currentCh string) int {
	var result = -1

	for idx, ch := range prevChs {
		if ch == currentCh {
			if result < idx {
				result = idx
			}
		}
	}

	return result
}

// 맵을 써서 뭉개버린다.: 이중 for 문 회피기법
func solution2(s string) []int {
	m := make(map[string]int)
	var res []int

	for i := range s {
		val, ok := m[string(s[i])]
		if !ok {
			m[string(s[i])] = i
			res = append(res, -1)
		} else {
			m[string(s[i])] = i
			res = append(res, i-val)
		}
	}
	return res
}
