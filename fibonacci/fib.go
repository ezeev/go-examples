package fibonacci

import (
	"fmt"
)

func PrintFib(length int) {
	var arr []int

	arr = append(arr, 0)
	arr = append(arr, 1)
	for i := 2; i <= length; i++ {
		val := arr[i-1] + arr[i-2]
		arr = append(arr, val)
	}
	fmt.Println(arr)
}

func SlowFib(n int) int {
	if n < 2 {
		return n
	}
	return SlowFib(n-1) + SlowFib(n-2)
}

func FibMemoization(n int) {
	fibMap := make(map[int]int)
	for i := 0; i <= n; i++ {
		if val, ok := fibMap[i]; ok {
			fmt.Println(val)
		} else {
			fibMap[i] = SlowFib(i)
			fmt.Println(fibMap[i])
		}
	}
}
