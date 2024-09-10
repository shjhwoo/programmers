package main_test

import (
	"math/rand"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

//BFS, DFS 가장 먼 노드 실습하기

type GraphTestCase struct {
	n      int
	vertex [][]int
	expect int
}

func TestSolution(t *testing.T) {
	var tests = []GraphTestCase{
		{
			n:      6,
			vertex: [][]int{{3, 6}, {4, 3}, {3, 2}, {1, 3}, {1, 2}, {2, 4}, {5, 2}},
			expect: 3,
		},
	}

	for _, test := range tests {
		ans := solution(test.n, test.vertex)
		assert.DeepEqual(t, test.expect, ans)
	}
}

func solution(n int, edge [][]int) int {
	graph := buildGraph(n, edge)
	distanceFromFirstMap := make(map[int]int)
	for i := 2; i < n+1; i++ {
		distance := getDistanceFrom1st(i, graph)
		distanceFromFirstMap[i] = distance
	}

	var maxDist int
	for _, distance := range distanceFromFirstMap {
		if maxDist < distance {
			maxDist = distance
		}
	}

	var cnt int
	for _, distance := range distanceFromFirstMap {
		if distance == maxDist {
			cnt++
		}
	}

	return cnt
}

// 2차원 그래프를 만들어 줘야 한다.
func buildGraph(n int, edge [][]int) [][]int {
	var result [][]int
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		result = append(result, row)
	}

	for _, rel := range edge {
		row := rel[0] - 1
		col := rel[1] - 1

		result[row][col] = 1
		result[col][row] = 1 //양방향 이므로.
	}

	return result
}

func getDistanceFrom1st(node int, graph [][]int) int {
	//BFS
	var queue []int
	queue = append(queue, 1) //일단 넣는다
	isVisited := map[int]bool{
		1: true,
	}

	var edges int

	for len(queue) > 0 { //큐에 뭐가 있는 동안은 반복을 해야 한다.
		firstNode := queue[0]
		queue = queue[1:]

		//이 정점에서 갈 수 있는 이웃한 정점을 구해서 싹 다 큐에 넣는다
		for i, row := range graph {
			for j, col := range row {
				if i+1 == firstNode && col == 1 && !isVisited[j+1] {
					//j 정점을 방문한 것으로 간주.
					isVisited[j+1] = true
					queue = append(queue, j+1)
					//여기서더하면안됨,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,edges++

					//만약에 방문한 정점이 목표로 하는 정점노드라면..
					if j+1 == node {
						return edges
					}
				}
			}
		}
	}

	return edges
}

// 여러가지 정렬 알고리즘을 GO로 구현해보자
type TestCase struct {
	numbers []int
	expect  []int
}

func TestBubbleSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3},
			expect:  []int{1, 3, 4, 5, 7},
		},
	}

	for _, test := range tests {
		ans := BubbleSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ { //총 배열의 길이 - 1만큼 반복
		for j := 0; j < len(numbers)-1-i; j++ {
			left := numbers[j]
			right := numbers[j+1]
			if left > right {
				numbers[j] = right
				numbers[j+1] = left
			}
		}
	}

	return numbers
}

func TestSelectSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3},
			expect:  []int{7, 5, 4, 3, 1},
		},
	}

	for _, test := range tests {
		ans := SelectSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func SelectSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		rmi := getHighestRemainderIdx(numbers[i+1:]) + i + 1

		if numbers[i] < numbers[rmi] {
			left := numbers[i]
			right := numbers[rmi]

			numbers[i] = right
			numbers[rmi] = left
		}
	}

	return numbers
}

func getHighestRemainderIdx(rmNumbers []int) int {
	var maxNum int
	var resIdx int
	for idx, num := range rmNumbers {
		if maxNum < num {
			maxNum = num
			resIdx = idx
		}
	}

	return resIdx
}

func TestInsertSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3},
			expect:  []int{1, 3, 4, 5, 7},
		},
	}

	for _, test := range tests {
		ans := InsertSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func InsertSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ { //배열의 요소가 5개라면, 4번만큼 반복할 수 있다.
		/*
			i == 1일때는 0과 비교
			i == 2일때는 1 => 1  vs 0
			...
			3 vs 2 , 2 vs 1 , 1 vs 0
			4 vs 3 , 3 vs 2 , 2 vs 1 , 1 vs 0

		*/

		for j := i; j >= 1; j-- {
			left := numbers[j-1]
			right := numbers[j]

			if left > right {
				numbers[j] = left
				numbers[j-1] = right
			}
		}
	}

	return numbers
}

func TestMergeSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3, 6, 11, 9, 20, 10},
			expect:  []int{1, 3, 4, 5, 6, 7, 9, 10, 11, 20},
		},
	}

	for _, test := range tests {
		ans := MergeSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

// MergeSort recursively sorts a slice of integers using merge sort algorithm
func MergeSort(arr []int) []int {
	// Base case: if the slice has 1 or 0 elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle index
	mid := len(arr) / 2

	//호출 스택 때문에 합쳐진걸 다시 나눌 일은 없음.
	//합쳐진 결과값은 아래 머지 함수에서만 쓰임

	// Recursively sort both halves
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge merges two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	// Create a result slice to store the merged values
	result := []int{}

	// Indices for left and right slices
	i, j := 0, 0

	// Merge while there are elements in both slices
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left or right slice (이미 정렬이 되어 있음.. 그리고 나뉘어져 있더라도 길이 차이가 1 차이이거나 길이가 같음.)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func TestQuickSort(t *testing.T) {
	var tests = []TestCase{
		{
			numbers: []int{7, 4, 5, 1, 3, 6, 11, 9, 20, 10},
			expect:  []int{1, 3, 4, 5, 6, 7, 9, 10, 11, 20},
		},
	}

	for _, test := range tests {
		ans := QuickSort(test.numbers)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIdx := len(arr) / 2

	left, right := getLeftAndRight(pivotIdx, arr)

	newLeft := QuickSort(left)
	newRight := QuickSort(right)

	return append(append(newLeft, arr[pivotIdx]), newRight...)
}

/*
매번 퀵 정렬이 호출될 때마다 고르는 피벗의 위치가 달라진다면?

원래 알고리즘에서는
1번째 호출에서 i번째를 피벗으로 고르고,
2번째 호출에서 i번째를 피벗으로 고르고,
3번째 호출에서도 i번째를 피벗으로 고른다.

하지만 만약 '난수'를 생성해 피벗을 고른다면

1번째 호출에서는 i번째를 피벗으로 고르고,
2번째 호출에서는 j번째를 피벗으로 고르고,
3번째 호출에서는 k번째를 피벗으로 고르게 된다!

만약 최악의 순서인 배열이 들어온다고 가정해보자.
그래도 난수를 사용하면 피벗을 고르는 위치가 계속 바뀐다.
피벗으로 고른 모든 숫자가 최대/최소값일 확률은 0에 수렴하게 된다.

최악의 인풋이 들어와도
우리는 평균적인 경우의 성능을 얻을 수 있다.

n개의 범위에서 난수를 생성할 때는 O(n)의 연산이 더 들어간다.
하지만 파티셔닝도 O(n)이기 때문에 전체 복잡도는 바뀌지 않는다.

난수를 생성해서 피벗을 고른다 해도 여전히 최악의 경우가 없어진 것은 아니다.
하지만 어떤 인풋이 들어와도 O(n^2) 연산을 하게 될 가능성은 0에 수렴한다.

마음놓고 퀵 정렬의 성능을 즐길 수 있게 된 것이다!

랜덤으로 알고리즘을 더 좋게 개선한 우아한 사례다.
*/
func BetterQuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIdx := getRandomIdx(len(arr))

	left, right := getLeftAndRight(pivotIdx, arr)

	newLeft := QuickSort(left)
	newRight := QuickSort(right)

	return append(append(newLeft, arr[pivotIdx]), newRight...)
}

func getRandomIdx(arrlen int) int {
	rand.NewSource(time.Now().UnixNano())

	return rand.Intn(arrlen)
}

func getLeftAndRight(pivotIdx int, arr []int) ([]int, []int) {
	left := []int{}
	right := []int{}
	for idx, num := range arr {
		if idx == pivotIdx {
			continue
		}

		if num <= arr[pivotIdx] {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return left, right
}

/*
goos: windows
goarch: amd64
pkg: sol
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
=== RUN   BenchmarkQuickSort
BenchmarkQuickSort
BenchmarkQuickSort-8

	884589              1142 ns/op             976 B/op         39 allocs/op

PASS
ok      sol     1.817s
*/
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	}
}

/*
실제 벤치마크에서도 더 뛰어난 성능을 보인다.
goos: windows
goarch: amd64
pkg: sol
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
=== RUN   BenchmarkBetterQuickSort
BenchmarkBetterQuickSort
BenchmarkBetterQuickSort-8

	122595              8777 ns/op            1109 B/op         40 allocs/op

PASS
ok      sol     1.552s
*/
func BenchmarkBetterQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BetterQuickSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	}
}
