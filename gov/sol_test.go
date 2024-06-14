package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	numbers []int
	expect  []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{2, 3, 3, 5},
			expect:  []int{3, 5, 5, -1},
		},
		{
			numbers: []int{9, 1, 5, 3, 6, 2},
			expect:  []int{-1, 5, 6, 6, -1, -1},
		},
	}

	for _, test := range tests {
		ans := solution(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(numbers []int) []int {
	var answer = make([]int, len(numbers))
	/*
		기본로직

		배열 순회는 해야한다
		자기 바로 뒤에 있는 숫자부터 비교를 시작한다 (최대한 빨리 탐색을 끝내는 것이 좋음)

		뒷 큰수가 자기 바로 뒤 말고 몇칸 떨어져서 있다면,

		자기 자신과 뒷 큰 수 사이에 있는 아이들은
		자기 자신보다는 분명히 작거나 같음을 알 수 있다.
		그리고 이 사이에 있는 아이들은 뒷 큰수보다도 작음.
		(이 범위에서만 비교를 하면 되는 것이다)
		그러면 배열의 끝까지 다 탐색을 하지않아도 된다
		예) 5,3,2,6
		예) 5,3,4,5,6
	*/

	for i := 0; i < len(numbers); i++ {
		leftNum := numbers[i]

		var found bool
		var findStartIdx = i + 1
		var findEndIdx = len(numbers)

		for j := findStartIdx; j < findEndIdx; j++ {
			rightNum := numbers[j]

			if leftNum < rightNum {
				answer[i] = rightNum
				found = true

				findEndIdx = j //이렇게 해야 하는데 다음 반복에서 덮어 씌워짐 ㅠㅠ

				break
			}
		}

		if !found {
			answer[i] = -1
		}
	}

	return answer
}
