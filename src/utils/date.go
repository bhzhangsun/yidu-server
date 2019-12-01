package utils

import (
	"time"
)

func FormatData(format string) string {
	return time.Now().Format(format)
}
