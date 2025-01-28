package gfx

// colorInter and functions intend to be system of working with RGB with no alpha.
// RGB can be imported from RGBA and exported back to RGBA
type colorRecoder interface {
	[]byte
	fromRGBA([]byte)
	toRGBA() []byte
}

// RGBi is a []byte that denotes it is a slice of RGB pixels, (only R,G,B values per pix)
type RGBi struct {
	c          []byte
	alpha      []byte
	storeAlpha bool
}

// FromRGBA populates RGBi with RGBA pixels, strips all Alpha values
func (ipx *RGBi) FromRGBA(pixels []byte) {
	if len(pixels)%4 > 0 {
		//idk but do somethin. trim? panic?
	}
	ipx.c = make([]byte, 3*(len(pixels)/4))
	inew := 0
	for i := 0; i < len(pixels); i += 4 {
		ipx.c[inew] = pixels[i]
		inew++
		ipx.c[inew] = pixels[i+1]
		inew++
		ipx.c[inew] = pixels[i+2]
		inew++
	}
}

// FromRGBAstore populates RGBi with RGBA pixels, while storing Alpha values
func (ipx *RGBi) FromRGBAstore(pixels []byte) {
	if len(pixels)%4 > 0 {
		//idk but do somethin. trim? panic?
	}
	ipx.c = make([]byte, 3*(len(pixels)/4))
	inew := 0
	for i := 0; i < len(pixels); i += 4 {
		ipx.c[inew] = pixels[i]
		inew++
		ipx.c[inew] = pixels[i+1]
		inew++
		ipx.c[inew] = pixels[i+2]
		inew++
		ipx.alpha[i/4] = pixels[i+3]
	}
	ipx.storeAlpha = true
}

// ToRGBA returns the RGBi image with alpha added back in
func (ipx *RGBi) ToRGBA() []byte {
	out := make([]byte, 4*len(ipx.c)/3)
	iold := 0
	for i := 0; i < len(out)-3; i += 4 {
		out[i] = ipx.c[iold]
		iold++
		out[i+1] = ipx.c[iold]
		iold++
		out[i+2] = ipx.c[iold]
		iold++
		out[i+3] = 255
	}
	return out
}
