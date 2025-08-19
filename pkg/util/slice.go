package util

// SliceUnique 对切片去重处理
func SliceUnique[T comparable](input []T) []T {
	uniqueMap := make(map[T]bool)
	var uniqueSlice []T
	for _, val := range input {
		if _, ok := uniqueMap[val]; !ok {
			uniqueMap[val] = true
			uniqueSlice = append(uniqueSlice, val)
		}
	}
	return uniqueSlice
}

// ChunkSlice 将切片按照指定size切割
func ChunkSlice[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// SliceToMap 切片转为map
func SliceToMap[T any](slice []T) map[any]bool {
	set := make(map[any]bool)
	for _, t := range slice {
		set[t] = true
	}
	return set
}
