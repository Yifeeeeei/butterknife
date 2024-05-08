# butterknife 🧈🗡️
A printing tool for Golang debugging. It prints out a variable and spread out its callers.

# functions

```go
func Print(varValue interface{}, varName ...string) 
```

Uses fmt to print formats and prints the value of the variable and its caller.

```go
func PrintConcurrent(varValue interface{}, varName ...string)
```

Uses fmt to print formats and prints the value of the variable and all its callers.

```go
func Log(varValue interface{}, varName ...string)
```

Uses log to print formats and prints the value of the variable and all its callers.

```go
func LogConcurrent(varValue interface{}, varName ...string)
```

# usage

Fetch the module

```bash
go get github.com/Yifeeeeei/butterknife
```

Import the package

```go
import "github.com/Yifeeeeei/butterknife"
```

Use the functions

```go
package main

import "test/butterknife"

func caller(num int) {
	butterknife.PrintConcurrent(num, "num")
}

func main() {
	a := 42
	butterknife.Print(a, "a")
	// [.../main.go:11 main.main] a = 42

	caller(a)
	// 	[.../main.go:14 main.main]
	//    [.../main.go:6 main.caller]
	//      num = 42
}

```

