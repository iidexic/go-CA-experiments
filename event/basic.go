package event

/*
don't need to get fancy with this.
*/

type Momentary interface {
	Activated() bool
}
type Toggleable interface {
	Activated() bool
	isActive() bool
	Deactivated() bool
}

type Event interface {
	Exec()
}
