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

// FindFirstBreakPoint finds a item that does not have any element between itself and a multiple of 2 * item.
func FindFirstBreakPoint[T int | uint](elems []T) (index int) {
	for i, elem := range elems {
		if !LessThanExist(elems, elem, elem*2) {
			return i
		}
	}
	return -1
}

func GetIndexByValue[T comparable](elems []T, value T) int {
	for i, s := range elems {
		if s == value {
			return i
		}
	}
	return -1
}
