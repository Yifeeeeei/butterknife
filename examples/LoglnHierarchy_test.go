package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExampleLoglnHierarchy() {
	mainHello := "hello from main"
	butterknife.LoglnHierarchy(mainHello)
	firstCallerLoglnHierarchy()
	// output:
}

func TestExampleLoglnHierarchy(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExampleLoglnHierarchy()
}

func firstCallerLoglnHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.LoglnHierarchy(firstCallerHello)
	secondCallerLoglnHierarchy()
}

func secondCallerLoglnHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.LoglnHierarchy(secondCallerHello)
}
