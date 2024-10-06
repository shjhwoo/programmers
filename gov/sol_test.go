package main_test

import (
	"container/heap"
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
		// {
		// 	N: 5,
		// 	road: [][]int{
		// 		{1, 2, 1},
		// 		{2, 3, 3},
		// 		{5, 2, 2},
		// 		{1, 4, 2},
		// 		{5, 3, 1},
		// 		{5, 4, 2},
		// 	},
		// 	k:      3,
		// 	expect: 4,
		// },
		// {
		// 	N: 6,
		// 	road: [][]int{
		// 		{1, 2, 1},
		// 		{1, 3, 2},
		// 		{2, 3, 2},
		// 		{3, 4, 3},
		// 		{3, 5, 2},
		// 		{3, 5, 3},
		// 		{5, 6, 1},
		// 	},
		// 	k:      4,
		// 	expect: 4,
		// },
		// {
		// 	N: 6,
		// 	road: [][]int{
		// 		{1, 2, 1},
		// 		{1, 3, 8},
		// 		{2, 3, 2},
		// 		{3, 4, 3},
		// 		{3, 5, 2},
		// 		{3, 5, 3},
		// 		{5, 6, 1},
		// 	},
		// 	k:      4,
		// 	expect: 3,
		// },
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
		assert.Equal(t, test.expect, solutionWithDijkstra(test.N, test.road, test.k))
	}
}

// 망한거.
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

//최소힙(우선순위 큐(Priority Queue)를 구현하기 위해 사용하는 자료구조 중 하나)을 사용해서 다익스트라 구현하고 적용 :

func solutionWithDijkstra(N int, road [][]int, k int) int {
	dist := Dijkstra(road, N)

	answer := 1 //같은 마을은 무조건 배달 가능하니까 포함시킨다.

	for idx, distance := range dist {
		//만약에, distance가 여전히 무한의 값이라면 그건 1번 지점에서 도달을 할 수 없다는 의미(방문가능한 경로 없다)
		if idx > 1 && distance <= k {
			answer++
		}
	}

	return answer
}

// 문제풀이에 적용 가능한 이유는: 특정지점까지의 최소 거리는 각 중간경유지까지의 최소거리를 먼저 구했기 때문에 이들을 합하면 가장 최소가 된다는 것.
/*
다익스트라 최단 경로 알고리즘은 그리디 알고리즘으로 분류됩니다.

 ; 매 상황에서 가장 비용이 적은 노드를 선택해 임의의 과정을 반복합니다.
*/
func Dijkstra(road [][]int, N int) []int {
	vheap := &VillageHeap{}
	heap.Init(vheap)
	heap.Push(vheap, Village{
		Name: 1,
		Cost: 0,
	})

	infiniteInt := math.MaxInt64
	var shortestDistance = []int{0, 0}
	for i := 2; i < N+1; i++ {
		shortestDistance = append(shortestDistance, infiniteInt)
	}

	for vheap.Len() > 0 {
		lowest := vheap.Pop()
		currentVill := lowest.(Village).Name
		currentCost := lowest.(Village).Cost

		for _, r := range road {
			start := r[0]
			dest := r[1]
			distance := r[2]

			nextCost := distance + currentCost

			if start == currentVill && nextCost < shortestDistance[dest] {
				shortestDistance[dest] = nextCost
				vheap.Push(Village{
					Name: dest,
					Cost: nextCost,
				})
			} else if dest == currentVill && nextCost < shortestDistance[start] {
				shortestDistance[start] = nextCost
				vheap.Push(Village{
					Name: start,
					Cost: nextCost,
				})
			}
		}
	}

	return shortestDistance
}

//최소힙부터 구현

type VillageHeap []Village

func (vh VillageHeap) Len() int {
	return len(vh)
}

func (vh VillageHeap) Less(i, j int) bool {
	return vh[i].Cost < vh[j].Cost
}

func (vh VillageHeap) Swap(i, j int) {
	vh[i], vh[j] = vh[j], vh[i]
}

// cap - len 에 여유공간 부족한데 새 요초 추가하면 새로운 배열 생성되므로 포인터 메서드로 선언해야 한다
func (vh *VillageHeap) Push(newNode any) {
	nv, ok := newNode.(Village)
	if ok {
		*vh = append(*vh, nv)
	}
}

func (vh *VillageHeap) Pop() any {

	if vh.IsEmpty() {
		return nil
	}

	old := *vh
	n := len(old)
	elem := old[n-1]
	*vh = old[0 : n-1]

	return elem
}

func (vh VillageHeap) IsEmpty() bool {
	return len(vh) == 0
}

type Village struct {
	Name int
	Cost int
}

//직접 최소힙을 구현하기

//힙은 배열

type MyMinHeap []MyNode

type MyNode struct {
	Value int
}

func NewMyMinHeap() MyMinHeap {
	return MyMinHeap{
		{
			Value: 0,
		},
	}
}

func (mmh *MyMinHeap) Len() int {
	return len(*mmh)
}

func (mmh *MyMinHeap) Push(newNode MyNode) {

	add := mmh
	fmt.Println("주소: ", add)

	*mmh = append(*mmh, newNode)

	currentIdx := mmh.Len() - 1
	parentIdx := currentIdx / 2

	for parentIdx != 0 && (*mmh)[parentIdx].Value > newNode.Value {
		temp := (*mmh)[parentIdx]
		(*mmh)[parentIdx] = newNode
		(*mmh)[currentIdx] = temp

		currentIdx = parentIdx
		parentIdx = currentIdx / 2
	}
}

func (mmh *MyMinHeap) Pop() MyNode {
	//뺄 값은 미리 복사를 해둔다
	returnVal := (*mmh)[1]

	//1번과 마지막 원소 자리를 바꾸기(뺄거를 맨 뒤로 해줘가지고 뺀다.) : 이렇게 해줘야 다음 뺄 값을 위로 올릴 수 있기 때문이다.
	temp := (*mmh)[mmh.Len()-1]

	//
	if mmh.Len() <= 3 && returnVal.Value >= temp.Value {
		returnVal = temp
	} else {
		(*mmh)[mmh.Len()-1] = (*mmh)[1]
		(*mmh)[1] = temp
	}

	(*mmh) = (*mmh)[:mmh.Len()-1]

	currentIdx := 1
	leftIdx := 2
	rightIdx := 3

	//go에서 직접 구현한거는 인덱스 경계값 주의하기
	for mmh.Len() > leftIdx && mmh.Len() > rightIdx && ((*mmh)[currentIdx].Value > (*mmh)[leftIdx].Value || (*mmh)[currentIdx].Value > (*mmh)[rightIdx].Value) {
		if (*mmh)[leftIdx].Value > (*mmh)[rightIdx].Value {
			temp := (*mmh)[currentIdx]
			(*mmh)[currentIdx] = (*mmh)[rightIdx]
			(*mmh)[rightIdx] = temp
			currentIdx = rightIdx
		} else {
			temp := (*mmh)[currentIdx]
			(*mmh)[currentIdx] = (*mmh)[leftIdx]
			(*mmh)[leftIdx] = temp
			currentIdx = leftIdx
		}

		leftIdx = currentIdx * 2
		rightIdx = currentIdx*2 + 1
	}

	return returnVal

}

func TestMyMinHeap(t *testing.T) {

	myHeap := NewMyMinHeap()

	myHeap.Push(MyNode{
		Value: 70,
	})

	myHeap.Push(MyNode{
		Value: 11,
	})

	myHeap.Push(MyNode{
		Value: 30,
	})

	myHeap.Push(MyNode{
		Value: 50,
	})

	myHeap.Push(MyNode{
		Value: 20,
	})

	myHeap.Push(MyNode{
		Value: 89,
	})

	myHeap.Push(MyNode{
		Value: 17,
	})

	myHeap.Push(MyNode{
		Value: 26,
	})

	t.Logf("myHeap:: %v", myHeap)

	t.Logf("myHeap 1st pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 2nd pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 3rd pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 4th pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 6th pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 7th pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 8th pop: %v, heap status: %v", myHeap.Pop(), myHeap)

	t.Logf("myHeap 9th pop: %v, heap status: %v", myHeap.Pop(), myHeap)
}

/*
append: cap - len 여유 공간이 있다면 같은 주소값에 해당하는 배열을 변경
그렇지 않고 공간이 모자라는 경우 cap을 2배(일반적으로는) 늘린 새로운 배열을 만든다.

슬라이스의 슬라이싱:
*/

func SliceSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:3]

	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2)) //cap = 배열의 총 길이 - 시작 인덱스 (5-1 = 4)

	s1[2] = 1

	fmt.Println("after change::")
	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2))

	s3 := make([]int, 10)
	copy(s3, s1)

	fmt.Println("s3", s3)

	//불필요한 메모리 사용 안하도록, 아래와 같이 중간에 새로운 요소를 추가하는 방법이있다.
	s4 := []int{1, 2, 3, 4, 5}
	s4 = append(s4, 0)

	copy(s4[2+1:], s4[2:])

	fmt.Println("s4:", s4)

	s4[2] = 100

	fmt.Println("s4:", s4)
}

func TestSliceSlice(t *testing.T) {
	SliceSlice()
}
