package gfx

import (
	"github.com/bytedance/gopkg/lang/fastrand"
)

// QuickRNG will provide a channel that loads/buffers random values
type QuickRNG struct {
	C         chan byte
	sl        []byte
	n, reload int
	mod       byte
}

// GetQuickRNG returns initialized QuickRNG with `buff` length channel
func GetQuickRNG(buff int) QuickRNG {
	q := QuickRNG{C: make(chan byte, buff), sl: make([]byte, buff),
		n: buff, reload: buff / 8, mod: 1}
	return q
}

// ROPcheck checks rand val reorder point
// will call it 0 for now - wonder if this will starve it
func (q *QuickRNG) ROPcheck() {
	amount := len(q.C)
	if amount == q.reload {
		go q.genc()
	}
}
func (q *QuickRNG) genc() {
	fillqty := q.n - len(q.C)
	_, e := fastrand.Read(q.sl[:fillqty])

	if e != nil {
		panic(e)
	}
	for _, v := range q.sl[:fillqty] {
		q.C <- v
	}
}

/*//> Started to get lost in the sauce
type nuber interface {
	~uint | ~int | ~float64 | ~float32 | ~byte
}

type seedvalue int

type rngsot struct {
	init bool
	seed seedvalue
}

// Seed makes a seed for rngsot from any nuber
func Seed[N nuber](input N) seedvalue {
	return seedvalue(math.Round(float64(input)))
}
*/
