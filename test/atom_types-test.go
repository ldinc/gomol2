package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/ldinc/gomol2/mol2"
)

func main() {
	for i := 0; i <= mol2.SN; i++ {
		fmt.Println(i, " : ", mol2.AtomTypeToString(i))
	}
	_bytes, err := ioutil.ReadFile("test.mol2")
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(_bytes)
	lex := mol2.NewLexer(buffer)
	fmt.Println(lex.NextAtom())
}
