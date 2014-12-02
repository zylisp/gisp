package core

import (
	"fmt"
	. "gopkg.in/check.v1"
	"io"
	"testing"
)

func Test_isInt(t *testing.T) {
	if isInt(1) != true {
		t.Error("1 wasn't an integer")
	}
	if isInt(4.0) == true {
		t.Error("4.0 tested as an int")
	}
}

func test_isFloat(t *testing.T) {
	if isFloat(1) == true {
		t.Error("1 tested as a float")
	}
	if isFloat(4.0) != true {
		t.Error("4.0 wasn't a float")
	}
}

func Test_MOD(t *testing.T) {
	if MOD(4.0, 2.0) != 0 {
		t.Error("4.0 % 2.0 != 0")
	}
	if MOD(7, 3.5) != 1 {
		t.Error(fmt.Printf(" 7 % 3.5 returned: \"%v\"", MOD(7, 3.5)))
	}
	if MOD(4, 2) != 0 {
		t.Error("4 % 2 != 0")
	}
}

func Test_ADD(t *testing.T) {
	if ADD(3, 4.1) != 7.1 {
		t.Error("can't add an int and a float")
	} else {
		t.Log("can add int and float")
	}
	if ADD(3, 4) != 7.0 {
		t.Error("can't add two ints")
	} else {
		t.Log("can add two int's")
	}
	if ADD(3.0, 4.1) != 7.1 {
		t.Error("Can't add two floats")
	} else {
		t.Log("can add two floats")
	}
}
