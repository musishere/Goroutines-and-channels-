// Simple Goroutine Print
//
// Launch 5 goroutines, each printing its ID (0–4).
//
// Ensure main waits until all prints finish.

// Key: sync.WaitGroup.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	go func() {
		for i := 0; i <= 5; i++ {
			fmt.Println(i)
		}
	}()
}
