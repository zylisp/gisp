package reader

// General reader constants
const (
	EOF                              = int32(-1)
	newline                     rune = '\n'
	carriageReturn              rune = '\r'
	tab                         rune = '\t'
	space                       rune = ' '
	leftParen                   rune = '('
	rightParen                  rune = ')'
	leftBracket                 rune = '['
	rightBracket                rune = ']'
	doubleQuote                 rune = '"'
	plusSign                    rune = '+'
	minusSign                   rune = '-'
	lowestDigit                 rune = '0'
	highestDigit                rune = '9'
	semiColon                   rune = ';'
	doubleSlash                 rune = '\\'
	complexNumberIdentifier     rune = 'i'
	floatDelimiter              rune = '.'
	optionalCollectionDelimiter rune = ','
)

// AdditionalAllowedRunes rune vars
var AdditionalAllowedRunes = map[rune]bool{
	'>': true,
	'<': true,
	'=': true,
	'-': true,
	'+': true,
	'*': true,
	'&': true,
	'_': true,
	'/': true,
	'?': true,
}
