package mol2

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	buf []byte
	line, column, pos int
	column_new, line_new, pos_new int
}

func NewLexer(buffer []byte) *Lexer {
	lex := new(Lexer)
	lex.line = 0
	lex.column = 0
	lex.pos = 0
	lex.buf = buffer

	return lex
}

func (lex *Lexer) Coords() (int, int) {
	return lex.line, lex.column
}

func (lex *Lexer) SkipWS() (bool, int) {
	return false, 0
}

func (lex *Lexer) readRune() rune {
	r, size := utf8.DecodeRune(lex.buf[lex.pos_new:])
	lex.pos_new += size
	if r == '\n' {
		lex.column_new = 0
		lex.line_new ++
	} else {
		lex.column_new ++
	}

	return r
}

func (lex *Lexer) fixCoords() {
	lex.pos = lex.pos_new
	lex.line = lex.line_new
	lex.column = lex.column_new
}

func (lex *Lexer) dropCoords() {
	lex.pos_new = lex.pos
	lex.line_new = lex.line
	lex.column_new = lex.column
}

type lexerState struct {
		pos, line, column int
		pos_new, line_new, column_new int
}

func (lex *Lexer) pushState() *lexerState {
	state := new(lexerState)
	state.pos = lex.pos
	state.pos_new = lex.pos_new
	state.line = lex.line
	state.line_new = lex.line_new
	state.column = lex.column
	state.column_new = lex.column_new

	return state
}

func (lex *Lexer) popState(state *lexerState) {
	lex.pos = state.pos
	lex.pos_new = state.pos_new
	lex.line = state.line
	lex.line_new = state.line_new
	lex.column = state.column
	lex.column_new = state.column_new
}

func (lex *Lexer) nextAtom() (bool, error) {
	pattern := "@<TRIPOS>ATOM"
	err_text := "was expected '" + pattern + "'"
	if (len(lex.buf) - lex.pos) < len(pattern) {
		return false, errors.New(err_text)
	}
	for _, r := range pattern {
		buf_rune := lex.readRune()
		if r != buf_rune {
			lex.dropCoords()
			return false, errors.New(err_text)
		}
	}
	lex.fixCoords()

	return true, nil
}

func (lex *Lexer) nextId() (bool, string, error) {
	buf := ""
	r := lex.readRune()
	for unicode.IsSpace(r) == false {
		buf += string(r)
		r = lex.readRune()
	}
	if len(buf) == 0 {
		lex.dropCoords()
		return false, "", errors.New("Id was expected")
	}
	lex.fixCoords()

	return true, buf, nil
}

func (lex *Lexer) NextReal() (bool, float64, error) {
	state := lex.pushState()
	err_text := "real was expected"
	ok, id, _ := lex.nextId()
	if ok {
		f, err := strconv.ParseFloat(id, 64)
		if err != nil {
			lex.popState(state)

			return false, 0.0, errors.New(err_text)
		}

		return true, f, nil
	}
	lex.popState(state)

	return false, 0.0, errors.New(err_text)
}
