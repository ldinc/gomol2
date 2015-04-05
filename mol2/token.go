package mol2

type Token int

const (
	TokInt  Token = iota
	TokReal Token = iota
	TokId   Token = iota
	TokAtom Token = iota
)
