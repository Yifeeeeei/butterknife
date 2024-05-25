package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExamplePrintf() {
	mainHello := "hello from main"
	butterknife.Printf("%s", mainHello)
	firstCallerPrintf()
	// output:
}

func TestExamplePrintf(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrintf()
}

func firstCallerPrintf() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Printf("%s", firstCallerHello)
	secondCallerPrintf()
}

func secondCallerPrintf() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Printf("%s", secondCallerHello)
}
