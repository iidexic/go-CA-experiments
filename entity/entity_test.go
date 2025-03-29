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
func TestCalcZone0(t *testing.T)     { testCalcZones(t, 5, 7, 23, 20) }
func TestCZSmallSquare(t *testing.T) { testCalcZones(t, 3, 3, 8, 8) }
func TestCZSmallRectX(t *testing.T)  { testCalcZones(t, 3, 3, 10, 8) }

func testCalcZones(t *testing.T, zx, zy, width, height int) {
	z1 := calculateZones(zx, zy, width, height)
	zlog := make([][]int, z1.count)
	//zassign := make([]int, len(z1.cellzone))
	t.Log("numX,numY:", z1.numX, z1.numY, "count:", z1.count, "zlog empty:", zlog)
	for i, v := range z1.cellzone { //=(i of px, zone #)
		if !(v >= uint16(zx*zy)) {
			z1.zsum[int(v)] += 3
			zlog[v] = append(zlog[v], i)
		}
	}
	t.Log("zlog filled:", zlog)
	var fails []int
	zoneqtyX := width / zx
	zoneqtyY := height / zy
	for i, val := range z1.zsum {
		if val != zoneqtyX*zoneqtyY*3 {
			fails = append(fails, i)
		}
	}
	if len(fails) > 0 {
		t.Logf("zone count=%d", z1.count)
		for _, v := range fails {
			t.Log("zone:", v, "sum:", z1.zsum[v], "zone px:", zlog[v])
		}
		t.Fail()
	}
	si := make([]int, len(z1.cellzone))
	for i := range z1.cellzone {
		si[i] = i
	}
	czi := to2D(si, width)
	cellzone2D := to2D(z1.cellzone, width)
	t.Log("all assigned zones:")
	for i := range cellzone2D {
		t.Log(czi[i])
		t.Log(cellzone2D[i])
	}

	//t.Log("px-i", czi, "cellzone:", cellzone2D)
}

func TestTo2D(t *testing.T) { //Passing
	schek := []int{4, 3, 2, 1, 8, 7, 6, 5, 0, 0, 0, 0, 16, 15, 14, 13, 20, 19, 18, 17}
	sresult := [][]int{{4, 3, 2, 1}, {8, 7, 6, 5}, {0, 0, 0, 0}, {16, 15, 14, 13}, {20, 19, 18, 17}}
	output := to2D(schek, 4)
	for y := range output {
		for x := range output[y] {
			if output[y][x] != sresult[y][x] {
				t.Logf("y:%d,x:%d,got:%d,wanted:%d", y, x, output[y][x], sresult[y][x])
				t.Fail()
			}
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

// AddBorder inline in MakeGridDefault, was to test zone edge case that no longer exists.
/*
func TestAddBorder(t *testing.T, bDiv, totW, totH int) []int {
	borderPX := totW / bDiv
	bhalf := borderPX / 2
	inW := totH - borderPX
	inH := totH - borderPX
	return []int{inW, inH, bhalf, bhalf, bhalf + inW, bhalf + inH}
}
*/
