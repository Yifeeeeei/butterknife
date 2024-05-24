package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExamplePrintfHierarchy() {
	mainHello := "hello from main"
	butterknife.PrintfHierarchy("%s", mainHello)
	firstCallerPrintfHierarchy()
}

func TestExamplePrintfHierarchy(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrintfHierarchy()
}

func firstCallerPrintfHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.PrintfHierarchy("%s", firstCallerHello)
	secondCallerPrintfHierarchy()
}

func secondCallerPrintfHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.PrintfHierarchy("%s", secondCallerHello)
}
