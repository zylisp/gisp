package lexer

import (
	. "github.com/rcarmo/check"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type LexerSuite struct{}

var _ = Suite(&LexerSuite{})

var testLexer Lexer

func (s *LexerSuite) SetUpTest(c *C) {
	testLexer = newTestLexer("()")
}

func newTestLexer(lexMe string) Lexer {
	return Lexer{
		name:  "tester",
		input: lexMe,
		items: make(chan Atom),
	}
}

func (s *LexerSuite) Test_next(c *C) {
	// returns and "consumes" the next rune
	c.Assert(int(testLexer.pos), Equals, 0)
	c.Assert(testLexer.next(), Equals, '(')
	c.Assert(int(testLexer.pos), Equals, 1)
	c.Assert(testLexer.next(), Equals, ')')

	// next should move the position forward but not
	// affect the start (of input)
	currentString := testLexer.input[testLexer.start:len(testLexer.input)]
	c.Assert(currentString, Equals, "()")
}

func (s *LexerSuite) Test_backup(c *C) {
	// returns and "consumes" the next rune
	c.Assert(testLexer.next(), Equals, '(')
	c.Assert(testLexer.next(), Equals, ')')
	// ^^^ consumes the ) so if we backup
	// we should get it again
	testLexer.backup()
	c.Assert(testLexer.next(), Equals, ')')
}

func (s *LexerSuite) Test_peek(c *C) {
	// peek is next, but without consuming the rune
	c.Assert(testLexer.peek(), Equals, '(')
	c.Assert(testLexer.peek(), Equals, '(')
	c.Assert(testLexer.next(), Equals, '(')
}

func (s *LexerSuite) Test_ignore(c *C) {
	// returns and "consumes" the next rune
	//c.Assert(testLexer.emit(AtomString), Equals, "BS")
	testLexer.next()   // moves position but not start. See Test_next
	testLexer.ignore() // moves start to current position.
	currentString := testLexer.input[testLexer.start:len(testLexer.input)]
	c.Assert(currentString, Equals, ")")
}

func (s *LexerSuite) Test_accept(c *C) {
	// accept consumes the next rune *if* it's from the valid set.

	acceptableSet := "abc(d"
	unacceptableSet := "nope"
	c.Assert(int(testLexer.pos), Equals, 0)
	c.Assert(testLexer.accept(unacceptableSet), IsFalse)
	c.Assert(int(testLexer.pos), Equals, 0)
	// position shouldn't have changed

	c.Assert(testLexer.accept(acceptableSet), IsTrue)
	c.Assert(int(testLexer.pos), Equals, 1)
}

func (s *LexerSuite) Test_acceptRuneRun(c *C) {

	acceptableSet := "12345"
	testLexer.input = "123a45"
	testLexer.acceptRuneRun(acceptableSet)
	// position should now be 3
	// e.g. it should have nexted to the a, then gone back one
	c.Assert(int(testLexer.pos), Equals, 3)
}

// TODO: complete Test_emit
// I, apparently, don't understand chanels enough to
// make it not blow up
// func (s *LexerSuite) Test_emit(c *C) {
// }

// TODO: Test_NextAtom
// TODO: Test_lexWhitespace
// TODO: Test_lexString
// TODO: Test_lexIdentifier
// TODO: Test_lexComment
// TODO: Test_lexNumber
// TODO: Test_scanNumber
func (s *LexerSuite) Test_scanNumber(c *C) {
	testLexer = newTestLexer("1234")
	c.Assert(testLexer.scanNumber(), IsTrue)
	testLexer = newTestLexer("-1234")
	c.Assert(testLexer.scanNumber(), IsTrue)
	testLexer = newTestLexer("+1234")
	c.Assert(testLexer.scanNumber(), IsTrue)
	testLexer = newTestLexer("1234.56")
	c.Assert(testLexer.scanNumber(), IsTrue)
	testLexer = newTestLexer("-1234.456e+78")
	c.Assert(testLexer.scanNumber(), IsTrue)
	testLexer = newTestLexer("0x1c8")
	c.Assert(testLexer.scanNumber(), IsTrue)
	//TODO: add a test for imaginary numbers (start with i)

	// totally not a number
	testLexer = newTestLexer("poopy")
	c.Assert(testLexer.scanNumber(), IsFalse)
}
