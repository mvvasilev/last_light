package engine

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func Profile(profileName string, what func()) {
	// Create a CPU profile file
	f, err := os.Create(fmt.Sprintf("%s.prof", profileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	what()
}
