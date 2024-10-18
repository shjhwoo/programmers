package main_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	n      int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			n:      4,
			expect: 2,
		},
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.n))
	}
}

func solution(n int) int {

	//처음 점 0,0부터 시작해야 한다
	//특정한 조건을 만족하는 것만 탐색하고 나머지는 탐색하지 않도록 조건문 작성
	//즉 절대로 답이 될 수 없는 것은 탐색을 종료한다.

	var candidates = make(map[string]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			startPos := []int{i, j}

			qPosList := getQPosList(startPos, n)

			if len(qPosList) == n {
				sort.Slice(qPosList, func(i, j int) bool {
					return qPosList[i][0] < qPosList[j][0]
				})

				posKey := makePosKey(qPosList)

				if _, exist := candidates[posKey]; !exist {
					candidates[posKey] = true
				}
			}

		}
	}

	return len(candidates)
}

func getQPosList(startPos []int, n int) [][]int {
	//계속 반복하면서 탐색!!!
	chessBoard := getChessBoard(n)

	var result = [][]int{startPos}

	queue := [][]int{startPos}

	for len(queue) > 0 {

		startRow := queue[0][0]
		startCol := queue[0][1]
		queue = queue[1:]

		//빼낸 지점에서부터 탐색 가능한 다음의 지점들을 넣는다.

		var firstFound bool
		var getResult bool
		for r, row := range chessBoard {
			for c, col := range row {
				if col == false {
					continue
				}

				if r == startRow || c == startCol || isOnDiagnoalLine(startRow, startCol, r, c) { // 피해야 하는 조건을 명시한다
					chessBoard[r][c] = false
					continue
				}

				if !firstFound {
					firstFound = true
					result = append(result, []int{r, c})

					if len(result) == n {
						getResult = true
						break
					}
				}

				if len(queue) == 0 {
					queue = append(queue, []int{r, c})
				}
			}

			if getResult {
				break
			}
		}
	}

	return result
}

func getChessBoard(n int) [][]bool {
	var board [][]bool

	for i := 0; i < n; i++ {
		line := []bool{}
		for j := 0; j < n; j++ {
			line = append(line, true)
		}

		board = append(board, line)
	}

	return board
}

func isOnDiagnoalLine(startR, startC, curR, curC int) bool {
	isNegativeD := float32(startR-curR)/float32(startC-curC) == float32(-1)

	isPositiveD := float32(startR-curR)/float32(startC-curC) == float32(1)

	return isNegativeD || isPositiveD
}

func makePosKey(qPosList [][]int) string {
	var result []string

	for _, qpos := range qPosList {
		result = append(result, fmt.Sprintf("(%d,%d)", qpos[0], qpos[1]))
	}

	return strings.Join(result, "-")
}
