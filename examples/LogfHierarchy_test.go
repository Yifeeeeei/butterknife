package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExampleLogfHierarchy() {
	mainHello := "hello from main"
	butterknife.LogfHierarchy("%s", mainHello)
	firstCallerLogfHierarchy()
}

func firstCallerLogfHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.LogfHierarchy("%s", firstCallerHello)
	secondCallerLogfHierarchy()
}

func secondCallerLogfHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.LogfHierarchy("%s", secondCallerHello)
}
