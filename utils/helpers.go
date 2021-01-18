package utils

func Remove(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func Contains(s []int, e int) (int, bool) {
	for idx, a := range s {
		if a == e {
			return idx, true
		}
	}
	return -1, false
}