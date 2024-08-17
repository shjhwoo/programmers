package main

import (
	"container/heap"
	"testing"

	//https://pkg.go.dev/container/heap
	/*
		A heap is a common way to implement a priority queue.
		To build a priority queue, implement the Heap interface with the (negative) priority as the ordering for the Less method, so Push adds items
		while Pop removes the highest-priority item from the queue.
	*/
	"fmt"

	"gotest.tools/v3/assert"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(elem any) {
	*h = append(*h, elem.(int))
}

func (h *IntHeap) Pop() any {
	copy := *h //슬라이스는 배열의 주소값이기 때문에 메서드 정의 시 포인터로 정의하지 않으면 원본이 훼손됨
	result := copy[len(copy)-1]
	*h = copy[0 : len(copy)-1]

	return result
}

// go에서 구현하는 힙?
func main() {
	h := &IntHeap{2, 1, 7}

	heap.Init(h)
	fmt.Println(*h)

	heap.Push(h, 4)

	heap.Push(h, 10)

	fmt.Println(*h)

	for h.Len() > 0 {
		m := heap.Pop(h)
		fmt.Println(m)
	}
}

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 7}

	heap.Init(h)
	fmt.Println(*h)

	heap.Push(h, 4)
	assert.DeepEqual(t, []int{1, 2, 4, 7}, *h)
}
