package main_test

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	k      []string
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			k:      []string{"aya", "yee", "u", "maa"},
			expect: 1,
		},
		{
			k:      []string{"ayaye", "uuu", "yeye", "yemawoo", "ayaayaa"},
			expect: 2,
		},
	}

	for _, test := range tests {
		ans := solution(test.k)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(babbling []string) int {

	var answer int

	for _, word := range babbling {
		if isAbleToSpeak(word) {
			answer++
		}
	}

	return answer
}

func isAbleToSpeak(word string) bool {
	//단어를 어떻게 쪼갤 수 있을까?

	/*
		경우의 수를 나누어 생각한다

		1) 단어이 길이가 1인 경우: 불가능

		2) 단어의 길이가 2인 경우: ye, ma만 가능

		3) 단어의 길이가 3인 경우: aya, woo 만 가능함

		4) 단어의 길이가 4인 경우: maye, yema 4가지만 가능하다!


		주어진 단어를 1글자 단위로 쪼개본다. 그러면, a, y, w, m 중 분명 한 글자는 나온다.
		a로 시작하면 a포함 3글자까지 분석해서 aya와 같은지 판단. 다르면 바로 탈락
		y로 시작하면 y포함 2글자까지 분석해서 ye와 같은지 판단. 다르면 바로 탈락
		w로 시작하면 w포함 3글자까지 분석해서 woo와 같은지 판단. 다르면 바로 탈락
		m으로 시작하면 m포함 2글자까지 분석해서 ma와 같은지 판단. 다르면 바로 탈락

		그러면, 같은 경우에는 바로 다음 글자로 건너뛰어 분석하는데, 또 같은 글자로 시작한다면 바로 탈락시킨다!!
	*/

	if len(word) == 1 {
		return false
	}

	if len(word) == 2 {
		if word == "ye" || word == "ma" {
			return true
		}
	}

	if len(word) == 3 {
		if word == "aya" || word == "woo" {
			return true
		}
	}

	var subWords []string
	var lastWord string
	splittedWord := strings.Split(word, "")

	var idx int

	for idx < len(splittedWord) {

		var char = splittedWord[idx]

		if char == "a" || char == "w" {
			var nextIdx = idx + 3
			if nextIdx >= len(splittedWord) {
				nextIdx = len(splittedWord)
			}
			if word[idx:nextIdx] == "aya" || word[idx:nextIdx] == "woo" {
				if lastWord != "" && lastWord == word[idx:nextIdx] {
					return false
				}
				subWords = append(subWords, word[idx:nextIdx])
				lastWord = subWords[len(subWords)-1]
				idx += 3
				continue
			}
		}

		if char == "y" || char == "m" {
			var nextIdx = idx + 2
			if nextIdx >= len(splittedWord) {
				nextIdx = len(splittedWord)
			}
			if word[idx:nextIdx] == "ye" || word[idx:nextIdx] == "ma" {
				if lastWord != "" && lastWord == word[idx:nextIdx] {
					return false
				}
				subWords = append(subWords, word[idx:nextIdx])
				lastWord = subWords[len(subWords)-1]
				idx += 2
				continue
			}
		}
		return false
	}

	return true
}
