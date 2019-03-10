package lexer

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/zylisp/gisp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var log = logging.MustGetLogger(gisp.ApplicationName)

type Pos int

type Atom struct {
	Type  AtomType
	Pos   Pos
	Value string
}

type AtomType int

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

	// Ok, these aren't really atoms but...
	AtomQuote
	AtomQuasiQuote
	AtomUnquote
	AtomUnquoteSplice
)

const EOF = -1

var AdditionalAlphaNumRunes map[rune]bool = map[rune]bool{
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

type stateFn func(*Lexer) stateFn

type Lexer struct {
	name    string
	input   string
	state   stateFn
	pos     Pos
	start   Pos
	width   Pos
	lastPos Pos
	items   chan Atom

	parenDepth int
	vectDepth  int
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
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
	l.pos -= l.width
}

// emit passes an Atom back to the client.
func (l *Lexer) emit(t AtomType) {
	l.items <- Atom{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRuneRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRuneRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Atom{AtomError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func (l *Lexer) NextAtom() Atom {
	item := <-l.items
	l.lastPos = item.Pos
	return item
}

func Lex(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		items: make(chan Atom),
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for l.state = lexWhitespace; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

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
	for r := l.next(); isSpace(r) || r == '\n'; l.next() {
		r = l.peek()
	}
	l.backup()
	l.ignore()

	switch r := l.next(); {
	case r == EOF:
		l.emit(AtomEOF)
		return nil
	case r == '(':
		return lexLeftParen
	case r == ')':
		return lexRightParen
	case r == '[':
		return lexLeftVect
	case r == ']':
		return lexRightVect
	case r == '"':
		return lexString
	case r == '+' || r == '-' || ('0' <= r && r <= '9'):
		return lexNumber
	case r == ';':
		return lexComment
	case isAlphaNumeric(r):
		return lexIdentifier
	default:
		msg := fmt.Sprintf(UnsupportedRuneError, r)
		log.Critical(msg)
		panic(msg)
	}
}

func lexString(l *Lexer) stateFn {
	for r := l.next(); r != '"'; r = l.next() {
		if r == '\\' {
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
	for r := l.next(); isAlphaNumeric(r); r = l.next() {
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
	i := strings.Index(l.input[l.pos:], "\n")
	l.pos += Pos(i)
	l.ignore()
	return lexWhitespace
}

func lexNumber(l *Lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf(BadNumberSyntax, l.input[l.start:l.pos])
	}

	if l.start+1 == l.pos {
		return lexIdentifier
	}

	if sign := l.peek(); sign == '+' || sign == '-' {
		// Complex: 1+2i. No spaces, must end in 'i'.
		if !l.scanNumber() || l.input[l.pos-1] != 'i' {
			return l.errorf(BadNumberSyntax, l.input[l.start:l.pos])
		}
		l.emit(AtomComplex)
	} else if strings.ContainsRune(l.input[l.start:l.pos], '.') {
		l.emit(AtomFloat)
	} else {
		l.emit(AtomInt)
	}

	return lexWhitespace
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
	if r := l.peek(); isAlphaNumeric(r) {
		l.next()
		return false
	}
	return true
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is a valid rune for an identifier.
func isAlphaNumeric(r rune) bool {
	return AdditionalAlphaNumRunes[r] == true || unicode.IsLetter(r) || unicode.IsDigit(r)
}
