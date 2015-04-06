package mol2

const (
	BOND_SINGLE byte = iota
	BOND_DOUBLE
	BOND_TRIPLE
	BOND_AMIDE
	BOND_AROMATIC
	BOND_DUMMY
	BOND_NOT_CONNECTED
	BOND_UNKNOWN
)

func BondTypeToString(bt byte) string {
	switch bt {
	case BOND_SINGLE:        return "1"
	case BOND_DOUBLE:        return "2"
	case BOND_TRIPLE:        return "3"
	case BOND_AMIDE:         return "am"
	case BOND_AROMATIC:      return "ar"
	case BOND_DUMMY:         return "du"
	case BOND_NOT_CONNECTED: return "nc"
	}

	return "un"
}

func BondTypeByString(s string) byte {
	switch s {
	case "1":  return BOND_SINGLE
	case "2":  return BOND_DOUBLE
	case "3":  return BOND_TRIPLE
	case "am": return BOND_AMIDE
	case "ar": return BOND_AROMATIC
	case "du": return BOND_DUMMY
	case "nc": return BOND_NOT_CONNECTED
	}

	return BOND_UNKNOWN
}

// TODO: add pls status bit for bond
