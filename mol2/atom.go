package mol2

import (
	"bytes"
	"strconv"
)

type Atom struct {
	Id int
	Name string
	X, Y, Z float64
	Type int
	Subst *AtomSubStructure
}

type AtomSubStructure struct {
	Id int
	Name string
	Charge float64
	Status AtomSubStatus
}

type AtomSubStatus byte

const (
	DSPMOD    AtomSubStatus = 1 << iota
	TYPECOL   AtomSubStatus = 1 << iota
	CAP       AtomSubStatus = 1 << iota
	BACKBONE  AtomSubStatus = 1 << iota
	DICT      AtomSubStatus = 1 << iota
	ESSENTIAL AtomSubStatus = 1 << iota
	WATER     AtomSubStatus = 1 << iota
	DIRECT    AtomSubStatus = 1 << iota
)

func NewAtomSubStructure(id int,
                         name string,
                         charge float64,
                         status AtomSubStatus) *AtomSubStructure {
	sub := new(AtomSubStructure)
	sub.Status = status
	sub.Name = name
	sub.Charge = charge
	sub.Id = id

	return sub
}

func (subst *AtomSubStructure) String() string {
	if subst == nil {
			return "none"
	}
	var buf bytes.Buffer
	buf.WriteString("{id = " + strconv.FormatInt(int64(subst.Id), 10) + "; ")
	buf.WriteString("name = " + subst.Name + "; ")
	buf.WriteString("charge = " + strconv.FormatFloat(subst.Charge, 'e', -1, 64))
	// TODO: add status to display
	buf.WriteString("}")

	return buf.String()
}

func (atom Atom) String() string {
	buffer := ""
	buffer += strconv.FormatInt(int64(atom.Id), 10)
	buffer += " [" + atom.Name + "]["
	buffer += strconv.FormatFloat(atom.X, 'e', -1, 64) + ","
	buffer += strconv.FormatFloat(atom.Y, 'e', -1, 64) + ","
	buffer += strconv.FormatFloat(atom.Z, 'e', -1, 64) + "]"
	buffer += ": " + AtomTypeToString(atom.Type) + " "
	buffer += "subst = " + atom.Subst.String()

	return buffer
}

func atomParse(lex *Lexer) *Atom {
	atom := new(Atom)
	ok, id, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	atom.Id = id
	ok, name, err := lex.nextId()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	atom.Name = name
	ok, f, err := lex.nextReal()
	if err != nil {
		panic(err)
	}
	if !ok {
			return nil
	}
	atom.X = f;
	ok, f, err = lex.nextReal()
	if err != nil {
		panic(err)
	}
	if !ok {
			return nil
	}
	atom.Y = f;
	ok, f, err = lex.nextReal()
	if err != nil {
		panic(err)
	}
	if !ok {
			return nil
	}
	atom.Z = f;
	ok, atype, err := lex.nextId()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	atom.Type = AtomTypeGetByString(atype)
	atom.Subst = atomSubStructureParse(lex)

	return atom
}

func atomSubStructureParse(lex *Lexer) *AtomSubStructure {
	ok, err := lex.nextNL()
	if ok {
		return nil
	}
	ok, id, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	ok, name, err := lex.nextId()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	ok, charge, err := lex.nextReal()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	// TODO: add status parsing

	return NewAtomSubStructure(id, name, charge, 0)
}
