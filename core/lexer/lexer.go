package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
)

/////////////////////////////////////////////////////////////////////////////
///   Constants, Vars, and Types   //////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

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

// Position position type
type Position struct {
	row    int
	column int
	// While the above give human-referencable position information, the lexer
	// needs to know, at any given point, the absolute position of a character
	// in the textual data of the given program.
	absolute int
}

// Atom object
type Atom struct {
	Type     AtomType
	Position Position
	Value    string
}

// AtomType atom type
type AtomType int

// Lexer atom constants
const (
	AtomError AtomType = iota
	AtomEOF

	AtomLeftParen
	AtomRightParen
	AtomLeftVect
	AtomRightVect

	AtomIdent
	AtomString
	AtomChar
	AtomFloat
	AtomInt
	AtomComplex

	// Ok, these aren't really atoms but ...
	AtomQuote
	AtomQuasiQuote
	AtomUnquote
	AtomUnquoteSplice
)

var atomNames = []string{
	AtomError:         "AtomError",
	AtomEOF:           "AtomEOF",
	AtomLeftParen:     "AtomLeftParen",
	AtomRightParen:    "AtomRightParen",
	AtomLeftVect:      "AtomLeftVect",
	AtomRightVect:     "AtomRightVect",
	AtomIdent:         "AtomIdent",
	AtomString:        "AtomString",
	AtomChar:          "AtomChar",
	AtomFloat:         "AtomFloat",
	AtomInt:           "AtomInt",
	AtomComplex:       "AtomComplex",
	AtomQuote:         "AtomQuote",
	AtomQuasiQuote:    "AtomQuasiQuote",
	AtomUnquote:       "AtomUnquote",
	AtomUnquoteSplice: "AtomUnquoteSplice",
}

// AdditionalAlphaNumRunes rune vars
var AdditionalAlphaNumRunes = map[rune]bool{
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

// stateFn state function type
type stateFn func(*Lexer) stateFn

/////////////////////////////////////////////////////////////////////////////
///   Object Definitions   //////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

// Lexer object
type Lexer struct {
	name             string
	input            string
	state            stateFn
	position         Position
	start            int
	currentRuneWidth int
	lastPosition     Position
	items            chan Atom

	// XXX currently unused; remove? or keep for later?
	// parenDepth int
	// vectDepth  int
}

/////////////////////////////////////////////////////////////////////////////
///   Constructors   ////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

// NewPosition is a variadic Position constructor
func NewPosition(args ...int) (Position, error) {
	defRow, defCol, defAbs := 1, 1, 0
	switch argCount := len(args); argCount {
	case 0:
		return NewPosition(defRow, defCol, defAbs)
	case 2:
		return NewPosition(args[0], args[1], defAbs)
	case 3:
		row, col, abs := args[0], args[1], args[2]
		if row == 0 {
			row = 1
		}
		if col == 0 {
			col = 1
		}
		return Position{row, col, abs}, nil
	default:
		return Position{}, fmt.Errorf(UnsupportedArgCount, argCount)
	}
}

// NewLexer returns a new lexer object
func NewLexer(name, input string) *Lexer {
	pos, _ := NewPosition()
	lastPos, _ := NewPosition(-1, -1, -1)
	l := &Lexer{
		name:         name,
		input:        input,
		position:     pos,
		lastPosition: lastPos,
		items:        make(chan Atom),
	}
	go l.run()
	return l
}

/////////////////////////////////////////////////////////////////////////////
///   Position Methods   ////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

// Column returns the column value of the associated item Position
func (p *Position) Column() int {
	return p.column
}

// Row returns the row value of the associated item Position
func (p *Position) Row() int {
	return p.row
}

/////////////////////////////////////////////////////////////////////////////
///   Atom Methods   ////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

// Column returns the column value of the associated item Position
func (a *Atom) Column() int {
	return a.Position.column
}

// Row returns the row value of the associated item Position
func (a *Atom) Row() int {
	return a.Position.row
}

// Absolute returns the absolute value of the associated item Position
func (a *Atom) Absolute() int {
	return a.Position.absolute
}

/////////////////////////////////////////////////////////////////////////////
///   Lexer Methods   ///////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

func (l *Lexer) codeSize() int {
	return utf8.RuneCountInString(l.input)
}

// Column returns the column value of the associated item Position
func (l *Lexer) Column() int {
	return l.position.column
}

// Row returns the row value of the associated item Position
func (l *Lexer) Row() int {
	return l.position.row
}

// Absolute returns the absolute value of the associated item Position
func (l *Lexer) Absolute() int {
	return l.position.absolute
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	log.Error("\nRUNNING NEXT ...")
	log.Warnf("Initial position: %#v", l.position)
	if l.Absolute() >= l.codeSize() {
		log.Error("Hit end of code string ...")
		log.Warnf("Last position: %#v", l.position)
		l.currentRuneWidth = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.input[l.Absolute():])
	log.Warnf("Rune: %s; width: %d", string(r), w)
	l.currentRuneWidth = w
	log.Warnf("Old position: %d", l.Absolute())
	l.position.absolute += l.currentRuneWidth
	log.Warnf("New position: %d", l.Absolute())
	log.Warnf("Updated position: %#v", l.position)

	if r == newline {
		log.Warn("Got newline; updating row and column ...")
		l.updatePositionNewLine()
	} else {
		log.Warn("Updating column ...")
		l.updatePositionNext()
	}
	log.Warnf("Final position update: %#v", l.position)
	log.Errorf("Lexed rune %s at position %#v", string(r), l.position)
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.position.absolute -= l.currentRuneWidth
	l.updatePositionBack()
}

// emit passes an Atom back to the client.
func (l *Lexer) emit(t AtomType) {
	pos, err := NewPosition(l.Row(), l.Column(), l.start)
	if err != nil {
		log.Error(err)
	}
	l.items <- Atom{t, pos, l.input[l.start:l.Absolute()]}
	l.start = l.Absolute()
}

func (l *Lexer) ignore() {
	l.start = l.Absolute()
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRuneRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRuneRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	pos, _ := NewPosition(l.Row(), l.Column(), l.start)
	l.items <- Atom{AtomError, pos, fmt.Sprintf(format, args...)}
	return nil
}

// NextAtom method
func (l *Lexer) NextAtom() Atom {
	item := <-l.items
	l.lastPosition = item.Position
	return item
}

func (l *Lexer) run() {
	for l.state = lexWhitespace; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

func (l *Lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRuneRun(digits)
	if l.accept(".") {
		l.acceptRuneRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRuneRun("0123456789")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if r := l.peek(); isAllowedIdentifierRune(r) {
		l.next()
		return false
	}
	return true
}

func (l *Lexer) String() string {
	lexedStrings := []string{""}
	for item := l.NextAtom(); item.Type != AtomEOF; {
		lexedStrings = append(lexedStrings,
			fmt.Sprintf("%-3s: row: %d, col: %2d, abs: %2d, type %s",
				item.Value, item.Row(), item.Column(),
				item.Absolute(), AtomName(item.Type)))
		item = l.NextAtom()
	}
	lexedStrings = append(lexedStrings, "")
	return strings.Join(lexedStrings, "\n")
}

// PrintAtoms prints the value of String()
func (l *Lexer) PrintAtoms() {
	fmt.Print(l.String())
}

func (l *Lexer) updatePositionNewLine() {
	l.lastPosition = l.position
	l.position.row++
	l.position.column = 1
}

func (l *Lexer) updatePositionNext() {
	log.Errorf("Column pre-update; last position: %#v; current position: %#v", l.lastPosition, l.position)
	l.lastPosition = l.position
	l.position.column++
	log.Errorf("Column post-update; last position: %#v; current position: %#v", l.lastPosition, l.position)
}

func (l *Lexer) updatePositionBack() {
	// back can only be called once per next call, so there's no need to
	// preserve a history of positions
	l.position.row = l.lastPosition.row
	l.position.column = l.lastPosition.column
}

/////////////////////////////////////////////////////////////////////////////
///   Support Functions   ///////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

func lexLeftVect(l *Lexer) stateFn {
	l.emit(AtomLeftVect)

	return lexWhitespace
}

func lexRightVect(l *Lexer) stateFn {
	l.emit(AtomRightVect)

	return lexWhitespace
}

// lexes an open parenthesis
func lexLeftParen(l *Lexer) stateFn {
	l.emit(AtomLeftParen)

	return lexWhitespace
}

func lexWhitespace(l *Lexer) stateFn {
	for r := l.next(); isWhiteSpace(r); l.next() {
		r = l.peek()
	}
	l.backup()
	l.ignore()

	switch r := l.next(); {
	case r == EOF:
		return lexEndOfFile
	case r == leftParen:
		return lexLeftParen
	case r == rightParen:
		return lexRightParen
	case r == leftBracket:
		return lexLeftVect
	case r == rightBracket:
		return lexRightVect
	case r == doubleQuote:
		return lexString
	case isCompoundNumber(r):
		return lexNumber
	case r == semiColon:
		return lexComment
	case isAllowedIdentifierRune(r):
		return lexIdentifier
	default:
		log.Panic(fmt.Sprintf(UnsupportedRuneError, r))
		return nil
	}
}

func lexString(l *Lexer) stateFn {
	for r := l.next(); r != doubleQuote; r = l.next() {
		if r == doubleSlash {
			r = l.next()
		}
		if r == EOF {
			return l.errorf(UnterminatedQuotedStringError)
		}
	}
	l.emit(AtomString)
	return lexWhitespace
}

func lexIdentifier(l *Lexer) stateFn {
	for r := l.next(); isAllowedIdentifierRune(r); r = l.next() {
	}
	l.backup()
	l.emit(AtomIdent)
	return lexWhitespace
}

// lex a close parenthesis
func lexRightParen(l *Lexer) stateFn {
	l.emit(AtomRightParen)

	return lexWhitespace
}

// lex a comment, comment delimiter is known to be already read
func lexComment(l *Lexer) stateFn {
	i := strings.Index(l.input[l.position.absolute:], "\n")
	l.position.absolute += i
	l.updatePositionNewLine()
	l.ignore()
	return lexWhitespace
}

func lexNumber(l *Lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf(BadNumberSyntax, l.input[l.start:l.position.absolute])
	}

	if l.start+1 == l.position.absolute {
		return lexIdentifier
	}

	if sign := l.peek(); sign == plusSign || sign == minusSign {
		// Complex: 1+2i. No spaces, must end in 'i'.
		if !l.scanNumber() || string(l.input[l.position.absolute-1]) != string(complexNumberIdentifier) {
			return l.errorf(BadNumberSyntax, l.input[l.start:l.position.absolute])
		}
		l.emit(AtomComplex)
	} else if strings.ContainsRune(l.input[l.start:l.position.absolute], floatDelimiter) {
		l.emit(AtomFloat)
	} else {
		l.emit(AtomInt)
	}

	return lexWhitespace
}

// lex end of a file
func lexEndOfFile(l *Lexer) stateFn {
	l.emit(AtomEOF)

	return nil
}

/////////////////////////////////////////////////////////////////////////////
///   Utility Functions   ///////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////

// isWhiteSpace reports whether r is a space character.
func isWhiteSpace(r rune) bool {
	return r == space || r == tab || r == newline
}

// XXX Currently unused; remove?
// isEndOfLine reports whether r is an end-of-line character.
// func isEndOfLine(r rune) bool {
// 	return r == carriageReturn || r == newLine
// }

// isAllowedIdentifierRune reports whether r is a valid rune for an identifier.
func isAllowedIdentifierRune(r rune) bool {
	return AdditionalAlphaNumRunes[r] || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// XXX Let's fix this so that + or - qualify ONLY if the next char is a number
func isNumber(r rune) bool {
	return (lowestDigit <= r && r <= highestDigit)
}

// isCompoundNumber reports whether r is a number or series of characters that
//                  represent a number.
// XXX Let's fix this so that + or - qualify ONLY if the next char is a number,
//     since this logic is currently breaking symbols that start with a + or -
func isCompoundNumber(r rune) bool {
	return r == plusSign || r == minusSign || isNumber(r)
}

// AtomName returns the name of the atom for a given atom value
func AtomName(atomType AtomType) string {
	return atomNames[atomType]
}
