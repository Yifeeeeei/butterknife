package butterknife_test

import (
	"testing"

	"github.com/Yifeeeeei/butterknife"
)

func ExampleLogf() {
	mainHello := "hello from main"
	butterknife.Logf("%s", mainHello)
	firstCallerLogf()
}

func TestExampleLogf(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExampleLogf()
}

func firstCallerLogf() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Logf("%s", firstCallerHello)
	secondCallerLogf()
}

func secondCallerLogf() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Logf("%s", secondCallerHello)
}
