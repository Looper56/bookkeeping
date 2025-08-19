package util

import "strings"

// ContainsStr 目标字符串是否存在于数组内
func ContainsStr(target string, arr []string) bool {
	for _, str := range arr {
		if strings.Contains(target, str) {
			return true
		}
	}
	return false
}
