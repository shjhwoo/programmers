package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	order  []int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			order:  []int{4, 3, 1, 2, 5},
			expect: 2,
		},
		{
			order:  []int{5, 4, 3, 2, 1},
			expect: 5,
		},
		{
			order:  []int{4, 3, 2, 5, 6, 1},
			expect: 6,
		},
		{
			order:  []int{4, 3, 2, 5, 1, 6},
			expect: 6,
		},
		{
			order:  []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 11, 12},
			expect: 12,
		},
	}

	for _, test := range tests {
		ans := solution(test.order)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

/*
택배 기사님이 미리 알려준 순서에 맞게 영재가 택배상자를 실어야 합니다.

바로 실을 수 없는 상자는 스택의 성격을 지닌 보조 컨테이너에다가 실어야 한다

보조 컨테이너 벨트를 이용해도 기사님이 원하는 순서대로 상자를 싣지 못 한다면,
더 이상 상자를 싣지 않습니다.

택배상자

컨베이어벨트

빠지는곳 ------------------ 들어오는곳
[1] [2] [3] [4] [5]   <-----------------------
--------------------------------------

택배기사님이 알려준 순서:
(먼저) [4] [3] [1] [2] [5] (나중)

보조 컨테이너(스택):
입구----------------------
[3] [2] [1]               |
--------------------------

order : 택배 기사님이 원하는 상자 순서

영재가 몇 개의 상자를 실을 수 있는지..


택배기사님이 원하는 순서가
5 4 3 2 1 인 경우에는

순서대로 다 실을 수 없기 때문에 일단 보조 컨테이너에 무조건 다 쑤셔넣고

다시 빼서 5 43 21 로 실으면 된다.

근데 우리가 원하는 건 일단 실을 수 있는 상자의 개수다

*/

func solution(order []int) int {

	var answer int

	var subConveyerStack []int

	var conveyerTop = 1

	for _, num := range order {

		if num == conveyerTop {
			answer++
			conveyerTop++
		} else {
			//다른 경우. 일단 보조 컨베이어에 넣어둔다.
			//넣기 전에 보조컨베이어에 첫번째 상자를 보고 일치하는게 있는지 확인한다.
			if len(subConveyerStack) != 0 {
				if num == subConveyerStack[0] {
					answer++
					subConveyerStack = subConveyerStack[1:]
					continue
				}

				if num < subConveyerStack[0] {
					break
				}
			}

			for i := num - 1; i > 0; i-- {
				subConveyerStack = append(subConveyerStack, i)
				conveyerTop++
			}
			answer++
		}
	}

	return answer
}
