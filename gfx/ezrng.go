package gfx

import "github.com/bytedance/gopkg/lang/fastrand"

// Fbytes returns func that reloads byt with rand values
func Fbytes(byt []byte) func() {
	return func() { fastrand.Read(byt) }
}
