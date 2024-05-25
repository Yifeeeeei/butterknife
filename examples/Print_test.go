package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExamplePrint() {
	mainHello := "hello from main"
	butterknife.Print(mainHello)
	firstCallerPrint()
	// output:
}

func TestExamplePrint(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrint()
}

func firstCallerPrint() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Print(firstCallerHello)
	secondCallerPrint()
}

func secondCallerPrint() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Print(secondCallerHello)
}
