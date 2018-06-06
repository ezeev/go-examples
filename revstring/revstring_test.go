package revstring_test

import (
	"testing"

	"github.com/ezeev/go-examples/revstring"
)

func TestRevString(t *testing.T) {
	str := "this is my input"
	out := revstring.Reverse(str)
	t.Log(out)
}

func TestPalindrome(t *testing.T) {
	in := "abba"
	ispal := revstring.IsPalindrome(in)
	if !ispal {
		t.Error("abba is palindrome!")
	} else {
		t.Logf("is abba a palindrome? %s", ispal)
	}
}
