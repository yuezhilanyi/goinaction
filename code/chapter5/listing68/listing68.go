// Sampler program to show how the program can access an
// unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/yuezhilanyi/goinaction/code/chapter5/listing68/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type and initialize
	// the value to 0.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}