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
		// { //for 문 조건에 대한 검증 테스트. < 로만 비교하게 되면 심사를 다 못볼 수 있다.
		// 	//한두명 더 보는 여유 시간이 남더라도 심사를 다 보는게 중요하기 때문이다.
		// 	n:      50,
		// 	times:  []int{3, 2, 1},
		// 	expect: 28,
		// },
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
			//범위를 더 늘린다: 최소의 시간이 조금 남더라도 심사를 다 보는게 중요하므로 이 값으로 리턴하는 게 맞다.
			start = mid + 1
		} else {
			//범위를 줄인다: end로 리턴하면, 주어진 시간 내에 모든 사람 심사 다 못볼수도 있다(줄이려다가 심사 다 못볼수도 있다)
			end = mid - 1
		}
	}

	return int64(end) //end말고 스타트값을 주는 것임에 유의해야 한다.
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
