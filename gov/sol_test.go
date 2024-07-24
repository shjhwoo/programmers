package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	x      int
	y      int
	n      int
	expect int
}

/*
문제 설명
자연수 x를 y로 변환하려고 합니다. 사용할 수 있는 연산은 다음과 같습니다.

x에 n을 더합니다
x에 2를 곱합니다.
x에 3을 곱합니다.
자연수 x, y, n이 매개변수로 주어질 때, x를 y로 변환하기 위해 필요한 최소 연산 횟수를 return하도록 solution 함수를 완성해주세요. 이때 x를 y로 만들 수 없다면 -1을 return 해주세요.
*/
func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	x:      10,
		// 	y:      40,
		// 	n:      5,
		// 	expect: 2,
		// },
		// {
		// 	x:      10,
		// 	y:      40,
		// 	n:      30,
		// 	expect: 1,
		// },
		// {
		// 	x:      2,
		// 	y:      5,
		// 	n:      4,
		// 	expect: -1,
		// },
		// {
		// 	x:      2,
		// 	y:      42,
		// 	n:      5,
		// 	expect: 3, //(2+5) * 3 * 2
		// },
		{
			x:      2,
			y:      21,
			n:      5,
			expect: 2, //(2+5) * 3 ..바로 앞의 케이스를 생각했을 때 재귀적으로 구현할 수 있다.
		},
	}

	for _, test := range tests {
		ans := solution(test.x, test.y, test.n)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(x int, y int, n int) int {

	if (y-x)%n == 0 {
		//x에 n만 여러 번 더해서 구할 수 있는 경우인듯.
		//덧셈 횟수를 1씩 줄이고 곱셈 수를 늘려나가며 원하는 연산횟수를 찾는다.
		var calcTime int 
		maxAddTimes := (y - x) / n
		calcTime = maxAddTimes //일단은 덧셈 횟수로 초기화

		for i := maxAddTimes -1; i >= 0; i-- {
			if (y - n * i) == x {
				cTime := 
			}
		}
	}

	return answer
}

func calcMulTimeBy2or3orAddN(num, x, n int) int {
	var calcTime int
	for num%2 == 0 {
		num = num / 2
		calcTime++
	}

	for num%3 == 0 {
		num = num / 3
		calcTime++
	}

	if num == x+n {
		calcTime++
	}

	return calcTime
}
