package lexer

// Lexer-specific errors
const (
	BadNumberSyntax               string = "Bad number syntax: %q"
	UnsupportedArgCount           string = "Unsupported number of args passed (%d)"
	UnsupportedRuneError          string = "Don't know what to do with rune: %q"
	UnterminatedQuotedStringError string = "Unterminated quoted string"
)
