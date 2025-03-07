package entity

// ^ This file is very early on (currently entirely unused)
//** initially intended as event handler to store conditions and subsequent movements taken
import "image"

type movement interface {
	nextStep()
	define()
}

type movetype int

const (
	recttest movetype = iota
	zonestest
)

type simplmover struct {
	x, y int
}

type testmover struct {
	moveType movetype
	istep    int
	zones    []simplmover
	r        image.Rectangle
}

/*
func newTestMover() testmover {
	mvr := testmover{moveType: zonestest, istep: 0}
	s := simplmover{}
	_ = mvr
	_ = s
	return mvr
}
*/
