package maxchar_test

import (
	"testing"

	"github.com/ezeev/go-examples/maxchar"
)

func TestMaxChar(t *testing.T) {
	in := "everyone sees seals, elephants and even eels"
	maxChar, count := maxchar.MaxChar(in)

	if string(maxChar) != "e" || count != 12 {
		t.Error("Wrong char or count!")
	}

	t.Logf("Max char: %s counted %d times", string(maxChar), count)

}

func TestAnagram(t *testing.T) {
	str1 := "state"
	str2 := "taste"

	if !maxchar.IsAnagram(str1, str2) {
		t.Errorf("%s and %s should be an anagram")
	}

	str1 = "cats"
	str2 = "act"

	if maxchar.IsAnagram(str1, str2) {
		t.Errorf("%s and %s should NOT be an anagram")
	}

}
