package main_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type DFSGraphTestCase struct {
	number string
	k      int //제거할 수!
	expect string
}

func TestSolution(t *testing.T) {
	var tests = []DFSGraphTestCase{
		{
			number: "1924",
			k:      2,
			expect: "94",
		},
		{
			number: "1231234",
			k:      3,
			expect: "3234",
		},
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

func solution(number string, k int) string {

	/*
		제거를 하면 몇자리 숫자가 되는지를 파악해야 한다?
		음 일단은 앞쪽에 있는 숫자가 최대한 크도록 제거해나가야한다.

		제일 작은 숫자들을 빼내야 한다.

		제일 작은 숫자 같은게 여러 개 있으면 그 중 제일 앞에 있는 숫자를 제거한다.

		그러면 일단은 숫자 메타정보 맵을 만들어보자.

		해당 숫자: 해당 숫자가 어느어느 인덱스에 있는지를 알려주기. map[int][]int

		조심해야하는경우... 예) 10009170 에서 숫자 3개만 빼라고 했을때 앞의 숫자 0을 빼면 19170, 뒤의 숫자 0을 빼면 10917
	*/

	idxToRmMap := buildNumberMeta(number, k)

	var answer string
	var splitted = strings.Split(number, "")
	for idx, numCh := range splitted {
		//제거조건에 안 맞으면은 answer에 가져다붙인다
		if _, needRm := idxToRmMap[idx]; !needRm {
			answer = answer + numCh
		}
	}

	return answer
}

func buildNumberMeta(number string, k int) map[int]bool {

	var foundNums []int

	var resMap = make(map[int][]int)
	for idx, numCh := range strings.Split(number, "") {
		n, _ := strconv.Atoi(numCh)

		foundNums = append(foundNums, n)

		resMap[n] = append(resMap[n], idx) //인덱스 순서대로 붙을거야..
	}

	sort.Ints(foundNums)

	//위 결과값을 인덱스 순서대로 정렬
	var res = make(map[int]bool)
	for _, nkey := range foundNums {

		for _, i := range resMap[nkey] {
			res[i] = true
			if len(res) == k {
				break
			}
		}

		if len(res) == k {
			break
		}
	}

	return res
}
