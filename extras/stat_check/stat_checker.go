package statcheck

type genable interface {
	float64 | int | uint | int16 | uint16 | int8 | byte
}
type rngen[G genable] interface {
	getN(n int) G
	makeN(n int)
	bounds(top, btm G)
	serve() G
}
type ttype int

const (
	tfloat ttype = iota
	tint
	tuint
	tint16
	tuint16
	tint8
	tbyte
)

type rngFast[G genable] struct {
	zero     G
	bounded  bool
	btm, top G
	nA       int
	genA     []G
}

// RNGenInit RETURNS AN RNG GENERATOR USING TYPE INDICATED BY TTYPE
// This function allows a non-shared generic struct to be created.
// Maybe there's a point for someone to use this, I don't know.
func RNGenInit(t ttype) interface{} {

	switch t {
	case tfloat:
		return rngFast[float64]{}
	case tint:
		return rngFast[int]{}
	case tuint:
		return rngFast[uint]{}
	case tint16:
		return rngFast[int16]{}
	case tuint16:
		return rngFast[uint16]{}
	case tint8:
		return rngFast[int8]{}
	case tbyte:
		return rngFast[byte]{}
	default:
		panic("Use a correct type. This function doesn't even work")
	}
}

func (rng rngFast[G]) makeN(n int) {
	//var get G
}
func (rng rngFast[G]) bounds(top, btm G) {
	rng.btm = btm
	rng.top = top
}
func (rng rngFast[G]) GetN(n int)
func main() {
	var gen rngFast[int]
	gen.makeN(5)
}
