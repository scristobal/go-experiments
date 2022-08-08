package median

func MedianWithQuick(arr []int) int {

	mid := (len(arr) - 1) / 2
	l, m := 0, len(arr)-1
	var pivot int

	for l != mid+1 {

		pivot = arr[l+(m-l)/2]

		i, j := l, m
		for i < j {
			for arr[i] <= pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

		if i < mid+1 {
			l = i + 1
		} else {
			m = j - 1
		}
	}

	return pivot

}
