package main

import (
	"flag"
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/core"
	"github.com/pedritoelcabra/projectx/gfx"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "profile/cpu.prof", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "profile/mem.prof", "write memory profile to `file`")
var doProfile = false

func main() {
	flag.Parse()
	if doProfile && *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	handleError(ebiten.Run(update, gfx.ScreenWidth, gfx.ScreenHeight, 1, "ProjectX"))

	if doProfile && *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func update(screen *ebiten.Image) error {
	return core.G().Update(screen)
}

func handleError(err ...interface{}) {
	if err[0] == nil {
		return
	}
	log.Fatal(err...)
}
