package butterknife

import (
	"fmt"
	"log"
	"runtime"
)

type writer interface {
	Writef(format string, args ...interface{})
}

type logger struct{}

func (l logger) Writef(format string, args ...interface{}) {
	log.Printf(format, args...)
}

type printer struct{}

func (p printer) Writef(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func write(varValue interface{}, wr writer, varName ...string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Failed to get caller info")
		return
	}
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	vn := "Var"
	if len(varName) > 0 {
		vn = varName[0]
	}
	wr.Writef("[%s:%d %s] %s = %v\n", file, line, funcName, vn, varValue)
}

func writeConcurrent(varValue interface{}, wr writer, varName ...string) {
	layers := []string{}
	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		pc, _, _, _ := runtime.Caller(i)
		funcName := runtime.FuncForPC(pc).Name()
		layers = append(layers, fmt.Sprintf("[%s:%d %s]", file, line, funcName))
	}
	spaces := ""
	for i := len(layers) - 3; i >= 2; i-- {
		wr.Writef("%s%s\n", spaces, layers[i])
		spaces += "  "
	}
	vn := "Var"
	if len(varName) > 0 {
		vn = varName[0]
	}
	wr.Writef("%s%s = %v\n", spaces, vn, varValue)
}

// Uses fmt to print formats and prints the value of the variable and its caller.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - name: the name of the variable (optional)
//
// It does not return any value.
func Print(varValue interface{}, varName ...string) {
	write(varValue, printer{}, varName...)
}

// Uses fmt to print formats and prints the value of the variable and all its callers.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - name: the name of the variable (optional)
//
// It does not return any value.
func PrintConcurrent(varValue interface{}, varName ...string) {
	writeConcurrent(varValue, printer{}, varName...)
}

// Uses log to print formats and prints the value of the variable and its caller.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - name: the name of the variable (optional)
//
// It does not return any value.
func Log(varValue interface{}, varName ...string) {
	write(varValue, logger{}, varName...)
}

// Uses log to print formats and prints the value of the variable and all its callers.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - name: the name of the variable (optional)
//
// It does not return any value.
func LogConcurrent(varValue interface{}, varName ...string) {
	writeConcurrent(varValue, logger{}, varName...)
}
