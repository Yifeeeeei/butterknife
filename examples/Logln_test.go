package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExampleLogln() {
	mainHello := "hello from main"
	butterknife.Logln(mainHello)
	firstCallerLogln()
	// output:
}

func TestExampleLogln(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExampleLogln()
}

func firstCallerLogln() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Logln(firstCallerHello)
	secondCallerLogln()
}

func secondCallerLogln() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Logln(secondCallerHello)
}
