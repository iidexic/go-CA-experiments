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
func main() {
	fmt.Println("we are trialing")
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

func trialFastRandbytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, e := fastrand.Read(b)
	return b, e
}
