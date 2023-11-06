package main_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	sequence []int
	k        int
	expect   []int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			sequence: []int{1, 2, 3, 4, 5},
			k:        7,
			expect:   []int{2, 3},
		},
		{
			sequence: []int{1, 1, 1, 2, 3, 4, 5},
			k:        5,
			expect:   []int{6, 6},
		},
		{
			sequence: []int{2, 2, 2, 2, 2},
			k:        6,
			expect:   []int{0, 2},
		},
	}

	for _, test := range tests {

		if !assert.True(t, slices.Compare(test.expect, Solution(test.sequence, test.k)) == 0) {
			t.Log(test.sequence, Solution(test.sequence, test.k), "@@@@@@@@@@@@@@@@@@")
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
	//슬라이딩 윈도우를 이용한 방법 => 고정된 크기라    안된다
	//투 포인터 패턴으로 접근하면 될거같다. https://butter-shower.tistory.com/226
	var leftPointer int
	var rightPointer int

	idx, found := slices.BinarySearch(sequence, k)
	if found {
		return []int{idx, idx}
	}

	var sum = sequence[0]

	for sum != k {
		if sum < k {
			rightPointer++
			sum = calculateSum(sum, sequence[rightPointer])
			continue
		}

		if sum == k {
			break
		}

		if sum > k {
			sum = calculateSum(sum, -1*sequence[leftPointer])
			leftPointer++
			continue
		}
	}

	return []int{leftPointer, rightPointer}
}

func calculateSum(sum int, addition int) int {
	return sum + addition
}
