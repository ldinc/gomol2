package mol2

// SYBYL atom types
const (
	C3 = iota // C.3     - carbon sp3
	C2        // C.2     - carbon sp2
	C1        // C.1     - carbon sp
	CAR       // C.ar    - carbon aromatic
	CCAT      // C.cat   - carbocation (C+) used only in a guadinium group
	N3        // N.3     - nitrogen sp3
	N2        // N.2     - nitrogen sp2
	N1        // N.1     - nitrogen sp
	NAR       // N.ar    - nitrogen aromatic
	NAM       // N.am    - nitrogen amide
	NPL3      // N.pl3   - nitrogen trigonal planar
	N4        // N.4     - nitrogen sp3 positively charged
	O3        // O.3     - oxygen sp3
	O2        // O.2     - oxygen sp2
	OCO2      // O.co2   - oxygen in carboxylate and phosphate groups
	OSPC      // O.spc   - oxygen in Single Point Charge (SPC) water model
	OT3P      // O.t3p   - oxygen in Transferable Intermolecular Potential (TIP3P) water model
	S3        // S.3     - sulfur sp3
	S2        // S.2     - sulfur sp2
	SO        // S.O     - sulfoxide sulfur
	SO2       // S.O2    - sulfone sulfur
	P3        // P.3     - phosphorous sp3
	F         // F       - fluorine
	H         // H       - hydrogen
	HSPC      // H.spc   - hydrogen in Single Point Charge (SPC) water model
	HT3P      // H.t3p   - hydrogen in Transferable intermolecular Potential (TIP3P) water model
	LP        // LP      - lone pair
	DUMMY     // Du      - dummy atom
	DUMMYC    // Du.C    - dummy carbon
	ANY       // Any     - any atom
	HAL       // Hal     - halogen
	HETERO    // Het     - heteroatom = N, O, S, P
	HEAVY     // Hev     - heavy atom (non hydrogen)
	LI        // Li      - lithium
	NA        // Na      - sodium
	MG        // Mg      - magnesium
	AL        // Al      - aluminum
	SI        // Si      - silicon
	K         // K       - potassium
	CA        // Ca      - calcium
	CRTH      // Cr.th   - chromium (tetrahedral)
	CROH      // Cr.oh   - chromium (octahedral)
	MN        // Mn      - manganese
	FE        // Fe      - iron
	COOH      // Co.oh   - cobalt (octahedral)
	CU        // Cu      - copper
	CL        // Cl      - chlorine
	BR        // Br      - bromine
	I         // I       - iodine
	ZN        // Zn      - zinc
	SE        // Se      - selenium
	MO        // Mo      - molybdenum
	SN        // Sn      - tin
)

var TypesNames []string = []string {
	"C.3",
	"C.2",
	"C.1",
	"C.ar",
	"C.cat",
	"N.3",
	"N.2",
	"N.1",
	"N.ar",
	"N.am",
	"N.pl3",
	"N.4",
	"O.3",
	"O.2",
	"O.co2",
	"O.spc",
	"O.t3p",
	"S.3",
	"S.2",
	"S.O",
	"S.O2",
	"P.3",
	"F",
	"H",
	"H.spc",
	"H.t3p",
	"LP",
	"Du",
	"Du.C",
	"Any",
	"Hal",
	"Het",
	"Hev",
	"Li",
	"Na",
	"Mg",
	"Al",
	"Si",
	"K",
	"Ca",
	"Cr.th",
	"Cr.oh",
	"Mn",
	"Fe",
	"Co.oh",
	"Cu",
	"Cl",
	"Br",
	"I",
	"Zn",
	"Se",
	"Mo",
	"Sn",
}

func AtomTypeGetByString(atype string) int {
	for code := 0; code < len(TypesNames); code++ {
			if TypesNames[code] == atype {
				return code
			}
	}

	return -1
}

func AtomTypeToString(id int) string {
	if id == -1 {
		return "error"
	}

	return TypesNames[id]
}
