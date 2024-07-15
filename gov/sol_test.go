package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

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
			n:      7,
			times:  []int{10, 6, 3, 1, 2},
			expect: 40,
		},
	}

	for _, test := range tests {
		ans := solution(test.n, test.times)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(n int, times []int) int64 {

	minTime := getMinTime(times)

	start := 1
	end := minTime * n

	for start <= end { //왜 같다는 것까지 조건에 포함하지,
		//=> 안그러면 start에 정답값인 mid를 할당할 수 없기 때문임.
		mid := (start + end) / 2 // ==주어진 시간

		sumOfCheckablePrson := 0
		for _, time := range times {
			sumOfCheckablePrson += mid / time
		}

		if sumOfCheckablePrson < n {
			//범위를 더 늘린다
			start = mid + 1
		} else {
			//범위를 줄인다
			end = mid - 1
		}
	}

	return int64(start)
}

func getMinTime(times []int) int {
	minTime := times[0]
	for _, time := range times {
		if minTime > time {
			minTime = time
		}
	}
	return minTime
}
