package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	numbers []int
	expect  []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{2, 3, 3, 5},
			expect:  []int{3, 5, 5, -1},
		},
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
	compareSlice, compareMap := makeCompareSliceAndMap(numbers)

	for idx, num := range numbers { //[9,1,5,3,6,2]
		var found bool
		var foundCnt int
		for _, compareNum := range compareSlice { //[6,5,1]

			if num >= compareNum || compareMap[compareNum] < idx {
				break
			}

			if num < compareNum {
				found = true
				foundCnt++
				if foundCnt > 1 {
					answer = answer[:len(answer)-1]
				}
				answer = append(answer, compareNum)
			}
		}

		if !found {
			answer = append(answer, -1)
		}
	}

	return answer
}

func makeCompareSliceAndMap(numbers []int) (resultSlice []int, resultMap map[int]int) {
	resultMap = make(map[int]int)

	//뒤에서부터, 최대값 모은다
	for i := len(numbers) - 1; i >= 0; i-- {
		if i == 0 {
			break
		}

		if i == len(numbers)-1 {
			resultSlice = append(resultSlice, numbers[i])
			resultMap[numbers[i]] = i
		} else {
			currentNum := numbers[i]
			recentNum := resultSlice[len(resultSlice)-1]

			if currentNum < recentNum {
				resultSlice = append(resultSlice, currentNum)
				resultMap[currentNum] = i
			} else if currentNum == recentNum {
				continue
			} else {
				//자기보다 작은 수가 있다면 pop.
				for {
					if currentNum > recentNum && len(resultSlice) > 0 {
						resultSlice = resultSlice[:len(resultSlice)-1]

						if len(resultSlice) == 0 {
							recentNum = currentNum
						} else {
							recentNum = resultSlice[len(resultSlice)-1]
						}
					} else {
						break
					}
				}
				resultSlice = append(resultSlice, currentNum)
				resultMap[currentNum] = i
			}
		}
	}

	return resultSlice, resultMap
}

//https://school.programmers.co.kr/questions/43218
//우선순위큐
