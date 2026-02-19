package lexer

import "go-interpreter/src/01/monkey/token"

// note: this only supports ASCII
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// have we reached the end of the input?
	if l.readPosition >= len(l.input) {
		// no read yet / end of file
		l.ch = 0 // "NUL"
	} else {
		// set to next char
		l.ch = l.input[l.readPosition]
	}
	// update last read position
	l.position = l.readPosition
	// increment next read position
	l.readPosition += 1
}
