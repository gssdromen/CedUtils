package CedArrayUtils

// Index 判断数组中是否存在某个元素, -1 表示不存在
func Index(array []string, item string) int {
	for index, element := range array {
		if element == item {
			return index
		}
	}
	return -1
}
