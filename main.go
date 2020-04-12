/**
 * Template golang project
 *
 * This sample will show how WaitGroups can be used to ensure child threads have completed.
 * We will also see how channels can be used to communicate between threads.
 */
package main

// We are going to require some builtin features - namely, text formatting, and thread synchronization
import (
	"fmt"
	"sync"
)

// Demonstrating a template dependency requires importing it first
import "./dependency"


// the main function is the entrypoint to the compiled go program
func main() {

	// Print 5 times using goroutines and a wait group:
	fmt.Println("Threading 5 Print calls in goroutines")
	goroutineExample()

	// Our sample dependency defines a print method. Let's defer printing 5 hello worlds:
	fmt.Println("Deferring 5 Print calls")
	deferredExample()
}

// This is an example of using WaitGroups to synchronize goroutine calls with their
// parent thread, ensuring completion before main thread sends an exit signal
func goroutineExample() {
	print_group := &sync.WaitGroup{}

	fmt.Println("Creating 5 goroutine threads (numbered 1-5), then waiting for them to complete in nondeterministic order:")
	for i := 1; i <= 5; i++ {

		//Register a new entry in the WaitGroup for each goroutine we run below.
		print_group.Add(1)

		// Let's use goroutines to insert these numbers in nondeterministic positions
		go dependency.Print(
			fmt.Sprintf("Hello World (Goroutine plus waitgroup %d)!", i),
			print_group,
		)
	}
	print_group.Wait()
	fmt.Println("All goroutine threads have completed")
}

// This is an example of using the defer keyword to stack up function calls
func deferredExample() {

	// Produce 5 deferred Print calls
	for i:= 5; i <= 10; i++ {
		defer dependency.Print(
			fmt.Sprintf(
				"Hello World (Deferred Thread %d)!",
				i,
			),
			// Since these are deferred, no need to pass a WaitGroup and synchronize:
			nil,
		)
	}
	fmt.Println("Deferred functions (numbered 5-10) should now (upon parent function completion) execute in reverse order:")
}
