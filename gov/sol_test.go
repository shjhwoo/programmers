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
			numbers: []int{2, 3, 3, 5}, //5,3,2
			expect:  []int{3, 5, 5, -1},
		},
		{
			numbers: []int{9, 1, 5, 3, 6, 2}, //6,5,1
			expect:  []int{-1, 5, 6, 6, -1, -1},
		},
		{
			numbers: []int{1, 1, 1, 1},
			expect:  []int{-1, -1, -1, -1},
		},
		{
			numbers: []int{2, 3, 4, 6, 3, 2, 7}, //7,6,4,3
			expect:  []int{3, 4, 6, 7, 7, 7, -1},
		},
		//이 반례만 해결하면 될거 같은데.
		{
			numbers: []int{8, 7, 6, 5, 4, 5, 6, 7, 8}, //8,7,6,5,7
			expect:  []int{-1, 8, 7, 6, 5, 6, 7, 8, -1},
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
		for j, compareNum := range compareSlice { //[6,5,1]

			if num >= compareNum || compareMap[j] < idx { //인덱스 검사 오류 수정할것 요거때문애 안됨
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

// 만드는 방법이 잘못됨: 뭉갤 때 끝까지 뭉개는 게 아니고 한번만 뭉개야할듯.. 이게맞나..;
func makeCompareSliceAndMap(numbers []int) (resultSlice []int, resultIdx []int) {

	//뒤에서부터, 최대값 모은다
	for i := len(numbers) - 1; i >= 0; i-- {
		if i == 0 {
			break
		}

		if i == len(numbers)-1 {
			resultSlice = append(resultSlice, numbers[i])
			resultIdx = append(resultIdx, i)
		} else {
			currentNum := numbers[i]
			recentNum := resultSlice[len(resultSlice)-1]

			if currentNum < recentNum {
				resultSlice = append(resultSlice, currentNum)
				resultIdx = append(resultIdx, i)
			} else if currentNum == recentNum {
				continue
			} else {
				//자기보다 작은 수가 있다면 pop.

				if currentNum > recentNum && len(resultSlice) > 0 {
					resultSlice = resultSlice[:len(resultSlice)-1]
					resultIdx = resultIdx[:len(resultIdx)-1]
				}

				resultSlice = append(resultSlice, currentNum)
				resultIdx = append(resultIdx, i) //인덱스 검사 오류 수정할것
			}
		}
	}

	return resultSlice, resultIdx
}

//https://school.programmers.co.kr/questions/43218
//우선순위큐
