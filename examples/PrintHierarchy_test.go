package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExamplePrintHierarchy() {
	mainHello := "hello from main"
	butterknife.PrintHierarchy(mainHello)
	firstCallerPrintHierarchy()
	// output:
}

func TestExamplePrintHierarchy(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrintHierarchy()
}

func firstCallerPrintHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.PrintHierarchy(firstCallerHello)
	secondCallerPrintHierarchy()
}

func secondCallerPrintHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.PrintHierarchy(secondCallerHello)
}
