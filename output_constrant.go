package om

import (
	"fmt"
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
)
