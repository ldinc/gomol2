package mol2

import (
	"bytes"
	"errors"
)

type Lexer struct {
	buf *bytes.Buffer
	line, column int
}

func NewLexer(buffer *bytes.Buffer) *Lexer {
	lex := new(Lexer)
	lex.line = 0
	lex.column = 0
	lex.buf = buffer

	return lex
}

func (lex *Lexer) Coords() (int, int) {
	return lex.line, lex.column
}

func (lex *Lexer) NextAtom() (bool, error) {
	pattern := "@<TRIPOS>ATOM"
	err_text := "was expected '" + pattern + "'"
	if lex.buf.Len() < len(pattern) {
		return false, errors.New(err_text)
	}
	for _, r := range pattern {
		buf_rune, _, err := lex.buf.ReadRune()
		if err != nil {
			return false, err
		}
		if r != buf_rune {
			return false, errors.New(err_text)
		}
	}
	lex.column += len(pattern)

	return true, nil
}
