package core

import (
	"errors"
)

//Number type
type Number struct {
	Value Any
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
