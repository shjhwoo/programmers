package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution(15))
	fmt.Println(solution(1))
}

func solution(n int) [][]int {
	//
	if n == 1 {
		return [][]int{{1, 3}}
	}

	// if n == 2 {
	// 	return [][]int{{1,2},{1,3},{2,3}}
	// }

	//홀수일때
	//짝수일때
	//배열의길이 = 이동하는 최소횟수
	//이동하는 최소횟수는 2pow n - 1

	answer := [][]int{{}}

	return answer
}

// func isEven(n int) bool {
// 	return n%2 == 0
// }

func buildLeftRoute() [][]int {

}

func buildMiddle(start, end int) []int {
	return []int{start, end}
}

func buildRightRoute() [][]int {

}

/*규칙성

f(1,3) = f(1,2) + (1,3) + f(2,3)
f(1,2) = f(1,3) + (1,2) + f(3,2)
f(2,3) = f(2,1) + (2,3) + f(1,3)
f(3,2) = f(3,1) + (3,2) + f(1,2)
f(2,1) = f(2,3) + (2,1) + f(3,1)

정리하자면
맨 밑의 판을 제외한 그 위의 나머지 판들을 목적지와는 다른 곳으로 이동을 시키고
그다음 맨 밑의 판을 목적지로 이동시키고
다른 곳으로 잠시 이동시켰던 나머지 판들을 다시 목적지로 이동시키는 방법과 같다

f 함수를 호출할 때 이동해야 하는 판의 개수가 1개뿐이면은 LR은 무시한다

즉 f(시작정, 도착점) = f(시작 , 도착아닌곳) + (시작, 도착) + f(도착아닌곳, 도착)
*/
