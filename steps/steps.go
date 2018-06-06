// basic recusion example

package steps

import (
	"fmt"
)

func LogSteps(n int) {
	LogStepRecur(0, n)
}

func LogStepRecur(start int, end int) {
	stair := "#"
	for i := 0; i < start; i++ {
		stair += "#"
	}
	fmt.Println(stair)
	if start+1 < end {
		LogStepRecur(start+1, end)
	}

}
