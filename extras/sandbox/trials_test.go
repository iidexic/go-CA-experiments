package sandbox

import "testing"

var lastval interface{}

func ttrialRandbytes(t *testing.T) {

}

//====|Benchmarks|====
func benchtrialRandbytes(w, h int, b *testing.B) {
	size := w * h * 4
	var out []byte = make([]byte, size)
	for range b.N {
		out, _ = trialRandbytes(size)
	}
	lastval = out
}

func benchtrialFastrandbytes(w, h int, b *testing.B) {
	size := w * h * 4
	var out []byte = make([]byte, size)
	for range b.N {
		out, _ = trialFastrandbytes(size)
	}
	lastval = out
}

func BenchmarkRand720(b *testing.B) {
	benchtrialRandbytes(1280, 720, b)
}
func BenchmarkFastrand720(b *testing.B) {
	benchtrialFastrandbytes(1280, 720, b)
}

func BenchmarkRand1080(b *testing.B) {
	benchtrialRandbytes(1920, 1080, b)
}
func BenchmarkFastrand1080(b *testing.B) {
	benchtrialFastrandbytes(1920, 1080, b)
}

//** ===============================================
func benchWrapScreenConditional(w, h int, b *testing.B) {
	screen, _ := trialFastrandbytes(w * h * 4)
	for range b.N {
		_ = screen
	}
}
func benchWrapScreenMod(w, h int, b *testing.B) {

}
