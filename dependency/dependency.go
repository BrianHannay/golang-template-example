/*
 * This package represents an example dependency included with your golang project.
 */
package dependency

// We need to output text and synchronize with parent processes
import (
	"fmt"
)

// Since we're testing, all inputs should be interfaces to permit mocking
type WaitGroup interface {
	Done()
}

// Let's define a simple print function that notifies a WaitGroup (if provided) when complete.
// This will help demonstrate concurrency patterns provided by golang
func Print(output string, group WaitGroup) {
	fmt.Println(output)

	// If the WaitGroup was provided, inform it that this function has completed.
	if group != nil {
		group.Done()
	}
}
