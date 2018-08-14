package om

// NewDefaultPrinter will create printer with default format and custom setting
func NewDefaultPrinter(name string, setting *Setting) *Output {
	setting.Type = Printer
	return New(name, nil, setting)
}

// NewCustomPrinter will create printer with custom format and custom setting
func NewCustomPrinter(name string, custom FormatMessage, setting *Setting) *Output {
	setting.Type = Printer
	return New(name, custom, setting)
}
