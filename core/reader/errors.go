package reader

// Reader-specific errors
const (
	BadNumberSyntax               string = "Bad number syntax: %q"
	UnsupportedRuneError          string = "Don't know what to do with rune: %q (%s)"
	UnterminatedQuotedStringError string = "Unterminated quoted string"
)
