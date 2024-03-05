package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/pointer"
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
		assert.Equal(t, test.expect, ans)
	}
}

func solution(k int, score []int) []int {
	var answer []int
	var podium = initPodium(score[:k])

	//처음 k개까지는, 가장 낮은 거 동일하ㅔ 채워넣는다.
	lowestScoreOfKdays := getLowestScore(score[:k])
	for len(answer) < k {
		answer = append(answer, lowestScoreOfKdays)
	}

	//그 다음부터는 1명씩 갈아치워야 한다.
	for i := k; i < len(score); i++ {
		answer = append(answer,  getLowestScore(podium))
		podium = updatePodium(podium, score[i])
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

func initPodium(scoresOfKdays []int) []int {

	var podium = []int{}
	podium[0] = scoresOfKdays[0]


	leftIdx := 0
	rightIdx := len(scoresOfKdays) - 1

	for i := 1; i < len(scoresOfKdays); i++ {
		firstVal := scoresOfKdays[0]
		lastVal := podium[i-1]

		//가운데 끼는 값인경우
		if firstVal < scoresOfKdays[i] && scoresOfKdays[i] < lastVal {

		}

		//firstVal보다도 작은경우
		if scoresOfKdays[i] =< firstVal {
			podium = append([]int{scoresOfKdays[i]}, podium...)
			continue
		}

		//lastVal보다도 큰경우
		if scoresOfKdays[i] >= lastVal {
			podium = append(podium, scoresOfKdays[i])
			continue
		}
	}

	return podium
}


func updatePodium(podium []int, score int) []int {
	lowestScore := podium[0]
	highestScore := podium[len(podium)-1]

	if score > highestScore {
		podium = append(podium, score)
		return podium[1:]
	}

	if score < lowestScore {
		return podium
	}

	//가운데에 끼는 경우..
	
}