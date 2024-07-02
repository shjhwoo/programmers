package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	numbers []int
	target  int
	expect  []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{2, 3, 4},
			target:  6,
			expect:  []int{0, 2},
		},
		{
			numbers: []int{9, 1, 5, 3, 6, 2},
			target:  11,
			expect:  []int{2, 4},
		},
	}

	for _, test := range tests {
		ans := solution(test.numbers, test.target)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(numbers []int, target int) []int {
	m := make(map[int]int)

	//자료의 검색과 저장을 동시에 진행한다. (이중 for문을 사용하지 않기 위해서는 한쪽을 해시테이블로 만들어놓는다 == 맵 자료구조.
	//따로 맵을 먼저 다 만들수도 있겠지만, 코드 간소화를 위해서는 이 방법이 맞다.)
	//직접 배열을 순회해보면서 아래 코드 구현 가능..

	//포인트는 나머지 짝이 되는 수를 뺄셈으로 찾아 낼 수 있었다는 것.. => 시간 단축..
	for i, num := range numbers {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}
