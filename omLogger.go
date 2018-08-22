package om

import (
	"errors"

	manager "github.com/kamontat/go-error-manager"
)

// NewLogger will create logger by setting
func NewLogger(setting *Setting) *Output {
	setting.Type = Logger
	return New(setting)
}

// LogWrapper for wrap the output to customize next chain
type LogWrapper struct {
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
