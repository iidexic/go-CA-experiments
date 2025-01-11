package main

import (
	"testing"
)

var resultByteSlice []byte

func BenchmarkRandpx(b *testing.B) {
	var w, h uint = 640, 480
	qty := w * h * 4 //px count * RGBA
	var res []byte
	for range b.N {
		res = Randpx(qty)
	}
	resultByteSlice = res
}
