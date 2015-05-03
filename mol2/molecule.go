package mol2

//import "fmt"
import "bytes"

type Molecule struct {
	name string
	mtype, ctype byte
	atoms []Atom
	bonds []Bond
}

//http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func (mol *Molecule) String() string {
	if mol == nil {
		return "empty molecule"
	}
/*	buffer := ""
	buffer += "molecule [" + mol.name + "]\n"
	buffer += "type: " + MoleculeTypeToString(mol.mtype) + "\n"
	buffer += "charge type: " + MoleculeChargesToString(mol.ctype) + "\n"
	buffer += "atoms:\n"
	for i := 0; i < len(mol.atoms); i++ {
		fmt.Println(i)
		buffer += mol.atoms[i].String() + "\n"
	}
	buffer += "bonds:\n"
	for i := 0; i < len(mol.bonds); i++ {
		buffer += mol.bonds[i].String() + "\n"
	}

	return buffer
*/
/*	buffer := make([]byte, 2048)
	bpos := copy(buffer[:], []byte("molecule ["))
	bpos += copy(buffer[bpos:], []byte(mol.name))
	bpos += copy(buffer[bpos:], []byte("]\ntype: " + MoleculeTypeToString(mol.mtype) + "\n" ))
	bpos += copy(buffer[bpos:], []byte("charge type: " + MoleculeChargesToString(mol.ctype) + "\n" ))
	bpos += copy(buffer[bpos:], []byte("atoms:\n"))
	for i := 0; i < len(mol.atoms); i++ {
			//fmt.Println(mol.atoms[i].String())
		bpos += copy(buffer[bpos:], []byte(mol.atoms[i].String() + "\n"))
		bpos += copy(buffer[bpos:], []byte("\n"))
	}
*/
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
/*
	buffer := ""
	buffer = append(buffer, "molecule [" + mol.name + "]\n")
	buffer = append(buffer, "type: " + MoleculeTypeToString(mol.mtype) + "\n" )
	buffer = append(buffer, "charge type: " + MoleculeChargesToString(mol.ctype) + "\n" )
	buffer = append(buffer, "atoms:\n")
	for i := 0; i < len(mol.atoms); i++ {
		buffer = append(buffer, mol.atoms[i].String() + "\n")
	}
	return buffer
*/
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
