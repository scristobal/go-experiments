package sqrt

func Bounded(x int) int {

	r := 0

	for r*r < x {
		r++
	}

	if r*r == x {
		return r
	}

	return r - 1
}
