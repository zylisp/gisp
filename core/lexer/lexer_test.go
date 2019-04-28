package lexer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func newTestLexer(data string) *Lexer {
	return NewLexer("tester", data)
}

type LexerSuite struct {
	suite.Suite
	testLexer *Lexer
}

func TestLexerSuite(t *testing.T) {
	suite.Run(t, new(LexerSuite))
}

func (s *LexerSuite) SetupTest() {
	s.testLexer = newTestLexer("()")
}

func (s *LexerSuite) TearDownTest() {
}

func (s *LexerSuite) TestNewPosition() {
	p, _ := NewPosition()
	s.Equal(1, p.row)
	s.Equal(1, p.column)
	s.Equal(0, p.absolute)
	p, _ = NewPosition(4, 2)
	s.Equal(4, p.row)
	s.Equal(2, p.column)
	s.Equal(0, p.absolute)
	p, _ = NewPosition(4, 2, 42)
	s.Equal(4, p.row)
	s.Equal(2, p.column)
	s.Equal(42, p.absolute)
}
func (s *LexerSuite) TestNewLexer() {
	s.Equal(1, s.testLexer.position.row)
	s.Equal(1, s.testLexer.position.column)
	s.Equal(0, s.testLexer.position.absolute)
}

func (s *LexerSuite) TestCodeSize() {
	s.Equal(2, s.testLexer.codeSize())
}
func (s *LexerSuite) TestNext() {
	s.Equal(0, s.testLexer.start)
	s.Equal(2, len(s.testLexer.input))
	// returns and "consumes" the next rune
	s.Equal(0, int(s.testLexer.position.absolute))
	s.Equal("(", string(s.testLexer.next()))
	s.Equal(1, int(s.testLexer.position.absolute))
	s.Equal(")", string(s.testLexer.next()))

	// next should move the position forward but not
	// affect the start (of input)
	s.Equal(0, s.testLexer.start)
	s.Equal(2, len(s.testLexer.input))
	currentString := s.testLexer.input[s.testLexer.start:len(s.testLexer.input)]
	s.Equal("()", currentString)
}

func (s *LexerSuite) TestBackup() {
	// returns and "consumes" the next rune
	s.Equal("(", string(s.testLexer.next()))
	s.Equal(")", string(s.testLexer.next()))
	// ^^^ consumes the ) so if we backup
	// we should get it again
	s.testLexer.backup()
	s.Equal(")", string(s.testLexer.next()))
}

func (s *LexerSuite) TestPeek() {
	// peek is next, but without consuming the rune
	s.Equal("(", string(s.testLexer.peek()))
	s.Equal("(", string(s.testLexer.peek()))
	s.Equal("(", string(s.testLexer.next()))
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
	s.Equal(int(s.testLexer.position.absolute), 0)
	s.False(s.testLexer.accept(unacceptableSet))
	s.Equal(int(s.testLexer.position.absolute), 0)
	// position shouldn't have changed

	s.True(s.testLexer.accept(acceptableSet))
	s.Equal(int(s.testLexer.position.absolute), 1)
}

func (s *LexerSuite) TestAcceptRuneRun() {

	acceptableSet := "12345"
	s.testLexer.input = "123a45"
	s.testLexer.acceptRuneRun(acceptableSet)
	// position should now be 3
	// e.g. it should have nexted to the a, then gone back one
	s.Equal(int(s.testLexer.position.absolute), 3)
}

// TODO: complete Test_emit
// I, apparently, don't understand chanels enough to
// make it not blow up
// func (s *LexerSuite) TestEmit() {
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
var expectedExampleFnTokensOneLine = `
(  : row: 1, col:  1, abs:  0, type AtomLeftParen
def: row: 1, col:  2, abs:  1, type AtomIdent
dbl: row: 1, col:  6, abs:  5, type AtomIdent
(  : row: 1, col: 10, abs:  9, type AtomLeftParen
fn : row: 1, col: 11, abs: 10, type AtomIdent
[  : row: 1, col: 14, abs: 13, type AtomLeftVect
x  : row: 1, col: 15, abs: 14, type AtomIdent
]  : row: 1, col: 16, abs: 15, type AtomRightVect
(  : row: 1, col: 18, abs: 17, type AtomLeftParen
*  : row: 1, col: 19, abs: 18, type AtomIdent
2  : row: 1, col: 21, abs: 20, type AtomIdent
x  : row: 1, col: 23, abs: 22, type AtomIdent
)  : row: 1, col: 24, abs: 23, type AtomRightParen
)  : row: 1, col: 25, abs: 24, type AtomRightParen
)  : row: 1, col: 26, abs: 25, type AtomRightParen
`

func (s *LexerSuite) TestExampleFn() {
	lexed := NewLexer("a-prog", exampleFnOneLine)
	s.Equal(expectedExampleFnTokensOneLine, lexed.String())
}

// var exampleFnManyLines = `
// (def dbl
//   (fn [x]
//     (* 2 x)))
// `
// var expectedExampleFnTokenManyLines = `
// (  : row: 1, col:  1, abs:  1, type AtomLeftParen
// def: row: 1, col:  2, abs:  2, type AtomIdent
// dbl: row: 1, col:  5, abs:  6, type AtomIdent
// (  : row: 2, col:  3, abs:  12, type AtomLeftParen
// fn : row: 2, col:  4, abs:  13, type AtomIdent
// [  : row: 2, col:  7, abs:  16, type AtomLeftVect
// x  : row: 2, col:  8, abs:  17, type AtomIdent
// ]  : row: 2, col:  9, abs:  18, type AtomRightVect
// (  : row: 3, col:  5, abs:  24, type AtomLeftParen
// *  : row: 3, col:  6, abs:  25, type AtomIdent
// 2  : row: 3, col:  8, abs:  27, type AtomIdent
// x  : row: 3, col:  10, abs:  29, type AtomIdent
// )  : row: 3, col:  11, abs:  30, type AtomRightParen
// )  : row: 3, col:  12, abs:  31, type AtomRightParen
// )  : row: 3, col:  13, abs:  32, type AtomRightParen
// `

// func (s *LexerSuite) TestExampleFnManyLines() {
// 	lexed := NewLexer("a-prog", exampleFnManyLines)
// 	s.Equal(expectedExampleFnTokenManyLines, lexed.String())
// }
