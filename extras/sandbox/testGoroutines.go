package sandbox

import (
	"fmt"
	"time"
)

func sumcount(nubmer, to int) {
	tnow := time.Now()
	var count int = 0

	for i := 1; i <= to; i++ {
		sum := count + i
		count = i
		if i == to-1 {
			f := float64(sum*sum) / float64(count*count)
			fmt.Printf("run %d to %d done: calc:%.f\ntook %dms\n", nubmer, to, f,
				time.Since(tnow).Milliseconds())
		}
	}
}
func sumcountTimer(nubmer, to int) time.Duration {
	tnow := time.Now()

	var count int = 0

	for i := 1; i <= to; i++ {
		sum := count + i
		count = i
		if i == to-1 {
			f := float64(sum) / float64(count*count)
			fmt.Printf("run %d to %d^5 done: calc:%.f\n", nubmer, to, f)

		}
	}
	return time.Since(tnow)
}
func goroutiner() {

	now := time.Now()
	runs := []int{44, 41, 38, 35, 32, 29, 26, 23}
	fmt.Println("=======================================")
	fmt.Printf("Starting %d runs\n", len(runs))
	fmt.Println("---------------------------------------")
	/* straight line
	for i, v := range runs {
		v5 := v * v * v * v * v
		t := sumcountTimer(i+1, v5)
		fmt.Printf("sumcount of %d^5 took %fs\n", v, t.Seconds())
	}
	*/
	/* check actual time of time.Duration
	for i := range 10 {
		t := time.Duration(1000000 * (i + 1))
		fmt.Printf("time 1000x%d = %f sec\n", (i + 1), t.Seconds())
	}
	*/
	for i, v := range runs {
		v5 := v * v * v * v * v
		go sumcount(i, v5)
		fmt.Printf("sent run %d\n", i)
		time.Sleep(2500000)
	}
	fmt.Printf("all sent took %dns\n", time.Since(now).Nanoseconds())
	time.Sleep(100000000)
	fmt.Println("---------------------------------------")
	fmt.Printf("total time %fs\n", time.Since(now).Seconds())
	fmt.Println("=======================================")
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
