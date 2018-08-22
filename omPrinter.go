package om

// NewPrinter will create printer by setting
func NewPrinter(setting *Setting) *Output {
	setting.Type = Printer
	return New(setting)
}
