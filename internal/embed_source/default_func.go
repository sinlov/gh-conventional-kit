package embed_source

import (
	"fmt"
	"github.com/aymerick/raymond"
	"math"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	DefaultFunctions = map[string]interface{}{
		"duration":       toDuration,
		"datetime":       toDatetime,
		"success":        isSuccess,
		"failure":        isFailure,
		"truncate":       truncate,
		"urlencode":      urlencode,
		"since":          since,
		"uppercasefirst": uppercaseFirst,
		"uppercase":      strings.ToUpper,
		"lowercase":      strings.ToLower,
		"regexReplace":   regexReplace,
	}
)

func toDuration(started, finished int64) string {
	return fmt.Sprint(time.Duration(finished-started) * time.Second)
}

func toDatetime(timestamp int64, layout, zone string) string {
	if len(zone) == 0 {
		return time.Unix(timestamp, 0).Format(layout)
	}

	loc, err := time.LoadLocation(zone)

	if err != nil {
		return time.Unix(timestamp, 0).Local().Format(layout)
	}

	return time.Unix(timestamp, 0).In(loc).Format(layout)
}

func isSuccess(conditional bool, options *raymond.Options) string {
	if !conditional {
		return options.Inverse()
	}

	switch options.ParamStr(0) {
	case "success":
		return options.Fn()
	default:
		return options.Inverse()
	}
}

func isFailure(conditional bool, options *raymond.Options) string {
	if !conditional {
		return options.Inverse()
	}

	switch options.ParamStr(0) {
	case "failure", "error", "killed":
		return options.Fn()
	default:
		return options.Inverse()
	}
}

func truncate(s string, len int) string {
	if utf8.RuneCountInString(s) <= int(math.Abs(float64(len))) {
		return s
	}

	runes := []rune(s)

	if len < 0 {
		len = -len
		return string(runes[len:])
	}

	return string(runes[:len])
}

func urlencode(options *raymond.Options) string {
	return url.QueryEscape(options.Fn())
}

func since(start int64) string {
	now := time.Unix(time.Now().Unix(), 0)
	return fmt.Sprint(now.Sub(time.Unix(start, 0)))
}

func uppercaseFirst(s string) string {
	a := []rune(s)

	a[0] = unicode.ToUpper(a[0])
	s = string(a)

	return s
}

func regexReplace(pattern string, input string, replacement string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(input, replacement)
}
