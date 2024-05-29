package main_test

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	k         int
	tangerine []int
	expect    int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			k:         6,
			tangerine: []int{1, 3, 2, 5, 4, 5, 2, 3},
			expect:    3,
		},
		{
			k:         4,
			tangerine: []int{1, 3, 2, 5, 4, 5, 2, 3},
			expect:    2,
		},
		{
			k:         2,
			tangerine: []int{1, 1, 1, 1, 2, 2, 2, 3},
			expect:    1,
		},
	}

	for _, test := range tests {
		ans := solution(test.k, test.tangerine)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(k int, tangerine []int) int {
	//일단 맵을 만든다. 크기 맵

	var tangerineSizeMap = make(map[int]int) //귤의 크기, 크기에 해당하는 귤의 개수

	for _, size := range tangerine {
		tangerineSizeMap[size]++
	}

	// fmt.Println(tangerineSizeMap, "tangerineSizeMap 확인..")
	//map[1:1 2:2 3:2 4:1 5:2] tangerineSizeMap 확인..

	//가짓수를 최소화 하고 싶으면, 크기별 갯수가 가장 큰 귤부터 담고, 나머지를 채워 나가야 함..
	//갯수 - 귤의 크기 목록 맵핑하기
	var cntSizeMap = make(map[int][]int) //갯수, 크기
	var cntSorted = []int{}
	for size, cnt := range tangerineSizeMap {
		cntSizeMap[cnt] = append(cntSizeMap[cnt], size)
	}

	for cnt := range cntSizeMap {
		cntSorted = append(cntSorted, cnt)
	}
	sort.Ints(cntSorted) //오름차순 정렬

	// fmt.Println(cntSizeMap, "cntSizeMap 확인..")
	// fmt.Println(cntSorted, "cntSorted 확인..")

	var answer int
	var emptySpace = k
	for i := len(cntSorted) - 1; i >= 0; i-- {
		if emptySpace == 0 {
			return answer
		}

		cnt := cntSorted[i]
		//이 값이 k보다 크거나 같고 처음이면 1 리턴
		if cnt >= k && i == len(cntSorted)-1 {
			return 1
		}

		for len(cntSizeMap[cnt]) > 0 {
			if emptySpace == 0 {
				return answer
			}

			emptySpace -= cnt
			answer++
			cntSizeMap[cnt] = cntSizeMap[cnt][1:]
		}
	}

	return answer
}
