package median

import (
	"fmt"
	"math"
)

func split(arr []int, pivot int, mid int) ([]int, int, int) {

	splitter := make([]int, len(arr))
	i, j := 0, len(arr)-1

	for _, v := range arr {
		if v < pivot {
			splitter[i] = v
			i++
		} else if v > pivot {
			splitter[j] = v
			j--
		}
	}

	if i <= mid && mid < j {
		return []int{pivot}, pivot, mid
	}

	if i > mid {
		arr = splitter[:i]
	} else {
		arr = splitter[j:]
		mid = mid - j
	}

	pivot = arr[(len(arr)-1)/2]
	return arr, pivot, mid
}

func MedianDivide(arr []int) int {

	pivot := arr[(len(arr)-1)/2]
	mid := (len(arr) - 1) / 2

	for len(arr) > 1 {
		arr, pivot, mid = split(arr, pivot, mid)
	}

	return pivot

}

func kSelect(arr []int, k int) []int {

	pivot := arr[(len(arr)-1)/2]
	mid := k

	for len(arr) > k {
		arr, pivot, mid = split(arr, pivot, mid)
	}

	return arr

}

func getStrongest(arr []int, k int) []int {

	m := MedianDivide(arr)

	dist := make([]int, len(arr))

	for i, v := range arr {
		dist[i] = -int(math.Abs(float64(v - m)))
	}

	sel := kSelect(dist, k)

	return sel
}

func Main(arr []int, k int) []int {
	return getStrongest(arr, k)
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	y := float64(x)

	r := y
	e := float64(0)

	fmt.Println("-----", y)

	for math.Abs(e-r) > 0.5 {
		r = e
		e = (r + (y / r)) / 2
		fmt.Println("guess", e)
	}

	return int(e)
}
