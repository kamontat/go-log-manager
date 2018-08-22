package om

import (
	manager "github.com/kamontat/go-error-manager"
)

// Log is logger object
var Log = &LogWrapper{out: manager.WrapNil()}

// Print is printer object
var Print = &PrintWrapper{out: manager.WrapNil()}

// Checker interface for check Output is exist
type Checker interface {
	Check() *manager.Throwable
	unwrap() *Output
}

// GetOutput will return Output if exist, or exit the program
func GetOutput(c Checker) *Output {
	t := c.Check()
	t.ShowMessage()

	if ConstrantExitOnError {
		t.Exit()
	}

	if t.CanBeThrow() {
		return &Output{setting: &Setting{Level: LLevelNone}}
	}
	return c.unwrap()
}

// SetupExistLogger will set log as output wrapper
func SetupExistLogger(output *Output) {
	output.setting.Type = Logger
	Log = &LogWrapper{out: manager.Wrap(output)}
}

// SetupExistPrinter will set print as output wrapper
func SetupExistPrinter(output *Output) {
	output.setting.Type = Printer
	Print = &PrintWrapper{out: manager.Wrap(output)}
}

// SetupNewLogger will create default logger with format and setting.
// note: format can be null
func SetupNewLogger(setting *Setting) {
	Log = &LogWrapper{out: manager.Wrap(NewLogger(setting))}
}

// SetupNewPrinter will create default printer with format and setting.
// note: format can be null
func SetupNewPrinter(setting *Setting) {
	Print = &PrintWrapper{out: manager.Wrap(NewPrinter(setting))}
}

// Setting will return current logger setting
func (p *PrintWrapper) Setting() *Setting {
	return GetOutput(p).setting
}
