package line

import (
	"bufio"
	"strings"
)

// Position ...
type Position struct {
	row      int
	column   int
	absolute int
}

// PositionReader ...
type PositionReader struct {
	positionStack []Position
	reader        *bufio.Reader
}

func defaultPositions() []Position {
	return []Position{Position{row: 1, column: 1}}
}

// NewPositionReader ...
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

// lastPositionIndex ...
func (r *PositionReader) lastPositionIndex() int {
	return len(r.positionStack) - 1
}

// lastPosition ...
func (r *PositionReader) lastPosition() Position {
	return r.positionStack[r.lastPositionIndex()]
}

// deleteLastPosition ...
func (r *PositionReader) deleteLastPosition() {
	r.positionStack = r.positionStack[:r.lastPositionIndex()]
}

// pushPosition ...
func (r *PositionReader) pushPosition(pos Position) {
	r.positionStack = append(r.positionStack, pos)
}

// pushPositions ...
func (r *PositionReader) pushPositions(pos ...Position) {
	r.positionStack = append(r.positionStack, pos...)
}

// popPosition ...
func (r *PositionReader) popPosition() Position {
	popped := r.lastPosition()
	r.deleteLastPosition()
	return popped
}

// nextRunePosition ...
func (r *PositionReader) nextRunePosition() Position {
	next := r.lastPosition()
	next.absolute++
	return next
}

// Row ...
func (r *PositionReader) Row() int {
	return r.lastPosition().row
}

// Column ...
func (r *PositionReader) Column() int {
	return r.lastPosition().column
}

// Absolute ...
func (r *PositionReader) Absolute() int {
	return r.lastPosition().absolute
}

// ReadRune ...
func (r *PositionReader) ReadRune() (rune, int, error) {
	rn, sz, err := r.reader.ReadRune()
	if err != nil {
		return rn, sz, err
	}
	r.pushPosition(r.nextRunePosition())
	return rn, sz, nil
}

// UnreadRune ...
func (r *PositionReader) UnreadRune() error {
	err := r.reader.UnreadRune()
	_ = r.popPosition()
	return err
}
