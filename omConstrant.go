package om

import (
	"fmt"
	"os"
	"time"
)

var (
	// ConstrantLogSuffixName is suffix of log file
	ConstrantLogSuffixName = fmt.Sprintf("-%d.log", time.Now().Unix())
	// ConstrantPrintSuffixName is suffix of print file
	ConstrantPrintSuffixName = fmt.Sprintf("-%d.print", time.Now().Unix())

	// ConstrantDateFormat format of date in logger
	ConstrantDateFormat = "02/Jan/06, 03:04:05.000"

	// ConstrantLogTitleSize is the size of title in logger
	ConstrantLogTitleSize = "15"

	// ConstrantPrintTitleSize is the size of title in printer
	ConstrantPrintTitleSize = "15"

	// ConstrantExitOnError is the behavior when error occurred
	ConstrantExitOnError = true

	// ConstrantStandartOutput is the writer to be write on log or print
	ConstrantStandartOutput = os.Stdout
	// ConstrantStandartOutputLevelList is the list of level that will print to stdout
	ConstrantStandartOutputLevelList = []*Level{LLevelVerbose, LLevelDebug, LLevelLog, LLevelInfo, PLevelNormal, PLevelOneLine, PLevelSilent, PLevelUnneccessary}

	// ConstrantStandartError is the writer to be write on error or warning
	ConstrantStandartError = os.Stderr
	// ConstrantStandartErrorLevelList is the list of level that will print to stderr
	ConstrantStandartErrorLevelList = []*Level{LLevelError, LLevelWarn}
)
