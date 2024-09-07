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

func TestMergeSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3, 6, 11, 9, 20, 10},
			expect:  []int{1, 3, 4, 5, 6, 7, 9, 10, 11, 20},
		},
	}

	for _, test := range tests {
		ans := MergeSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

// MergeSort recursively sorts a slice of integers using merge sort algorithm
func MergeSort(arr []int) []int {
	// Base case: if the slice has 1 or 0 elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle index
	mid := len(arr) / 2

	//호출 스택 때문에 합쳐진걸 다시 나눌 일은 없음.
	//합쳐진 결과값은 아래 머지 함수에서만 쓰임

	// Recursively sort both halves
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge merges two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	// Create a result slice to store the merged values
	result := []int{}

	// Indices for left and right slices
	i, j := 0, 0

	// Merge while there are elements in both slices
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left or right slice (이미 정렬이 되어 있음.. 그리고 나뉘어져 있더라도 길이 차이가 1 차이이거나 길이가 같음.)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
