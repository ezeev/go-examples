package capitalize_test

import (
	"testing"

	"github.com/ezeev/go-examples/capitalize"
)

func TestCaps(t *testing.T) {

	capped := capitalize.CapWords("this is a test")
	t.Log(capped)
	if capped != "This Is A Test" {
		t.Error("words were not capitalized!")
	}

}
