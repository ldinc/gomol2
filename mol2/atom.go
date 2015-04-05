package mol2

import (
	"strconv"
)

type Atom struct {
	id int
	name string
	x, y, z float64
	atype int
	subst *AtomSubStructure
}

type AtomSubStructure struct {
	id int
	name string
	charge float64
	status AtomSubStatus
}

type AtomSubStatus byte

const (
	DSPMOD    AtomSubStatus = 1 << iota
	TYPECOL   AtomSubStatus = 1<< iota
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
	sub.status = status
	sub.name = name
	sub.charge = charge
	sub.id = id

	return sub
}

func (atom Atom) String() string {
		buffer := ""
		buffer += "id = " + strconv.FormatInt(int64(atom.id), 10) + "\n"
		buffer += "name = " + atom.name + "\n"
		buffer += "coords = (" + strconv.FormatFloat(atom.x, 'e', -1, 64) + ","
		buffer += strconv.FormatFloat(atom.y, 'e', -1, 64) + ","
		buffer += strconv.FormatFloat(atom.z, 'e', -1, 64) + ")\n"


		return buffer
}

func AtomParse(lex *Lexer) *Atom {
	atom := new(Atom)
	ok, err := lex.NextAtom()
	if err != nil {
		panic(err)
	}
	if ok {
		ok, id, err := lex.nextInt()
		if err != nil {
			panic(err)
		}
		if !ok {
			return nil
		}
		atom.id = id
		ok, name, err := lex.nextId()
		if err != nil {
			panic(err)
		}
		if !ok {
			return nil
		}
		atom.name = name
		ok, f, err := lex.nextReal()
		if err != nil {
			panic(err)
		}
		if !ok {
				return nil
		}
		atom.x = f;
		ok, f, err := lex.nextReal()
		if err != nil {
			panic(err)
		}
		if !ok {
				return nil
		}
		atom.y = f;
		ok, f, err := lex.nextReal()
		if err != nil {
			panic(err)
		}
		if !ok {
				return nil
		}
		atom.z = f;

		return atom
	}

	return nil
}
