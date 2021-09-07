package rhelper

func IsZero(numb float64) bool {
	return numb == 0
}

func IsNegatif(numb float64) bool {
	return numb < 0
}

func IsPositif(numb float64) bool {
	return numb > 0
}

func MinInt32(numbs ...int32) int32 {
	if len(numbs) == 0 {
		panic("")
	} else if len(numbs) == 1 {
		return numbs[0]
	}

	SortInt32(numbs)
	return numbs[0]
}

func MaxInt32(numbs ...int32) int32 {
	if len(numbs) == 0 {
		panic("")
	} else if len(numbs) == 1 {
		return numbs[0]
	}

	SortInt32(numbs)
	return numbs[len(numbs)-1]
}
