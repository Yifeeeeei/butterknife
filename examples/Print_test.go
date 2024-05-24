package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExamplePrint() {
	mainHello := "hello from main"
	butterknife.Print(mainHello)
	firstCallerPrint()
}

func firstCallerPrint() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Print(firstCallerHello)
	secondCallerPrint()
}

func secondCallerPrint() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Print(secondCallerHello)
}
