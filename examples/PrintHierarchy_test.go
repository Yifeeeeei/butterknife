package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExamplePrintHierarchy() {
	mainHello := "hello from main"
	butterknife.PrintHierarchy(mainHello)
	firstCallerPrintHierarchy()
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
