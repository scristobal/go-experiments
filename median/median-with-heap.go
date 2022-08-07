package median

import (
	"container/heap"
)

// min priority queue definition
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Max priority queue definition
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MedianWithHeap(s []int) int {

	var minHeap MinHeap = []int{}
	highValues := &minHeap

	var maxHeap MaxHeap = []int{}
	lowerValues := &maxHeap

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return s[0]
	}

	if s[0] < s[1] {
		lowerValues.Push(s[0])
		highValues.Push(s[1])
	} else {
		lowerValues.Push(s[1])
		highValues.Push(s[0])
	}

	for i := 2; i < len(s); i++ {

		heap.Push(highValues, s[i])

		for lowerValues.Len() < highValues.Len() {
			heap.Push(lowerValues, heap.Pop(highValues).(int))
		}
	}

	return (*lowerValues)[0]
}
