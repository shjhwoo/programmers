package main_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputS string
	inputN int
	expect string
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			inputS: "AB",
			inputN: 1,
			expect: "BC",
		},
		{
			inputS: "z",
			inputN: 1,
			expect: "a",
		},
		{
			inputS: "a B z",
			inputN: 4,
			expect: "e F d",
		},
	}

	for _, test := range tests {
		ans := solution(test.inputS, test.inputN)
		assert.Equal(t, test.expect, ans)
	}
}

func solution(s string, n int) string {
	var runes = strings.Split(s, "")

	var answerRunes []string

	for _, char := range runes {
		if char == " " {
			answerRunes = append(answerRunes, char)
			continue
		}
		answerRunes = append(answerRunes, encryptChar(char, n))
	}

	return strings.Join(answerRunes, "")
}

var weakAlphabets = "abcdefghijklmnopqrstuvwxyz"

var alpMap = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
	"i": 8,
	"j": 9,
	"k": 10,
	"l": 11,
	"m": 12,
	"n": 13,
	"o": 14,
	"p": 15,
	"q": 16,
	"r": 17,
	"s": 18,
	"t": 19,
	"u": 20,
	"v": 21,
	"w": 22,
	"x": 23,
	"y": 24,
	"z": 25,
}

/*
x + 8 인경우 => yzabcdef : 26에서 x의 인덱스를 뺀 값* 이거를 8에서 뺌 : n - (26 - xidx)

z + 25 == 51인 경우 => 25 - (26-26) = 25
z + 1 == 27 =>
*/

func encryptChar(s string, n int) string {

	oriIdx := alpMap[strings.ToLower(s)]

	newIdx := oriIdx + n

	if newIdx < 26 {
		c := weakAlphabets[newIdx : newIdx+1]
		if IsStrongChar(s) {
			return strings.ToUpper(c)
		}
		return c
	}

	c := weakAlphabets[n-(26-oriIdx) : n-(26-oriIdx)+1]
	if IsStrongChar(s) {
		return strings.ToUpper(c)
	}
	return c
}

func IsStrongChar(char string) bool {
	return strings.ToUpper(char) == char
}

func IsWeakChar(char string) bool {
	return strings.ToLower(char) == char
}
