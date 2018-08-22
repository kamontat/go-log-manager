package om

import (
	"errors"

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

// LogWrapper for wrap the output to customize next chain
type LogWrapper struct {
	out *manager.ResultWrapper
}

// PrintWrapper for wrap the output to customize next chain
type PrintWrapper struct {
	out *manager.ResultWrapper
}

// Check will check output object in wrapper
func (l *LogWrapper) Check() *manager.Throwable {
	return l.out.Catch(func() error {
		return errors.New("logger is not exist, please call SetupLogger first")
	}, nil)
}

// unwrap will call force unwrap so be aware
func (l *LogWrapper) unwrap() *Output {
	return l.out.ForceUnwrap().(*Output)
}

// Check will check output object in wrapper
func (p *PrintWrapper) Check() *manager.Throwable {
	return p.out.Catch(func() error {
		return errors.New("printer is not exist, please call SetupPrinter first")
	}, nil)
}

// unwrap will call force unwrap so be aware
func (p *PrintWrapper) unwrap() *Output {
	return p.out.ForceUnwrap().(*Output)
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
func (l *LogWrapper) Setting() *Setting {
	return GetOutput(l).setting
}

// ToError log as error
func (l *LogWrapper) ToError(title string, message interface{}) {
	GetOutput(l).Print(LLevelError, title, message)
}

// ToWarn log as warn
func (l *LogWrapper) ToWarn(title string, message interface{}) {
	GetOutput(l).Print(LLevelWarn, title, message)
}

// ToWarning log as warn
func (l *LogWrapper) ToWarning(title string, message interface{}) {
	GetOutput(l).Print(LLevelWarn, title, message)
}

// ToInfo log as info
func (l *LogWrapper) ToInfo(title string, message interface{}) {
	GetOutput(l).Print(LLevelInfo, title, message)
}

// ToLog log as log
func (l *LogWrapper) ToLog(title string, message interface{}) {
	GetOutput(l).Print(LLevelLog, title, message)
}

// ToDebug log as debug
func (l *LogWrapper) ToDebug(title string, message interface{}) {
	GetOutput(l).Print(LLevelDebug, title, message)
}

// ToVerbose log as verbose
func (l *LogWrapper) ToVerbose(title string, message interface{}) {
	GetOutput(l).Print(LLevelVerbose, title, message)
}

// Setting will return current logger setting
func (p *PrintWrapper) Setting() *Setting {
	return GetOutput(p).setting
}

// ToOneLine print as oneline
func (p *PrintWrapper) ToOneLine(title string, message interface{}) {
	GetOutput(p).Print(PLevelOneLine, title, message)
}

// ToSilent print as slient
func (p *PrintWrapper) ToSilent(title string, message interface{}) {
	GetOutput(p).Print(PLevelSilent, title, message)
}

// ToNormal print as normal
func (p *PrintWrapper) ToNormal(title string, message interface{}) {
	GetOutput(p).Print(PLevelNormal, title, message)
}

// ToUnneccessary print as unneccesssary
func (p *PrintWrapper) ToUnneccessary(title string, message interface{}) {
	GetOutput(p).Print(PLevelUnneccessary, title, message)
}
