package slice

// Depulicate removes duplicate elements from a string slice
// use a map to store the unique elements
func Depulicate(s []string) []string {
	r := make([]string, 0)
	m := make(map[string]bool)

	for _, v := range s {
		if !m[v] {
			r = append(r, v)
		}
	}
	return r
}
