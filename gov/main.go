package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var result = solution([]string{"XXX", "XXX", "XXX", "XXX"})
	fmt.Println("최종 결과값", result)
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isValidSpot(maps []string, row, col int) bool {
	return row < len(maps) && col < len(maps[0]) && row >= 0 && col >= 0
}

func solution(maps []string) []int {
	//주어진 맵에 전부다 X밖에 없는 상황이면은..
	if hasNoIsland(maps) {
		return []int{-1}
	}

	//그 외의 상황에서는 각 섬에서 최대 며칠 머무를 수 있는지 알아야 한다.
	//일단은.. 먼저 섬이 몇개인지를 알아야 할까?
	//그런 다음에 각 섬에 대해서 얻을 수 있는 식량을 알아야 할지도 .

	//일단은 섬이 어디 있는지를 알랴면 무조건 첫번쨰줄부터 한칸씩 찾아야 한다.
	//X가 아닌 칸을 찾았으면 거기서부터 상하좌우 탐색하고, 숫자가 나오면 더해준다. 그리고 그 칸을 0으로 바꾼다. 이미 방문했다는 표시로.
	//그리고 식량정보는 어딘가에 저장을 해야 함.
	result := []int{}

	food := 0
	for _, row := range maps {
		cols := strings.Split(row, "")
		for _, col := range cols {
			if col != "X" && col != "0" {
				//0이 아닌 숫자에 대해서만 탐색을 시작한다. 0은 이미 방문해서 식량을 다 소진했다는 뜻.
				//문제는 하나의 섬이라는 걸 어떻게 알아야 하는지..
				//더 이상 탐색할 수 있는, 식량이 없다고 판단되었을 때 섬 전체를 다 돌았다고 판단하고,
				//그 합산한 결과를 result에 푸시해주는 방식으로 가야할거같은덴
				//그런 다음에 다음 섬을 찾아야할거같아.
				colInt, _ := strconv.Atoi(col)
				food += colInt

				spot := []int{row, col}

				//상하좌우 탐색
				for i := 0; i < 4; i++ {

				}
			}
		}
	}

	return result
}

func hasNoIsland(maps []string) bool {
	for _, row := range maps {
		if !isAllX(row) {
			return false
		}
	}
	return true
}

func isAllX(row string) bool {
	pList := strings.Split(row, "")

	for _, p := range pList {
		if p != "X" {
			return false
		}
	}
	return true
}
