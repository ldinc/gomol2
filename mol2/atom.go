package mol2

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
