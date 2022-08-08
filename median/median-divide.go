package median

func split(arr []int, pivot int, mid int) ([]int, int, int) {
	smaller, s := make([]int, len(arr)), 0
	larger, l := make([]int, len(arr)), 0
	equal, e := make([]int, len(arr)), 0

	for _, v := range arr {
		if v < pivot {
			smaller[s] = v
			s++
		} else if v > pivot {
			larger[l] = v
			l++
		} else {
			equal[e] = v
			e++
		}
	}

	if s <= mid && mid < (s+e) {
		return []int{pivot}, pivot, mid
	}

	if s > mid {
		arr = smaller[:s]
	} else {
		arr = larger[:l]
		mid = mid - s - e
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
