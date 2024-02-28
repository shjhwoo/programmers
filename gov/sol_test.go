package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	a      int
	b      int
	n      int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			a:      2,
			b:      1,
			n:      20,
			expect: 19,
		},
		{
			a:      3,
			b:      1,
			n:      20,
			expect: 9,
		},
		{
			a:      7,
			b:      2,
			n:      20,
			expect: 6, //(20-14+4) = 10 => (10-7+2) = 5  , 4 + 2
		},
	}

	for _, test := range tests {
		ans := solution(test.a, test.b, test.n)
		t.Log(ans, "계산값")
		assert.Equal(t, test.expect, ans)
	}
}

func solution(a int, b int, n int) int {
	//얼마만큼을 반복해야 하는지가 문제이다.

	//a개만큼의 빈병을 주면 콜라 b병을 받는다.
	//처음에 가지고 있는 빈병은 n개이다.

	// n개의 병을 a로 나눈다. => 몫 : 받을 수 있는 콜라의 수,
	// 이후 가지게 되는 콜라의 수 => 나머지 + 몫 .... 이 값을 x라고 하자.

	// 그러면 x가 a보다 작으면 이때 정답을 반환하면 된다

	//반복은 x가 a보다 크거나, 같은 동안만 할 수 있다.

	var answer = (n / a) * b

	var numberOfColaBottles = (n % a) + (n/a)*b

	for numberOfColaBottles >= a {
		answer += (numberOfColaBottles / a) * b
		numberOfColaBottles = (numberOfColaBottles % a) + (numberOfColaBottles/a)*b
	}

	return answer
}
