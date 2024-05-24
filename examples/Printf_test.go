package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExamplePrintf() {
	mainHello := "hello from main"
	butterknife.Printf("%s", mainHello)
	firstCallerPrintf()
}

func firstCallerPrintf() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Printf("%s", firstCallerHello)
	secondCallerPrintf()
}

func secondCallerPrintf() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Printf("%s", secondCallerHello)
}
