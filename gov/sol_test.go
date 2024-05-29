package main_test

import (
	"math"
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	N      int
	A      int
	B      int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	N:      8,
		// 	A:      4,
		// 	B:      7,
		// 	expect: 3,
		// },
		// {
		// 	N:      4,
		// 	A:      1,
		// 	B:      2,
		// 	expect: 1,
		// },
		// {
		// 	N:      8,
		// 	A:      4,
		// 	B:      5,
		// 	expect: 3,
		// },
		// {
		// 	N:      16,
		// 	A:      1,
		// 	B:      9,
		// 	expect: 4,
		// },
		// {
		// 	N:      16,
		// 	A:      9,
		// 	B:      13,
		// 	expect: 3,
		// },
		// {
		// 	N:      8,
		// 	A:      5,
		// 	B:      8,
		// 	expect: 2,
		// },
		// {
		// 	N:      16,
		// 	A:      7,
		// 	B:      8,
		// 	expect: 1,
		// },
		{
			N:      16,
			A:      12,
			B:      13,
			expect: 3,
		},
	}

	for _, test := range tests {
		ans := solution(test.N, test.A, test.B)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(n int, a int, b int) int {
	answer := 0

	var aNum int = a
	var bNum int = b

	for {
		ns := []int{aNum, bNum}
		sort.Ints(ns)

		if int(math.Abs(float64(aNum-bNum))) == 1 && ns[1]%2 == 0 {
			break
		}

		arem := aNum % 2

		if arem == 0 {
			aNum = aNum / 2
		}

		if arem == 1 {
			aNum = aNum/2 + 1
		}

		brem := bNum % 2

		if brem == 0 {
			bNum = bNum / 2
		}

		if brem == 1 {
			bNum = bNum/2 + 1
		}

		answer++
	}

	answer++

	return answer
}
