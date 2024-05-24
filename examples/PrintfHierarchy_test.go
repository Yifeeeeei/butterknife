package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExamplePrintfHierarchy() {
	mainHello := "hello from main"
	butterknife.PrintfHierarchy("%s", mainHello)
	firstCallerPrintfHierarchy()
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
