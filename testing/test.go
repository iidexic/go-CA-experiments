package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	u := rand.New(rand.Source())
	fmt.Printf("rand u64: %d\n", u)
}

//
//*=========================================================================
//==========================================================================
//! I decided this really is not worth it to be doing
// Was testing manual byte extraction from a uint64
/*

	var bytUse uint
	for bu := 512; bu != 0; bu -= 8 {
		if u>>bu > 0 {
			bytUse = uint(bu)
			break
		}
	}

	fmt.Printf("using %d bytes\n", bytUse)

	var s uint
	if bytUse > 64 {
		bytUse = 64
	}
	for i := bytUse;i!=0;i-=8 {
		s = bytUse - i
		fmt.Printf("extract byte %d: %d\n", s, u>>s)
		if i > 66 {
			break
		}
	}
*/
