package main_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputT string
	inputP string
	expect int
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			inputT: "3141592",
			inputP: "271",
			expect: 2,
		},
		{
			inputT: "500220839878",
			inputP: "7",
			expect: 8,
		},
		{
			inputT: "10203",
			inputP: "15",
			expect: 3,
		},
	}

	for _, test := range tests {
		ans := solution(test.inputT, test.inputP)
		assert.Equal(t, test.expect, ans)
	}
}

func solution(t string, p string) int {
	var answer int

	subStringSize := len(p)

	pn := stringToNum(p)

	for i, ns := range strings.Split(t, "") {
		if i == (len(t) - subStringSize + 1) {
			break
		}

		//첫번째 자리 숫자가 크면 패스
		if stringToNum(string(ns)) > stringToNum(p[0:1]) {
			continue
		}

		//그 외에는 탐색가능.
		if stringToNum(t[i:i+subStringSize]) <= pn {
			answer++
		}

	}

	return answer
}

func stringToNum(strNum string) int {
	n, _ := strconv.Atoi(strNum)
	return n
}
