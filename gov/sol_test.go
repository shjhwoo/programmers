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

	//특이 케이스 처리하기
	if x+n == y || x*2 == y || x*3 == y {
		return 1
	}

	/*
		x * (2, 3, 6의 배수) = y - (n을 더하는 횟수) * n
		위 식을 만족하는 어떤수 + n을 더하는 횟수 최소값을 찾아야함.

		그래프의 기울기를 사용.
	*/

	//초기 덧셈연산을 최대값으로 하고
	var nAddTime int
	var mulTime int

	if (y-x)%n == 0 {
		nAddTime = (y - x) / n
	} else if (y-2*x)%n == 0 {
		nAddTime = (y - 2*x) / n
		mulTime = 1
	} else if (y-3*x)%n == 0 {
		nAddTime = (y - 3*x) / n
		mulTime = 1
	}

	var answer = nAddTime + mulTime
	if answer == 0 {
		return -1
	}

	//nAddTime을 .. 1씩 줄여나가면서 절충점을 찾는다..?
	for i := nAddTime - 1; i >= 0; i-- {
		mul := (y - i*n) / x //mul은 2 또는 3의 배수인데, 2와 3을 몇번 곱해서 나오는 수인지 알아야 한다.

		if (y-i*n)%x != 0 {
			continue
		}

		mulTime = calcMulTimeBy2or3orAddN(mul, x, n) //여기도 잘못됨.
		if mulTime == 0 {
			continue
		}

		calcTime := i + mulTime
		if answer > calcTime {
			answer = calcTime
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
