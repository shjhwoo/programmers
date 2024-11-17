package main_test

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	strs   []string
	t      string
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	strs:   []string{"ba", "na", "n", "a"},
		// 	t:      "banana",
		// 	expect: 3,
		// },
		{
			strs:   []string{"app", "ap", "p", "l", "e", "ple", "pp"},
			t:      "apple",
			expect: 2,
		},
		// {
		// 	strs:   []string{"ba", "an", "nan", "ban", "n"},
		// 	t:      "banana",
		// 	expect: -1,
		// },
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.strs, test.t))
	}
}

var cache = make(map[string]int)

func solution(strs []string, t string) int {

	var strsMap = make(map[string]bool)
	for _, str := range strs {
		strsMap[str] = true
	}

	for i := 0; i < len(t); i++ {

		start := 0
		end := i + 1
		subT := t[start:end]

		//subT에 대해서 나누어 쪼개어 생각
		for j := 1; j < end+1; j++ {
			left := subT[0:j]
			right := subT[j:]

			fmt.Println("left: ", left, "right: ", right)
			//단 안 되었던 조합은 continue 해야 한다
			if !strsMap[left] {
				continue
			} else {
				//있는경우: 되었던 조합을 캐시에서 찾기
				cache[subT] = cache[left] + cache[right]
			}
		}
	}

	return 0
}
