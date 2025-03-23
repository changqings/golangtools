package slice

import "cmp"

// Depulicate removes duplicate elements from a string slice
// use a map to store the unique elements
func Depulicate[T cmp.Ordered](s []T) []T {
	r := make([]T, 0)
	m := make(map[T]bool)

	for _, v := range s {
		if !m[v] {
			r = append(r, v)
		}
	}
	return r
}
