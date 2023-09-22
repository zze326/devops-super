package util

func GetPointer[T any](v T) *T {
	return &v
}

func SqlLikeStr(str string) string {
	return "%" + str + "%"
}

func InSlice[T comparable](slice []T, v T) bool {
	for _, item := range slice {
		if item == v {
			return true
		}
	}
	return false
}
