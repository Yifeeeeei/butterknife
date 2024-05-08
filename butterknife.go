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

func write(varValue interface{}, wr writer, comments ...string) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller info")
		return
	}
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	if len(comments) > 0 {
		vn := ""
		for _, v := range comments {
			vn += v + " "
		}
		wr.Writef("[%s:%d %s] %s: %v\n", file, line, funcName, vn, varValue)
	} else {
		wr.Writef("[%s:%d %s] %v\n", file, line, funcName, varValue)
	}
}

func writeConcurrent(varValue interface{}, wr writer, comments ...string) {
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
	if len(comments) > 0 {
		vn := ""
		for _, v := range comments {
			vn += v + " "
		}
		wr.Writef("%s%s: %v\n", spaces, vn, varValue)
	} else {
		wr.Writef("%s %v\n", spaces, varValue)
	}
}

// Uses fmt to print formats and prints the value of the variable and its caller.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - comments: you might want to put the name of the variable here (optional)
//
// It does not return any value.
func Print(varValue interface{}, comments ...string) {
	write(varValue, printer{}, comments...)
}

// Uses fmt to print formats and prints the value of the variable and all its callers.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - comments: you might want to put the name of the variable here (optional)
//
// It does not return any value.
func PrintConcurrent(varValue interface{}, comments ...string) {
	writeConcurrent(varValue, printer{}, comments...)
}

// Uses log to print formats and prints the value of the variable and its caller.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - comments: you might want to put the name of the variable here (optional)
//
// It does not return any value.
func Log(varValue interface{}, comments ...string) {
	write(varValue, logger{}, comments...)
}

// Uses log to print formats and prints the value of the variable and all its callers.
// The function takes two parameters:
//   - varValue: the value of the variable.
//   - comments: you might want to put the name of the variable here (optional)
//
// It does not return any value.
func LogConcurrent(varValue interface{}, comments ...string) {
	writeConcurrent(varValue, logger{}, comments...)
}
