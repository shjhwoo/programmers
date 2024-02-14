package main_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputS string
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			inputS: "one4seveneight",
			expect: 1478,
		},
		{
			inputS: "23four5six7",
			expect: 234567,
		},
		{
			inputS: "2three45sixseven",
			expect: 234567,
		},
		{
			inputS: "123",
			expect: 123,
		},
		{
			inputS: "twozerozero0000zero9ninezerozerozero78six",
			expect: 2000000099000786,
		},
	}

	for _, test := range tests {
		ans := solution(test.inputS)

		t.Log(ans, "계산값")

		assert.Equal(t, test.expect, ans)
	}
}

/*
수도:

s의 길이가 1,2이라는건 무조건 숫자라는 거니까 그대로 리턴하면 된다

길이가 3일때 문자가 있다면 그건 무조건 영단어가 3글자인 거임. => 한자리 숫자임을 보장한다

길이가 4, 5일때도 문자가 있다면 같은 원리로 한자리 숫자임을 보장한다.

길이가 6인 경우부터, 숫자거나 문자일 수 있다. oneone, 111111
*/

type NumStInfo struct {
	NumSt  string
	NextTo int
}

var numStringMap = map[string]NumStInfo{
	"ze": {NumSt: "0", NextTo: 4}, // --현재 인덱스에 4를 더한 인덱스부터 검사해라
	"on": {NumSt: "1", NextTo: 3}, // -- 현재 인덱스에 3을 더한 인덱스부터 검사해라
	"tw": {NumSt: "2", NextTo: 3},
	"th": {NumSt: "3", NextTo: 5},
	"fo": {NumSt: "4", NextTo: 4},
	"fi": {NumSt: "5", NextTo: 4},
	"si": {NumSt: "6", NextTo: 3},
	"se": {NumSt: "7", NextTo: 5},
	"ei": {NumSt: "8", NextTo: 5},
	"ni": {NumSt: "9", NextTo: 4}, // ... 현재 인덱스에 글자수 길이만큼을 더한 인덱스로 가서 검사해라
}

func solution(s string) int {

	var numst string
	for idx := 0; idx < len(s); idx++ {
		if idx == len(s)-1 {
			numst += s[idx : idx+1]
		}

		if idx+1 <= len(s)-1 {
			nst, exist := numStringMap[s[idx:idx+2]]
			if exist {
				numst += nst.NumSt
				idx += nst.NextTo - 1
				continue
			}

			//둘 다 숫자인경우
			_, err := strconv.Atoi(s[idx : idx+2])
			if err == nil {
				numst += s[idx : idx+2]
				idx += 1
				continue
			} else {

				_, err := strconv.Atoi(s[idx : idx+1])
				//첫번째 자리가 숫자고 두번째가 문자인경우
				if err == nil {
					numst += s[idx : idx+1]
					continue
				} else {
					//둘다 문자인 경우
					// fmt.Println(s[idx:idx+2], "둘다 문자래요")
				}

			}
		}

	}

	nst, err := strconv.Atoi(numst)
	if err != nil {
		return 0
	}

	return nst
}
