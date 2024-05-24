package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExampleLogln() {
	mainHello := "hello from main"
	butterknife.Logln(mainHello)
	firstCallerLogln()
}

func firstCallerLogln() {
	firstCallerHello := "hello from firstCaller"
	butterknife.Logln(firstCallerHello)
	secondCallerLogln()
}

func secondCallerLogln() {
	secondCallerHello := "hello from secondCaller"
	butterknife.Logln(secondCallerHello)
}
