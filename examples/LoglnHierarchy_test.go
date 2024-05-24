package butterknife_test

import "github.com/Yifeeeeei/butterknife"

func ExampleLoglnHierarchy() {
	mainHello := "hello from main"
	butterknife.LoglnHierarchy(mainHello)
	firstCallerLoglnHierarchy()
}

func firstCallerLoglnHierarchy() {
	firstCallerHello := "hello from firstCaller"
	butterknife.LoglnHierarchy(firstCallerHello)
	secondCallerLoglnHierarchy()
}

func secondCallerLoglnHierarchy() {
	secondCallerHello := "hello from secondCaller"
	butterknife.LoglnHierarchy(secondCallerHello)
}
