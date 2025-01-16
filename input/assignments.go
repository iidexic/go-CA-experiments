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

// CallKey will return key's bound func()
//or maybe just make it happen
//?Also: Is it better to do individual keys for concurrency?
func CallKey(ikey int) func() {
	switch ikey {
	case 0 /*A*/ :
		return func() {}
	case 4 /*E*/ :
		return func() {}
	case 16 /*Q*/ :
		return func() {}
	case 17 /*R*/ :
		return func() {}
	case 54 /*enter*/ :
		return func() {}
	case 116 /*space*/ :
		return func() {}
	default:
		return func() {}
	}

}
