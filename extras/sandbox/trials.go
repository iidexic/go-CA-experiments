package sandbox

import (
	"fmt"
	v2 "math/rand/v2"

	"github.com/bytedance/gopkg/lang/fastrand"
)

var seedset [16]string = [16]string{"g4vJqSiKnYRi6378JFOV1rczgpMkZkCn", "f6I9rcA9JsEBx7pUm1UsZTYZMFCgQumO",
	"wjkSLsE5GSGtuYBpOoNx8yUQ24aEEiZw", "EqvtN8q0jEd9ezX9zacXaBA0DnnucFN6", "jvobxzh0fQ8DUxahKKpRAhiIYX14K8Kh",
	"05uhKGOTvEK1veIn5q74zOmTHUzSxWVQ", "BT0kq8BwC4CDlOIziflUCgyLb9qOj4MG", "SoE17aSitgTAlqIa5AAC4gpsnKFOK8WB",
	"QD53tmI9NYFNXo38rW51WezhxE3kyPHe", "eYAmffFaLVy8xdqfFkvMhZIRZEUhcgDZ", "VyeCwXR87YJgP2W4YETPBBsF8SHxlXW9",
	"HDga84PIr24iSNmQ5cmQxk5RqyNQZG8j", "SwAt592mmwk8v2YkbEwbwAdoQfAUtkZ9", "678IiEUOmzLPWfN6qDoLYabhQTTA2ABm"}

// ^ This file used to test different approaches.
// * using main to make things easier. cd into testing to run

func checkAverageKnown(b []byte) int {
	return (int(b[0]) + int(b[1]) + int(b[2])) / 3
}
func checkAverageUnknown(b []byte) int {
	var sum int = 0
	var i int
	for i = range b {
		sum += int(b[i])
	}
	return sum / (i + 1)
}

func byteseed(s32 string) [32]byte {
	var bs [32]byte
	for i, c := range s32 {
		bs[i] = byte(c)
	}
	return bs
}
func trialRandbytes(n int) ([]byte, error) {
	b := make([]byte, n)
	seed := byteseed(seedset[0])
	cc := v2.NewChaCha8(seed)
	_, e := cc.Read(b)
	return b, e
}

func trialFastrandbytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, e := fastrand.Read(b)
	return b, e
}

func trialSliceRaw(n int, b0, b1 byte) []byte {
	var slc []byte = []byte{}
	slc[0] = b0
	slc[1] = b1
	for i := 2; i < n+1; i++ {
		slc[i] = slc[i-1] << slc[i-2]
	}
	return slc
}
func trialSliceMake(n int, b0, b1 byte) []byte {
	slc := make([]byte, n+1)
	slc[0] = b0
	slc[1] = b1
	for i := 2; i < n+1; i++ {
		slc[i] = slc[i-1] << slc[i-2]
	}
	return slc
}
func trialSliceRawInverse(n int, b0, b1 byte) []byte {
	var slc []byte = []byte{}
	slc[n+1] = b0
	slc[n] = b1
	for i := n - 1; i >= 0; i-- {
		slc[i] = slc[i-1] << slc[i-2]
	}
	return slc
}
func trialSliceMakeInverse(n int, b0, b1 byte) []byte {
	slc := make([]byte, n+2)
	slc[n+1] = b0
	slc[n] = b1
	for i := n - 1; i >= 0; i-- {
		slc[i] = slc[i-1] << slc[i-2]
	}
	return slc
}

func trialShifting() {
	var b []byte = make([]byte, 36)
	for i := range b {
		b[i] = byte(i)
		//fmt.Printf("%d ",b[i])
	}
	fmt.Printf("\nlen = %d\n", len(b))
	for i := range b {
		fmt.Printf("(%d -> ", b[i])
		b[i] = (b[i]) << 3
		fmt.Printf("%d) ", b[i])
		if i%10 == 0 {
			fmt.Println("")
		}
	}
	fmt.Println(b[24])
	fmt.Printf("\nlen = %d\n", len(b))
}

//** ================================================

func WrapScreenConditional(wlen int, scr []byte) {

}
func WrapScreenHyperConcurrency(wlen int, scr []byte, c1, c2, c3 chan int) {

}
