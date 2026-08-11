package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*SECTION - HOW INPUT HANDLING/KB ASSIGN/FUNCTION TRIGGER GOING WORK

ok so two kinds of input check available:

1. Check Individual Keys to see if pressed, press duration, etc.

2. Get All Keys just pressed, All keys just released, do some shit.
*/

type KeyBind interface {
	Exec()
	Triggered() bool
	BoundKey() ebiten.Key
}

type KeyBindHeld interface {
	KeyBind
	Reset()
	state() bool
}
type KeyBindManager interface {
	Update()
	Add(KeyBind) bool
	IsBound(ebiten.Key) bool
	Unbind(ebiten.Key)
	Overwrite(KeyBind)
}

type KeyMap struct {
	binds map[ebiten.Key]KeyBind
}

func (k *KeyMap) Update() {
	for _, b := range k.binds {
		if b.Triggered() {
			if lb, ok := b.(Latching); ok && lb.state() {
				lb.Reset()
			} else {
				b.Exec()
			}

		}
	}
}

func (k *KeyMap) IsBound(key ebiten.Key) bool {
	_, ok := k.binds[key]
	return ok
}

func (k *KeyMap) Add(bind KeyBind) bool {
	if k.IsBound(bind.BoundKey()) {
		return false
	}
	k.binds[bind.BoundKey()] = bind
	return true
}
func (k *KeyMap) Overwrite(bind KeyBind) {
	k.binds[bind.BoundKey()] = bind
}

func (k *KeyMap) Unbind(key ebiten.Key) { delete(k.binds, key) }

//TODO: Check if non-ptr func works

type bindchanb struct {
	ebiten.Key
	c chan bool
}
type bindptrb struct {
	ebiten.Key
	pressed *bool
}

type bindfn struct {
	ebiten.Key
	exec *func()
}

func (b *bindptrb) Exec()                { *b.pressed = true }
func (b *bindptrb) Reset()               { *b.pressed = false }
func (b *bindptrb) state() bool          { return *b.pressed }
func (b *bindptrb) BoundKey() ebiten.Key { return b.Key }
func (b *bindptrb) Triggered() bool      { return inpututil.IsKeyJustPressed(b.Key) }

func (b *bindchanb) Exec()                { b.c <- true }
func (b *bindchanb) BoundKey() ebiten.Key { return b.Key }
func (b *bindchanb) Triggered() bool      { return inpututil.IsKeyJustPressed(b.Key) }

func (b *bindfn) Exec()                { (*b.exec)() }
func (b *bindfn) BoundKey() ebiten.Key { return b.Key }
func (b *bindfn) Triggered() bool      { return inpututil.IsKeyJustPressed(b.Key) }
