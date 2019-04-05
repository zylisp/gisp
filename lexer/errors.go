package lexer

const (
	UnsupportedRuneError          string = "Don't know what to do with rune: %q"
	UnterminatedQuotedStringError string = "Unterminated quoted string"
	BadNumberSyntax               string = "Bad number syntax: %q"
)
