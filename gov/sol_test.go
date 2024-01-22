package main_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  string
	expect string
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			input:  "try hello world",
			expect: "TrY HeLlO WoRlD",
		},
		{
			input:  "try he           llo wor      ld",
			expect: "TrY He           LlO WoR      Ld",
		},
	}

	for _, test := range tests {
		ans := solution(test.input)
		assert.Equal(t, test.expect, ans)
	}
}

func solution(s string) string {
	var reversedWordSlice = []string{}

	words := strings.Split(s, " ")
	for _, word := range words {
		var reversed string
		for idx, char := range strings.Split(word, "") {
			if char == " " {
				reversed += char
				idx--
			} else {
				if idx == 0 || idx%2 == 0 {
					reversed += strings.ToUpper(char)
				} else {
					reversed += strings.ToLower(char)
				}
			}
		}
		reversedWordSlice = append(reversedWordSlice, reversed)
	}

	return strings.Join(reversedWordSlice, " ")
}
