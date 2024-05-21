package main_test

import (
	"encoding/json"
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

type GiftTradeInfo struct {
	TotalGiveCnt int
	TotalTakeCnt int
	GiftIndex    int
	NewGift      int
	GiftTrade    map[string]*GiftCnt
}

type GiftCnt struct {
	GiveCnt int
	TakeCnt int
}

func solution(friends []string, gifts []string) int {
	//선물 주고받은 기록 생성하기

	var giftTradeHistory = make(map[string]*GiftTradeInfo)
	for _, friend := range friends {
		giftTradeHistory[friend] = &GiftTradeInfo{
			TotalGiveCnt: 0,
			TotalTakeCnt: 0,
			GiftTrade:    make(map[string]*GiftCnt),
		}
	}

	//선물 거래가 없었던 친구들 목록을 만들어야 한다.
	var tradedPairs = make(map[string]bool)
	for _, gift := range gifts {
		tradedPairs[gift] = true

		giver := strings.Split(gift, " ")[0]
		taker := strings.Split(gift, " ")[1]

		giftTradeHistory[giver].TotalGiveCnt++
		giftTradeHistory[taker].TotalTakeCnt++

		if giftTradeHistory[giver].GiftTrade[taker] == nil {
			giftTradeHistory[giver].GiftTrade[taker] = &GiftCnt{}
		}
		giftTradeHistory[giver].GiftTrade[taker].GiveCnt++

		if giftTradeHistory[taker].GiftTrade[giver] == nil {
			giftTradeHistory[taker].GiftTrade[giver] = &GiftCnt{}
		}
		giftTradeHistory[taker].GiftTrade[giver].TakeCnt++

	}
	//이걸 기준으로 판단한다
	//일단 쌍을 무조건 다 만들어서 거래 안한 리스트에 올린다.
	var untradedPairs = make(map[string]bool)
	for i, friend := range friends {
		for j, friend2 := range friends {
			if i == j {
				continue
			}

			pairKey := friend + " " + friend2
			pairKey2 := friend2 + " " + friend
			if !tradedPairs[pairKey] && !tradedPairs[pairKey2] { //일절 거래 없었던 경우
				untradedPairs[friend+" "+friend2] = true
				giftTradeHistory[friend].GiftTrade[friend2] = &GiftCnt{
					GiveCnt: 0,
					TakeCnt: 0,
				}
			}
		}
	}

	bytes, _ := json.MarshalIndent(giftTradeHistory, "", "  ")

	fmt.Println(string(bytes), "giftTradeHistory 확인하기!! == 맨 처음")

	for _, history := range giftTradeHistory {
		//선물지수 먼저 계산
		history.GiftIndex = history.TotalGiveCnt - history.TotalTakeCnt
	}

	bytes, _ = json.MarshalIndent(giftTradeHistory, "", "  ")

	fmt.Println(string(bytes), "giftTradeHistory 확인하기!! == 선물지수 계산 후")

	var answer = 0
	for _, history := range giftTradeHistory {
		for pairName, trade := range history.GiftTrade {
			//두 사람이 선물을 주고받은 기록이 있다면, 이번 달까지 두 사람 사이에 더 많은 선물을 준 사람이 다음 달에 선물을 하나 받습니다
			if trade.GiveCnt > trade.TakeCnt {
				history.NewGift++
			}

			if (trade.GiveCnt == 0 && trade.TakeCnt == 0) || (trade.GiveCnt == trade.TakeCnt) {
				if history.GiftIndex > giftTradeHistory[pairName].GiftIndex {
					history.NewGift++
				}
			}
		}
	}

	bytes, _ = json.MarshalIndent(giftTradeHistory, "", "  ")

	fmt.Println(string(bytes), "giftTradeHistory 확인하기!! == 다음 달 받는 선물 수 계산 후")

	for _, history := range giftTradeHistory {
		if history.NewGift >= answer {
			answer = history.NewGift
		}
	}

	return answer
}
