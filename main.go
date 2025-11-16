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
	PixWidth    int  = 1600
	PixHeight   int  = 900
	GameWidth   int  = 800
	GameHeight  int  = 450
	tick, frame uint = 0, 0
	layoutCount int  = 0
)

// Profiling
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to 'file'")
var memprofile = flag.String("memprofile", "", "write memory profile to 'file'")

func windowSetup() { //TODO: Move into core?
	ebiten.SetWindowSize(PixWidth, PixHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("CA Experimentor")
	ebiten.SetWindowPosition(0, 80)

}
func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func setupProfiling() *os.File {
	flag.Parse()
	//-CPU Profiling-
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		checkFatal(err, "could not create CPU profile: ")
		return f
	} //-------------
	return nil
}

// =================================
func main() {
	//-CPU Profiling---
	if f := setupProfiling(); f != nil {
		defer f.Close()
		checkFatal(pprof.StartCPUProfile(f), "could not start CPU profile:")
		defer pprof.StopCPUProfile()
	} //---------------

	windowSetup()

	g := core.GameSimInit(GameWidth, GameHeight)
	// ╭────────────────────────── launch game loop ────────────────────────╮

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	// ╰────────────────────────────────────────────────────────────────────╯
	//-Memory Profiling-

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		checkFatal(err, "could not create memory profile: ")
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
