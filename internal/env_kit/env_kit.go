package env_kit

import (
	"github.com/bar-counter/slog"
	"os"
	"strconv"
	"strings"
)

// FetchOsEnvBool
//
//	fetch os env by key.
//	if not found will return defValue.
//	return env not same as true (will be lowercase, so TRUE is same)
func FetchOsEnvBool(key string, defValue bool) bool {
	if os.Getenv(key) == "" {
		return defValue
	}
	return strings.ToLower(os.Getenv(key)) == "true"
}

// FetchOsEnvInt
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to int, return defValue
func FetchOsEnvInt(key string, defValue int) int {
	if os.Getenv(key) == "" {
		return defValue
	}
	outNum, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defValue
	}

	return outNum
}

// FetchOsEnvStr
//
//	fetch os env by key.
//	return not found will return defValue.
func FetchOsEnvStr(key, defValue string) string {
	if os.Getenv(key) == "" {
		return defValue
	}
	return os.Getenv(key)
}

// FetchOsEnvArray
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
func FetchOsEnvArray(key string) []string {
	var defValueStr []string
	if os.Getenv(key) == "" {
		return defValueStr
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	for _, item := range splitVal {
		defValueStr = append(defValueStr, strings.TrimSpace(item))
	}

	return defValueStr
}

// SetEnvStr
//
//	set env by key and val
func SetEnvStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		slog.Fatalf(err, "set env key [%v] string err: %v", key, err)
	}
}

// SetEnvBool
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvBool(key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		slog.Fatalf(err, "set env key [%v] bool err: %v", key, err)
	}
}

// SetEnvU64
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvU64(key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		slog.Fatalf(err, "set env key [%v] uint64 err: %v", key, err)
	}
}

// SetEnvInt64
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvInt64(key string, val int64) {
	err := os.Setenv(key, strconv.FormatInt(val, 10))
	if err != nil {
		slog.Fatalf(err, "set env key [%v] int64 err: %v", key, err)
	}
}
