package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	n      int
	left   int64
	right  int64
	expect []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			n:      3,
			left:   2,
			right:  5,
			expect: []int{3, 2, 2, 3},
		},
		{
			n:      4,
			left:   7,
			right:  14,
			expect: []int{4, 3, 3, 3, 4, 4, 4, 4},
		},
	}

	for _, test := range tests {
		ans := solution(test.n, test.left, test.right)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(n int, left int64, right int64) []int {

	/*
		주기: n
		left에 들어가는 숫자를 알아낸다
			=> (left + 1)를 n으로 일단 나눠봄.
			나머지가 0이면 몫 그 자체가 몇번째 행인지를 알랴준다.
			나머지가 0이 아니면 몫 + 1 값이 몇번째 행인지를 알 수 있다.

			나머지가 0이라는 것은 맨 끝 열임을 뜻한다
			나머지가 0보다 큰 숫자면은 몇번째 열인지를 알 수 있다.

			( 7 + 1 ) 나누기 4 = 몫 2 나머지 0 : 2행 4열. => 4가 들어간다.
			( 2 + 1 ) 나누기 3 = 몫 1 나머지 0 : 1행 3열. => 3이 들어간다.

			(6 + 1) 나누기 4 = 몫 1 나머지 3 : 2행 3열. => 3이 들어간다.

		right에 들어가는 숫자를 알아낸다
		left에서 right 까지 차지하는 칸 수를 세본다 = right - left + 1 개의 숫자가 들어간다.
		=> answer의 길이가 right - left + 1 이 될때까지 for 문을 실행시키고,
		for 문 안에서 다음 작업을 진행한다
	*/

	var answer []int
	for i := left; i <= right; i++ {
		row, col := findRowColFromIdx(n, i)
		answer = append(answer, findNumByRowCol(row, col))
	}

	return answer
}

//행과 열 정보를 주면 그 자리에 들어가야 하는 숫자가 뭔지 알아내면 되고, 그 규칙을 만들면 된다
//left 의 행과 열 정보부터 파악 => ... => right의 행과 열 정보 파악한다.
//그 사이에 있는 값들의 행과 열 정보를 파악 => 그 사이 숫자들에 대한

func findRowColFromIdx(n int, left int64) (int, int) {
	var rowNo int
	var columnNo int

	remainder := (left + 1) % int64(n)

	if remainder == 0 {
		columnNo = n
		rowNo = int((left + 1) / int64(n))
	} else {
		columnNo = int(remainder)
		rowNo = int((left+1)/int64(n)) + 1
	}

	return rowNo, columnNo
}

func findNumByRowCol(row int, col int) int {
	if row == col {
		return row
	}

	if row > col {
		return row
	}

	if row < col {
		return col
	}

	return 0
}
