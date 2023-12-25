package main_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
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
		if !assert.True(t, slices.Compare(test.expect, solution(test.numbers)) == 0) {
			t.Log(solution(test.numbers))
		}
	}
}

func solution(numbers []int) []int {
	/*
		모든 경우의 수
		원소가 1개인 경우 [-1]
		마지막 원소는 항상 -1이다
		이 원소가 1이면 무조건 바로 뒷 수가 뒷 큰수가 된다


		배열을 순회할 수밖엔는 없다

		그런데 그 원소에 대해서 뒤 나머지 숫자와 비교를 어떻게 하느냐가 문제

		n번째 원소에 대해서:

		뒤 큰 수 찾는 법
		다음 수 - 현재 수 = 양수 .. 로 나오는 게 처음이다. 그러면 이게 뒷 큰수
		그게 처음이라는거 어떻게 알아.
		카운터를 만들어줘야겠지
		카운터가 0에서 1이 되는 순간에 스탑하고. 다음 순회로 넘어가면 되지않니?


		n+1번부터 마지막까지의 원소 중 가장 큰 수를 저장한다.
		그러면, 그 큰 수
		이전 지점까지는:
			이전 지점 중, 큰수 직전까지 계속 하강하는 구간에서는 가장 큰수로 통일 가능
		이후 지점부터는:
			계속 하강하는 구간이라면 -1로 통일 가능
	*/

	if len(numbers) == 1 {
		return []int{-1}
	}

	var answer []int

	for i := 0; i < len(numbers)-1; i++ {
		found := 0

		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				found = 1
				answer = append(answer, numbers[j])
				break
			}
		}

		if found == 0 {
			answer = append(answer, -1)
		}

	}

	answer = append(answer, -1)

	return answer
}

/*
테스트 20 〉	실패 (시간 초과)
테스트 21 〉	실패 (시간 초과)
테스트 22 〉	실패 (시간 초과)
테스트 23 〉	실패 (시간 초과)
*/
