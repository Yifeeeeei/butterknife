package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExampleLogHierarchy() {
	mainHello := "hello from main"
	butterknife.LogHierarchy(mainHello)
	firstCallerLogHierarchy()
}

func firstCallerLogHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.LogHierarchy(firstCallerHello)
	secondCallerLogHierarchy()
}

func secondCallerLogHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.LogHierarchy(secondCallerHello)
}
