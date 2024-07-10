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
		{
			numbers: []int{80, 20, 4, 5, 79, 60, 10},
			expect:  []int{-1, 79, 5, 79, -1, -1, -1},
		},
	}

	for _, test := range tests {
		ans := solution(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(numbers []int) []int {
	var answer []int

	var compareSlice []int
	var compareMap = make(map[int]int) //숫자와 그 index

	//일단 순회하는 건 필요하다.모든 숫자에 대해서 뒷 큰 수를 알아야 하니깐.!!
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			/*
				일단 처음에는 비교배열을 만들어야 한다
				그러므로 일단 비교할 숫자를 뒤에서부터 순회한다
			*/
			for j := len(numbers) - 1; j > i; j-- {
				if j == len(numbers)-1 {
					compareSlice = append(compareSlice, numbers[j])
					compareMap[numbers[j]] = j
				} else {
					//이제는 append를 앞에 해야 한다: 규칙 - 붙이는 수가 현재 맨 앞의 수보다 같거나 작으면 그대로 붙이고, 크다면 한번 shift해주고 붙인다.
					currentNum := compareSlice[0]
					numberToShift := numbers[j]
					if currentNum >= numberToShift {
						compareSlice = append([]int{numberToShift}, compareSlice...)
					}

					if currentNum < numberToShift {
						compareSlice = append([]int{numberToShift}, compareSlice[1:]...)
						delete(compareMap, currentNum)
					}

					compareMap[numberToShift] = j
				}
			}
		}

		//비교배열을 만들었으니까 그것만 가지고 뒤 큰 수를 찾아야 한다.
		for c := 0; c < len(compareSlice); c++ {

			compareIdx := compareMap[compareSlice[c]]

			if numbers[i] < compareSlice[c] && i < compareIdx {
				answer = append(answer, compareSlice[c])
				break
			}

			if c == len(compareSlice)-1 {
				answer = append(answer, -1)
			}
		}

	}

	return answer
}
