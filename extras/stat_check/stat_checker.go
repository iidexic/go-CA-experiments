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

/* this not work
func genInit(tt ttype) rngFast {

	switch tt {
	case tfloat:
		return rngFast[float64]{zero: float64(0.00)}
	case tint:
		return rngFast[int]{zero: int(0)}
	case tuint:
		return rngFast[uint]{zero: uint(0)}
	case tint16:
		return rngFast[int16]{zero: int16(0)}
	case tuint16:
		return rngFast[uint16]{zero: uint16(0)}
	case tint8:
		return rngFast[int8]{zero: int8(0)}
	case tbyte:
		return rngFast[byte]{zero: byte(0)}
	default:
		panic("Use a correct type. This function doesn't even work")
	}
}

func (rng rngFast[G]) makeN(n int) {
	var get G
}
func (rng rngFast[G]) bounds(top, btm G) {
	rng.btm = btm
	rng.top = top
}
func (rng rngFast[G]) getN(n int)
func main() {
	var gen rngFast[int]
	gen.makeN(5)
}
*/
