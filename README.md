# butterknife 🧈🗡️
A printing tool for Golang debugging. Prints out the function call hierarchy alongside the content. By hierarchy, it means

- the function call that wraps over butterknife's print
- the function call that wraps over the function call that wraps over butterknife's print
- ... all the way up

Spreads the whole thing onto the terminal smoothly.

It also prints out the line number of the code in the file, cmd + click to navigate to the line of code.

# usage

After initiating the project

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

import (
	"github.com/Yifeeeeei/butterknife"
)

func main() {
	mainHello := "hello from main"
	butterknife.PrintlnHierarchy(mainHello)
	// [main.go:9 main.main]
	//   hello from main
	firstCaller()
}

func firstCaller() {
	firstCallerHello := "hello from firstCaller"
	butterknife.PrintlnHierarchy(firstCallerHello)
	// [main.go:12 main.main]
	//   [main.go:17 main.firstCaller]
	//     hello from firstCaller
	secondCaller()
}

func secondCaller() {
	secondCallerHello := "hello from secondCaller"
	butterknife.PrintlnHierarchy(secondCallerHello)
	// [main.go:12 main.main]
	//   [main.go:21 main.firstCaller]
	//     [main.go:26 main.secondCaller]
	// 	     hello from secondCaller
}

```

Full outputs:

```bash
[main.go:9 main.main]
  hello from main
[main.go:12 main.main]
  [main.go:17 main.firstCaller]
    hello from firstCaller
[main.go:12 main.main]
  [main.go:21 main.firstCaller]
    [main.go:26 main.secondCaller]
      hello from secondCaller
```

# functions

- If you want to use the everyday print, choose functions start with "Print". If you want to use log's print, choose functions start with "Log".
- For normal print, formatted print or print stuffs with nice little spaces between them, use functions that have "" or "f" or "ln" in the middle.
- to print out all layers of the caller functions, choose the ones ends with "Hierarchy", if you only want to see one layer of caller function, use the other ones.

```go
func Print(args ...any)
func Printf(format string, args ...any)
func Println(args ...any)
func PrintHierarchy(args ...any)
func PrintfHierarchy(format string, args ...any)
func PrintlnHierarchy(args ...any)
func Log(args ...any)
func Logf(format string, args ...any)
func Logln(args ...any)
func LogHierarchy(args ...any)
func LogfHierarchy(format string, args ...any)
func LoglnHierarchy(args ...any)
```

