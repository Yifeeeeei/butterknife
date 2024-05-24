package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExampleLog() {
	mainHello := "hello from main"
	butterknife.Log(mainHello)
	firstCallerLog()
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
