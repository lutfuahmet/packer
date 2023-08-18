package utils

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func LessThanExist[T int | uint](elems []T, from, to T) bool {
	for _, s := range elems {
		if s > from && s < to {
			return true
		}
	}
	return false
}

func FindFirstBreakPoint[T int | uint](elems []T) (index int) {
	for i, elem := range elems {
		if !LessThanExist(elems, elem, elem*2) {
			return i
		}
	}
	return -1
}
