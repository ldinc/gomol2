# gomol2
Go MOL2 simple parser

# Description
Simple Go package for parsing [Tripos Mol2 File Format](http://www.tripos.com/data/support/mol2.pdf).

[GoMOL2 Wiki](https://github.com/ldinc/gomol2/wiki)

Feature parsing:

1. Atom
  - atom_id 
  - atom_name 
  - x, y, z 
  - atom_type
  - substructure 
    * subst_name
    * charge
    * status_bit ***TODO***

2. Bond
  - bond_id
  - origin_atom_id
  - target_atom_id
  - bond_type
  - status_bits ***TODO***

3. Molecule
  - mol_name
  - num_atoms
  - num_bonds
  - num_subst ***TODO***
  - num_feat ***TODO***
  - num_sets ***TODO***
  - mol_type
  - charge_type
  - status_bits
  - mol_comment ***TODO***

# Install & Usage
Install: `go get github.com/ldinc/gomol2/mol2`

Example:

__Parsing from file:__

```go
mol, err := mol2.ParseFile(filename)
fmt.Println(mol)
```

__Parsing from byte-slice__

```go
var buffer []byte
...
mol, err := mol2.ParseText(buffer)
fmt.Println(mol)
```
