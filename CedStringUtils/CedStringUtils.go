package CedStringUtils

import "strings"

// IndexAll 获得字符串中所有target的Index
func IndexAll(str string, target string) []int {
	indexes := []int{}

	for strings.Contains(str, target) {
		index := strings.Index(str, target)
		str = str[index+1:]
		if len(indexes) > 0 {
			lastIndex := indexes[len(indexes)-1]
			index += lastIndex + 1
		}
		indexes = append(indexes, index)
	}
	return indexes
}

// IndexOf 获得第几个target的Index
func IndexOf(str string, target string, targetIndex int) int {
	indexes := IndexAll(str, target)
	if targetIndex < len(indexes) {
		return indexes[targetIndex]
	}
	return -1
}
