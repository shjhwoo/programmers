package main_test

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	k      int
	score  []int
	expect []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			k:      3,
			score:  []int{10, 100, 20, 150, 1, 100, 200},
			expect: []int{10, 10, 10, 20, 20, 100, 100},
		},
		{
			k:      4,
			score:  []int{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000},
			expect: []int{0, 0, 0, 0, 20, 40, 70, 70, 150, 300},
		},
	}

	for _, test := range tests {
		ans := solution(test.k, test.score)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(k int, score []int) []int {
	var answer []int
	var podium = score[:k]

	//처음 k개까지는, 가장 낮은 거 동일하ㅔ 채워넣는다.
	lowestScoreOfKdays := getLowestScore(score[:k])
	for len(answer) < k {
		answer = append(answer, lowestScoreOfKdays)
	}

	//그 다음부터는 1명씩 갈아치워야 한다.
	for i := k; i < len(score); i++ {
		podium = append(podium, score[i])
		sort.IntSlice(podium).Sort()
		podium = podium[1:]
		answer = append(answer, getLowestScore(podium))

	}

	return answer
}

func getLowestScore(scores []int) int {
	lowest := scores[0]
	for _, v := range scores {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}
