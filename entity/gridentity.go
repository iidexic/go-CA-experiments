package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// GridEntity intended basis of cellular automata grid
type GridEntity struct {
	Img                   *ebiten.Image
	X, Y                  uint
	modAdd, modMult, Area int
	Pixels, PixLum        []byte
	Op                    ebiten.DrawImageOptions
	Set, Draw             bool
}

// MakeGridDefault generates base CA grid
func MakeGridDefault(gWidth, gHeight int) *GridEntity { // ? no ptr ok on GridEntity????
	width := (3 * gWidth) / 4
	height := (3 * gHeight) / 4
	grid := GridEntity{
		Img: ebiten.NewImage(width, height),
		Op:  ebiten.DrawImageOptions{},
		X:   uint(width), Y: uint(height), Area: width * height,
		Set:     true,
		Pixels:  make([]byte, width*height*4),
		modAdd:  0,
		modMult: 1,
	}
	grid.Op.GeoM.Translate(float64((gWidth-width)/2), float64((gHeight-height)/2))
	grid.Img.Fill(color.RGBA{R: 155, G: 155, B: 165, A: 255})
	return &grid
}

// SetMod sets modulation values modAdd and modMult
func (grid *GridEntity) SetMod(modAdd, modMult int) {
	grid.modAdd = modAdd
	grid.modMult = modMult
}

// SimstepLVSD performs one cycle/screen of checks and updates
// for the center-distance intensity comparison sim ("Light VS Dark")
func (grid *GridEntity) SimstepLVSD() {
	for i := 0; i < len(grid.Pixels); i += 4 {
		//? any benefit to using i++ in range area *4 for index?
		//toIndex := i * 2 % int(len(grid.Pixels)) //original index modulate
		toIndex := functionMod(i, grid.modAdd, grid.modMult, len(grid.Pixels))
		grid.pxGoToward(i, grid.Pixels[toIndex:toIndex+3])
	}
}

func (grid *GridEntity) pxGoToward(indexR int, toPx []byte) {
	for i := range 3 {
		if grid.Pixels[indexR+i] > toPx[i] {
			grid.Pixels[indexR+i] -= (grid.Pixels[indexR+i] - toPx[i]) / 2
		} else {
			grid.Pixels[indexR+i] += (toPx[i] - grid.Pixels[indexR+i]) / 2
		}
	}
}
func (grid *GridEntity) pxReplace(indexR int, new []byte) {
	grid.Pixels[indexR] = new[0]
	grid.Pixels[indexR+1] = new[1]
	grid.Pixels[indexR+2] = new[2]
}

// pxTransplant overwrites 1px (3 indices) in grid.Pixels, starting at index
// write uses values at indices R,G,B without changes made during the function call
func (grid *GridEntity) pxTransplant(index int, R, G, B int) {
	var nR byte = grid.Pixels[R]
	var nG byte = grid.Pixels[G]
	var nB byte = grid.Pixels[B]
	grid.Pixels[index] = nR
	grid.Pixels[index+1] = nG
	grid.Pixels[index+2] = nB
}

func functionMod(start, add, mult int, limit int) int {
	return wrap(((start + add) * mult), limit)
}

//TODO fix this error:
/*
github.com/iidexic/go-CA-experiments/entity.(*GridEntity).SimstepLVSD(0xc00039c080?)
        D:/Coding/github/go-CA-experiments/entity/gridentity.go:50 +0xf0
main.(*Game).Update(0xc00078a600)
        D:/Coding/github/go-CA-experiments/main.go:106 +0xe7
github.com/hajimehoshi/ebiten/v2.(*gameForUI).Update(0xc0000f6000)
        C:/Users/derek/go/pkg/mod/github.com/hajimehoshi/ebiten/v2@v2.8.6/gameforui.go:112 +0x23
github.com/hajimehoshi/ebiten/v2/internal/ui.(*context).updateFrameImpl(0xc00057ad00, {0xeb3790, 0xc0004500b0}, 0x1, 0x4094000000000000, 0x4086800000000000, 0x3ff8000000000000, 0xc00010a308, 0x0)
        C:/Users/derek/go/pkg/mod/github.com/hajimehoshi/ebiten/v2@v2.8.6/internal/ui/context.go:154 +0x2f0
github.com/hajimehoshi/ebiten/v2/internal/ui.(*context).updateFrame(0xc00057ad00, {0xeb3790, 0xc0004500b0}, 0x4094000000000000, 0x4086800000000000, 0x3ff8000000000000, 0xc00010a308)
        C:/Users/derek/go/pkg/mod/github.com/hajimehoshi/ebiten/v2@v2.8.6/internal/ui/context.go:73 +0x85
github.com/hajimehoshi/ebiten/v2/internal/ui.(*UserInterface).updateGame(0xc00010a308)
        C:/Users/derek/go/pkg/mod/github.com/hajimehoshi/ebiten/v2@v2.8.6/internal/ui/ui_glfw.go:1491 +0x138
github.com/hajimehoshi/ebiten/v2/internal/ui.(*UserInterface).loopGame(0xc00010a308)
        C:/Users/derek/go/pkg/mod/github.com/hajimehoshi/ebiten/v2@v2.8.6/internal/ui/ui_glfw.go:1448 +0x8b
github.com/hajimehoshi/ebiten/v2/internal/ui.(*UserInterface).runMultiThread.func2()
        C:/Users/derek/go/pkg/mod/github.com/hajimehoshi/ebiten/v2@v2.8.6/internal/ui/run.go:71 +0x12c
golang.org/x/sync/errgroup.(*Group).Go.func1()
        C:/Users/derek/go/pkg/mod/golang.org/x/sync@v0.10.0/errgroup/errgroup.go:78 +0x50
created by golang.org/x/sync/errgroup.(*Group).Go in goroutine 1
        C:/Users/derek/go/pkg/mod/golang.org/x/sync@v0.10.0/errgroup/errgroup.go:75 +0x96
exit status 2
*/
