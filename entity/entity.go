package entity

import (
	"image"

	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/hajimehoshi/ebiten/v2"
)

// =======================
const (
	simShiftMod int = iota
	simLightVSDark
	simFloorCycle
)
const halfL int = 3 * 0xff / 2

// BaseEntity default entity type/debug entity
type BaseEntity struct {
	img        *ebiten.Image
	r          image.Rectangle
	op         *ebiten.DrawImageOptions
	set, drawn bool
}

// wrap val to remain within limit
func wrap(val, limit int) int {
	return ((val % limit) + limit) % limit
}

// TestSimulate is a testing simulation logic step
// will crash if grid.Pixel is not initialized.
func (grid *GridEntity) TestSimulate(shift, modifier int) {
	gX := int(grid.X)
	//	gY := int(grid.Y)
	testSimType := simLightVSDark
	imax := len(grid.Pixels)
	switch testSimType {
	case simShiftMod:
		for i := range grid.Pixels {
			grid.Pixels[i] = grid.Pixels[((i+shift)<<modifier)%imax]
		}
	case simLightVSDark:
		rng := make([]byte, grid.Area)
		_, _ = fastrand.Read(rng) //lazy, but fastrand directly returns nil
		//[Light vs Dark colorSum]
		//!Fails Immediately. Writing into 
		for i := 0; i < len(grid.Pixels); i += 4 {

			up := i - gX
			left := i - 4
			//Wrap: Wrap top-bottom and left-right
			if up < 0 {
				up = wrap(up, imax)
			}
			if left%gX == gX-1 {
				left += gX
			}
			//up-left color avgs
			cavg := (int(grid.Pixels[i]) + int(grid.Pixels[i+1]) + int(grid.Pixels[i+2])) / 3

			cavgU := coloravg(grid.Pixels[up : up+3])
			cavgL := coloravg(grid.Pixels[left : left+3])
			isLight := cavg > 127
			var iwinvU, iwinvL bool
			// calculting VS. Note that rng is same for all  checks at 1 position(i)
			if bu := cavgU > 127; bu != isLight { //if opposites
				if bu { // if up light
					iwinvU = !colorDistanceVS(cavgU, cavg, 127, int(rng[i%grid.Area]))
				} else {
					iwinvU = colorDistanceVS(cavg, cavgU, 127, int(rng[i%grid.Area]))
				}
			}
			if bl := cavgL > 127; bl != isLight {
				if bl { // if left light
					iwinvL = !colorDistanceVS(cavgL, cavg, 127, int(rng[i%grid.Area]))
				} else {
					iwinvL = colorDistanceVS(cavg, cavgL, 127, int(rng[i%grid.Area]))
				}
			}
			if iwinvU {

			} else {

			}
			if iwinvL {
				for cw := range grid.Pixels[left : left+2] {
					grid.Pixels[left+cw] = grid.Pixels[left+cw] + moveBhalf(grid.Pixels[left+cw], grid.Pixels[i+cw])
				}
			}
		}
	case simFloorCycle:

	}
}
func colorApplyWin() {

}
func moveBhalf(from, to byte) byte {
	return (to - from) / 2
}
func colorDistanceVS(vHi, vLo, center, rng int) bool {
	return vHi+vLo-center > rng // true == vHi win
}
func coloravg(b []byte) int {
	return (int(b[0]) + int(b[1]) + int(b[2])) / 3
}
func bavg(b []byte) int {
	var sum int = 0
	var i int
	for i = range b {
		sum += int(b[i])
	}
	return sum / (i + 1)
}
