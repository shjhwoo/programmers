package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var result = Solution([]string{"XXX", "XXX", "XXX", "XXX"})
	fmt.Println("최종 결과값", result)

	var res2 = Solution([]string{"X591X", "X1X5X", "X231X", "1XXX1"})
	fmt.Println(res2, "두번째 테스트")
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Solution(maps []string) []int {
	//주어진 맵에 전부다 X밖에 없는 상황이면은..
	if hasNoIsland(maps) {
		return []int{-1}
	}

	//그 외의 상황에서는 각 섬에서 최대 며칠 머무를 수 있는지 알아야 한다.

	//map을 쪼개준다
	pixmap := getPixeledMap(maps)
	result := []int{}

	for r, row := range pixmap {
		for c, col := range row {
			//탐색가능한 지점이다(완전히 새로운 섬이다!!)
			if col != "X" && col != "0" {
				food, _ := strconv.Atoi(col)
				func() {
					var dfs func(int, int)
					dfs = func(startRow, startCol int) {
						pixmap[r][c] = "0"
						for _, d := range directions {
							newSpot := []int{r + d[0], r + d[1]}
							if isValidSpot(pixmap, newSpot[0], newSpot[1]) {
								colInt, _ := strconv.Atoi(pixmap[newSpot[0]][newSpot[1]])
								food += colInt
								pixmap[newSpot[0]][newSpot[1]] = "0"
								dfs(newSpot[0], newSpot[1])
							}
						}
					}
					dfs(r, c)
				}()
				result = append(result, food)
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
		newRow := append([]string{}, strings.Split(row, "")...)
		result = append(result, newRow)
	}

	return result
}

func isValidSpot(maps [][]string, row, col int) bool {
	return row < len(maps) && col < len(maps[0]) && row >= 0 && col >= 0 && maps[row][col] != "0"
}
