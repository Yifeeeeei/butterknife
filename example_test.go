package butterknife_test

import (
	"github.com/Yifeeeeei/butterknife"
)

func Example() {
	mainHello := "hello from main"
	butterknife.PrintlnHierarchy(mainHello)
	firstCaller()
}

func firstCaller() {
	firstCallerHello := "hello from firstCaller"
	butterknife.PrintlnHierarchy(firstCallerHello)
	secondCaller()
}

func secondCaller() {
	secondCallerHello := "hello from secondCaller"
	butterknife.PrintlnHierarchy(secondCallerHello)
}
