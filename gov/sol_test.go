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

	leftRow, leftCol := findLeftRowCol(n, left)

	var answer []int
	answer = append(answer, leftCol)

	goalLength := int(right - left + 1)

	for len(answer) < goalLength {
		/*
			경우의 수를 생각
			leftCol 우선 기준으로 했을 때,

			leftCol이 n과 동일하다면 다음 행으로 넘어가야 한다.
				=> 다음 행 채우기: leftRow값을 보고 결정해야 한다
				leftRow + 1 로 된 숫자로 그만큼 반복해서 채운다. (길이가 n이 될때까지)
				만약에 모자라면 1씩 증가시키면서 1칸씩 채운다.

			leftCol이 n보다 작다면 leftRow가 일단 몇번째 행인지를 확인한다.


			X행 Y열인 경우, Y값이 X보다 작거나 같으면 X값을 그대로 쓰면 된다. 그 외에는 1씩 증가시키며 채운다

		*/

		if leftCol == n {

			var filler []int

			for len(filler) < n {
				for i := 0; i < leftRow+1; i++ {
					filler = append(filler, leftRow+1)
				}

				if leftRow+1 < n {
					//모자란 만큼 더 채워야 한다. 길이가 n이 될때까지(마지막에 채우는 숫자가 n이 될때까지)
					var fillNum = leftRow + 2

				}
			}

		}

		if leftCol < n {
		}

	}

	return []int{}
}

func findLeftRowCol(n int, left int64) (int, int) {
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
