package main_test

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

type Case struct {
	n      int
	costs  [][]int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []Case{
		{
			n: 4,
			costs: [][]int{
				{0, 1, 1},
				{0, 2, 2},
				{1, 2, 5}, //사이클 발생
				{1, 3, 1},
				{2, 3, 8},
			},
			expect: 4,
		},
	}

	for _, test := range tests {
		assert.DeepEqual(t, test.expect, solution(test.n, test.costs))
	}
}

type Node struct {
	Connection []int
	IsCycle    bool
}

func solution(n int, costs [][]int) int {

	///간선을 가중치 기준 오름차순 정렬. 그래야 최소를 계산가능
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	//부모 정보를 저장할 슬라이스선언.
	var parentInfo []int
	for i := 0; i < n; i++ {
		parentInfo = append(parentInfo, i)
	}

	var answer int
	for _, cost := range costs {
		if !compare(parentInfo, cost[0], cost[1]) { //두 원소의 부모가 같으면 사이클
			answer += cost[2] //같지 않은 경우에만 하나의 집합으로 만들수있다 (결국 하나의 부모로 귀결된다는 의미)
			union(parentInfo, cost[0], cost[1])
		}
	}

	return answer
}

func compare(parentInfo []int, start int, dest int) bool {
	sp := findParent(start, parentInfo)
	dp := findParent(dest, parentInfo)

	return sp == dp
}

func union(parentInfo []int, start int, dest int) {
	start = findParent(start, parentInfo)
	dest = findParent(dest, parentInfo)

	if start < dest {
		parentInfo[dest] = start
	} else {
		parentInfo[start] = dest
	}
}

func findParent(node int, parentInfo []int) int {
	if parentInfo[node] == node {
		return node
	}

	parentInfo[node] = findParent(parentInfo[node], parentInfo)
	return parentInfo[node]
}
