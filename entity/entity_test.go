package entity

import (
	"fmt"
	"slices"
	"testing"
)

func TestSidewrap(t *testing.T) {
	out1 := sidewrap(10, -1, 10)
	out2 := sidewrap(158976, -780, 768)
	out3 := sidewrap(255, 1, 128)
	out4 := sidewrap(0, -1, 128)
	switch {
	case out1 != 19:
		t.Logf("out1=%d, should be 19\n", out1)
		t.Fail()
	case out2 != 159732:
		t.Logf("out1 = %d, should be 159732\n", out2)
		t.Fail()
	case out3 != 128:
		t.Logf("out3 = %d, should be 128\n", out3)
		t.Fail()
	case out4 != 127:
		t.Logf("out4 = %d, should be 127\n", out4)
		t.Fail()
	}
}
func TestSort(t *testing.T) {
	var vals [][]byte = [][]byte{{7, 8, 9}, {7, 9, 8}, {8, 7, 9}, {8, 9, 7}, {9, 7, 8}, {9, 8, 7}}
	var result [][]int = [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {2, 0, 1}, {1, 2, 0}, {2, 1, 0}}
	for i := range 6 {
		out := pxisort(vals[i])
		if !slices.Equal(out, result[i]) {
			t.Log("out:", out, "len=", len(out), "\nresult:", result[i], "len=", len(result[i]))
			t.Fail()
		}
	}
	fmt.Println("test end")
}
func TestMto(t *testing.T) {
	in1 := [][]byte{{255, 128, 200, 0}}
	in2 := [][]byte{{33, 230, 180, 255}}
	expected := [][]byte{{33, 230, 180, 255}}
	tempresult := make([]byte, 4)
	for i := range in1 {
		mto(in1[i], in2[i], tempresult)
		if !slices.Equal(tempresult, expected[i]) {
			t.Log("expected vs actual ->", expected[i], "vs", tempresult)
			t.Fail()
		}
	}
}

func benchmarkGridLVSD(b *testing.B, x, y int) {
	g := MakeGridDefault(x, y)
	for b.Loop() {
		g.SimstepLVSD(true)
	}
}

func BenchmarkGridLVSD_100x100(b *testing.B)   { benchmarkGridLVSD(b, 100, 100) }
func BenchmarkGridLVSD_200x200(b *testing.B)   { benchmarkGridLVSD(b, 200, 200) }
func BenchmarkGridLVSD_500x500(b *testing.B)   { benchmarkGridLVSD(b, 500, 500) }
func BenchmarkGridLVSD_1000x1000(b *testing.B) { benchmarkGridLVSD(b, 1000, 1000) }
func BenchmarkGridLVSD_2500x2500(b *testing.B) { benchmarkGridLVSD(b, 2500, 2500) }

func BenchmarkAllGridLVSD(b *testing.B) {
}
