package main

import (
	"fmt"
	"mol2"
)

func main() {
	for i := 0; i <= mol2.SN; i++ {
		fmt.Println(i, " : ", mol2.AtomTypeToString(i))
	}
}
