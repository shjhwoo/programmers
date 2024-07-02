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
	}

	for _, test := range tests {
		ans := solution(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

/*
                  i
   9, 1, 5, 3, 6, 2
   -1 5  6  6 -1 -1 j
*/

func solution(numbers []int) []int {
	var answer []int

	var dictionary = make(map[int]int) //특정 인덱스 위치 - 그 인덱스 위치에 해당하는 수의 뒷 큰수의 인덱스
	//일단 순회하는 건 필요하다.모든 숫자에 대해서 뒷 큰 수를 알아야 하니깐.!!
	for i := 0; i < len(numbers); i++ {
		leftNum := numbers[i]
		var found bool

		var boundaryIdx int

		if dictionary[i-1] > 0 && i != dictionary[i-1] {
			//끝까지 돌 필요가 없다.
			boundaryIdx = dictionary[i-1] + 1
		} else {
			//끝까지 돌아야 한다.
			boundaryIdx = len(numbers)
		}

		for j := i + 1; j < boundaryIdx; j++ {
			rightNum := numbers[j]
			if leftNum < rightNum {
				dictionary[i] = j
				answer = append(answer, rightNum)
				found = true
				break
			}
		}

		if !found {
			dictionary[i] = -1
			answer = append(answer, -1)
		}
		//뒷 큰 수를 빨리 찾아야 한다.
		/*
			일단 맵이 비어 있을 때는 다 돌 수 밖에 없음..
			일단 발견하면 넣는다.

			5,3,1,1,6 같은 케이스라면 => 5의 뒷 큰수 6에 대한 인덱스를 저장.. == 이 말뜻은 결국에는 5와 6 사이에는,
			5보다 작거나 같은 숫자만 있다는 것을 의미한다.
			그럼 3의 차례에서는 굳이 전체 배열의 끝까지 다 볼 필요가 없이, 6까지만 보면 된다는 소리다.

			4,3,2,1,1 케이스의 경우:
		*/
	}

	return answer
}
