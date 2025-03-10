package gfx

import (
	"fmt"
	"testing"
	"time"
)

func TestHW(t *testing.T) {
	println("hi hello")
	tic := time.NewTicker(8 * time.Millisecond)
	t0 := <-tic.C
	for wt := range 10 {
		wait := time.NewTimer(10 * time.Millisecond)
		<-wait.C
		tnow := <-tic.C
		if tnow.After(t0.Add(50 * time.Millisecond)) {
			fmt.Println("TIC IS OVER IT DIED :(:(")
		}
		fmt.Printf("#%d: tic is %s \n", wt, tnow)
	}
}
func castCoinToss(i int) int {
	tobyte := byte(i)
	tosign := int8(tobyte)
	return int(tosign)
}
func subtractCoinToss(i int) int {
	tobyte := byte(i)
	return int(tobyte) - 127
}

func benchmarkTossMethod(b *testing.B, fnTossInt func(int) int) {
	var dump int = 0
	for i := 0; i < b.N; i++ {
		dump += fnTossInt(i)
	}
	_ = (dump)
}

// ==== the tests ====
func BenchmarkTossSub(b *testing.B) { benchmarkTossMethod(b, subtractCoinToss) }

func BenchmarkTossCast(b *testing.B) { benchmarkTossMethod(b, castCoinToss) }

//===================

func TestTimeout(t *testing.T) {

	tmr := time.NewTimer(50 * time.Millisecond)
	go catchTimeout(tmr, t)
	for i := range 50 {
		time.Sleep(10 * time.Millisecond)
		if i%2 == 0 {
			tmr.Reset(50 * time.Millisecond)
		}
		if i%10 == 0 {
			fmt.Printf("%d", i)
		}
	}

}
func TestGen(t *testing.T) {

	tmr := time.NewTimer(50 * time.Millisecond)

	c := GetQuickRNG(64)
	go c.ROPcheck()
	outl := make([]byte, 10)
	go catchTimeout(tmr, t)
	for i := range 100 {
		tmr.Reset(20 * time.Millisecond)
		cc := <-c.C
		outl[i%10] = cc
		if i%10 == 9 {
			fmt.Println(outl)
		}
	}
}

func catchTimeout(timeout *time.Timer, t *testing.T) {
	<-timeout.C
	panic(t)
	//t.Fatal("timeout")
}
