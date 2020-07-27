package topic

import "strings"

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	beginIndex := strings.Index(t, string(s[0]))
	if beginIndex == -1 {
		return false
	}
	var fromIndex = beginIndex + 1
	for i := 1; i < len(s); i++ {
		index := Index(t, s[i], fromIndex)
		if index == -1 {
			return false
		}
		fromIndex = index + 1
	}
	return true
}

func Index(s string, char uint8, fromIndex int) int {
	for j := fromIndex; j < len(s); j++ {
		if s[j] == char {
			return j
		}
	}
	return -1
}
