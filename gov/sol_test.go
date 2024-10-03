package main_test

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/v3/assert"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func setTree() *Node {
	tree := &Node{
		Value: 1,
	}

	tree.Left = NewNode(2)
	tree.Right = NewNode(3)

	tree.Left.Left = NewNode(4)
	tree.Left.Right = NewNode(5)

	tree.Right.Left = NewNode(6)
	tree.Right.Right = NewNode(7)

	tree.Right.Right.Left = NewNode(8)
	tree.Right.Right.Right = NewNode(9)

	return tree
}

func (n *Node) PrintPreOrder() []int {
	mid := []int{n.Value}

	if n.Left == nil && n.Right == nil {
		return mid
	}

	lefts := n.Left.PrintPreOrder()
	rights := n.Right.PrintPreOrder()

	return append(append(mid, lefts...), rights...)
}

func (n *Node) PrintInOrder() []int {
	mid := []int{n.Value}
	if n.Left == nil && n.Right == nil {
		return mid
	}

	lefts := n.Left.PrintInOrder()
	rights := n.Right.PrintInOrder()

	return append(append(lefts, mid...), rights...)
}

func (n *Node) PrintPostOrder() []int {
	mid := []int{n.Value}
	if n.Left == nil && n.Right == nil {
		return mid
	}

	lefts := n.Left.PrintPostOrder()
	rights := n.Right.PrintPostOrder()

	return append(append(lefts, rights...), mid...)
}

func TestPreOrderTree(t *testing.T) {
	tree := setTree()

	preOrder := tree.PrintPreOrder()

	assert.DeepEqual(t, []int{1, 2, 4, 5, 3, 6, 7, 8, 9}, preOrder)
}

func TestInOrder(t *testing.T) {
	tree := setTree()

	inOrder := tree.PrintInOrder()

	assert.DeepEqual(t, []int{4, 2, 5, 1, 6, 3, 8, 7, 9}, inOrder)

}

func TestPostOrder(t *testing.T) {
	tree := setTree()

	inOrder := tree.PrintPostOrder()

	assert.DeepEqual(t, []int{4, 5, 2, 6, 8, 9, 7, 3, 1}, inOrder)

}

//순열과 조합
//성능, 콜 스택의 위험성 때문에 성능상 스택으로 구현하는 것이 좋다.
//하지만 코테에서는 시간복잡도 자체가 크기 때문에 코테에서는 N이 크게 나오는 경우는 드물다
//이에 재귀로 외우는게 직관적이고 편하다.

/*
순열이란?

"서로 다른" N개의 대상 중에서 M개를 골라 나열하는 방법을 구하는 것

시간복잡도는 O(N!) 이 된다. 매우크다
*/
func Permutations(arr []int, picks int) int {
	//탈출 조건
	if picks == 1 {
		return len(arr)
	}

	//재귀
	return Permutations(arr[:len(arr)-1], picks-1) * len(arr)
}

func TestPermutations(t *testing.T) {
	assert.Equal(t, 24, Permutations([]int{4, 3, 2, 1}, 3))
	assert.Equal(t, 12, Permutations([]int{4, 3, 2, 1}, 2))
}

// 순열 실제 경우를 프린트아웃하는 함수
func PrintPermutationCases(arr []int, n int) [][]int {
	if n == 1 {
		var result [][]int
		for _, num := range arr {
			result = append(result, []int{num})
		}

		return result
	}

	var result [][]int

	for idx, num := range arr {
		var subArr []int
		for i := 0; i < len(arr); i++ {
			if i == idx {
				continue
			}
			subArr = append(subArr, arr[i])
		}

		pre := PrintPermutationCases(subArr, n-1)

		for _, c := range pre {
			result = append(result, append(c, num))
		}
	}

	return result
}

func TestPrintPermutationCases(t *testing.T) {
	cases := PrintPermutationCases([]int{4, 3, 2, 1}, 3)

	t.Log(cases)

	assert.DeepEqual(t, 24, len(cases))

	for _, c := range cases {
		assert.Equal(t, 3, len(c))
	}
}

/*
조합이란?

서로 다른 N개의 대상 중에서 M개를 고르는 방법의 수를 구하는 것
순서가 중요하지 않다.
*/
func Combinations(arr []int, n int) int {
	//탈출 조건
	if n == 1 {
		return len(arr)
	}

	var cases int

	for idx, _ := range arr {
		subArr := arr[idx+1:]
		subCases := Combinations(subArr, n-1)
		cases += subCases
	}

	//재귀 조건
	return cases //Combinations(arr[:len(arr)-1], n-1) * len(arr) / n
}

func TestCombinations(t *testing.T) {
	assert.Equal(t, 4, Combinations([]int{4, 3, 2, 1}, 3))
	assert.Equal(t, 6, Combinations([]int{4, 3, 2, 1}, 2))
}

// 조합 실제 경우를 프린트아웃하는 함수
func PrintCombinationCases(arr []int, n int) [][]int {
	//탈출 조건은 동일하다
	if n == 1 {
		var result [][]int
		for _, num := range arr {
			result = append(result, []int{num})
		}

		return result
	}

	var result [][]int

	for idx, num := range arr {
		var subArr = arr[idx+1:] //중간 요소만 빼는게 아니고 이후 요소부터 고르는 이유는 조합이기 때문. 앞전 요소를 다시 고를 이유가 없다

		pre := PrintCombinationCases(subArr, n-1)

		for _, c := range pre {
			result = append(result, append(c, num))
		}
	}

	return result
}

func TestPrintCombinationCases(t *testing.T) {
	cases := PrintCombinationCases([]int{4, 3, 2, 1}, 3)

	t.Log(cases)

	assert.DeepEqual(t, 4, len(cases))

	for _, c := range cases {
		assert.Equal(t, 3, len(c))
	}
}

/*

다익스트라 알고리즘을 적용한 코테 문제
*/

type TestCase struct {
	N      int
	road   [][]int
	k      int
	expect int
}

func TestCountDeliveryAbleTown(t *testing.T) {
	tests := []TestCase{
		{
			N: 5,
			road: [][]int{
				{1, 2, 1},
				{2, 3, 3},
				{5, 2, 2},
				{1, 4, 2},
				{5, 3, 1},
				{5, 4, 2},
			},
			k:      3,
			expect: 4,
		},
		{
			N: 6,
			road: [][]int{
				{1, 2, 1},
				{1, 3, 2},
				{2, 3, 2},
				{3, 4, 3},
				{3, 5, 2},
				{3, 5, 3},
				{5, 6, 1},
			},
			k:      4,
			expect: 4,
		},
		{
			N: 6,
			road: [][]int{
				{1, 2, 1},
				{1, 3, 8},
				{2, 3, 2},
				{3, 4, 3},
				{3, 5, 2},
				{3, 5, 3},
				{5, 6, 1},
			},
			k:      4,
			expect: 3,
		},
		{
			N: 6,
			road: [][]int{
				{1, 2, 29},
				{1, 6, 2},
				{2, 3, 2},
				{3, 4, 3},
				{4, 5, 7},
				{5, 6, 1},
			},
			k:      20,
			expect: 6,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, solution(test.N, test.road, test.k))
	}
}

func solution(N int, road [][]int, k int) int {
	var isVisitedMap = make(map[string]bool)

	infiniteInt := math.MaxInt64
	var shortestDistance = []int{0, 0}
	for i := 2; i < N+1; i++ {
		shortestDistance = append(shortestDistance, infiniteInt)
	}

	var queue = []int{1}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		for _, route := range road {
			start := route[0]
			dest := route[1]
			distance := route[2]

			if start == first {
				visitKey := fmt.Sprintf("%d_%d_%d", start, dest, distance)
				if isVisitedMap[visitKey] {
					continue
				}

				//1번 정점에서 dest까지 가는 거리를 계산
				distanceToDest := distance + shortestDistance[start]

				//더 짧은 경로가 있다면 업데이트.
				if distanceToDest < shortestDistance[dest] {
					queue = append(queue, dest)
					isVisitedMap[visitKey] = true
					shortestDistance[dest] = distanceToDest
				}
			} else if dest == first {
				visitKey := fmt.Sprintf("%d_%d_%d", dest, start, distance)
				if isVisitedMap[visitKey] {
					continue
				}

				//1번 정점에서 dest까지 가는 거리를 계산
				distanceToDest := distance + shortestDistance[dest]

				//더 짧은 경로가 있다면 업데이트.
				if distanceToDest < shortestDistance[start] {
					queue = append(queue, start)
					isVisitedMap[visitKey] = true
					shortestDistance[start] = distanceToDest
				}
			}
		}
	}

	answer := 1 //같은 마을은 무조건 배달 가능하니까 포함시킨다.

	for idx, distance := range shortestDistance {
		if idx > 1 && distance <= k {
			answer++
		}
	}

	return answer
}
