package core

import (
	"errors"
)

type Number struct {
	Value Any
}

func (n Number) IsInt() bool {
	return isInt(n.Value)
}
func (n Number) IsFloat() bool {
	return isFloat(n.Value)
}

func (n Number) ToInt() int {
	if n.IsInt() {
		return n.Value.(int)
	} else {
		return int(n.Value.(float64))
	}
}

func (n Number) ToFloat() float64 {
	if n.IsInt() {
		return float64(n.Value.(int))
	} else {
		return n.Value.(float64)
	}
}
func AddNumbers(args ...Number) (newNum Number, err error) {
	temp := Number{Value: 0}
	newNum, err = temp.add(args...)
	return
}

func (n Number) add(args ...Number) (newNum Number, err error) {
	newNum = Number{Value: n.Value}
	for i := 0; i < len(args); i++ {
		x := args[i]
		if isInt(newNum) && isInt(x) {
			newNum.Value = newNum.Value.(int) + x.Value.(int)
		} else if isInt(newNum) {
			newNum.Value = float64(newNum.Value.(int)) + x.Value.(float64)
		} else if isFloat(newNum) && isFloat(x) {
			newNum.Value = newNum.Value.(float64) + x.Value.(float64)
		} else if isFloat(newNum) {
			newNum.Value = newNum.Value.(float64) + float64(x.Value.(int))
		} else {
			// can only happen if someone has constructed
			// a Number object with something other than a number
			// in Value
			err = errors.New("ADD requires int and/or float64 objects")
		}
	}
	return
}
