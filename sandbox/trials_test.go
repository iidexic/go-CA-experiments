package sandbox

import "testing"

var lastval interface{}

func BenchmarktrialRandbytes(w, h int, b *testing.B) {
	size := w * h * 4
	var out []byte = make([]byte, size)
	for range b.N {
		out, _ = trialRandbytes(size)
	}
	lastval = out
}
