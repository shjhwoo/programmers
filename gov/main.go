package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution(15))
	fmt.Println(solution(1))
}

func solution(n int) [][]int {
	// if n == 1 {
	// 	return [][]int{{1, 3}}
	// }

	// if n == 2 {
	// 	return [][]int{{1,2},{1,3},{2,3}}
	// }

	//홀수일때
	//짝수일때
	//배열의길이 = 이동하는 최소횟수
	//이동하는 최소횟수는 2pow n - 1

	answer := [][]int{{}}
	if isEven(n) {
		answer = append(answer, []int{1, 2})
	} else {
		answer = append(answer, []int{1, 3})
	}

	movNum := math.Pow(2, float64(n)) - 1
	for len(answer) < int(movNum) {

	}

	return answer
}

func isEven(n int) bool {
	return n%2 == 0
}

/*
1,3
1,2 -> 1,3 -> 2,3
1,3 -> 1,2 -> 3,2 -> 1,3 -> 2,1 -> 2,3 -> 1,3 : (1,2) 3 이 2개의 원판을 1-> 3으로 이동하는 방법임
1,2 -> 1,3 -> 2,3 -> 1,2 -> 3,1 -> 3,2 -> 1,2 -> 1,3 -> 2,3 -> 2,1 -> 3,1 -> 2,3 -> 1,2 -> 1,3 -> 2,3
1,3 -> 1,2 -> 3,2 -> 1,3 -> 2,1 -> 2,3 -> 1,3 -> ...
*/
