package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExampleLogHierarchy() {
	mainHello := "hello from main"
	butterknife.LogHierarchy(mainHello)
	firstCallerLogHierarchy()
	// output:
}

func TestExampleLogHierarchy(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExampleLogHierarchy()
}

func firstCallerLogHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.LogHierarchy(firstCallerHello)
	secondCallerLogHierarchy()
}

func secondCallerLogHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.LogHierarchy(secondCallerHello)
}
