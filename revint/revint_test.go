package revint_test

import (
	"testing"

	"github.com/ezeev/go-examples/revint"
)

func TestRevInt(t *testing.T) {

	num := 23
	rev, _ := revint.ReverseInt(num)
	if rev != 32 {
		t.Error("Expected 32, got %d", rev)
	}

	num = -987
	rev, _ = revint.ReverseInt(num)
	if rev != -789 {
		t.Error("Expected -789, got %d", rev)
	}

}
