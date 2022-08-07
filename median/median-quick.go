package median

import "fmt"

func MedianWithQuick(arr []int) int {

	fmt.Println("input", arr)

	mid := (len(arr) - 1) / 2

	fmt.Println("mid", mid)

	if mid < 0 {
		return 0 // happens when len(arr) == 0, should throw
	}

	i, j := 0, len(arr)-1

	for {

		fmt.Println("looking at range", i, j)

		x := arr[(len(arr)-1)/2]
		fmt.Println("pivot", x)

		for i < j {
			for arr[i] < x {
				i++
			}
			for arr[j] > x {
				j--
			}

			arr[i], arr[j] = arr[j], arr[i]
			i++
		}

		fmt.Println("arr after swapping", arr)
		fmt.Println("pivot location after swapping", j)

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
