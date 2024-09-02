package main_test

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

func FuzzSolution(f *testing.F) {
	// Seed corpus with example values
	// Since []int is not allowed, we use []byte and convert it to []int in the fuzz function
	f.Add(6, []byte{7, 10}) // Corresponds to n=6 and times=[7,10]

	f.Fuzz(func(t *testing.T, n int, timesBytes []byte) {
		// Convert []byte to []int
		times := make([]int, len(timesBytes))
		for i, b := range timesBytes {
			times[i] = int(b)
		}

		// Validate inputs
		if n < 0 {
			t.Skip() // Skip invalid test case
		}
		if len(times) == 0 {
			t.Skip() // Skip test case with empty times slice
		}
		for _, time := range times {
			if time <= 0 {
				t.Skip() // Skip invalid times
			}
		}

		// Call the solution function
		res := solution(n, times)

		// Optionally, you can perform assertions if you have expected behavior
		// For example, ensuring the result is non-negative
		if res < 0 {
			t.Errorf("Expected non-negative result, got %d", res)
		}

		// Log the result for debugging purposes
		t.Logf("n: %d, times: %v, result: %d", n, times, res)
	})
}

type TestCase struct {
	n      int
	times  []int
	expect int64
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			n:      6,
			times:  []int{7, 10},
			expect: 28,
		},
		{
			n:      6,
			times:  []int{7},
			expect: 42,
		},
	}

	for _, test := range tests {
		ans := solution(test.n, test.times)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

// 결정문제 => parametric search.
func solution(n int, times []int) int64 {

	shortestExamTime := getShortestExamTime(times)

	var start = int64(1)
	var end = int64(shortestExamTime * n)

	for start <= end {
		var mid = int64(math.Round(float64((start + end) / 2)))

		if ableToFinishExamIn(n, mid, times) {
			//더 최소값을 찾을 수 있을지.. (주어진 시간 내에 주어진 인원을 다 심사가능하단 소리)
			end = mid - 1
		} else {
			//주어진 시간 내 n명 모두 다 보기는 힘드니까 시간을 늘려달라.
			start = mid + 1
		}
	}

	return int64(end)
}

func getShortestExamTime(times []int) int {
	var result = times[0]
	for _, time := range times[1:] {
		if result > time {
			result = time
		}
	}

	return result
}

func ableToFinishExamIn(n int, examTime int64, times []int) bool {
	var examCnt int64
	for _, time := range times {
		examCnt += examTime / int64(time)
	}

	return int64(n) <= examCnt
}
