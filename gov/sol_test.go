package main_test

import (
	"fmt"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	number string
	k      int //제거할 수!
	expect string
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		// {
		// 	number: "1924",
		// 	k:      2,
		// 	expect: "94",
		// },
		// {
		// 	number: "1231234",
		// 	k:      3,
		// 	expect: "3234",
		// },
		{
			number: "4177252841",
			k:      4,
			expect: "775841",
		},
	}

	for _, test := range tests {
		ans := solution(test.number, test.k)
		assert.DeepEqual(t, test.expect, ans)
	}
}

// 최대한 앞에 있는 것 => 그 다음으로 제일 작은 수부터 빼낸다.
func solution(number string, k int) string { //결과값에는 뺴야 하는 숫자의 인덱스 위치를 저장해둔다
	var afterBiggerNumMap = make(map[int]bool)

	// var nums []int
	// for idx, numCh := range strings.Split(number, "") {
	// 	n, _ := strconv.Atoi(numCh)
	// 	nums = append(nums, n)
	// }

	// sort.Ints(nums) //오름차순 정렬이다. 상대적으로

	for idx, numCh := range strings.Split(number, "") {
		for _, afterNum := range strings.Split(number, "")[idx+1:] {
			if numCh < afterNum { //여기에다가 조건을 하나 더 넣던지 해야함. 현재 수 뒤에, 더 작은 숫자들이 존재하고 있다면 빼기에 곤란하다.
				afterBiggerNumMap[idx] = true
				break
			}
		}
	}

	fmt.Println("afterBiggerNumMap: ", afterBiggerNumMap)

	var answer string
	var rmCnt int
	for idx, numCh := range strings.Split(number, "") {
		//현재 숫자의 인덱스 이후의 뒤에 더 큰 숫자들이 존재한다면 빼도 된다. => 그러면 현재 인덱스 --- 자기보다 더 큰 숫자들이 존재
		//빼낸 갯수 다 채웠으면은 스탑한다.
		//조건 부족..
		if _, afterBiggerNumExists := afterBiggerNumMap[idx]; afterBiggerNumExists {
			rmCnt++
		} else {
			answer += numCh
		}

		if rmCnt == k {
			break
		}
	}

	return answer
}
