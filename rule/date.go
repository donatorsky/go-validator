package rule

import (
	"time"
)

func Date() *dateFormatRule {
	return DateFormat(time.RFC3339Nano)
}
