package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExampleLogf() {
	mainHello := "hello from main"
	butterknife.Logf("%s", mainHello)
	firstCallerLogf()
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
