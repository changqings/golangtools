package slice

import (
	"slices"
)

func Depulicate(s []string) []string {
	result := make([]string, 0)
	for _, v := range s {
		if !slices.Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}
