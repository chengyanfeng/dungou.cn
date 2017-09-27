package def

import "time"
var Md5Salt string = "Dh@)!^o5l3!%Op0"
var MODE string = ""
var UPTIME = time.Now().UnixNano() / int64(time.Millisecond)
var PUNCTUATION []string = []string{".", ";", ",", "(", ")", "%"}

const (
	ROW_LIMIT_PREVIEW int = 50
	ROW_LIMIT_MAX     int = 1000
)

const (
	IP_REGEX string = "((?:(?:25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(?:25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d))))"
)

const (
	GENERAL_ERR int = 400
)
