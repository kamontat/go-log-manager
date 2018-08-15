package om

var logger *LogWrapper
var printer *PrintWrapper

// LogWrapper for wrap the output to customize next chain
type LogWrapper struct {
	out *Output
}

// PrintWrapper for wrap the output to customize next chain
type PrintWrapper struct {
	out *Output
}

// LoadLogger will set log as output wrapper
func LoadLogger(output *Output) {
	logger = &LogWrapper{out: output}
}

// LoadPrinter will set print as output wrapper
func LoadPrinter(output *Output) {
	printer = &PrintWrapper{out: output}
}

// SetupLogger will create default logger with format and setting.
// note: format can be null
func SetupLogger(setting *Setting, format FormatMessage) {
	logger = &LogWrapper{out: NewCustomLogger("default", format, setting)}
}

// SetupPrinter will create default printer with format and setting.
// note: format can be null
func SetupPrinter(setting *Setting, format FormatMessage) {
	printer = &PrintWrapper{out: NewCustomPrinter("default", format, setting)}
}

// Log will return LogWrapper
func Log() *LogWrapper {
	return logger
}

// Setting will return current logger setting
func (l *LogWrapper) Setting() *Setting {
	return l.out.setting
}

// ToError log as error
func (l *LogWrapper) ToError(title string, message interface{}) {
	l.out.Print(LLevelError, title, message)
}

// ToWarn log as warn
func (l *LogWrapper) ToWarn(title string, message interface{}) {
	l.out.Print(LLevelWarn, title, message)
}

// ToWarning log as warn
func (l *LogWrapper) ToWarning(title string, message interface{}) {
	l.out.Print(LLevelWarn, title, message)
}

// ToInfo log as info
func (l *LogWrapper) ToInfo(title string, message interface{}) {
	l.out.Print(LLevelInfo, title, message)
}

// ToLog log as log
func (l *LogWrapper) ToLog(title string, message interface{}) {
	l.out.Print(LLevelLog, title, message)
}

// ToDebug log as debug
func (l *LogWrapper) ToDebug(title string, message interface{}) {
	l.out.Print(LLevelDebug, title, message)
}

// ToVerbose log as verbose
func (l *LogWrapper) ToVerbose(title string, message interface{}) {
	l.out.Print(LLevelVerbose, title, message)
}

// Print will return PrintWrapper
func Print() *PrintWrapper {
	return printer
}

// Setting will return current logger setting
func (p *PrintWrapper) Setting() *Setting {
	return p.out.setting
}

// ToOneLine print as oneline
func (p *PrintWrapper) ToOneLine(title string, message interface{}) {
	p.out.Print(PLevelOneLine, title, message)
}

// ToSilent print as slient
func (p *PrintWrapper) ToSilent(title string, message interface{}) {
	p.out.Print(PLevelSilent, title, message)
}

// ToNormal print as normal
func (p *PrintWrapper) ToNormal(title string, message interface{}) {
	p.out.Print(PLevelNormal, title, message)
}

// ToUnneccessary print as unneccesssary
func (p *PrintWrapper) ToUnneccessary(title string, message interface{}) {
	p.out.Print(PLevelUnneccessary, title, message)
}
