package reader

import (
	"bufio"
	"strings"
)

// Position defines the data tracked in the PositionReader
type Position struct {
	row      int
	column   int
	absolute int
}

// PositionReader wraps bufio.Reader and embeds the Position struct
type PositionReader struct {
	positionStack []Position
	reader        *bufio.Reader
}

// defaultPositions defines the defaults for use by the constructor
func defaultPositions() []Position {
	return []Position{Position{row: 1, column: 1}}
}

// NewPositionReader creates a PositionReader for the given string and optional
//                   position stack
func NewPositionReader(stringData string, opts ...Position) *PositionReader {
	defaultPos := defaultPositions()

	if len(opts) == 0 {
		opts = defaultPos
	}
	if opts[0].row == 0 {
		opts[0].row = defaultPos[0].row
	}
	if opts[0].column == 0 {
		opts[0].column = defaultPos[0].column
	}
	return &PositionReader{
		opts,
		bufio.NewReader(strings.NewReader(stringData))}
}

// lastPositionIndex returns the index in the postition stack for the most
//                   recently added position
func (r *PositionReader) lastPositionIndex() int {
	return len(r.positionStack) - 1
}

// lastPosition returns the most recently added position from the position
//              stack
func (r *PositionReader) lastPosition() Position {
	return r.positionStack[r.lastPositionIndex()]
}

// deleteLastPosition deletes the most recently added position in the position
//                    stack
func (r *PositionReader) deleteLastPosition() {
	r.positionStack = r.positionStack[:r.lastPositionIndex()]
}

// pushPosition adds a new position to the position stack
func (r *PositionReader) pushPosition(pos Position) {
	r.positionStack = append(r.positionStack, pos)
}

// pushPositions adds any number of positions to the position stack; note that
//               this is an append operation, so the last item in the passed
//               positions will be interpreted as the most recently added
//               position
func (r *PositionReader) pushPositions(pos ...Position) {
	r.positionStack = append(r.positionStack, pos...)
}

// popPosition remove and returns the most recently added position from the
//             position stack
func (r *PositionReader) popPosition() Position {
	popped := r.lastPosition()
	r.deleteLastPosition()
	return popped
}

// nextRunePosition copies the most recently added position from the position
//                  stack and updates it with new values; the position is then
//                  returned. The rune passed to this method is used to
//                  determine if a newline has been read, and if so, to update
//                  the row and column apropriately
func (r *PositionReader) nextRunePosition(rn rune) Position {
	next := r.lastPosition()
	next.absolute++
	if rn == newline {
		next.column = 1
		next.row++
	} else {
		next.column++
	}
	return next
}

// Row returns the row number for the most recently added position
func (r *PositionReader) Row() int {
	return r.lastPosition().row
}

// Column returns the column number for the most recently added position
func (r *PositionReader) Column() int {
	return r.lastPosition().column
}

// Absolute returns the absolute rune location in the string data provided to the reader
func (r *PositionReader) Absolute() int {
	return r.lastPosition().absolute
}

// ReadRune calls the ReadRune function of bufio.Reader and then applies
//          position-tracking logic
func (r *PositionReader) ReadRune() (rune, int, error) {
	rn, sz, err := r.reader.ReadRune()
	if err != nil {
		return rn, sz, err
	}
	r.pushPosition(r.nextRunePosition(rn))
	return rn, sz, nil
}

// UnreadRune calls the UnreadRune function of bufio.Reader and then applies
//            position-tracking logic
func (r *PositionReader) UnreadRune() error {
	err := r.reader.UnreadRune()
	_ = r.popPosition()
	return err
}
