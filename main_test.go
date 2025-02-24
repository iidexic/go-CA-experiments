package main

import (
	"fmt"
	"testing"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/iidexic/go-CA-experiments/gfx"
)

var resultByteSlice []byte

func BenchmarkRandpx(b *testing.B) {
	var w, h uint = 640, 480
	qty := w * h * 4 //px count * RGBA
	var res []byte
	for range b.N {
		res = gfx.Randpx(qty)
	}
	resultByteSlice = res
}

func TestSandbox(t *testing.T) {
	rv := make([]byte, 128)
	_, _ = fastrand.Read(rv)
	for _, v := range rv {
		if v%16 == 0 {
			fmt.Println(v)

		}
	}
}
