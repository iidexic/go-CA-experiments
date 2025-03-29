package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/iidexic/go-CA-experiments/core"

	_ "net/http/pprof"
)

// globals and Structs
// ==================================
var ( //16 by 9: 1920x1080, 960x540
	PixWidth    int  = 500
	PixHeight   int  = 1000
	GameWidth   int  = 300
	GameHeight  int  = 600
	tick, frame uint = 0, 0
	layoutCount int  = 0
)

// ?Profiling
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to 'file'")
var memprofile = flag.String("memprofile", "", "write memory profile to 'file'")

// ==================================
func main() {
	flag.Parse()
	//-CPU Profiling-
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() //error handling omitted for example?
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile:", err)
		}
		defer pprof.StopCPUProfile()
	} //-------------

	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("CA Experimentor")
	ebiten.SetWindowPosition(0, 80)
	g := core.GameSimInit(GameWidth, GameHeight)

	//>>>/ launch game loop /<<<//
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	//-Memory Profiling-
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() //error handling omitted for example?
		runtime.GC()    // get up-to-date statistics
		/* From pprof documentation:
		Lookup("allocs") creates a profile similar to `go test -memprofile`
		or use Lookup("heap") for profile that has inuse_space as default index.
		*/
		if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
			log.Fatal("could not write memory profile:", err)
		}

	}
	//------------------
}
