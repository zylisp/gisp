package line

import (
	"bufio"
	"strings"
)

// PositionOpts ...
type PositionOpts struct {
	row      int
	column   int
	absolute int
}

// PositionReader ...
type PositionReader struct {
	PositionOpts
	reader *bufio.Reader
}

func defaultPositionOpts() []PositionOpts {
	return []PositionOpts{PositionOpts{row: 1, column: 1}}
}

// NewPositionReader ...
func NewPositionReader(stringData string, opts ...PositionOpts) *PositionReader {
	defaultOpts := defaultPositionOpts()

	if len(opts) == 0 {
		opts = defaultOpts
	}
	if opts[0].row == 0 {
		opts[0].row = defaultOpts[0].row
	}
	if opts[0].column == 0 {
		opts[0].column = defaultOpts[0].column
	}
	return &PositionReader{
		opts[0],
		bufio.NewReader(strings.NewReader(stringData))}
}

// ReadRune ...
func (r *PositionReader) ReadRune() (rune, int, error) {
	rn, sz, err := r.reader.ReadRune()
	if err != nil {
		return rn, sz, err
	}
	r.absolute++
	return rn, sz, nil
}

// UnreadRune ...
func (r *PositionReader) UnreadRune() error {
	err := r.reader.UnreadRune()
	r.absolute--
	return err
}
