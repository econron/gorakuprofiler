package gorakuprofiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
 
func ShowCPU(port string) error {
	if err := startPprofServer("cpu.prof", "8081"); err != nil {
		return err
	}
	return nil
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

func ShowMemory() error {
	if err := startPprofServer("mem.prof", "8082"); err != nil {
		return err
	}
	return nil
}

func startPprofServer(path string, port string) error {
	listenPort := ":" + port
	cmd := exec.Command("go", "tool", "pprof", fmt.Sprintf("-http=%s", listenPort), path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	log.Printf("Started HTTP server on background")
	log.Printf("If you want to stop server, please stop in your terminal.")

	return nil
}