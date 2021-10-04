package intslc

// Contains returns true if the int is in the slice
func Contains(slc []int, match int) bool {
	for _, s := range slc {
		if s == match {
			return true
		}
	}
	return false
}

// Where returns the subset slice that returns true from the matchFunc
func Where(slc []int, matchFunc func(int) bool) []int {
	newslc := []int{}
	for _, s := range slc {
		if matchFunc(s) {
			newslc = append(newslc, s)
		}
	}
	return newslc
}
