package median

func MedianWithQuick(arr []int) int {

	//fmt.Println("input", arr)

	mid := (len(arr) - 1) / 2

	//fmt.Println("mid", mid)

	if mid < 0 {
		return 0 // happens when len(arr) == 0, should throw
	}

	i, j := 0, len(arr)-1

	for {

		//fmt.Println("looking at range", i, j)

		x := arr[(len(arr)-1)/2]
		//fmt.Println("pivot", x)

	outer:
		for {
			for arr[i] < x {
				i++
			}
			for arr[j] > x {
				j--
			}

			if arr[i] != arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			} else {

				for k := 1; i+k < j; k++ {
					if arr[i] > arr[i+k] {
						arr[i], arr[i+k] = arr[i+k], arr[i]
						break
					}
				}

				for k := 1; j-k > i; k++ {
					if arr[j] < arr[j-k] {
						arr[j], arr[j-k] = arr[j-k], arr[j]
						break
					}
				}

			}

			if i == j {
				break
			}

			for k := i; k <= j; k++ {
				if arr[k] != arr[j] {
					continue outer
				}
			}
			break

		}

		//fmt.Println("arr after swapping", arr)
		//fmt.Println("pivot location after swapping", j)

		if i == mid {
			return arr[i]
		}

		if i < mid {
			i++
			j = len(arr) - 1
		} else {
			j--
			i = 0
		}
	}

}
