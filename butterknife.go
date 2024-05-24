package butterknife

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func getRelativeFilePath(filePath string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		return filePath
	}
	relativePath, err := filepath.Rel(currentDir, filePath)
	if err != nil {
		return filePath
	}
	return relativePath
}

type writer interface {
	writef(format string, args ...any)
	write(args ...any)
	writeln(args ...any)
}

type logger struct{}

func (logger) writef(format string, args ...any) {
	log.Printf("\n"+format, args...)
}

func (logger) write(args ...any) {
	msg := fmt.Sprint(args...)
	log.Printf("\n%s", msg)
}

func (logger) writeln(args ...any) {
	msg := fmt.Sprintln(args...)
	log.Printf("\n%s", msg)
}

type printer struct{}

func (printer) writef(format string, args ...any) {
	fmt.Printf(format, args...)
}

func (printer) write(args ...any) {
	fmt.Print(args...)
}

func (printer) writeln(args ...any) {
	fmt.Println(args...)
}

func writef(wr writer, format string, args ...any) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller info")
		return
	}
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	file = getRelativeFilePath(file)
	header := fmt.Sprintf("[%s:%d %s]  ", file, line, funcName)
	msg := fmt.Sprintf(format, args...)
	wr.write(header + msg + "\n")
}

func write(wr writer, args ...any) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller info")
		return
	}
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	file = getRelativeFilePath(file)
	header := fmt.Sprintf("[%s:%d %s]  ", file, line, funcName)
	msg := fmt.Sprint(args...)
	wr.write(header + msg + "\n")
}

func writeln(wr writer, args ...any) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller info")
		return
	}
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	file = getRelativeFilePath(file)
	header := fmt.Sprintf("[%s:%d %s]  ", file, line, funcName)
	msg := fmt.Sprintln(args...)
	wr.write(header + msg)
}

func writelnHierarchy(wr writer, args ...any) {
	layers := []string{}
	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		pc, _, _, _ := runtime.Caller(i)
		funcName := runtime.FuncForPC(pc).Name()
		file = getRelativeFilePath(file)
		layers = append(layers, fmt.Sprintf("[%s:%d %s]", file, line, funcName))
	}
	msg := ""
	spaces := ""
	for i := len(layers) - 3; i >= 2; i-- {
		msg += fmt.Sprintf("%s%s\n", spaces, layers[i])
		spaces += "  "
	}
	msg += spaces
	msg += fmt.Sprintln(args...)
	wr.write(msg)
}

func writeHierarchy(wr writer, args ...any) {
	layers := []string{}
	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		pc, _, _, _ := runtime.Caller(i)
		funcName := runtime.FuncForPC(pc).Name()
		file = getRelativeFilePath(file)
		layers = append(layers, fmt.Sprintf("[%s:%d %s]", file, line, funcName))
	}
	msg := ""
	spaces := ""
	for i := len(layers) - 3; i >= 2; i-- {
		msg += fmt.Sprintf("%s%s\n", spaces, layers[i])
		spaces += "  "
	}
	msg += spaces
	msg += fmt.Sprint(args...)
	wr.write(msg + "\n")
}

func writefHierarchy(wr writer, format string, args ...any) {
	layers := []string{}
	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		pc, _, _, _ := runtime.Caller(i)
		funcName := runtime.FuncForPC(pc).Name()
		file = getRelativeFilePath(file)
		layers = append(layers, fmt.Sprintf("[%s:%d %s]", file, line, funcName))
	}
	msg := ""
	spaces := ""
	for i := len(layers) - 3; i >= 2; i-- {
		msg += fmt.Sprintf("%s%s\n", spaces, layers[i])
		spaces += "  "
	}
	msg += spaces
	msg += fmt.Sprintf(format, args...)
	wr.write(msg + "\n")
}

// Use it like fmt.Print. It will print the args and the caller info.
//   - args: the values to print.
//
// It does not return any value.
func Print(args ...any) {
	write(printer{}, args...)
}

// Use it like fmt.Printf. It will format and print the args and the caller info.
//   - format: the format string.
//   - args: the values to print.
//
// It does not return any value.
func Printf(format string, args ...any) {
	writef(printer{}, format, args...)
}

// Use it like fmt.Println. It will print the args and the caller info.
//   - args: the values to print.
//
// It does not return any value.
func Println(args ...any) {
	writeln(printer{}, args...)
}

// Use it like fmt.Print. It will print the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to print.
//
// It does not return any value.
func PrintHierarchy(args ...any) {
	writeHierarchy(printer{}, args...)
}

// Use it like fmt.Printf. It will format and print the args and its caller, as well as its caller's caller, and so on.
//   - format: the format string.
//   - args: the values to print.
//
// It does not return any value.
func PrintfHierarchy(format string, args ...any) {
	writefHierarchy(printer{}, format, args...)
}

// Use it like fmt.Println. It will print the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to print.
//
// It does not return any value.
func PrintlnHierarchy(args ...any) {
	writelnHierarchy(printer{}, args...)
}

// Use it like log.Print. It will log the args and the caller info.
//   - args: the values to log.
//
// It does not return any value.
func Log(args ...any) {
	write(logger{}, args...)
}

// Use it like log.Printf. It will format and log the args and the caller info.
//   - format: the format string.
//   - args: the values to log.
//
// It does not return any value.
func Logf(format string, args ...any) {
	writef(logger{}, format, args...)
}

// Use it like log.Println. It will log the args and the caller info.
//   - args: the values to log.
//
// It does not return any value.
func Logln(args ...any) {
	writeln(logger{}, args...)
}

// Use it like log.Print. It will log the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to log.
//
// It does not return any value.
func LogHierarchy(args ...any) {
	writeHierarchy(logger{}, args...)
}

// Use it like log.Printf. It will format and log the args and its caller, as well as its caller's caller, and so on.
//   - format: the format string.
//   - args: the values to log.
//
// It does not return any value.
func LogfHierarchy(format string, args ...any) {
	writefHierarchy(logger{}, format, args...)
}

// Use it like log.Println. It will log the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to log.
//
// It does not return any value.
func LoglnHierarchy(args ...any) {
	writelnHierarchy(logger{}, args...)
}
