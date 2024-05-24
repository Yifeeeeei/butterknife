package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExampleLog() {
	mainHello := "hello from main"
	butterknife.Log(mainHello)
	firstCallerLog()
}

func TestExampleLog(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExampleLog()
}

func firstCallerLog() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Log(firstCallerHello)
	secondCallerLog()
}

func secondCallerLog() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Log(secondCallerHello)
}
