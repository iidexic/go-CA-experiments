package gfx

import "github.com/bytedance/gopkg/lang/fastrand"

type Frng struct {
	bb []byte
	n  int
}



func Fbytes(byt []byte) func() {
	return func(){fastrand.Read(byt)}
}

