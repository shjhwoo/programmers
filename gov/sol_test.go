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
		// {
		// 	friends: []string{"joy", "brad", "alessandro", "conan", "david"},
		// 	gifts:   []string{"alessandro brad", "alessandro joy", "alessandro conan", "david alessandro", "alessandro david"},
		// 	expect:  4,
		// },
		// {
		// 	friends: []string{"a", "b", "c"},
		// 	gifts:   []string{"a b", "b a", "c a", "a c", "a c", "c a"},
		// 	expect:  0,
		// },
		// {
		// 	friends: []string{"a", "b", "c"},
		// 	gifts:   []string{"a b"},
		// 	expect:  2,
		// },
	}

	for _, test := range tests {
		ans := solution(test.friends, test.gifts)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

type GiftInfo struct {
	GiveMap   map[string]int
	TotalGive int
	TotalTake int
	GiftIndex int
	WillGet   int
}

func solution(friends []string, gifts []string) int {
	//일단 친구들을 가지고 모든 짝 조합을 만들어내고 값으로는 주고받은 선물 수를 저장한다.
	var GiftPairs = make(map[string]int)
	for _, giftTradePair := range gifts {
		GiftPairs[giftTradePair]++
	}

	var giftInfoMap = make(map[string]*GiftInfo)
	for _, friend := range friends {
		giftInfoMap[friend] = &GiftInfo{
			GiveMap: make(map[string]int),
		}
	}

	var AllFriendsFairs = make(map[string]int)
	for i, friend := range friends {
		for j, friend2 := range friends {
			if i == j {
				continue
			}
			pairKey := fmt.Sprintf("%s %s", friend, friend2)
			AllFriendsFairs[pairKey] = 0

			if GiftPairs[pairKey] > 0 {
				AllFriendsFairs[pairKey] = GiftPairs[pairKey]

				//선물을 주고받은 기록이 있었다.
				//준 선물, 받은 선물, 선물지수를 계산할 수 있다.
				giver := strings.Split(pairKey, " ")[0]
				taker := strings.Split(pairKey, " ")[1]

				giftInfoMap[giver].GiveMap[taker] = GiftPairs[pairKey]
				giftInfoMap[giver].TotalGive += GiftPairs[pairKey]
				giftInfoMap[giver].GiftIndex += GiftPairs[pairKey]

				giftInfoMap[taker].TotalTake += GiftPairs[pairKey]
				giftInfoMap[taker].GiftIndex -= GiftPairs[pairKey]
			}
		}
	}

	//일단 여기까지 주고받았던 선물 정보를을 맵으로 저장한다

	//전체 조합을 돈다

	var answer = 0

	for pair := range AllFriendsFairs {
		giver := strings.Split(pair, " ")[0]
		taker := strings.Split(pair, " ")[1]

		//선물지수에 따른, 다음달에 받을 선물 개수를 계산해야한다.
		//두 사람이 선물을 주고받은 기록이 있다면, 이번 달까지 두 사람 사이에 더 많은 선물을 준 사람이 다음 달에 선물을 하나 받습니다.
		if giftInfoMap[giver].GiveMap[taker] > giftInfoMap[taker].GiveMap[giver] {
			giftInfoMap[giver].WillGet++

			if answer < giftInfoMap[giver].WillGet {
				answer = giftInfoMap[giver].WillGet
			}
		}

		if giftInfoMap[giver].GiveMap[taker] == giftInfoMap[taker].GiveMap[giver] || (giftInfoMap[giver].GiveMap[taker] == 0 && giftInfoMap[taker].GiveMap[giver] == 0) {
			giverGiftIndex := giftInfoMap[giver].GiftIndex
			takerGiftIndex := giftInfoMap[taker].GiftIndex

			if giverGiftIndex > takerGiftIndex {
				giftInfoMap[giver].WillGet++
				if answer < giftInfoMap[giver].WillGet {
					answer = giftInfoMap[giver].WillGet
				}
			}

			if giverGiftIndex < takerGiftIndex {
				giftInfoMap[taker].WillGet++
				if answer < giftInfoMap[taker].WillGet {
					answer = giftInfoMap[taker].WillGet
				}
			}
		}
	}

	return answer
}
