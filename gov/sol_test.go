package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  int
	expect int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			input:  16,
			expect: 6,
		},
		{
			input:  2554,
			expect: 16,
		},
		{
			input:  6628,
			expect: 13,
		},
		{
			input:  99999,
			expect: 2,
		},
		{
			input:  10110,
			expect: 3,
		},
		{
			input:  1273,
			expect: 10,
		},
		{
			input:  1580, //
			expect: 8,
		},
		{
			input:  90909,
			expect: 6,
		},
		{
			input:  90807,
			expect: 9,
		},
		{
			input:  9807,
			expect: 7,
		},
		{
			input:  9007,
			expect: 6,
		},
		{
			input:  4545,
			expect: 18,
		},
		{
			input:  5555,
			expect: 18,
		},
	}

	for _, test := range tests {
		if !assert.Equal(t, test.expect, Solution2(test.input)) {
			t.Log(test.input, "@@@@@@@@@@@@@@@@@@")
		}
	}
}

//오름차순으로 정렬되어있음
// * 최소길이
// * 가장 빠른 인덱스
// * 연속된 수열
// * 리턴: 시작인덱스 끝 인덱스(두 경계 모두 포함)

// 우선순위?
// 최소 길이 -> 최소 인덱스

func Solution(sequence []int, k int) []int {
	//단순하게 생각하자
	//배열 순회
	//각 인덱스에 대해서 다음 숫자를 하나씩 더해서 K값이 되면 중지하고 그때의 배열을 복사해서
	//candidate에 저장한다
	//이중 포문이 되겠구나.

	//슬라이딩 윈도우를 이용한 방법 => 고정된 크기라 안된다
	//투 포인터 패턴으로 접근하면 될거같다. https://butter-shower.tistory.com/226

}
