package main

import (
	"fmt"
)

func main() {
	// fmt.Println(solution(15))
	fmt.Println(solution(3))
}

// 키: start-end-panelNum 이렇게 만들자!!
var cache = map[string][][]int{}

func solution(n int) [][]int {
	return move(1, 3, n)
}

func move(start, end, panelNum int) [][]int {
	k := fmt.Sprintf("%d-%d-%d", start, end, panelNum)
	if panelNum == 1 {
		cache[k] = [][]int{{start, end}}
	}

	if cache[k] != nil {
		return cache[k]
	}

	op := findOtherPillar(start, end)

	left := move(start, op, panelNum-1)
	middle := [][]int{{start, end}}
	right := move(op, end, panelNum-1)

	result := [][]int{}
	result = append(append(append(result, left...), middle...), right...)
	cache[k] = result

	return result
}

func findOtherPillar(start, end int) int {
	p := []int{1, 2, 3}
	for _, item := range p {
		if item != start && item != end {
			return item
		}
	}
	return 0
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

재귀 종료 조건: 이동해야 하는 판의 개수가 한개뿐일때 재귀조건은 종료된다,
*/
