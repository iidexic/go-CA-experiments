package main

import (
	"fmt"
)

func main() {
	for f := range 5 {
		fmt.Printf("3 mod %d = %d\n", f+1, 3%(f+1))
	}
}

//
//*=========================================================================
//==========================================================================

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
/*
experimenting:
func main() {
	byte1 := byte(240)
	byte2 := byte(183)
	bpos := byte1 - byte2
	bneg := byte2 - byte1
	fmt.Printf("bigger-larger=%d(%08b)\nlarger-bigger = %d(%08b)\n", int(bpos), bpos, int(bneg), bneg)
	for f := range 5 {
		fmt.Printf("num %d\n", f)
	}
}*/
