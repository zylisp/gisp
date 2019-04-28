package reader

// stateFn state function type
type stateFn func(*LispReader) stateFn

// LispReader embeds the PositionReader struct
type LispReader struct {
	name  string
	state stateFn
	items chan Atom
	*PositionReader
}

// NewLispReader creates a LispReader for the given string and optional
//                   position stack
func NewLispReader(programName string, programData string) *LispReader {
	return &LispReader{
		programName,
		nil,
		make(chan Atom),
		NewPositionReader(programData, initPosition()),
	}
}
