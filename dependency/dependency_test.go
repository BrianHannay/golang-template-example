// Packages prefixed with _test are run by "go test".
package dependency_test

// We need to import the package being tested. In this case, we're in the package
// being tested, so import everything from "."
import "."

// We'll also need the testing package, for, you know... testing.
import (
	"testing"
	"sync"
)

// Here's an example of testing a package using many subtests:
func TestDependency(t *testing.T) {

	// Map out all the tests we intend to run
	tests := make(map[string]func(*testing.T))
	tests["Dependency.Print accepts nil"]		= testDependencyPrintAcceptsNil
	tests["Dependency.Print updates WaitGroup"]	= testDependencyPrintUpdatesWaitGroup
	for name, test := range tests {
		t.Run(name, test)
	}
}

// This will be used to determine wheather dependency.Print calls WaitGroup.Done exactly once
type WaitGroupSpy struct {
	sync.WaitGroup
	DoneCalls	int
}

// Implement the functions we'll be spying on
func (spy *WaitGroupSpy) Done() {
	spy.DoneCalls++
}

// Test that nil is an acceptable value to pass to dependency.Print
func testDependencyPrintAcceptsNil(t *testing.T) {
	dependency.Print("test message", nil)
}

// Test that Dependency.Print calls Done exactly once
func testDependencyPrintUpdatesWaitGroup(t *testing.T) {
	g := WaitGroupSpy{}
	dependency.Print("test message", &g)
	if g.DoneCalls != 1 {
		t.Errorf("Failed to update waitgroup - Done() called %d times", g.DoneCalls)
	}
}
