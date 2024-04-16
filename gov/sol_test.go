package main_test

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	friends []string
	gifts   []string
	expect  int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			friends: []string{"muzi", "ryan", "frodo", "neo"},
			gifts:   []string{"muzi frodo", "muzi frodo", "ryan muzi", "ryan muzi", "ryan muzi", "frodo muzi", "frodo ryan", "neo muzi"},
			expect:  2,
		},
		{
			friends: []string{"joy", "brad", "alessandro", "conan", "david"},
			gifts:   []string{"alessandro brad", "alessandro joy", "alessandro conan", "david alessandro", "alessandro david"},
			expect:  4,
		},
		{
			friends: []string{"a", "b", "c"},
			gifts:   []string{"a b", "b a", "c a", "a c", "a c", "c a"},
			expect:  0,
		},
		{
			friends: []string{"a", "b", "c"},
			gifts:   []string{"a b"},
			expect:  2,
		},
	}

	for _, test := range tests {
		ans := solution(test.friends, test.gifts)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

type GiftInfo struct {
	GiveCnt   int //준 선물 수
	GetCnt    int //받은 선물 수
	GiftIndx  int //선물지수
	GiftToGet int //받게 될 선물 수
}

func solution(friends []string, gifts []string) int {

	var giftMap = make(map[string]bool)
	for _, gift := range gifts {
		giftMap[gift] = true
	}

	var tradeCountMap = make(map[string]int)
	for _, giver := range friends {
		for _, taker := range friends {
			pairName := fmt.Sprintf("%s %s", giver, taker)

			if _, ok := giftMap[pairName]; ok {
				tradeCountMap[pairName]++
			}
		}
	}

	var friendGiftInfo = make(map[string]GiftInfo)

	for _, giver := range friends {
		//친구 한명한명 따져본다.

		var giveCnt int
		var takeCnt int

		for _, taker := range friends {
			pairName := fmt.Sprintf("%s %s", giver, taker)
			opPairName := fmt.Sprintf("%s %s", taker, giver)

		}
	}
}
