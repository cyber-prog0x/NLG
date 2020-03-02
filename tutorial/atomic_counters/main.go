package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// An Atomic Counter is a GLSL variable type who's storage comes from a Buffer Object.
// Atomic counters, as the name suggests, can have atomic memory operations performed on them.
// They can be thought of as a very limited form of buffer image variable.
// Ref: https://www.khronos.org/opengl/wiki/Atomic_Counter
func main() {
	var ops uint64

	var wg sync.WaitGroup

	for i:= 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c:= 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)

			}
			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("ops:", ops)

}
