# butterknife 🧈🗡️
A printing tool for Golang debugging. It works like "fmt" or "log", but it prints out something else alongside, including:

- the line number and file that calls it
- the function that wraps over it
- the function that wraps over the function that wraps over it
- ...

Spreads the whole thing onto the terminal smoothly.

Note: cmd + click on the line number to jump right to it

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
	a := "a"
	b := "b"
	butterknife.Print(a, b)
	// [main.go:10 main.main]  ab

	butterknife.Println(a, b)
	// [main.go:13 main.main]  a b

	butterknife.Printf("a:%s ,b:%s\n", a, b)
	// [main.go:16 main.main]  a:a ,b:b

	caller()
}

func caller() {
	c := "c"
	d := "d"
	butterknife.PrintHierarchy(c, d)
	// [main.go:19 main.main]
	//   [main.go:25 main.caller]
	//     cd

	butterknife.PrintlnHierarchy(c, d)
	// [main.go:19 main.main]
	//   [main.go:30 main.caller]
	//     c d

	butterknife.PrintfHierarchy("c:%s ,d:%s\n", c, d)
	// [main.go:19 main.main]
	//	[main.go:35 main.caller]
	//	  c:c ,d:d
}

```

# functions

"Print" / "Log": use fmt/log

+

"" / "f" / "ln": normal output / formatted output / line output

+

"" / "Hierarchy": just print one caller / and all callers' callers

```go
func Print(args ...interface{})
func Printf(format string, args ...interface{}) 
func Println(args ...interface{})
func PrintHierarchy(args ...interface{})
func PrintfHierarchy(format string, args ...interface{})
func PrintlnHierarchy(args ...interface{})
func Log(args ...interface{})
func Logf(format string, args ...interface{})
func Logln(args ...interface{}) 
func LogHierarchy(args ...interface{})
func LogfHierarchy(format string, args ...interface{})
func LoglnHierarchy(args ...interface{})
```

