package mol2

//import "fmt"
import "bytes"

type Molecule struct {
	name string
	mtype, ctype byte
	atoms []Atom
	bonds []Bond
}

func (mol *Molecule) String() string {
	if mol == nil {
		return "empty molecule"
	}

	var buf bytes.Buffer
	buf.WriteString("molecule [" + mol.name + "]\n")
	buf.WriteString("type: " + MoleculeTypeToString(mol.mtype) + "\n")
	buf.WriteString("charge type: " + MoleculeChargesToString(mol.ctype) + "\n")
	buf.WriteString("atoms:\n")
	for i := 0; i < len(mol.atoms); i++ {
		buf.WriteString("\t");
		buf.WriteString(mol.atoms[i].String())
		buf.WriteString("\n")
	}
	buf.WriteString("bonds:\n")
	for i := 0; i < len(mol.bonds[i].String()); i++ {
		buf.WriteString("\t");
		buf.WriteString(mol.bonds[i].String())
		buf.WriteString("\n")
	}

	return buf.String()
}

func moleculeParse(lex *Lexer) *Molecule {
	molecule := new(Molecule)
	ok, err := lex.nextMolecule()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	lex.skipWS()
	molecule.name = lex.nextLine()
	ok, length, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	molecule.atoms = make([]Atom, length)
	ok, length, err = lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	molecule.bonds = make([]Bond, length)
	// TODO: parse other
	lex.nextInt()
	lex.nextInt()
	lex.nextInt()
	//
	ok, mtype, err := lex.nextId()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	molecule.mtype = MoleculeTypeByString(mtype)
	ok, ctype, err := lex.nextId()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	ok, err = lex.nextNL()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	molecule.ctype = MoleculeChargesByString(ctype)
	ok, err = lex.nextAtom()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	for i := 0; i < len(molecule.atoms); i ++ {
		molecule.atoms[i] = *atomParse(lex)
	}
	ok, err = lex.nextBond()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	for i := 0; i < len(molecule.bonds); i ++ {
		molecule.bonds[i] = *bondParse(lex)
	}

	return molecule
}
