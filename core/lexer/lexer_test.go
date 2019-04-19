package lexer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func newTestLexer(data string) Lexer {
	return Lexer{
		name:  "tester",
		input: data,
		items: make(chan Atom),
	}
}

type LexerSuite struct {
	suite.Suite
	testLexer Lexer
}

func TestLexerSuite(t *testing.T) {
	suite.Run(t, new(LexerSuite))
}

func (s *LexerSuite) SetupTest() {
	s.testLexer = newTestLexer("()")
}

func (s *LexerSuite) TestNext() {
	// returns and "consumes" the next rune
	s.Equal(int(s.testLexer.pos), 0)
	s.Equal(s.testLexer.next(), '(')
	s.Equal(int(s.testLexer.pos), 1)
	s.Equal(s.testLexer.next(), ')')

	// next should move the position forward but not
	// affect the start (of input)
	currentString := s.testLexer.input[s.testLexer.start:len(s.testLexer.input)]
	s.Equal(currentString, "()")
}

func (s *LexerSuite) TestBackup() {
	// returns and "consumes" the next rune
	s.Equal(s.testLexer.next(), '(')
	s.Equal(s.testLexer.next(), ')')
	// ^^^ consumes the ) so if we backup
	// we should get it again
	s.testLexer.backup()
	s.Equal(s.testLexer.next(), ')')
}

func (s *LexerSuite) TestPeek() {
	// peek is next, but without consuming the rune
	s.Equal(s.testLexer.peek(), '(')
	s.Equal(s.testLexer.peek(), '(')
	s.Equal(s.testLexer.next(), '(')
}

func (s *LexerSuite) TestIgnore() {
	// returns and "consumes" the next rune
	//s.Equal(s.testLexer.emit(AtomString), "BS")
	s.testLexer.next()   // moves position but not start. See Test_next
	s.testLexer.ignore() // moves start to current position.
	currentString := s.testLexer.input[s.testLexer.start:len(s.testLexer.input)]
	s.Equal(currentString, ")")
}

func (s *LexerSuite) TestAccept() {
	// accept consumes the next rune *if* it's from the valid set.

	acceptableSet := "abc(d"
	unacceptableSet := "nope"
	s.Equal(int(s.testLexer.pos), 0)
	s.False(s.testLexer.accept(unacceptableSet))
	s.Equal(int(s.testLexer.pos), 0)
	// position shouldn't have changed

	s.True(s.testLexer.accept(acceptableSet))
	s.Equal(int(s.testLexer.pos), 1)
}

func (s *LexerSuite) TestAcceptRuneRun() {

	acceptableSet := "12345"
	s.testLexer.input = "123a45"
	s.testLexer.acceptRuneRun(acceptableSet)
	// position should now be 3
	// e.g. it should have nexted to the a, then gone back one
	s.Equal(int(s.testLexer.pos), 3)
}

// TODO: complete Test_emit
// I, apparently, don't understand chanels enough to
// make it not blow up
// func (s *LexerSuite) Test_emit() {
// }

// TODO: Test_NextAtom
// TODO: Test_lexWhitespace
// TODO: Test_lexString
// TODO: Test_lexIdentifier
// TODO: Test_lexComment
// TODO: Test_lexNumber
// TODO: Test_scanNumber
func (s *LexerSuite) TestScanNumber() {
	testLexer := newTestLexer("1234")
	s.True(testLexer.scanNumber())
	testLexer = newTestLexer("-1234")
	s.True(testLexer.scanNumber())
	testLexer = newTestLexer("+1234")
	s.True(testLexer.scanNumber())
	testLexer = newTestLexer("1234.56")
	s.True(testLexer.scanNumber())
	testLexer = newTestLexer("-1234.456e+78")
	s.True(testLexer.scanNumber())
	testLexer = newTestLexer("0x1c8")
	s.True(testLexer.scanNumber())
	//TODO: add a test for imaginary numbers (start with i)

	// totally not a number
	testLexer = newTestLexer("poopy")
	s.False(testLexer.scanNumber())
	// we want to support names with leading +/-
	testLexer = newTestLexer("-pvt-fn")
	s.False(testLexer.scanNumber())
	testLexer = newTestLexer("+silly-fn")
	s.False(testLexer.scanNumber())
}

func (s *LexerSuite) TestAtomName() {
	s.Equal(AtomName(0), "AtomError")
	s.Equal(AtomName(1), "AtomEOF")
	s.Equal(AtomName(2), "AtomLeftParen")
	s.Equal(AtomName(5), "AtomRightVect")
	s.Equal(AtomName(10), "AtomInt")
	s.Equal(AtomName(15), "AtomUnquoteSplice")
}

var exampleFnOneLine = "(def dbl (fn [x] (* 2 x)))"
var expectedExampleFnTokens = `
(  : position  0, type AtomLeftParen
def: position  1, type AtomIdent
dbl: position  5, type AtomIdent
(  : position  9, type AtomLeftParen
fn : position 10, type AtomIdent
[  : position 13, type AtomLeftVect
x  : position 14, type AtomIdent
]  : position 15, type AtomRightVect
(  : position 17, type AtomLeftParen
*  : position 18, type AtomIdent
2  : position 20, type AtomIdent
x  : position 22, type AtomIdent
)  : position 23, type AtomRightParen
)  : position 24, type AtomRightParen
)  : position 25, type AtomRightParen
`

func (s *LexerSuite) TestExampleFn() {
	lexed := Lex("a-prog", exampleFnOneLine)
	s.Equal(lexed.String(), expectedExampleFnTokens)
}

var exampleFnManyLines = `
(def dbl 
	(fn [x] 
		(* 2 x)))
`
