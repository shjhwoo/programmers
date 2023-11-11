package main_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	weights []int
	expect  int64
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			weights: []int{100, 180, 360, 100, 270},
			expect:  int64(4),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, Solution(test.weights))
	}
}

func Solution(weights []int) int64 {

	var answer int64

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] < weights[j]
	})

	//순회한다
	//현재 보고 있는 숫자 * 2, 숫자 * 3, 숫자 * 4  연산한다 ---- a,b,c
	//a,b,c 숫자를 바이너리 서치로 찾는다
	//찾을때마다 앤서에다가 1을 증가시켜준다

	//그 숫자 이하의 범위 내에서만 for 문을 돈다.
	for i := 0; i < len(weights); i++ {
		target1 := 1 * weights[i]

		fmt.Println(weights[i], "========================")

		//찾을 때 이렇게만 찾으면 안되고 어디서 찾니 저 부분에 있는 숫자들도 2,3,4배수해서 찾아봐야해,,

		//이진탐색
		if binarySearch(target1, weights[i+1:]) {
			answer++
		}

	}
	return answer
}

func binarySearch(pair int, weights []int) bool {
	low := 0
	high := len(weights) - 1

	for low <= high {
		median := (low + high) / 2

		if weights[median] < pair {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	fmt.Println(pair, "@@", weights, low)

	if low == len(weights) || /*최소공배수를 구했고, 두 수가 모두 최소공배수가 되기 위해 어느 한쪽이 4보다 큰 수를 곱해야 한다면 탈락*/ {
		return false
	}

	/*최소공배수를 구했고, 두 수가 모두 최소공배수가 되기 위해 2,3,4 중 하나의 숫자로 곱해서 만들 수 있을 경우에는*/
	fmt.Println("찾음!")
	return true
}

// func LCM(a, b int) int {
// 	return a * b / GCD(a, b)
// }

// func GCD(a, b int) int {
// 	for b != 0 {
// 		t := b
// 		b = a % b
// 		a = t
// 	}
// 	return a
// }
