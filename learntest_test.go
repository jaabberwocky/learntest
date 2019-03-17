package learntest

import (
	//"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	testVar := Plusone(2)
	if testVar != 3 {
		t.Errorf("This is wrong!")
	}
}

func TestOneSquareRoot(t *testing.T) {
	num, _ := SquareRoot(49)
	if num != 7 {
		t.Errorf("Square root is not correct")
	}
}

func TestTwoSquareRoot(t *testing.T) {
	_, err := SquareRoot(-15)
	if err == nil {
		t.Errorf("Error expected here")
	}
}
