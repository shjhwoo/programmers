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
		assert.Equal(t, test.expect, solution(test.weights))
	}
}

func solution(weights []int) int64 {

	var answer int64

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] < weights[j]
	})

	for i := 0; i < len(weights)-1; i++ {
		cp1 := weights[i]
		cp2 := weights[i] * 2
		cp3 := weights[i] * 3
		cp4 := weights[i] * 4

		//바이너리서치.
		if binarySearch(cp1, weights[i+1:]) {
			fmt.Println("페어 찾음", weights[i], cp1)
			answer++
		}

		if binarySearch(cp2, weights[i+1:]) {
			fmt.Println("페어 찾음", weights[i], cp2)
			answer++
		}

		if binarySearch(cp3, weights[i+1:]) {
			fmt.Println("페어 찾음", weights[i], cp3)
			answer++
		}

		if binarySearch(cp4, weights[i+1:]) {
			fmt.Println("페어 찾음", weights[i], cp4)
			answer++
		}
	}
	return answer
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func binarySearch(pair int, weights []int) bool {
	low := 0
	high := len(weights) - 1

	tar1 := pair
	tar2 := pair * 2
	tar3 := pair * 3
	tar4 := pair * 4

	for low <= high {
		median := (low + high) / 2

		if weights[median] < tar1 {
			low = median + 1
		} 
		// else if weights[median] < tar2 {
		// 	low = median + 1
		// } else if weights[median] < tar3 {
		// 	low = median + 1
		// } else if weights[median] < tar4 {
		// 	low = median + 1
		// } 
		else {
			high = median - 1
		}
	}

	if low == len(weights) { //다 뒤져도 못 찾음.
		return false
	}

	/*최소공배수를 구했고, 두 수가 모두 최소공배수가 되기 위해 2,3,4 중 하나의 숫자로 곱해서 만들 수 있을 경우에는*/
	fmt.Println("찾음!")
	return true
}
