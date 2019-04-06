package parser

// Errors specific to the parser package
const (
	RightCurvedBracketError string = "Unexpected \")\" [%d]"
	RightSquareBracketError string = "Unexpected \"]\" [%d]"
	AtomTypeError           string = "Bad Atom type"
	UnspecifiedAtomError    string = "Unspecified Atom error %#v: %#v"
)
