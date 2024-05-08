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
