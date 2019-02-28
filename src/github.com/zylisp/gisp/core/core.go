package core

import (
	"fmt"
	"reflect"
)

type Any interface{}

func MOD(a, b Any) int {
	var n, m int

	switch {
	case isInt(a):
		n = a.(int)
	case isFloat(a):
		n = int(a.(float64))
	default:
		panic("need int/float argument to mod. Received " +
			fmt.Sprintf("%v", reflect.TypeOf(a)))
	}

	switch {
	case isInt(b):
		m = b.(int)
	case isFloat(b):
		m = int(b.(float64))
	default:
		panic("need int/float argument to mod. Received " +
			fmt.Sprintf("%v", reflect.TypeOf(b)))
	}
	return n % m
}

// ADD any number of int or float64 args
// ALWAYS returns a float64
func ADD(args ...Any) float64 {
	length := len(args)
	nums := make([]Number, length, length)
	for i := 0; i < len(args); i++ {
		switch n := args[i]; {
		case isInt(n):
			nums[i] = Number{Value: n.(int)}
		case isFloat(n):
			nums[i] = Number{Value: n.(float64)}
		default:
			panic("ADD requires int and/or float64 objects")
		}
	}
	sumNum, err := AddNumbers(nums...)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return sumNum.ToFloat()
}

func SUB(args ...Any) float64 {
	var result float64
	if isInt(args[0]) {
		result = float64(args[0].(int))
	} else if isFloat(args[0]) {
		result = args[0].(float64)
	} else {
		panic("need int/float for SUB")
	}

	for i := 1; i < len(args); i++ {
		switch n := args[i]; {
		case isInt(n):
			result -= float64(n.(int))
		case isFloat(n):
			result -= n.(float64)
		}
	}

	return result
}

func MUL(args ...Any) float64 {
	var prod float64 = 1

	for i := 0; i < len(args); i++ {
		switch n := args[i]; {
		case isInt(n):
			prod *= float64(n.(int))
		case isFloat(n):
			prod *= n.(float64)
		}
	}

	return prod
}
func DIV() {}

// TODO: can only compare ints and slice lens for now.
func LT(args ...Any) bool {
	if len(args) < 2 {
		panic("can't compare less than 2 values!")
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if isInt(args[i]) {
			n = float64(args[i].(int))
		} else if isFloat(args[i]) {
			n = args[i].(float64)
		} else {
			panic("you can't compare that!")
		}

		var m float64
		if isInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if isFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			panic("you can't compare that!")
		}

		if n >= m {
			return false
		}
	}

	return true
}

// TODO: can only compare ints and slice lens for now.
func GT(args ...Any) bool {
	if len(args) < 2 {
		panic("can't compare less than 2 values!")
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if isInt(args[i]) {
			n = float64(args[i].(int))
		} else if isFloat(args[i]) {
			n = args[i].(float64)
		} else {
			panic("you can't compare that!")
		}

		var m float64
		if isInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if isFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			panic("you can't compare that!")
		}

		if n <= m {
			return false
		}
	}

	return true
}

func EQ(args ...Any) bool {
	if len(args) < 2 {
		panic("can't compare less than 2 values!")
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if isInt(args[i]) {
			n = float64(args[i].(int))
		} else if isFloat(args[i]) {
			n = args[i].(float64)
		} else {
			panic("you can't compare that!")
		}

		var m float64
		if isInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if isFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			panic("you can't compare that!")
		}

		if n != m {
			return false
		}
	}

	return true
}

// greater than or equal
func GTEQ(args ...Any) bool {
	if GT(args...) || EQ(args...) {
		return true
	}

	return false
}

// less than or equal
func LTEQ(args ...Any) bool {
	if LT(args...) || EQ(args...) {
		return true
	}

	return false
}

func isFloat(n Any) bool {
	nType := reflect.TypeOf(n)
	if nType == reflect.TypeOf(Number{}) {
		return isFloat(n.(Number).Value)
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

func isInt(n Any) bool {
	nType := reflect.TypeOf(n)
	if nType == reflect.TypeOf(Number{}) {
		return isInt(n.(Number).Value)
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

func Get(args ...Any) Any {
	if len(args) != 2 && len(args) != 3 {
		panic(fmt.Sprintf("get needs 2 or 3 arguments %d given.", len(args)))
	}

	if len(args) == 2 {
		if a, ok := args[1].([]Any); ok {
			return a[args[0].(int)]
		} else if a, ok := args[1].(string); ok {
			return a[args[0].(int)]
		} else {
			panic("arguments to get must include slice/vector/string")
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
			panic("arguments to get must include slice/vector/string")
		}
	}
}
