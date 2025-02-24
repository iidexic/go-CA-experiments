package gfx

type waveshape int

const (
	tri waveshape = iota
	sin
	rampdown
	rampup
	square
)

/*
	Wave represents a waveform

Segments: minimum 4 [c/a][a\c][c\0][0/c]
over 4 segments
*/
type idwave struct {
	segments     int
	tphase, tamp []int
}

// Wave will be used to drive smooth controlled parameter changes
type Wave struct {
	amplitude                       byte
	period, segments, samples, sseg int
	shape                           waveshape
	phase                           []int
	phaseVal                        []int
}

func idWave(shape waveshape) {
	w := idwave{segments: 4}
	switch shape {
	case tri:
		w.tphase = []int{0, 90, 180, 270}
		w.tamp = []int{127, 255, 128, 0}
	}
}
