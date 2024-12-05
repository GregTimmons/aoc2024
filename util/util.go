package util

type Number interface {
    int64 | int
}

func SlicePop[K Number](s []K) (K, []K, bool) {
    if len(s) == 0 {
		return 0, s, false
    }
    return s[len(s)-1], s[:len(s)-1], true
}

func SliceIntersects[K Number](lhs, rhs []K) bool {
	for _, l := range lhs {
		for _, r := range rhs {
			if l == r { return true }
		}
	}
	return false
}

func SliceContains[K Number](arr []K, targ K) bool {
	for _, val := range arr {
		if val == targ { return true}
	}
	return false
}
