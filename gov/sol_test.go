package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	progresses []int
	speeds     []int
	expect     []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			progresses: []int{93, 30, 55},
			speeds:     []int{1, 30, 5},
			expect:     []int{2, 1},
		},
		{
			progresses: []int{95, 90, 99, 99, 80, 99},
			speeds:     []int{1, 1, 1, 1, 1, 1},
			expect:     []int{1, 3, 2},
		},
		{
			progresses: []int{90, 91, 92, 93, 94, 95},
			speeds:     []int{1, 1, 1, 1, 1, 1}, // 10, 9, 8, 7, 6, 5
			expect:     []int{6},
		},
		{
			progresses: []int{90, 91, 92, 93, 94, 95},
			speeds:     []int{1, 1, 1, 1, 1, 1}, // 10, 9, 8, 7, 6, 5
			expect:     []int{6},
		},
		{
			progresses: []int{90, 91, 92, 93, 94, 90},
			speeds:     []int{1, 1, 1, 1, 1, 1}, // 10, 9, 8, 7, 6, 10
			expect:     []int{6},
		},
	}

	for _, test := range tests {
		ans := solution(test.progresses, test.speeds)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(progresses []int, speeds []int) []int {

	//일단 배열을 돌면서 작업 완료까지 필요한 기간을 계산해본다
	var remainDaysForDeploy []int
	for i := 0; i < len(progresses); i++ {
		rm := (100 - progresses[i]) % speeds[i]
		daysNeed := (100 - progresses[i]) / speeds[i]
		if rm > 0 {
			daysNeed++
		}

		remainDaysForDeploy = append(remainDaysForDeploy, daysNeed)
	}

	//순회하면서 뭉개버리기.::
	var answer []int

	for i, rmDays := range remainDaysForDeploy {
		if i == 0 {
			answer = append(answer, 1)
			continue
		}

		//정확한 조건: 거꾸로 올라가면서 모두 다 훑어봤을 때 자기보다 더 큰 수가 진짜 하나도 없을 때만 추가를 하고,
		//자기와 같거나 더 큰 수가 있으면은 그냥 더해주면 된다.
		//자기보다 같거나 더 큰 수를 발견한 그 순간 탐색을 중지함으로서 속도를 빨리할 수 있겠다

		if hasBiggerOrEqualNumBeforeMe(remainDaysForDeploy[:i], rmDays) {
			answer[len(answer)-1]++
		} else {
			answer = append(answer, 1)
		}
	}

	return answer
}

func hasBiggerOrEqualNumBeforeMe(rmDaysList []int, rmDays int) bool {
	for i := len(rmDaysList) - 1; i >= 0; i-- {
		pNum := rmDaysList[i]
		if pNum >= rmDays {
			return true
		}
	}

	return false
}
