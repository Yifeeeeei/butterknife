package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExamplePrintlnHierarchy() {
	mainHello := "hello from main"
	butterknife.PrintlnHierarchy(mainHello)
	firstCallerPrintlnHierarchy()
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
