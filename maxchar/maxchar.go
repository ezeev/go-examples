package maxchar

import (
	"reflect"
)

func runeCountMap(in string) map[rune]int {
	runeMap := make(map[rune]int)
	for _, char := range in {
		runeMap[char] += 1
	}
	return runeMap
}

func MaxChar(in string) (rune, int) {
	charMap := runeCountMap(in)

	var maxChar rune
	var maxCount int
	for k, v := range charMap {
		if v > maxCount && k != maxChar {
			maxCount = v
			maxChar = k
		}
	}
	return maxChar, maxCount
}

func IsAnagram(in1 string, in2 string) bool {
	map1 := runeCountMap(in1)
	map2 := runeCountMap(in2)
	return reflect.DeepEqual(map1, map2)
}
