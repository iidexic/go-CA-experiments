package core

// cycler counts cycles to be sent to anywhere game timing is needed
type cycler struct {
	ticks, frames int
}
