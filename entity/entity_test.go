package entity

import (
	"fmt"
	"testing"
)

// TestSidewrap does that
func TestSidewrap(t *testing.T) {
	/*
	   [0,1,2,3
	    4,5,6,7
	    8,9,A,B
	    C,D,E,F]
	*/
	// out1: 10x10, i = 10, move = -1, width = 10 | out1->19
	out1 := sidewrap(10, -1, 10)
	// out2: 768x512, i=158976 (0 on row 207), move = -780, width = 768,
	// out2->159732
	out2 := sidewrap(158976, -780, 768)
	// out3:128x128,, i=255(last on row 2), move=1, width=128 | out3->128
	out3 := sidewrap(255, 1, 128)
	//out4: 128x test 0 <- left | out4=127
	out4 := sidewrap(0, -1, 128)
	//out5: 128x test max +513 | out5= 16256? wait on this

	if out1 != 19 {
		fmt.Printf("out1 = %d, should be 19\n", out1)
		t.Fail()
	}
	if out2 != 159732 {
		fmt.Printf("out1 = %d, should be 159732\n", out2)
		t.Fail()
	}
	if out3 != 128 {
		fmt.Printf("out3 = %d, should be 128\n", out3)
		t.Fail()
	}
	if out4 != 127 {
		fmt.Printf("out4 = %d, should be 127\n", out4)
		t.Fail()
	}
}
