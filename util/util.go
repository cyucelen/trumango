package util

// Difference finds the difference of given string slice a, from string slice b
// in linear time using map
func Difference(a, b []string) []string {
	bContains := make(map[string]bool)
	for _, x := range b {
		bContains[x] = true
	}

	diff := make([]string, 0)
	for _, x := range a {
		if _, ok := bContains[x]; !ok {
			diff = append(diff, x)
		}
	}
	return diff
}
