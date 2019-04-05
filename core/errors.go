package core

const (
	ModArgTypeError               string = "mod function requires need int/float argument; received incorrect type: "
	AddArgTypeError               string = "add function requires int and/or float64 objects"
	SubArgTypeError               string = "sub function requires need int/float argument"
	CompareArgsCountError         string = "You cannot compare less than two values"
	IncompatibleCompareTypesError string = "The comparison of types of args provided is not allowed"
	GetNot2Or3ArgsError           string = "The get function needs 2 or 3 arguments %d given."
	GetArgsTypesError             string = "Arguments to the get function must include slice/vector/string"
)
