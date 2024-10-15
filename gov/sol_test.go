package main_test

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	gems   []string
	expect []int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			gems:   []string{"DIA", "RUBY", "RUBY", "DIA", "DIA", "EMERALD", "SAPPHIRE", "DIA"},
			expect: []int{3, 7},
		},
		{
			gems:   []string{"AA", "AB", "AC", "AA", "AC"},
			expect: []int{1, 3},
		},
		{
			gems:   []string{"XYZ", "XYZ", "XYZ"},
			expect: []int{1, 1},
		},
		{
			gems:   []string{"ZZZ", "YYY", "NNNN", "YYY", "BBB"},
			expect: []int{1, 5},
		},
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.gems))
	}
}

//진열된 모든 종류의 보석을
//적어도 1개 이상 포함하는
//가장 짧은 구간을 찾아서 구매

func solution(gems []string) []int {

	gemMap := make(map[string]bool)
	for _, gem := range gems {
		gemMap[gem] = true
	}

	gems = append([]string{""}, gems...)

	//처음 콜렉션을 만들어 둔다..
	collected := make(map[string]int)

	startIdx := 1
	endIdx := 1

	collected[gems[startIdx]] = 1

	var answer = []int{1, len(gems) - 1}

	for startIdx < len(gems) && endIdx < len(gems) {
		if len(gemMap) == len(collected) {

			//더 줄일 수 있는지..
			if endIdx-startIdx < answer[1]-answer[0] {
				answer = []int{startIdx, endIdx}
			}

			collected[gems[startIdx]]--

			if collected[gems[startIdx]] == 0 {
				delete(collected, gems[startIdx])
			}

			startIdx++
		} else {
			endIdx++
			if endIdx == len(gems) {
				break
			}
			collected[gems[endIdx]]++
		}
	}

	return answer
}

/*
javascript:

포인트:
- 부분합을 구하여 저장해두는 맵을 loop 밖에 둬야 한다 => 맵끼리 빠른 비교 유도 (길이만으로 알 수 있으니까.)
- 경계값에 유의하여야 한다.

function solution(gems) {

    const gemKind = new Set(gems);

    const collection = new Map();

    let start = 0
    let end = 0

    collection.set(gems[start], 1)

    var answer = [1, gems.length];

    while (start < gems.length && end < gems.length){
        if (gemKind.size === collection.size) {

            if (end - start < answer[1] - answer[0]){
                answer = [start+1, end+1]
            }

            collection.set(gems[start], collection.get(gems[start]) - 1)

            if ( collection.get(gems[start]) === 0) {
                collection.delete(gems[start])
            }

            start++
        } else {
            end++
            collection.set(gems[end], (collection.get(gems[end]) || 0) + 1)
        }
    }

    return answer;
}
*/
