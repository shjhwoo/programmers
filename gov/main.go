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

func isValidSpot(maps [][]string, row, col int) bool {
	return row < len(maps) && col < len(maps[0]) && row >= 0 && col >= 0 && maps[row][col] != "0"
}

func solution(maps []string) []int {
	//주어진 맵에 전부다 X밖에 없는 상황이면은..
	if hasNoIsland(maps) {
		return []int{-1}
	}

	//그 외의 상황에서는 각 섬에서 최대 며칠 머무를 수 있는지 알아야 한다.

	//map을 쪼개준다
	pixmap := getPixeledMap(maps)

	//일단은 섬이 어디 있는지를 알랴면 무조건 첫번쨰줄부터 한칸씩 찾아야 한다.
	//X가 아닌 칸을 찾았으면 거기서부터 상하좌우 탐색하고, 숫자가 나오면 더해준다.
	//그리고 그 칸을 0으로 바꾼다. 이미 방문했다는 표시로.
	//그리고 식량정보는 어딘가에 저장을 해야 함.
	result := []int{}

	food := 0
	for r, row := range pixmap {
		for c, col := range row {
			if col != "X" && col != "0" {
				//0이 아닌 숫자에 대해서만 탐색을 시작한다. 0은 이미 방문해서 식량을 다 소진했다는 뜻.
				//문제는 하나의 섬이라는 걸 어떻게 알아야 하는지..
				//더 이상 탐색할 수 있는, 식량이 없다고 판단되었을 때 섬 전체를 다 돌았다고 판단하고,
				//그 합산한 결과를 result에 푸시해주는 방식으로 가야할거같은덴
				//그런 다음에 다음 섬을 찾아야할거같아.
				colInt, _ := strconv.Atoi(col)
				food += colInt

				pixmap[r][c] = "0" //방문했다.

				//상하좌우 탐색
				for _, d := range directions {
					newSpot := []int{r + d[0], c + d[1]}
					if isValidSpot(pixmap, newSpot[0], newSpot[1]) {
						//식량값 더한다
						//방문표기한다
						colInt, _ := strconv.Atoi(pixmap[newSpot[0]][newSpot[1]])
						food += colInt
						pixmap[newSpot[0]][newSpot[1]] = "0"

						//여기서 다시 상하좌우 탐색을 들어간다...재귀
					}
				}
				//재귀 종료 조건) 4방향 모두 다 탐색했는데 더 이상 갈 곳이 없다.
				//DFS 최대한 갈 수 있는 곳 다 탐색.
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

func getPixeledMap(maps []string) [][]string {
	result := [][]string{}
	for _, row := range maps {
		newRow := []string{}
		for _, col := range strings.Split(row, "") {
			newRow = append(newRow, col)
		}
		result = append(result, newRow)
	}

	return result
}
