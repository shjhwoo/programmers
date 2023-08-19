package main

import "fmt"

func main() {
	var result = Solution([][]int{{1, 4}, {3, 2}, {4, 1}}, [][]int{{3, 3}, {3, 3}})
	fmt.Println("최종 결과값", result)

	var result2 = Solution([][]int{{2, 3, 2}, {4, 2, 4}, {3, 1, 4}}, [][]int{{5, 4, 3}, {2, 4, 1}, {3, 1, 1}})
	fmt.Println("최종 결과값", result2)
}

func Solution(arr1 [][]int, arr2 [][]int) [][]int {
	result := [][]int{}
	for i := range arr1 {
		resRow := []int{}
		for h := 0; h < len(arr2[0]); h++ {
			sum := 0
			for j := range arr1[i] {
				sum += arr1[i][j] * arr2[j][h]
				//fmt.Println(i, "번째 줄일때: ", i, j, j, h, arr1[i][j], arr2[j][h])
			}
			resRow = append(resRow, sum)
		}
		result = append(result, resRow)
	}
	return result
}

/*
arr1의 원소 개수만큼 행이 생기고
arr2에 있는 첫번째 원소의 원소개수만큼 열이생긴다

1차 반복>
arr1에 있는 원소 개수만큼 반복

2차반복> arr1에 있는 첫번째 원소가 가진 원소 개수만큼 반복.
	arr2의 원소의 각 n번째 요소랑 각각 곱한뒤 더해준다!
*/
