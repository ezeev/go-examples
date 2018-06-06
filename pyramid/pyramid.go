package pyramid

import (
	"fmt"
)

func Pyramid(start int, end int) {
	space := []byte(" ")
	hash := []byte("#")
	str := make([]byte, end)

	mid := end / 2

	for i := 0; i < end; i++ {
		if i <= mid+start && i >= mid-start {
			str = append(str, hash...)
		} else {
			str = append(str, space...)
		}
	}
	fmt.Println(string(str))
	if start < mid {
		Pyramid(start+1, end)
	}
}
