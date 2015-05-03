package mol2

import (
	"bytes"
	"strconv"
)

type Bond struct {
	id int
	origin int
	target int
	btype byte
	/*
	 * TODO: add supporting for status_bit
	 * at current, it is ignored
	 */
	status byte
}

func (bond *Bond) String() string {
	if bond == nil {
		return "nil"
	}
	var buf bytes.Buffer
	buf.WriteString("[" + strconv.FormatInt(int64(bond.id), 10) + "]")
	buf.WriteString("(" + strconv.FormatInt(int64(bond.origin), 10) + ")")
	buf.WriteString("->(" + strconv.FormatInt(int64(bond.target), 10) + ")")
	buf.WriteString("{" + BondTypeToString(bond.btype) + "}")

	return buf.String()
}

func bondParse(lex *Lexer) *Bond {
	b := new(Bond)
	ok, id, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.id = id
	ok, origin, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.origin = origin
	ok, target, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.target = target
	ok, btype, err := lex.nextId()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.btype = BondTypeByString(btype)

	return b
}
