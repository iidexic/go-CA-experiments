package entity

// wrap val to remain within limit
func wrap(val, limit int) int {
	return ((val % limit) + limit) % limit
}

func sidewrap(index, move, width int) int {
	//[Treat as if wrapping around 1 single row, also this code could use cleanup]
	irownum := index / width
	iWrap := index % width
	wrapMoved := ((iWrap+(move%width))%width + width) % width
	return wrapMoved + (irownum * width)

}

func bavg(r, g, b byte) byte {
	return byte((int(r) + int(g) + int(b)) / 3)
}

// absolute center difference byte average.
func acdbavg(b ...byte) byte {
	tot := 0
	for _, v := range b {
		tot += int(v) - 127
	}
	a := tot / len(b)
	if a < 0 {
		a = -a
	}
	return byte(a)
}

// byte slice limit add (probably should combine with bslsub)
func bsladd(b []byte, add byte) {
	for i := range b {
		if add > 255-b[i] {
			b[i] = 255
		} else {
			b[i] += add
		}
	}
}

func bslsub(b []byte, subtract byte) {
	for i := range b {
		if b[i] < subtract {
			b[i] = 0
		} else {
			b[i] -= subtract
		}
	}
}
