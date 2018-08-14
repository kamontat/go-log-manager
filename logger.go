package om

// NewDefaultLogger will create logger with default format and custom setting
func NewDefaultLogger(name string, setting *Setting) *Output {
	setting.Type = Logger
	return New(name, nil, setting)
}

// NewCustomLogger will create logger with custom format and custom setting
func NewCustomLogger(name string, custom FormatMessage, setting *Setting) *Output {
	setting.Type = Logger
	return New(name, custom, setting)
}
