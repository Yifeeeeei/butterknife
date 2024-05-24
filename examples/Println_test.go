package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExamplePrintln() {
	mainHello := "hello from main"
	butterknife.Println(mainHello)
	firstCallerPrintln()
}

func TestExamplePrintln(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExamplePrintln()
}

func firstCallerPrintln() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Println(firstCallerHello)
	secondCallerPrintln()
}

func secondCallerPrintln() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Println(secondCallerHello)
}
