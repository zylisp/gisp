package core

import (
	"errors"
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

//Number type
type Number struct {
	Value Any
}

// IsFloat boolean check
func IsFloat(n Any) bool {
	nType := reflect.TypeOf(n)
	if nType == reflect.TypeOf(Number{}) {
		return IsFloat(n.(Number).Value)
	}

	switch n.(type) {
	case float64:
		return true
	case float32:
		return true
	default:
		return false
	}
}

// IsInt boolean check
func IsInt(n Any) bool {
	nType := reflect.TypeOf(n)
	if nType == reflect.TypeOf(Number{}) {
		return IsInt(n.(Number).Value)
	}
	switch n.(type) {
	case int:
		return true
	case int64:
		return true
	default:
		return false
	}
}

// MOD modulus operator
func MOD(a, b Any) int {
	var n, m int

	switch {
	case IsInt(a):
		n = a.(int)
	case IsFloat(a):
		n = int(a.(float64))
	default:
		log.Error(ModArgTypeError + fmt.Sprintf("%v", reflect.TypeOf(a)))
	}

	switch {
	case IsInt(b):
		m = b.(int)
	case IsFloat(b):
		m = int(b.(float64))
	default:
		log.Error(ModArgTypeError + fmt.Sprintf("%v", reflect.TypeOf(b)))
	}
	return n % m
}

// ADD any number of int or float64 args
// ALWAYS returns a float64
func ADD(args ...Any) float64 {
	length := len(args)
	nums := make([]Number, length)
	for i := 0; i < len(args); i++ {
		switch n := args[i]; {
		case IsInt(n):
			nums[i] = Number{Value: n.(int)}
		case IsFloat(n):
			nums[i] = Number{Value: n.(float64)}
		default:
			log.Error(AddArgTypeError)
		}
	}
	sumNum, err := AddNumbers(nums...)
	if err != nil {
		log.Errorf("%v", err)
	}
	return sumNum.ToFloat()
}

// SUB any number of int or float64 args
// ALWAYS returns a float64
func SUB(args ...Any) float64 {
	var result float64
	if IsInt(args[0]) {
		result = float64(args[0].(int))
	} else if IsFloat(args[0]) {
		result = args[0].(float64)
	} else {
		log.Error(SubArgTypeError)
	}

	for i := 1; i < len(args); i++ {
		switch n := args[i]; {
		case IsInt(n):
			result -= float64(n.(int))
		case IsFloat(n):
			result -= n.(float64)
		}
	}

	return result
}

// MUL any number of int or float64 args
// ALWAYS returns a float64
func MUL(args ...Any) float64 {
	var prod float64 = 1

	for i := 0; i < len(args); i++ {
		switch n := args[i]; {
		case IsInt(n):
			prod *= float64(n.(int))
		case IsFloat(n):
			prod *= n.(float64)
		}
	}

	return prod
}

// DIV any number of int or float64 args
// ALWAYS returns a float64
func DIV() {
	// Not implemented
}

// IsInt method on number type
func (n Number) IsInt() bool {
	return IsInt(n.Value)
}

// IsFloat method on number type
func (n Number) IsFloat() bool {
	return IsFloat(n.Value)
}

// ToInt method on number type
func (n Number) ToInt() int {
	if n.IsInt() {
		return n.Value.(int)
	}
	return int(n.Value.(float64))
}

// ToFloat method on number type
func (n Number) ToFloat() float64 {
	if n.IsInt() {
		return float64(n.Value.(int))
	}
	return n.Value.(float64)
}

// AddNumbers adds a series of numbers
func AddNumbers(args ...Number) (newNum Number, err error) {
	temp := Number{Value: 0}
	newNum, err = temp.add(args...)
	return
}

func (n Number) add(args ...Number) (newNum Number, err error) {
	newNum = Number{Value: n.Value}
	for i := 0; i < len(args); i++ {
		x := args[i]
		if IsInt(newNum) && IsInt(x) {
			newNum.Value = newNum.Value.(int) + x.Value.(int)
		} else if IsInt(newNum) {
			newNum.Value = float64(newNum.Value.(int)) + x.Value.(float64)
		} else if IsFloat(newNum) && IsFloat(x) {
			newNum.Value = newNum.Value.(float64) + x.Value.(float64)
		} else if IsFloat(newNum) {
			newNum.Value = newNum.Value.(float64) + float64(x.Value.(int))
		} else {
			// can only happen if someone has constructed
			// a Number object with something other than a number
			// in Value
			err = errors.New(AddArgTypeError)
		}
	}
	return
}
