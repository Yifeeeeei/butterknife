package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExamplePrintln() {
	mainHello := "hello from main"
	butterknife.Println(mainHello)
	firstCallerPrintln()
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
