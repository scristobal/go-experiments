package median

import (
	"container/heap"
	"math/rand"
	"sort"
	"testing"
)

func TestMinQueue(t *testing.T) {

	s := []int{3, 1, 2, 4}

	var mh MinHeap = s
	h := &mh

	heap.Init(h)

	if h.Len() != len(s) {
		t.Errorf("not all items are in the queue after init")
	}

	m := heap.Pop(h).(int)

	if m != 1 {
		t.Errorf("Pop did not return the right value")
	}

	if h.Len() != 3 {
		t.Errorf("Pop did not remove the item")
	}

	heap.Push(h, 0)

	if h.Len() != 4 {
		t.Errorf("Push did not add the item")
	}

	m = heap.Pop(h).(int)

	if m != 0 {
		t.Errorf("Pop did not return the value added by push")
	}

}

func TestMaxQueue(t *testing.T) {

	s := []int{3, 1, 2, 4}

	var mh MaxHeap = s
	h := &mh

	heap.Init(h)

	if h.Len() != len(s) {
		t.Errorf("not all items are in the queue after init")
	}

	m := heap.Pop(h).(int)

	if m != 4 {
		t.Errorf("Pop did not return the right value")
	}

	if h.Len() != 3 {
		t.Errorf("Pop did not remove the item")
	}

	heap.Push(h, 5)

	if h.Len() != 4 {
		t.Errorf("Push did not add the item")
	}

	m = heap.Pop(h).(int)

	if m != 5 {
		t.Errorf("Pop did not return the value added by push")
	}

}

func MedianWithSort(arr []int) int {

	sort.Ints(arr)

	return arr[(len(arr)-1)/2]
}

func TestMedian(t *testing.T) {

	s := []int{3, 1, 2, 4}

	m := Median(s)

	withSort := MedianWithSort(s)

	if m != withSort {
		t.Errorf("Median did not return the right value, expected %d returned %d", withSort, m)
	}

}

func Rand(max int, size int) []int {
	b := make([]int, size)
	for i := range b {
		b[i] = rand.Intn(max)
	}
	return b
}

var randomSeq = Rand(100, 100)

func BenchmarkHeap(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Median(randomSeq)
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MedianWithSort(randomSeq)
	}
}
