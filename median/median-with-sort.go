package median

import "sort"

func MedianWithSort(arr []int) int {

	sort.Ints(arr)

	return arr[(len(arr)-1)/2]
}
