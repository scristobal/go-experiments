package median

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
