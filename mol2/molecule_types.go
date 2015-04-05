package mol2

const (
	MOL_SMALL byte = iota
	MOL_BIOPOLYMER
	MOL_PROTEIN
	MOL_NUCLEIC_ACID
	MOL_SACCHARIDE
	MOL_ERR
)

func MoleculeTypeToString(mtype byte) string {
	switch mtype {
	case MOL_SMALL:        return "SMALL"
	case MOL_BIOPOLYMER:   return "BIOPOLYMER"
	case MOL_PROTEIN:      return "PROTEIN"
	case MOL_NUCLEIC_ACID: return "NUCLEIC_ACID"
	case MOL_SACCHARIDE:   return "SACCHARIDE"
	}

	return "empty"
}

func MoleculeTypeByString(s string) byte {
	switch s {
	case "SMALL":        return MOL_SMALL
	case "BIOPOLYMER":   return MOL_BIOPOLYMER
	case "PROTEIN":      return MOL_PROTEIN
	case "NUCLEIC_ACID": return MOL_NUCLEIC_ACID
	case "SACCHARIDE":   return MOL_SACCHARIDE
	}

	return MOL_ERR
}

const (
	CHARGES_NO byte = iota
	CHARGES_DEL_RE
	CHARGES_GASTEIGER
	CHARGES_GAST_HUCK
	CHARGES_HUCKEL
	CHARGES_PULLMAN
	CHARGES_GAUSS80
	CHARGES_AMPAC
	CHARGES_MULLIKEN
	CHARGES_DICT
	CHARGES_MMFF94
	CHARGES_USER
	CHARGES_ERR
)

func MoleculeChargesToString(charges byte) string {
	switch charges {
	case CHARGES_NO:        return "NO_CHARGES"
	case CHARGES_DEL_RE:    return "DEL_RE"
	case CHARGES_GASTEIGER: return "GASTEIGER"
	case CHARGES_GAST_HUCK: return "GAST_HUCK"
	case CHARGES_HUCKEL:    return "HUCKEL"
	case CHARGES_PULLMAN:   return "PULLMAN"
	case CHARGES_GAUSS80:   return "GAUSS80_CHARGES"
	case CHARGES_AMPAC:     return "AMPAC_CHARGES"
	case CHARGES_MULLIKEN:  return "MULLIKEN_CHARGES"
	case CHARGES_DICT:      return "DICT_CHARGES"
	case CHARGES_MMFF94:    return "MMFF94_CHARGES"
	case CHARGES_USER:      return "USER_CHARGES"
	}

	return "none"
}

func MoleculeChargesByString(s string) byte {
	switch s {
	case "NO_CHARGES":       return CHARGES_NO
	case "DEL_RE":           return CHARGES_DEL_RE
	case "GASTEIGER":        return CHARGES_GASTEIGER
	case "GAST_HUCK":        return CHARGES_GAST_HUCK
	case "HUCKEL":           return CHARGES_HUCKEL
	case "PULLMAN":          return CHARGES_PULLMAN
	case "GAUSS80_CHARGES":  return CHARGES_GAUSS80
	case "AMPAC_CHARGES":    return CHARGES_AMPAC
	case "MULLIKEN_CHARGES": return CHARGES_MULLIKEN
	case "DICT_CHARGES":     return CHARGES_DICT
	case "MMFF94_CHARGES":   return CHARGES_MMFF94
	case "USER_CHARGES":     return CHARGES_USER
	}

	return CHARGES_ERR
}
