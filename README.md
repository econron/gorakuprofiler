# gorakuprofiler

I want a SIMPLE USE tool.

I don't want to think or remenber "HOW TO USE".

I JUST WANT TO PROFILE EASILY.

# HOW TO USE

```go
package main

import (
	"log"
	"fmt"
	"github.com/econron/gorakuprofiler"
)

func sum() int32 {
	var val int32 = 20
	return val + 200 * 300 / 100
}

func main() {
  // This is a example use. JUST COPY AND PASTE THIS.
	fcpu, err := gorakuprofiler.StartCPU()
	if err != nil {
		log.Fatal("failed to init goraku cpu profiler. ", err)
	}
	fmem, err := gorakuprofiler.StartMemory()
	if err != nil {
		log.Fatal("failed to init goraku memory profiler. ", err)
	}
	gorakuprofiler.StopCPU(fcpu)
	gorakuprofiler.StopMemory(fmem)
  // This is end.

	ret := sum()
	fmt.Println(ret)
}
```