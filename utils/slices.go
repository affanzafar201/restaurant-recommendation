package utils

func FetchTopNElements[T any](slice []T, n int) []T {
	if len(slice) > n {
		return slice[:n]
	}

	return slice
}
