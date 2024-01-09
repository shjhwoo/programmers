package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputs [2]int
	expect string
}

func TestSolution(t *testing.T) {

	var tests = []TestCase{
		{
			inputs: [2]int{5, 24},
			expect: "TUE",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, solution(test.inputs[0], test.inputs[1]))
	}
}

// 2016년 1월 1일은 금요일 (SUN,MON,TUE,WED,THU,FRI,SAT)
// 윤년이므로 366일이라고 가정한다.

var cache = map[int]string{
	1: "SUN",
}

func solution(a int, b int) string {
	if a == 1 {
		if b == 1 {
			return cache[a]
		} else {
			return getDateString((b - 1) % 7)
		}
	} else {
		//2월 ~ 12월!

		//구하려는 일자의 달의 이전 달 마지막 날짜에 대한 요일 정보를 조회한다.
		prevMonthLastDateString := solution(a-1, getDaysInMonth(a-1))
		currentMonthFirstDateString := getNextDateString(prevMonthLastDateString)

		if b == 1 {
			return currentMonthFirstDateString
		} else {
			return getDateString((b - 1) % 7)
		}
	}
}

/*
1월1일 금요일

1월 27일 => 27 - 1 = 26일 차이남. 26 나누기 7 : 나머지 5임 (8, 15, 22 모두 금요일) 금요일로부터 5일 이후니까 토일월화수. 수요일 .

2월 14일 => 2월 1일이 무슨 요일인지 알아내자. 윤년은 2월이 29일까지임.

1월 마지막 31일임. => 31 - 1 = 30 나누기 7 ...나머지 2임. 금요일로부터 2일후는 일요일.
즉 2월 1일은 월요일. 그러면 2월 14일은 ? 14 - 1 = 13 나누기 7 . 나머지 6임 토일월화수목. 목요일임

3월 28일 => 3월 1일의 요일을 구하자,, => 2월의 마지막 날짜 구하기 => 2월 1일 요일구헤야함
4월 x일 => 4월 1일의 요일을 구하자 => 3월의 마지막 날짜 => 3월 1일 요일 구해야함 => ... 무한반복 ;;

=> 캐시로 쌓아둔다 => 그러면은 각 월에 대한 1일의 요일을 캐시로 저장하자.

1월: 31
2월: 29
3월 31
4월 30
5월 31
6월 30
7월 31
8월 31
9월 30
10월 31
11월 30
12월 31

*/

func getDateString(remainder int) string {
	switch remainder {
	case 0:
		return "FRI"
	case 1:
		return "SAT"
	case 2:
		return "SUN"
	case 3:
		return "MON"
	case 4:
		return "TUE"
	case 5:
		return "WED"
	case 6:
		return "THU"
	}
	return ""
}

func getNextDateString(date string) string {
	switch date {
	case "MON":
		return "TUE"
	case "TUE":
		return "WED"
	case "WED":
		return "THU"
	case "THU":
		return "FRI"
	case "FRI":
		return "SAT"
	case "SAT":
		return "SUN"
	case "SUN":
		return "MON"
	}
	return ""
}

func getDaysInMonth(month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 2:
		return 29
	case 4, 6, 9, 11:
		return 30
	}

	return 0
}
