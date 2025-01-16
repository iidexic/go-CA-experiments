package input

/*SECTION - HOW INPUT HANDLING/KB ASSIGN/FUNCTION TRIGGER GOING WORK
var KEYMAPassignments []func() = []func(){} //this aint happening atm
- this is like first thing I can think of, actually havent thought of it yet
- whatever comes out comes out
- will need to write settings to a file at some point

STRUCTURE
-----------------

For now I am just going to hardcode functions in a switch case as they come up.
Later I can look at big picture to see if theres a better way
*/ //!SECTION

//Started setting up rebinds but that is so not necessary right now.
//check ebiten's keys.go OR .\util_external\kbdefaults.txt (most likely correct)

// CallKey is gonna do some wack shit
/* Can we make this stupid shit work somehow
func (g *Game) GetCallKey() *func() {
	return (g * Game) & func() {

	}
}
*/
