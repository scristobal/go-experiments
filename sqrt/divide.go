package sqrt

func Divide(x int) int {

	low := 1
	high := x

	var result int
	for low <= high {
		mid := (low + high) / 2

		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			low = mid + 1
			result = mid
		} else {
			high = mid - 1
		}
	}

	return result

}
