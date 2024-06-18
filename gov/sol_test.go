package main_test

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	numbers []int
	expect  []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	numbers: []int{2, 3, 3, 5},
		// 	expect:  []int{3, 5, 5, -1},
		// },
		{
			numbers: []int{9, 1, 5, 3, 6, 2},
			expect:  []int{-1, 5, 6, 6, -1, -1},
		},
	}

	for _, test := range tests {
		ans := solution(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(numbers []int) []int {
	var answer = []int{}

	//비교 배열을 만든다
	compareSlice := makeCompareSlice(numbers)

	for i := len(numbers) - 1; i >= 0; i-- {
		var number = numbers[i]
		var found bool
		var foundCnt int
		for _, compareNum := range compareSlice {
			if number < compareNum {
				found = true
				foundCnt++
				if foundCnt > 1 {
					answer = answer[1:]
				}
				answer = append([]int{compareNum}, answer...)
			}
		}

		if !found {
			answer = append([]int{-1}, answer...)
		}
	}

	return answer
}

func makeCompareSlice(numbers []int) (result []int) {
	//뒤에서부터, 최대값 모은다
	for i := len(numbers) - 1; i >= 0; i-- {
		if i == len(numbers)-1 {
			result = append(result, numbers[i])
		} else {
			currentNum := numbers[i]
			recentNum := result[len(result)-1]

			if currentNum < recentNum {
				result = append(result, currentNum)
			} else if currentNum == recentNum {
				continue
			} else {
				//자기보다 작은 수가 있다면 pop.
				for {
					if currentNum > recentNum && len(result) > 0 {
						fmt.Println("res::", result)
						result = result[:len(result)-1]

						if len(result) == 0 {
							recentNum = currentNum
							result = append(result, currentNum)
						} else {
							recentNum = result[len(result)-1]
						}
					} else {
						break
					}
				}
			}
		}
	}

	return result
}

//https://school.programmers.co.kr/questions/43218
//우선순위큐
