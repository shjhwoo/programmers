package main_test

import (
	"fmt"
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

	/*

		일단 주어진 order 배열을 순회해야 한다.

		4 3 1 2 5 를 가지고 생각해보자.

		0번째 인덱스: 4를 넣어야 한다.
		하지만 배열은 1부터 시작한다. => 4라는 숫자가 나올때까지 스택에 밀어 넣는다,
		즉 subconveyerStack = append(reverse([]int{1,2,3}), subconveyerStack...)
		즉, 이 과정 후에 스택에서 제일 먼저 뺄 수 있는 숫자는 3이 된다!
		택배상자를 실었기 때문에 카운트 1 한다

		1번째 인덱스: 3을 넣어야 한다.
		즉 1번째 인덱스 부터는, 기존 컨베이어벨트와 보조 벨트를 모두 봐야 한다.
		보조벨트를 먼저 본다. 첫번째 원소로서 3이 있기 때문에 빼온다.
		즉 subconveyerStack = subconveyerStack[1:] 이 된다.
		즉 이 과정 후에 스택에서 제일 먼저 뺄 수 있는 숫자는 2가 된다!
		택배상자를 실었기 때문에 카운트 1 한다

		2번째 인덱스: 1을 넣어야 한다.
		숫자가 앞의 숫자 3보다 작기 때문에 분명히 얘는 보조 컨테이너에 들어가 있을거다.
		하지만 보조 컨베이어벨트는 2부터 뺄수있기 때문에, 여기서 stop해야 한다.
		break!!
	*/

	var answer int

	var firstBoxThatCanBegetFromConveyer = 1
	firstBoxThatCanBegetFromSubConveyer := 0

	fmt.Println("firstBoxThatCanBegetFromSubConveyer", firstBoxThatCanBegetFromSubConveyer)

	for idx, num := range order {
		if idx == 0 {
			if num > 1 {
				firstBoxThatCanBegetFromConveyer = num + 1
				firstBoxThatCanBegetFromSubConveyer = num - 1
			}
			answer++
		} else {

			//체크조건: 현재 컨베이어에서 뺄 수 있는 첫번째 박스의 번호가 현재 트럭에 실어야 하는 박스의 번호와 같다/다르다 여부

			//다른경우
			if firstBoxThatCanBegetFromConveyer != num {

				//스택에서 찾아본다.
				if firstBoxThatCanBegetFromSubConveyer == num {
					firstBoxThatCanBegetFromSubConveyer--
					answer++
				}

				//더 큰 경우
				if firstBoxThatCanBegetFromConveyer < num {
					//또 스택에 넣고 연산해야함
					firstBoxThatCanBegetFromSubConveyer = num - 1
					firstBoxThatCanBegetFromConveyer = num + 1

					answer++
				}
				//더 작은 경우는 있을 수 없다

			} else {
				//같으면 그대로 실으면 된다
				firstBoxThatCanBegetFromConveyer = num + 1
				answer++
			}
		}
	}

	return answer
}
