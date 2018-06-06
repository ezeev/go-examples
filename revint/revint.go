package revint

import (
	"strconv"

	"github.com/ezeev/go-examples/revstring"
)

func ReverseInt(in int) (int, error) {

	negative := in < 0
	inStr := strconv.Itoa(in)

	if negative {
		inStr = inStr[1:]
	}
	//reverse
	newStr := revstring.Reverse(inStr)
	if negative {
		newStr = "-" + newStr
	}
	return strconv.Atoi(newStr)

}
