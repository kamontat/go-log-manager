package main

import (
	"github.com/fatih/color"
	"github.com/kamontat/go-log-manager"
)

func main() {
	om.LoadLogger(om.NewDefaultLogger("default", &om.Setting{
		Color: true,
		Level: om.LLevelVerbose,
	}))

	om.LoadPrinter(om.NewDefaultPrinter("default", &om.Setting{
		Color: true,
		Level: om.PLevelUnneccessary,
	}))

	om.Log().ToDebug("debug", "debug message")

	om.Log().ToVerbose("verbose", "verbose message")
	om.Log().ToError("error", "error message")
	om.Log().ToWarn("warn", "warn message")
	om.Log().ToInfo("info", "info message")
	om.Log().ToLog("log", "log message")
	om.Log().ToDebug("debug", "debug message")

	om.Print().ToOneLine("one line", "one line message")
	om.Print().ToSilent("silent", "silent message")
	om.Print().ToNormal("normal", "normal "+color.BlueString("message"))
	om.Print().ToUnneccessary("unneccessary", "unneccessary message")

	om.Log().Setting().SetColor(false)
	om.Log().ToError("error", "no color error message")
	om.Print().ToNormal("normal", "no color normal "+color.BlueString("message"))

	om.Log().Setting().SetColor(true)
	om.Log().Setting().SetMaximumLevel(om.LLevelInfo)
	om.Log().ToError("error", "show error message")
	om.Log().ToWarn("warn", "show warn message")
	om.Log().ToInfo("info", "show info message")
	om.Log().ToLog("log", "not show log message")
	om.Log().ToDebug("debug", "not show debug message")
	om.Log().ToVerbose("verbose", "not show verbose message")
}
