package main_test

import (
	"fmt"
	"strings"
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

	var realTradeInfo = make(map[string]int)
	for _, trade := range gifts {
		realTradeInfo[trade]++
	}

	var alltradeInfo = make(map[string]int)
	for _, giver := range friends {
		for _, taker := range friends {

			if giver == taker {
				continue
			}

			pairName := fmt.Sprintf("%s %s", giver, taker)
			if _, ok := realTradeInfo[pairName]; ok {
				alltradeInfo[pairName]++
			} else {
				alltradeInfo[pairName] = 0
			}
		}
	}

	var giftTradeInfo = make(map[string]*GiftInfo)
	for _, giver := range friends {
		for _, taker := range friends {
			if giver == taker {
				continue
			}

			if _, ok := giftTradeInfo[giver]; !ok {
				giftTradeInfo[giver] = &GiftInfo{}
			}

			if _, ok := giftTradeInfo[taker]; !ok {
				giftTradeInfo[taker] = &GiftInfo{}
			}

			pairName := fmt.Sprintf("%s %s", giver, taker)
			if _, ok := realTradeInfo[pairName]; ok {
				//일단 위와 같이 초기화를 해주고 규칙에 따라서 주고받은 선물 내역 계산해준다
				giftTradeInfo[giver].GiveCnt++
				giftTradeInfo[taker].GetCnt++
			}
		}
	}

	//선물지수 계산
	for _, trade := range giftTradeInfo {
		trade.GiftIndx = trade.GiveCnt - trade.GetCnt
	}

	for namePair, giveCnt := range alltradeInfo {
		prnsl := strings.Split(namePair, " ")
		giver := prnsl[0]
		taker := prnsl[1]

		op := fmt.Sprintf("%s %s", taker, giver)

		takeCnt := alltradeInfo[op]

		//더 많이 준 경우
		if giveCnt > takeCnt {
			giftTradeInfo[giver].GiftToGet++
		}

		//같은 경우
		if (giveCnt == 0 && giftTradeInfo[giver].GetCnt == 0) || giveCnt == takeCnt {
			if giftTradeInfo[giver].GiftIndx > giftTradeInfo[taker].GiftIndx {
				giftTradeInfo[giver].GiftToGet++
			} else if giftTradeInfo[giver].GiftIndx < giftTradeInfo[taker].GiftIndx {
				giftTradeInfo[taker].GiftToGet++
			}
		}

		if takeCnt > giveCnt {
			giftTradeInfo[taker].GiftToGet++
		}

		delete(alltradeInfo, namePair)
		delete(alltradeInfo, op)
	}

	var maxGiftN int
	for _, giftInfo := range giftTradeInfo {
		if giftInfo.GiftToGet > maxGiftN {
			maxGiftN = giftInfo.GiftToGet
		}
	}

	//이때, 다음달에 가장 많은 선물을 받는 친구가 받을 선물의 수
	return maxGiftN
}
