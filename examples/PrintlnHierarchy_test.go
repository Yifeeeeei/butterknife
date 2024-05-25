package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExamplePrintlnHierarchy() {
	mainHello := "hello from main"
	butterknife.PrintlnHierarchy(mainHello)
	firstCallerPrintlnHierarchy()
	// output:
}

func TestExamplePrintlnHierarchy(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrintlnHierarchy()
}

func firstCallerPrintlnHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.PrintlnHierarchy(firstCallerHello)
	secondCallerPrintlnHierarchy()
}

func secondCallerPrintlnHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.PrintlnHierarchy(secondCallerHello)
}
