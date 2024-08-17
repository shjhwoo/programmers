package main_test

import (
	"container/heap"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

type TestCase struct {
	no     int
	works  []int
	expect int
}

// 포인트는 인덱스를 같이 저장하는 것..!!
func TestSolution(t *testing.T) {
	var tests = []TestCase{
		{
			no:     4,
			works:  []int{4, 3, 3},
			expect: 12,
		},
		{
			no:     2,
			works:  []int{3, 3, 3},
			expect: 17,
		},
	}

	for _, test := range tests {
		ans := solution(test.no, test.works)
		t.Log(ans, "계산값")
		assert.DeepEqual(t, test.expect, ans)
	}
}

// 남은 일의 작업량을 숫자로 매기고 배상비용을 최소화
// 배상 비용은 각 선박의 완성까지 남은 일의 작업량을 제곱하여 모두 더한 값이 됩니다.
func solution(no int, works []int) int {
	//일단 배열의 모든 인수가 다 최소화가 되어야 한다.

	//가장 간단히 하려면 빼는 경우의 수 모두 다 찾고, 각각 계산해서 최소값을 구함.
	//그러면 일단은 어찌 되었든 간에 works 배열을 최대 힙으로 정렬하고,
	// 가장 상위의 원소에서 무조건 1을 뺀 다음에, 그 다음으로 최대값이 되는 원소를 찾아
	// 루트로 올리는 정렬을 한다.
	// 이런 과정을 반복하면 최소값을 구할 수 있다.
	// 최대가 되는 숫자들을 어찌되었든 간에 최소로 하려는 것이다.

	worksHeap := &IntHeap{}
	for _, work := range works {
		worksHeap.Push(work)
	}

	heap.Init(worksHeap)

	for no > 0 {
		longestWork := worksHeap.Pop().(int)
		longestWork--

		//다시 힙에 넣어준다.. 정렬은 알아서 함.
		worksHeap.Push(longestWork)

		//뺀다.
		no--
	}

	fmt.Println(*worksHeap)

	var answer int
	for worksHeap.Len() > 0 {
		elem := worksHeap.Pop().(int)

		answer += elem * elem
	}

	return answer
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(elem any) {
	*h = append(*h, elem.(int))
}

func (h *IntHeap) Pop() any {
	copy := *h //슬라이스는 배열의 주소값이기 때문에 메서드 정의 시 포인터로 정의하지 않으면 원본이 훼손됨
	result := copy[0]
	*h = copy[1:]

	return result
}

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 7}

	heap.Init(h)
	fmt.Println(*h)

	heap.Push(h, 4)
	assert.DeepEqual(t, IntHeap{1, 2, 7, 4}, *h)
}
