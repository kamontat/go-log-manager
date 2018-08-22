package om

// NewLogger will create logger by setting
func NewLogger(setting *Setting) *Output {
	setting.Type = Logger
	return New(setting)
}
