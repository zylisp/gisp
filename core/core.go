package core

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// Any is an interface for any ZYLISP type
type Any interface{}

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

// LT less-than
// TODO: can only compare ints and slice lens for now.
func LT(args ...Any) bool {
	if len(args) < 2 {
		log.Error(CompareArgsCountError)
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if IsInt(args[i]) {
			n = float64(args[i].(int))
		} else if IsFloat(args[i]) {
			n = args[i].(float64)
		} else {
			log.Error(IncompatibleCompareTypesError)
		}

		var m float64
		if IsInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if IsFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			log.Error(IncompatibleCompareTypesError)
		}

		if n >= m {
			return false
		}
	}

	return true
}

// GT greater-than
// TODO: can only compare ints and slice lens for now.
func GT(args ...Any) bool {
	if len(args) < 2 {
		log.Error(CompareArgsCountError)
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if IsInt(args[i]) {
			n = float64(args[i].(int))
		} else if IsFloat(args[i]) {
			n = args[i].(float64)
		} else {
			log.Error(IncompatibleCompareTypesError)
		}

		var m float64
		if IsInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if IsFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			log.Error(IncompatibleCompareTypesError)
		}

		if n <= m {
			return false
		}
	}

	return true
}

// EQ equal
func EQ(args ...Any) bool {
	if len(args) < 2 {
		log.Error(CompareArgsCountError)
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if IsInt(args[i]) {
			n = float64(args[i].(int))
		} else if IsFloat(args[i]) {
			n = args[i].(float64)
		} else {
			log.Error(IncompatibleCompareTypesError)
		}

		var m float64
		if IsInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if IsFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			log.Error(IncompatibleCompareTypesError)
		}

		if n != m {
			return false
		}
	}

	return true
}

// GTEQ greater-than or equal
func GTEQ(args ...Any) bool {
	if GT(args...) || EQ(args...) {
		return true
	}

	return false
}

// LTEQ less-than or equal
func LTEQ(args ...Any) bool {
	if LT(args...) || EQ(args...) {
		return true
	}

	return false
}

// Get ...
func Get(args ...Any) Any {
	if len(args) != 2 && len(args) != 3 {
		log.Errorf(GetNot2Or3ArgsError, len(args))
		return nil
	}

	if len(args) == 2 {
		if a, ok := args[1].([]Any); ok {
			return a[args[0].(int)]
		} else if a, ok := args[1].(string); ok {
			return a[args[0].(int)]
		} else {
			log.Error(GetArgsTypesError)
			return nil
		}
	} else {
		if a, ok := args[2].([]Any); ok {
			if args[1].(int) == -1 {
				return a[args[0].(int):]
			}

			return a[args[0].(int):args[1].(int)]
		} else if a, ok := args[2].(string); ok {
			if args[1].(int) == -1 {
				return a[args[0].(int):]
			}

			return a[args[0].(int):args[1].(int)]
		} else {
			log.Error(GetArgsTypesError)
			return nil
		}
	}
}
