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
	writef(format string, args ...interface{})
	write(args ...interface{})
	writeln(args ...interface{})
}

type logger struct{}

func (l logger) writef(format string, args ...interface{}) {
	log.Printf("\n"+format, args...)
}
func (l logger) write(args ...interface{}) {
	msg := fmt.Sprint(args...)
	log.Printf("\n%s", msg)
}
func (l logger) writeln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	log.Printf("\n%s", msg)
}

type printer struct{}

func (p printer) writef(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (p printer) write(args ...interface{}) {
	fmt.Print(args...)
}
func (p printer) writeln(args ...interface{}) {
	fmt.Println(args...)
}

func writef(wr writer, format string, args ...interface{}) {
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

func write(wr writer, args ...interface{}) {
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

func writeln(wr writer, args ...interface{}) {
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

func writelnHierarchy(wr writer, args ...interface{}) {
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

func writeHierarchy(wr writer, args ...interface{}) {
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

func writefHierarchy(wr writer, format string, args ...interface{}) {
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
func Print(args ...interface{}) {
	write(printer{}, args...)
}

// Use it like fmt.Printf. It will format and print the args and the caller info.
//   - format: the format string.
//   - args: the values to print.
//
// It does not return any value.
func Printf(format string, args ...interface{}) {
	writef(printer{}, format, args...)
}

// Use it like fmt.Println. It will print the args and the caller info.
//   - args: the values to print.
//
// It does not return any value.
func Println(args ...interface{}) {
	writeln(printer{}, args...)
}

// Use it like fmt.Print. It will print the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to print.
//
// It does not return any value.
func PrintHierarchy(args ...interface{}) {
	writeHierarchy(printer{}, args...)
}

// Use it like fmt.Printf. It will format and print the args and its caller, as well as its caller's caller, and so on.
//   - format: the format string.
//   - args: the values to print.
//
// It does not return any value.
func PrintfHierarchy(format string, args ...interface{}) {
	writefHierarchy(printer{}, format, args...)
}

// Use it like fmt.Println. It will print the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to print.
//
// It does not return any value.
func PrintlnHierarchy(args ...interface{}) {
	writelnHierarchy(printer{}, args...)
}

// Use it like log.Print. It will log the args and the caller info.
//   - args: the values to log.
//
// It does not return any value.
func Log(args ...interface{}) {
	write(logger{}, args...)
}

// Use it like log.Printf. It will format and log the args and the caller info.
//   - format: the format string.
//   - args: the values to log.
//
// It does not return any value.
func Logf(format string, args ...interface{}) {
	writef(logger{}, format, args...)
}

// Use it like log.Println. It will log the args and the caller info.
//   - args: the values to log.
//
// It does not return any value.
func Logln(args ...interface{}) {
	writeln(logger{}, args...)
}

// Use it like log.Print. It will log the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to log.
//
// It does not return any value.
func LogHierarchy(args ...interface{}) {
	writeHierarchy(logger{}, args...)
}

// Use it like log.Printf. It will format and log the args and its caller, as well as its caller's caller, and so on.
//   - format: the format string.
//   - args: the values to log.
//
// It does not return any value.
func LogfHierarchy(format string, args ...interface{}) {
	writefHierarchy(logger{}, format, args...)
}

// Use it like log.Println. It will log the args and its caller, as well as its caller's caller, and so on.
//   - args: the values to log.
//
// It does not return any value.
func LoglnHierarchy(args ...interface{}) {
	writelnHierarchy(logger{}, args...)
}
