package main_test

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	number string
	k      int //제거할 수!
	expect string
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			number: "1924",
			k:      2,
			expect: "94",
		},
		{
			number: "1231234",
			k:      3,
			expect: "3234",
		},
		{
			number: "4177252841",
			k:      4,
			expect: "775841",
		},
	}

	for _, test := range tests {
		ans := solution(test.number, test.k)
		assert.DeepEqual(t, test.expect, ans)
	}
}

// 최대한 앞에 있는 것 => 그 다음으로 제일 작은 수부터 빼낸다.
func solution(number string, k int) string { //결과값에는 뺴야 하는 숫자의 인덱스 위치를 저장해둔다

	var stack []string

	var rmCnt int
	var lastIdx int
	for idx, numCh := range strings.Split(number, "") {
		if idx == 0 {
			stack = append(stack, numCh)
		} else {

			top := stack[len(stack)-1]

			if top >= numCh {
				stack = append(stack, numCh)
			} else {
				for top < numCh {
					stack = stack[:len(stack)-1]
					rmCnt++
					if rmCnt == k {
						lastIdx = idx
						break
					}

					if len(stack) > 0 {
						top = stack[len(stack)-1]
					}

					if len(stack) == 0 {
						break
					}
				}

				stack = append(stack, numCh)

			}

			if rmCnt == k {
				break
			}

		}
	}

	return strings.Join(stack, "") + number[lastIdx+1:]
}
