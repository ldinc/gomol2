package mol2

//import "fmt"
import "bytes"

type Molecule struct {
	Name string
	MType, CType byte
	Atoms []Atom
	Bonds []Bond
}

func (mol *Molecule) String() string {
	if mol == nil {
		return "empty molecule"
	}

	var buf bytes.Buffer
	buf.WriteString("molecule [" + mol.Name + "]\n")
	buf.WriteString("type: " + MoleculeTypeToString(mol.MType) + "\n")
	buf.WriteString("charge type: " + MoleculeChargesToString(mol.CType) + "\n")
	buf.WriteString("atoms:\n")
	for i := 0; i < len(mol.Atoms); i++ {
		buf.WriteString("\t");
		buf.WriteString(mol.Atoms[i].String())
		buf.WriteString("\n")
	}
	buf.WriteString("bonds:\n")
	for i := 0; i < len(mol.Bonds); i++ {
		buf.WriteString("\t");
		buf.WriteString(mol.Bonds[i].String())
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
	molecule.Name = lex.nextLine()
	ok, length, err := lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	molecule.Atoms = make([]Atom, length)
	ok, length, err = lex.nextInt()
	if err != nil {
		panic(err)
	}
	if !ok {
		return nil
	}
	molecule.Bonds = make([]Bond, length)
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
	molecule.MType = MoleculeTypeByString(mtype)
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
	molecule.CType = MoleculeChargesByString(ctype)
	ok, err = lex.nextAtom()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	for i := 0; i < len(molecule.Atoms); i ++ {
		molecule.Atoms[i] = *atomParse(lex)
	}
	ok, err = lex.nextBond()
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	for i := 0; i < len(molecule.Bonds); i ++ {
		molecule.Bonds[i] = *bondParse(lex)
	}

	return molecule
}
