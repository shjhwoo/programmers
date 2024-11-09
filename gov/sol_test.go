package main_test

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	strs   []string
	t      string
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			strs:   []string{"ba", "na", "n", "a"},
			t:      "banana",
			expect: 3,
		},

		{
			strs:   []string{"app", "ap", "p", "l", "e", "ple", "pp"},
			t:      "apple",
			expect: 2,
		},
		{
			strs:   []string{"ba", "an", "nan", "ban", "n"},
			t:      "banana",
			expect: -1,
		},
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.strs, test.t))
	}
}

func solution(strs []string, t string) int {

	/*
		사용 가능한 단어 조각들을 담고 있는 배열 strs와
		 완성해야 하는 문자열 t가 매개변수로 주어질 때,
		 주어진 문장을 완성하기 위해 사용해야 하는 단어조각 개수의
		 최솟값을 return

		 만약 주어진 문장을 완성하는 것이 불가능하면 -1을 return
	*/

	var answer = 0

	/*

		ㅁㅁㅁ ...                 .... ㅁㅁ
		ㄴ 단어를 구성하는 글자들의 뭉침 덩어리들

		최소라고 했기 때문에..

		일단 0번째 인덱스에 있는 부분문자열이 t의 prefix인지 검사. , 그러면 남은 뒷 부분에 대해서도 같은 방식으로 알아서 검사가능
		만약에 prefix가 아니라면 다음 인덱스로 넘어가서 검사한다.
		prefix가 나올때까지 반복하고 찾았으면 answer에 카운팅을 해둔다. 그리고 t에서 prefix를 잘라내고 남은 부분에 대해 같은 방법으로 반복한다.
		단, prefix를 찾더라도, 최대한 길이가 긴 prefix를 찾아야 한다.
		=> 그런데 이렇게 한 것 때문에 뒤에 남은 부분에 맞은 prefix를 못 찾았다면, 카운팅을 취소한다.




	*/
	var prefixCandidates []string
	for _, pre := range strs {
		if strings.HasPrefix(t, pre) {
			prefixCandidates = append(prefixCandidates, pre)
		}
	}

	if len(prefixCandidates) == 0 {
		return -1
	}

	//후보군들을 찾았다
	for _, pre := range prefixCandidates {
		newT := strings.TrimPrefix(t, pre)
		cnt := solution(strs, newT)

		answer += cnt
		if cnt == -1 {
			//find other prefix again.
			//then repeat same thing
			continue //?
		} else {
			return 1 //?
		}
	}

	return answer
}
