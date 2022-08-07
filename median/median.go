package median

import (
	"container/heap"
	"fmt"
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

func Median(s []int) int {

	var minHeap MinHeap = []int{}
	minH := &minHeap

	var maxHeap MaxHeap = []int{}
	maxH := &maxHeap

	if len(s) == 0 {
		return 0
	}

	minH.Push(s[0])

	for i := 1; i < len(s); i++ {

		med := minH.Pop().(int)

		if s[i] >= med {
			heap.Push(minH, s[i])
			heap.Push(maxH, med)
		} else {
			heap.Push(maxH, s[i])
			heap.Push(minH, med)
		}

		if minH.Len() < maxH.Len() {
			heap.Push(minH, heap.Pop(maxH).(int))
		}
	}

	fmt.Println("min", minH)
	fmt.Println("max", maxH)

	return minHeap[0]
}
