package mol2

import (
	"bytes"
	"strconv"
)

type Bond struct {
	Id int
	Origin int
	Target int
	Type byte
	/*
	 * TODO: add supporting for status_bit
	 * at current, it is ignored
	 */
	Status byte
}

func (bond *Bond) String() string {
	if bond == nil {
		return "nil"
	}
	var buf bytes.Buffer
	buf.WriteString("[" + strconv.FormatInt(int64(bond.Id), 10) + "]")
	buf.WriteString("(" + strconv.FormatInt(int64(bond.Origin), 10) + ")")
	buf.WriteString("->(" + strconv.FormatInt(int64(bond.Target), 10) + ")")
	buf.WriteString("{" + BondTypeToString(bond.Type) + "}")

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
	b.Id = id
	ok, origin, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.Origin = origin
	ok, target, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.Target = target
	ok, btype, err := lex.nextId()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	b.Type = BondTypeByString(btype)

	return b
}
