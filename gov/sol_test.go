package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	inputS []int
	expect string
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			inputS: []int{1, 3, 4, 6},
			expect: "1223330333221",
		},
		{
			inputS: []int{1, 7, 1, 2},
			expect: "111303111",
		},
	}

	for _, test := range tests {
		ans := solution(test.inputS)

		t.Log(ans, "계산값")

		assert.Equal(t, test.expect, ans)
	}
}

func solution(food []int) string {
	var player1 string

	for idx, fo := range food {
		if idx == 0 || fo == 1 {
			continue
		}

		repeat := fo / 2
		var i int
		for i < repeat {
			player1 += fmt.Sprintf("%d", idx)
			i++
		}
	}

	player2 := getReversedString(player1)

	return fmt.Sprintf("%s0%s", player1, player2)
}

func getReversedString(pst string) string {
	var answer string
	for i := len(pst) - 1; i > -1; i-- {
		answer += pst[i : i+1]
	}

	return answer
}
