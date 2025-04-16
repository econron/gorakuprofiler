package gorakuprofiler

import (
	"os"
	"runtime/pprof"
)

func StartCPU() (*os.File, error) {
	fCpu, err := os.Create("cpu.prof")
	if err != nil {
		return nil, err
	}
	if err = pprof.StartCPUProfile(fCpu); err != nil {
		defer fCpu.Close()
		return nil, err
	}
	return fCpu, nil
}

// stop all
func StopCPU(fCpu *os.File) {
	defer fCpu.Close()
	defer pprof.StopCPUProfile()
}

// show 
func ShowCPU() {
	// TODO
}

func StartMemory() (*os.File, error) {
	fMem, err := os.Create("mem.prof")
	if err != nil {
		return nil, err
	}
	if err = pprof.WriteHeapProfile(fMem); err != nil {
		defer fMem.Close()
		return nil, err
	}
	return fMem, nil
}

func StopMemory(fMem *os.File) {
	defer fMem.Close()
}

func ShowMemory() {
	// TODO
}