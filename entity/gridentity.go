package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/gfx"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img     *ebiten.Image
	X, Y    uint   // (X,Y) -> Width,Height of grid
	Bounds  []int  // image bounds on screen
	zone    *zones // zone obj, holds all zone data
	Area    int    // Area(width*Height), total nbr of cells
	nticks  byte
	Px      []byte // main color slice for the grid
	rng     []byte // Slice of rng bytes, refreshed by grid.reload()
	Op      *ebiten.DrawImageOptions
	Visible bool   // grid visibility toggle
	Debug   bool   // grid debug  toggle
	reload  func() // func called to refresh rng
	ruleset Ruleset
}

// Subsections of full Grid
/* Tracks zone stats and zone number of all pixels in grid */
type zones struct {
	numX, numY, count             int
	zsum, zrange, zcenter, zx, zy []int
	cellzonebig                   []int
	cellzone                      []uint16
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity {
	borderTot := (gWidth / 16)
	width := gWidth - borderTot
	height := gHeight - borderTot
	grid := GridEntity{
		Img:    ebiten.NewImage(width, height),
		Bounds: make([]int, 4),
		Op:     &ebiten.DrawImageOptions{},
		X:      uint(width), Y: uint(height), Area: width * height,
		Px:   make([]byte, width*height*4),
		rng:  make([]byte, width*height),
		zone: calculateZones(3, 3, width, height),
	}
	grid.reload = gfx.Fbytes(grid.rng)
	grid.Bounds[0] = (gWidth - width) / 2
	grid.Bounds[2] = grid.Bounds[0] + width

	grid.Bounds[1] = (gHeight - height) / 2
	grid.Bounds[3] = grid.Bounds[1] + height
	grid.Op.GeoM.Translate(float64(grid.Bounds[0]), float64(grid.Bounds[1]))
	grid.Img.Fill(gfx.PaletteGP[gfx.Dark])

	return &grid
}

// calculateZones returns a zones struct with the zones for the grid
func calculateZones(x, y, width, height int) *zones {
	z := zones{
		numX: x, numY: y, count: x * y,
	}
	z.zsum = make([]int, z.count)
	z.cellzone = make([]uint16, width*height)
	xpix := width / z.numX
	ypix := height / z.numY
	xmod := width % x  //x pixel remainder
	ymod := height % y //y pixel remainder
	xoffs := xmod / 2
	yoffs := ymod / 2
	oob := uint16(z.numX*z.numY + 1)
	for i := range z.cellzone {
		zposx := i % width //pixel x in (x,y)
		zposy := i / width //pixel y in (x,y)
		if zposx < xoffs || zposy < yoffs {
			z.cellzone[i] = oob
			continue
		}
		zx := (zposx - xoffs) / xpix //x-position, offset removed every time(oh ok)
		zy := (zposy - yoffs) / ypix

		//handle the only condition where end px would be zoneless:
		if zx == z.numX || zy == z.numY {
			z.cellzone[i] = oob * 100
			continue
		}
		z.cellzone[i] = uint16((zy * z.numX) + zx)
	}
	return &z
}

// pxisort EXCLUSIVELY sorts RGB bytes: returns indices of RGB smallest-to-largest.
func pxisort(pix []byte) [3]int {
	a, b, c := pix[0], pix[1], pix[2]
	if a <= b {
		if b <= c {
			return [3]int{0, 1, 2}
		}
		if a <= c {
			return [3]int{0, 2, 1}
		}
		return [3]int{2, 0, 1}
	}
	if a <= c {
		return [3]int{1, 0, 2}
	}
	if b <= c {
		return [3]int{1, 2, 0}
	}
	return [3]int{2, 1, 0}
}

// XY returns grid GeoM tx, ty screen location
func (grid *GridEntity) XY() (int, int) {
	return int(grid.Op.GeoM.Element(0, 2)), int(grid.Op.GeoM.Element(1, 2))
}

func (grid *GridEntity) getrng(i int) byte { return grid.rng[i%grid.Area] }

func (grid *GridEntity) pxtozone(i int) {
	zn := int(grid.zone.cellzone[i])
	if !(zn >= grid.zone.numX*grid.zone.numY) {
		add := int(grid.Px[i*4]) + int(grid.Px[i*4+1]) + int(grid.Px[i*4+2])
		grid.zone.zsum[zn] += add
	}
}

// bmov moves up to amt from src to dest, clamped so neither wraps.
func bmov(src, dest byte, amt byte) (byte, byte) {
	lim := min(src, 255-dest, amt)
	src -= lim
	dest += lim
	return src, dest
}

func moveToward(from, to byte, amount byte) byte {
	dist := int(to) - int(from)
	dx := (dist * int(amount)) / 255
	return byte(int(from) + dx)
}

// sliceToward moves each byte in from toward the corresponding byte in to.
// to loops when from is longer.
func sliceToward(from, to []byte, amount byte) {
	lto := len(to)
	for i := range from {
		from[i] = moveToward(from[i], to[i%lto], amount)
	}
}

func mto(bs1, bs2, dest []byte) {
	for i := range dest[:3] {
		dest[i] = moveToward(bs1[i], bs2[i], bs2[3])
	}
	dest[3] = 255
}
