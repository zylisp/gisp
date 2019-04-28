package reader

// General lexer constants
const (
	EOF                          = -1
	newline                 rune = '\n'
	carriageReturn          rune = '\r'
	tab                     rune = '\t'
	space                   rune = ' '
	leftParen               rune = '('
	rightParen              rune = ')'
	leftBracket             rune = '['
	rightBracket            rune = ']'
	doubleQuote             rune = '"'
	plusSign                rune = '+'
	minusSign               rune = '-'
	lowestDigit             rune = '0'
	highestDigit            rune = '9'
	semiColon               rune = ';'
	doubleSlash             rune = '\\'
	complexNumberIdentifier rune = 'i'
	floatDelimiter          rune = '.'
)
