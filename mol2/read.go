package mol2

import "io/ioutil"

func ParseFile(filename string) (*Molecule, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lex := newLexer(buffer)

	return moleculeParse(lex), nil
}

func ParseText(text  []byte) (*Molecule, error) {
	lex := newLexer(text)
	return moleculeParse(lex), nil
}
