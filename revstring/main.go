package revstring

import (
	"strings"
)

func Reverse(input string) string {
	bytes := []byte(input)
	revbytes := make([]byte, len(bytes))
	pos := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		revbytes[pos] = bytes[i]
		pos++
	}
	return string(revbytes)
}

func IsPalindrome(input string) bool {
	out := Reverse(input)
	if strings.TrimRight(input, " ") == out {
		return true
	}
	return false
}
