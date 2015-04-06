package mol2

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
	buffer := ""
	buffer += "molecule [" + mol.name + "]\n"
	buffer += "type: " + MoleculeTypeToString(mol.mtype) + "\n"
	buffer += "charge type: " + MoleculeChargesToString(mol.ctype) + "\n"
	buffer += "atoms:\n"
	for _, atom := range mol.atoms {
		buffer += atom.String() + "\n"
	}
	buffer += "bonds:\n"
	for _, bond := range mol.bonds {
		buffer += bond.String() + "\n"
	}

	return buffer
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
