package strslc

// Contains returns true if the string is in the slice
func Contains(slc []string, match string) bool {
	for _, s := range slc {
		if s == match {
			return true
		}
	}
	return false
}

// Where returns the subset slice that returns true from the matchFunc
func Where(slc []string, matchFunc func(string) bool) []string {
	newslc := []string{}
	for _, s := range slc {
		if matchFunc(s) {
			newslc = append(newslc, s)
		}
	}
	return newslc
}
