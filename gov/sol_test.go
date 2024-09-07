package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

//여러가지 정렬 알고리즘을 GO로 구현해보자

type TestCase struct {
	numbers []int
	expect  []int
}

func TestBubbleSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3},
			expect:  []int{1, 3, 4, 5, 7},
		},
	}

	for _, test := range tests {
		ans := BubbleSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func BubbleSort(numbers []int) []int {

	for i := 0; i < len(numbers)-1; i++ { //총 배열의 길이 - 1만큼 반복
		for j := 0; j < len(numbers)-1-i; j++ {
			left := numbers[j]
			right := numbers[j+1]
			if left > right {
				numbers[j] = right
				numbers[j+1] = left
			}
		}
	}

	return numbers
}

func TestSelectSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3},
			expect:  []int{7, 5, 4, 3, 1},
		},
	}

	for _, test := range tests {
		ans := SelectSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func SelectSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		rmi := getHighestRemainderIdx(numbers[i+1:]) + i + 1

		if numbers[i] < numbers[rmi] {
			left := numbers[i]
			right := numbers[rmi]

			numbers[i] = right
			numbers[rmi] = left
		}
	}

	return numbers
}

func getHighestRemainderIdx(rmNumbers []int) int {
	var maxNum int
	var resIdx int
	for idx, num := range rmNumbers {
		if maxNum < num {
			maxNum = num
			resIdx = idx
		}
	}

	return resIdx
}

func TestInsertSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3},
			expect:  []int{1, 3, 4, 5, 7},
		},
	}

	for _, test := range tests {
		ans := InsertSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func InsertSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ { //배열의 요소가 5개라면, 4번만큼 반복할 수 있다.
		/*
			i == 1일때는 0과 비교
			i == 2일때는 1 => 1  vs 0
			...
			3 vs 2 , 2 vs 1 , 1 vs 0
			4 vs 3 , 3 vs 2 , 2 vs 1 , 1 vs 0

		*/

		for j := i; j >= 1; j-- {
			left := numbers[j-1]
			right := numbers[j]

			if left > right {
				numbers[j] = left
				numbers[j-1] = right
			}
		}
	}

	return numbers
}
