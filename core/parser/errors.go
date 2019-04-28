package parser

// Errors specific to the parser package
const (
	RightCurvedBracketError string = "Unexpected \")\" [row: %d, column: %d]"
	RightSquareBracketError string = "Unexpected \"]\" [row: %d, column: %d]"
	AtomTypeError           string = "Bad Atom type"
	UnspecifiedAtomError    string = "Unspecified Atom error for Atom %s [row: %d, column: %d]"
)
