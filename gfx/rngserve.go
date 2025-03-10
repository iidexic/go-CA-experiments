package gfx

import (
	"github.com/bytedance/gopkg/lang/fastrand"
)

// QuickRNG will provide a channel that loads/buffers random values
type QuickRNG struct {
	C               chan byte
	sl              []byte
	size, n, reload int
	mod             byte
}

// GetQuickRNG returns initialized QuickRNG with `buff` length channel
func GetQuickRNG(buff int) QuickRNG {
	q := QuickRNG{C: make(chan byte, buff), sl: make([]byte, buff),
		size: buff, n: buff, reload: buff / 2, mod: 1}
	return q
}

// ROPcheck checks rand val reorder point
// will call it 0 for now - wonder if this will starve it
func (q *QuickRNG) ROPcheck() {
	q.n--
	if q.n < q.reload {
		go q.genc()
		q.n = q.size
	}
}

/*
// Coin flip, return pseudo-rand int centered around 0

	func (q *QuickRNG) coin() int {
		flip := <-q.C
		return int(flip) - 127
	}
*/
func (q *QuickRNG) genc() {
	_, e := fastrand.Read(q.sl)

	if e != nil {
		panic(e)
	}
	for _, v := range q.sl {
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
