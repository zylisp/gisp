package generator

const (
	DefInExpressionError           string = "A def within an expression is not allowed"
	NSInExpressionError            string = "A namespace defined in an expression is not allowed"
	LoopWithoutRecurError          string = "There was no recur found in the loop"
	AssertArgsCountError           string = "The assert function requires two arguments"
	AssrtArgTypeError              string = "The assert function's first argument must be a type"
	MissingCallNodeError           string = "Expected call node is missing in root scope"
	CalleeIndentifierMismatchError string = "Expecting call to identifier (i.e. def, defconst, etc.)"
	MissingAssgnmentArgsError      string = "Expecting expression to be assigned to variable: %q"
	NSPackageTypeMismatch          string = "ns package name needs to be an identifier"
	InvalidImportError             string = "Import declaration is invalid"
	InvalidImportUseError          string = "Use of import is invalid"
	ExpectingAsInImportError       string = "Use of import is invalid; expecting: \":as\""
	BinaryArgsCountError           string = "Use of binary operator with only one argument is not allowed"
	UnaryArgsCountError            string = "Use of unary operator requires exactly on argument"
	TooManyArgsError               string = "Too many args (%d) were passed to %s; expected %d"
	TooFewArgsError                string = "Too few args (%d) were passed to %s; expected %d"
)
