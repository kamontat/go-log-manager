package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/kamontat/go-log-manager"
)

func exec(i func(), desc string) {
	om.Print = nil

	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fnName := strings.Replace(fn, ".", " ", -1)
	fnName = strings.Replace(fnName, "case", "case ", -1)
	final := strings.Title(fnName)

	fmt.Println(final + ") " + desc)
	fmt.Println("-------------------------")

	i()

	fmt.Println()
}

// Log without setup
func case1() {
	om.ConstrantExitOnError = false

	om.Log.ToVerbose("verbose", "verbose message")
}

// Normal log
func case2() {
	om.SetupNewLogger(&om.Setting{
		Color: true,
		Level: om.LLevelVerbose,
	})

	om.Log.ToVerbose("verbose", "verbose message")
	om.Log.ToDebug("debug", "debug message")
	om.Log.ToLog("log", "log message")
	om.Log.ToInfo("info", "info message")
	om.Log.ToWarn("warn", "warn message")
	om.Log.ToError("error", "error message")
}

func case3() {
	om.SetupNewPrinter(&om.Setting{
		Color: true,
		Level: om.PLevelUnneccessary,
	})

	om.Print.ToOneLine("one line", "one line message")
	om.Print.ToSilent("silent", "silent message")
	om.Print.ToNormal("normal", "normal "+color.BlueString("message"))
	om.Print.ToUnneccessary("unneccessary", "unneccessary message")
}

func main() {
	exec(case1, "log without setup")

	exec(case2, "normal log")

	exec(case3, "normal print")
}
