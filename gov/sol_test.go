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
		{
			numbers: []int{80, 20, 4, 5, 79, 60, 10},
			expect:  []int{-1, 79, 5, 79, -1, -1, -1},
		},
	}

	for _, test := range tests {
		ans := solution(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(numbers []int) []int {
	var answer []int

	for i := 0; i < len(numbers); i++ {
		answer = append(answer, -1)
	}

	var stack []int

	for i := len(numbers) - 1; i >= 0; i-- {

		for len(stack) > 0 && numbers[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && numbers[i] < stack[len(stack)-1] {
			answer[i] = stack[len(stack)-1]
		}

		stack = append(stack, numbers[i])
	}

	return answer
}

//https://velog.io/@rvbear/%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%A8%B8%EC%8A%A4-%EB%92%A4%EC%97%90-%EC%9E%88%EB%8A%94-%ED%81%B0-%EC%88%98-%EC%B0%BE%EA%B8%B0Java
/*

이중 for문이 맞지만 배열 전체 길이를 이중으로 순회하지 않는다. 물론, 두 수 중간에, 작은 수가 너무 많다면 시간적 제한이 있을 수 있다. 분명.
스택을 사용 => 더 큰 수가 나오면 밑에 수가 작을때 pop 시켜버린다 => 자신보다 작은 수를 없애버리는거니깐 가까운 뒤 큰 수를 빨리 찾을 수 있다.

즉 이중포문을 도는 시간을 더 줄이는 방법임

-1로 초기화는 => 초기값 설정으로 뒤 큰수 없는 경우 더 빨리 처리.

*/
