package main_test

import (
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {

}

type Node struct {
	Connection []int
	IsCycle    bool
}

func solution(n int, costs [][]int) int {
	/*
		1. 일단 제일 작은 간선 값을 가지는 경우가 앞에 오도록 정렬한다.

		2. 1에서 정렬한 배열의 0번째 인덱스 요소부터 탐색을한다 => 서로소 집합을 만든다.

	*/

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] > costs[j][2]
	})

	var nodes []Node
	for _, cost := range costs {
		nodes = append(nodes, Node{
			Connection: cost,
		})
	}

	//일단은 자기자신으로 초기화
	parentMap := make(map[int]int)
	for i := 0; i < n; i++ {
		parentMap[i] = i
	}

	for _, node := range nodes {
		start := node.Connection[0]
		dest := node.Connection[1]

		//부모 찾기(재귀적으로-- 자기자신이 부모면 탈출)

	}

	return 0
}
