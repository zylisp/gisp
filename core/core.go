package core

import (
	log "github.com/sirupsen/logrus"
)

// Any is an interface for any ZYLISP type
type Any interface{}

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
